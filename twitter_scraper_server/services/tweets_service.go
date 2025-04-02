package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"twitter_scraper_server/config"

	twitterscraper "github.com/imperatrona/twitter-scraper"
	"github.com/sirupsen/logrus"
)

const bearerToken2 = "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

type TweetsService struct {
	cfg      *config.Config
	scrapers []*twitterscraper.Scraper
	as       *AccountService
}

func NewTweetsService(cfg *config.Config, as *AccountService) *TweetsService {
	scrapers := make([]*twitterscraper.Scraper, 0)
	scrapers = append(scrapers, twitterscraper.New())

	ts := &TweetsService{cfg, scrapers, as}

	return ts
}

func (cs *TweetsService) GetUserTweetsAndReplies(encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/bt4TKuFz4T7Ckk-VvQVSow/UserTweetsAndReplies", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetUserTweets(encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/UGi7tjRPr-d_U3bCPIko5Q/UserTweets", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) FetchTweetsByUserIDLegacy(userID string, encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://api.twitter.com/2/timeline/profile/"+userID+".json", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetTweet1(id string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://api.twitter.com/2/timeline/conversation/"+id+".json", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetTweet2(encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/VWFGPVAGkZMGRKGe3GFFnA/TweetDetail", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery
	curBearerToken := s.GetBearerToken()
	defer s.SetBearerToken(curBearerToken)

	if curBearerToken != bearerToken2 {
		s.SetBearerToken(bearerToken2)
	}

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetTweetResultByRestId(encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/xBtHv5-Xsk268T5ng_OGNg/TweetResultByRestId", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery
	curBearerToken := s.GetBearerToken()
	defer s.SetBearerToken(curBearerToken)

	if curBearerToken != bearerToken2 {
		s.SetBearerToken(bearerToken2)
	}

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetHomeLatestTimeline(encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/9EwYy8pLBOSFlEoSP2STiQ/HomeLatestTimeline", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetHomeTimeline(encodeQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/1u0Wlkw6Ru1NwBUD-pDiww/HomeTimeline", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodeQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (cs *TweetsService) GetBookmarks(encodeQuery string) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/-IyJFt9_jS_9d_vS3NN-fA/Bookmarks", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (cs *TweetsService) GetFollowing(encodeQuery string) (*twitterscraper.TimelineV2, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/g5P4cbXR4ta4oCeE7y2vLQ/Following", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (cs *TweetsService) GetFollowers(encodeQuery string) (*twitterscraper.TimelineV2, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/jwbfbSzn0FRL_AMZGsYDag/Followers", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (cs *TweetsService) GetUserMedia(encodeQuery string) (*twitterscraper.TimelineV2, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/2tLOJWwGuCTytDrGBg8VwQ/UserMedia", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (cs *TweetsService) GetUserByScreenName(encodeQuery string) (*twitterscraper.User, error) {
	var jsn twitterscraper.User

	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://api.twitter.com/graphql/Yka-W8dz7RaEuQNkroPkYw/UserByScreenName", nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return nil, err
	}
	return &jsn, nil
}

func (cs *TweetsService) GetUserByRestId(encodeQuery string) (*twitterscraper.User, error) {
	var jsn twitterscraper.User

	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://twitter.com/i/api/graphql/Qw77dDjp9xCpUY-AXwt-yQ/UserByRestId", nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return nil, err
	}

	return &jsn, nil
}

func (cs *TweetsService) GetTweetDetail(encodeQuery string) (*twitterscraper.ThreadedConversation, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/ldqoq5MmFHN1FhMGvzC9Jg/TweetDetail", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	var threads twitterscraper.ThreadedConversation

	err = s.RequestAPI(req, &threads)
	if err != nil {
		return nil, err
	}

	return &threads, nil
}

func (cs *TweetsService) GetFetchScheduledTweets(encodeQuery string) (*twitterscraper.ScheduleTweets, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/ITtjAzvlZni2wWXwf295Qg/FetchScheduledTweets", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodeQuery

	var timeline twitterscraper.ScheduleTweets
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

// DeleteScheduledTweet removes tweet from scheduled.
func (cs *TweetsService) DeleteScheduledTweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}
	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/CTOVqej0JBXAZSwkp1US0g/DeleteScheduledTweet", true)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (cs *TweetsService) CreateScheduledTweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/LCVzRQGxOaGnOnYH01NQXg/CreateScheduledTweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) SearchTimeline(encodedQuery string) (*twitterscraper.SearchTimeline, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/nK1dw4oV3k4w5TdtcAdSww/SearchTimeline", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline twitterscraper.SearchTimeline
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}
	return &timeline, nil
}

func (cs *TweetsService) GetAudioSpaceById(encodedQuery string) (*twitterscraper.SpaceData, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/d03OdorPdZ_sH9V3D1_yWQ/AudioSpaceById", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var spaceData twitterscraper.SpaceData
	err = s.RequestAPI(req, &spaceData)
	if err != nil {
		return nil, err
	}
	return &spaceData, nil
}

func (cs *TweetsService) GetGuide(encodedQuery string) (*twitterscraper.TimelineV1, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return nil, err
	}

	req, err := newRequest("GET", "https://api.twitter.com/2/guide.json", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline twitterscraper.TimelineV1
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}
	return &timeline, nil
}

func (cs *TweetsService) CreateTweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/oB-5XsHNAbjvARJEc8CZFw/CreateTweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	b, _ := json.Marshal(payload)
	req.Body = io.NopCloser(bytes.NewReader(b))

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) DeleteTweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/VaenaVgh5q5ih7kvyVjgtg/DeleteTweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	b, _ := json.Marshal(payload)
	req.Body = io.NopCloser(bytes.NewReader(b))

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) CreateRetweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/ojPdsZsimiJrUGLR1sjUtA/CreateRetweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	b, _ := json.Marshal(payload)
	req.Body = io.NopCloser(bytes.NewReader(b))

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) DeleteRetweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/iQtK4dl5hBmXewYZuEOKVw/DeleteRetweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	b, _ := json.Marshal(payload)
	req.Body = io.NopCloser(bytes.NewReader(b))

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) FavoriteTweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	b, _ := json.Marshal(payload)
	req.Body = io.NopCloser(bytes.NewReader(b))

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) UnfavoriteTweet(payload map[string]interface{}) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/ZYKSe-w7KEslx3JhSIk5LA/UnfavoriteTweet", true)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	b, _ := json.Marshal(payload)
	req.Body = io.NopCloser(bytes.NewReader(b))

	var response interface{}
	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cs *TweetsService) GetRetweeters(encodedQuery string) (interface{}, error) {
	s, account, err := cs.as.GetAuthenticatedScraper()
	if err != nil {
		logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
		return twitterscraper.ScheduleTweets{}, err
	}

	req, err := newRequest("POST", "https://twitter.com/i/api/graphql/8019obfgnveiPiJuS2Rtow/Retweeters", true)
	if err != nil {
		return "", err
	}

	req.URL.RawQuery = encodedQuery
	var response interface{}

	err = s.RequestAPI(req, &response)
	if err != nil {
		return "", err
	}

	return response, nil
}
