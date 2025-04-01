package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"twitter_scraper_server/config"

	twitterscraper "github.com/imperatrona/twitter-scraper"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type TwitterAccount struct {
	Username         string
	Password         string
	TwoFACode        string
	RateLimitedUntil time.Time
	LastScraped      time.Time
	LoginStatus      string
}

const (
	minSleepDuration  = 500 * time.Millisecond
	maxSleepDuration  = 2 * time.Second
	RateLimitDuration = 15 * time.Minute
	consumerKey       = "3nVuSoBZnx6U4vzUxf5w"
	consumerSecret    = "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys"
)

var (
	rng *rand.Rand
)

type TwitterAccountManager struct {
	accounts []*TwitterAccount
	index    int
	mutex    sync.Mutex
}

func NewTwitterAccountManager(accounts []*TwitterAccount) *TwitterAccountManager {
	return &TwitterAccountManager{
		accounts: accounts,
		index:    0,
	}
}

func (manager *TwitterAccountManager) GetNextAccount() *TwitterAccount {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	for i := 0; i < len(manager.accounts); i++ {
		account := manager.accounts[manager.index]
		manager.index = (manager.index + 1) % len(manager.accounts)
		if time.Now().After(account.RateLimitedUntil) {
			return account
		}
	}
	return nil
}

func (manager *TwitterAccountManager) MarkAccountRateLimited(account *TwitterAccount) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	account.RateLimitedUntil = time.Now().Add(GetRateLimitDuration())
}

func (manager *TwitterAccountManager) getAccountsCount() int {
	return len(manager.accounts)
}

func (manager *TwitterAccountManager) handleRateLimit(err error, account *TwitterAccount) bool {
	if strings.Contains(err.Error(), "Rate limit exceeded") {
		manager.MarkAccountRateLimited(account)
		logrus.Warnf("rate limited: %s", account.Username)
		return true
	}
	return false
}

func loadAccountsFromConfig() []*TwitterAccount {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading .env file: %v", err)
	}

	accountsEnv := os.Getenv("TWITTER_ACCOUNTS")
	if accountsEnv == "" {
		logrus.Fatal("TWITTER_ACCOUNTS not set in .env file")
	}

	return parseAccounts(strings.Split(accountsEnv, ","))
}

func parseAccounts(accountPairs []string) []*TwitterAccount {
	return filterMap(accountPairs, func(pair string) (*TwitterAccount, bool) {
		credentials := strings.Split(pair, ":")
		if len(credentials) != 2 {
			logrus.Warnf("invalid account credentials: %s", pair)
			return nil, false
		}
		return &TwitterAccount{
			Username: strings.TrimSpace(credentials[0]),
			Password: strings.TrimSpace(credentials[1]),
		}, true
	})
}

// AccountState holds the state of a Twitter account
type AccountState struct {
	Username         string
	IsRateLimited    bool
	RateLimitedUntil time.Time
	LastScraped      time.Time
	LoginStatus      string // e.g., "Successful", "Please verify", "Failed - [Reason]"
}

// GetAccountStates returns the state of all Twitter accounts
func (manager *TwitterAccountManager) GetAccountStates() []AccountState {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	states := make([]AccountState, len(manager.accounts))
	for i, account := range manager.accounts {
		states[i] = AccountState{
			Username:         account.Username,
			IsRateLimited:    time.Now().Before(account.RateLimitedUntil),
			RateLimitedUntil: account.RateLimitedUntil,
			LastScraped:      account.LastScraped,
			LoginStatus:      account.LoginStatus,
		}
	}
	return states
}

func (manager *TwitterAccountManager) GetAccountByUsername(username string) *TwitterAccount {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	for _, account := range manager.accounts {
		if account.Username == username {
			return account
		}
	}
	return nil
}

type AccountService struct {
	cfg            *config.Config
	scrapers       []*twitterscraper.Scraper
	accountManager *TwitterAccountManager
}

func NewAccountService(cfg *config.Config) *AccountService {
	scrapers := make([]*twitterscraper.Scraper, 0)
	scrapers = append(scrapers, twitterscraper.New())
	accounts := loadAccountsFromConfig()
	accountManager := NewTwitterAccountManager(accounts)

	as := &AccountService{cfg, scrapers, accountManager}

	return as
}

func (as *AccountService) GetNextAuthenticatedScraper() (*twitterscraper.Scraper, *TwitterAccount, error) {
	baseDir := "/root/.masa"

	account := as.accountManager.GetNextAccount()
	if account == nil {
		return nil, nil, fmt.Errorf("all accounts are rate-limited")
	}
	scraper := NewScraper(account, baseDir)
	if scraper == nil {
		err := fmt.Errorf("twitter authentication failed for %s", account.Username)
		logrus.Error(err)
		return nil, account, err
	}
	account.LastScraped = time.Now()
	return scraper, account, nil
}

func (as *AccountService) GetAccountSettings() (twitterscraper.AccountSettings, error) {
	var settings twitterscraper.AccountSettings
	s, account, err := as.GetNextAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.AccountSettings{}, err
	}
	req, err := newRequest("GET", "https://api.twitter.com/1.1/account/settings.json", true)
	if err != nil {
		return settings, err
	}

	err = s.RequestAPI(req, &settings)
	return settings, err
}

