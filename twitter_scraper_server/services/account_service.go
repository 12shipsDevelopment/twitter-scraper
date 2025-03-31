package services

import (
	"twitter_scraper_server/config"

	twitterscraper "github.com/imperatrona/twitter-scraper"
)

type AccountService struct {
	cfg      *config.Config
	scrapers []*twitterscraper.Scraper
}

func NewAccountService(cfg *config.Config) *AccountService {
	scrapers := make([]*twitterscraper.Scraper, 0)
	scrapers = append(scrapers, twitterscraper.New())

	as := &AccountService{cfg: cfg, scrapers: scrapers}

	return as
}

func (as *AccountService) GetNextScraper() *twitterscraper.Scraper {
	return as.scrapers[0]
}

func (as *AccountService) GetAccountSettings() (twitterscraper.AccountSettings, error) {
	var settings twitterscraper.AccountSettings
	s := as.GetNextScraper()
	req, err := newRequest("GET", "https://api.twitter.com/1.1/account/settings.json", true)
	if err != nil {
		return settings, err
	}

	err = s.RequestAPI(req, &settings)
	return settings, err
}

func (as *AccountService) GetAccountList() ([]twitterscraper.Account, error) {
	var list twitterscraper.AccountList
	s := as.GetNextScraper()
	req, err := newRequest("GET", "https://api.twitter.com/1.1/account/multi/list.json", true)
	if err != nil {
		return list.Users, err
	}

	err = s.RequestAPI(req, &list)
	return list.Users, err
}
