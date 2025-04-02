package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"twitter_scraper_server/config"

	twitterscraper "github.com/imperatrona/twitter-scraper"
)

const bearerToken2 = "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

type TweetsService struct {
	cfg      *config.Config
	scrapers []*twitterscraper.Scraper
}

func NewTweetsService(cfg *config.Config) *TweetsService {
	scrapers := make([]*twitterscraper.Scraper, 0)
	scrapers = append(scrapers, twitterscraper.New())

	ts := &TweetsService{cfg, scrapers}

	return ts
}

func (ts *TweetsService) GetUserTweetsAndReplies(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/bt4TKuFz4T7Ckk-VvQVSow/UserTweetsAndReplies", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (ts *TweetsService) GetUserTweets(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/UGi7tjRPr-d_U3bCPIko5Q/UserTweets", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (ts *TweetsService) FetchTweetsByUserIDLegacy(userID string, s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://api.twitter.com/2/timeline/profile/"+userID+".json", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (ts *TweetsService) GetTweet1(s *twitterscraper.Scraper, id string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
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

func (ts *TweetsService) GetTweet2(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/VWFGPVAGkZMGRKGe3GFFnA/TweetDetail", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery
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

func (ts *TweetsService) GetTweetResultByRestId(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/xBtHv5-Xsk268T5ng_OGNg/TweetResultByRestId", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery
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

func (ts *TweetsService) GetHomeLatestTimeline(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/9EwYy8pLBOSFlEoSP2STiQ/HomeLatestTimeline", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (ts *TweetsService) GetHomeTimeline(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	// get from cache

	// if not in cache, fetch from twitter
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/1u0Wlkw6Ru1NwBUD-pDiww/HomeTimeline", true)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	req.URL.RawQuery = encodedQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	// save to cache

	// return value
	return &timeline, nil
}

func (ts *TweetsService) GetBookmarks(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/-IyJFt9_jS_9d_vS3NN-fA/Bookmarks", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline interface{}
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (ts *TweetsService) GetFollowing(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.TimelineV2, error) {
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/g5P4cbXR4ta4oCeE7y2vLQ/Following", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (ts *TweetsService) GetFollowers(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.TimelineV2, error) {
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/jwbfbSzn0FRL_AMZGsYDag/Followers", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (ts *TweetsService) GetUserMedia(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.TimelineV2, error) {
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/2tLOJWwGuCTytDrGBg8VwQ/UserMedia", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

func (ts *TweetsService) GetUserByScreenName(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.User, error) {
	var jsn twitterscraper.User

	req, err := http.NewRequest("GET", "https://api.twitter.com/graphql/Yka-W8dz7RaEuQNkroPkYw/UserByScreenName", nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return nil, err
	}
	return &jsn, nil
}

func (ts *TweetsService) GetUserByRestId(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.User, error) {
	var jsn twitterscraper.User

	req, err := http.NewRequest("GET", "https://twitter.com/i/api/graphql/Qw77dDjp9xCpUY-AXwt-yQ/UserByRestId", nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return nil, err
	}

	return &jsn, nil
}

func (ts *TweetsService) GetTweetDetail(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.ThreadedConversation, error) {
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/ldqoq5MmFHN1FhMGvzC9Jg/TweetDetail", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var threads twitterscraper.ThreadedConversation

	err = s.RequestAPI(req, &threads)
	if err != nil {
		return nil, err
	}

	return &threads, nil
}

func (ts *TweetsService) GetFetchScheduledTweets(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.ScheduleTweets, error) {
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/ITtjAzvlZni2wWXwf295Qg/FetchScheduledTweets", true)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = encodedQuery

	var timeline twitterscraper.ScheduleTweets
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}

// DeleteScheduledTweet removes tweet from scheduled.
func (ts *TweetsService) DeleteScheduledTweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) CreateScheduledTweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) SearchTimeline(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.SearchTimeline, error) {
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

func (ts *TweetsService) GetAudioSpaceById(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.SpaceData, error) {
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

func (ts *TweetsService) GetGuide(s *twitterscraper.Scraper, encodedQuery string) (*twitterscraper.TimelineV1, error) {
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

func (ts *TweetsService) CreateTweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) DeleteTweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) CreateRetweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) DeleteRetweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) FavoriteTweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {
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

func (ts *TweetsService) UnfavoriteTweet(s *twitterscraper.Scraper, payload map[string]interface{}) (interface{}, error) {

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

func (ts *TweetsService) GetRetweeters(s *twitterscraper.Scraper, encodedQuery string) (interface{}, error) {
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