func (as *AccountService) GetAccountList() ([]twitterscraper.Account, error) {
	var list twitterscraper.AccountList
	s, account, err := as.GetNextAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return []twitterscraper.Account{}, err
	}
	req, err := newRequest("GET", "https://api.twitter.com/1.1/account/multi/list.json", true)
	if err != nil {
		return list.Users, err
	}

	err = s.RequestAPI(req, &list)
	return list.Users, err
}

func (as *AccountService) GetFlow(data map[string]interface{}) (*twitterscraper.Flow, error) {
	s, account, err := as.GetNextAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	headers := http.Header{
		"Authorization":             []string{"Bearer " + s.GetBearerToken()},
		"Content-Type":              []string{"application/json"},
		"User-Agent":                []string{s.GetUserAgent()},
		"X-Guest-Token":             []string{s.GetGuestTokenStr()},
		"X-Twitter-Auth-Type":       []string{"OAuth2Client"},
		"X-Twitter-Active-User":     []string{"yes"},
		"X-Twitter-Client-Language": []string{"en"},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/onboarding/task.json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header = headers
	s.SetCSRFToken(req)

	var flow twitterscraper.Flow
	err = s.RequestAPI(req, &flow)
	return &flow, err
}

func (as *AccountService) Logout() error {
	logrus.Warn("no need to logout...")
	return nil
}

func (as *AccountService) GetAccessToken() (string, error) {
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(consumerKey, consumerSecret)

	s, account, err := as.GetNextAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return "", err
	}

	var a struct {
		AccessToken string `json:"access_token"`
	}
	err = s.RequestAPI(req, &a)

	return a.AccessToken, err
}

func NewScraper(account *TwitterAccount, cookieDir string) *twitterscraper.Scraper {
	scraper := twitterscraper.New()
	http_proxy := os.Getenv("http_proxy")
	if http_proxy != "" {
		scraper.SetProxy(http_proxy)
	}
	if err := LoadCookies(scraper, account, cookieDir); err == nil {
		logrus.Debugf("Cookies loaded for user %s.", account.Username)
		if scraper.IsLoggedIn() {
			logrus.Debugf("Already logged in as %s.", account.Username)
			return scraper
		}
	}

	RandomSleep()

	if err := scraper.Login(account.Username, account.Password); err != nil {
		account.LoginStatus = fmt.Sprintf("Failed - %v", err)
		logrus.WithError(err).Warnf("Login failed for %s", account.Username)
		return nil
	}

	RandomSleep()

	if err := SaveCookies(scraper, account, cookieDir); err != nil {
		logrus.WithError(err).Errorf("Failed to save cookies for %s", account.Username)
	}

	logrus.Infof("Login successful for %s", account.Username)

	return scraper
}

func SaveCookies(scraper *twitterscraper.Scraper, account *TwitterAccount, baseDir string) error {
	logrus.Debugf("Saving cookies for user %s", account.Username)
	cookieFile := filepath.Join(baseDir, fmt.Sprintf("%s_twitter_cookies.json", account.Username))
	cookies := scraper.GetCookies()
	logrus.Debugf("Got %d cookies to save", len(cookies))

	data, err := json.Marshal(cookies)
	if err != nil {
		return fmt.Errorf("error marshaling cookies: %v", err)
	}

	logrus.Debugf("Writing cookies to file: %s", cookieFile)
	if err = os.WriteFile(cookieFile, data, 0644); err != nil {
		return fmt.Errorf("error saving cookies: %v", err)
	}
	logrus.Debug("Successfully saved cookies")
	return nil
}

func LoadCookies(scraper *twitterscraper.Scraper, account *TwitterAccount, baseDir string) error {
	logrus.Debugf("Loading cookies for user %s", account.Username)
	cookieFile := filepath.Join(baseDir, fmt.Sprintf("%s_twitter_cookies.json", account.Username))

	logrus.Debugf("Reading cookie file: %s", cookieFile)
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return fmt.Errorf("error reading cookies: %v", err)
	}

	var cookies []*http.Cookie
	if err = json.Unmarshal(data, &cookies); err != nil {
		return fmt.Errorf("error unmarshaling cookies: %v", err)
	}
	logrus.Debugf("Loaded %d cookies from file", len(cookies))

	// Verify critical cookies are present
	var hasAuthToken, hasCSRFToken bool
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			hasAuthToken = true
			logrus.Debug("Found auth_token cookie")
		}
		if cookie.Name == "ct0" {
			hasCSRFToken = true
			logrus.Debug("Found CSRF token cookie")
		}
	}

	if !hasAuthToken || !hasCSRFToken {
		logrus.Debug("Missing critical authentication cookies")
		return fmt.Errorf("missing critical authentication cookies")
	}

	logrus.Debug("Setting cookies in scraper")
	scraper.SetCookies(cookies)
	logrus.Debug("Successfully loaded and set cookies")
	return nil
}
