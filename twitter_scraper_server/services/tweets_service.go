package services

import (
	"net/url"
	"twitter_scraper_server/config"

	twitterscraper "github.com/imperatrona/twitter-scraper"
)

type TweetsService struct {
	cfg      *config.Config
	scrapers []*twitterscraper.Scraper
}

func NewTweetsService(cfg *config.Config) *TweetsService {
	scrapers := make([]*twitterscraper.Scraper, 0)
	scrapers = append(scrapers, twitterscraper.New())

	ts := &TweetsService{cfg: cfg, scrapers: scrapers}

	return ts
}

func (cs *TweetsService) GetNextScraper() *twitterscraper.Scraper {
	return cs.scrapers[0]
}

func (cs *TweetsService) GetUserTweets(userID string, maxTweetsNbr int, cursor string) (twitterscraper.TimelineV2, error) {
	// get from cache

	// if not in cache, fetch from twitter
	s := cs.GetNextScraper()
	req, err := newRequest("GET", "https://twitter.com/i/api/graphql/UGi7tjRPr-d_U3bCPIko5Q/UserTweets", true)
	if err != nil {
		return twitterscraper.TimelineV2{}, err
	}

	variables := map[string]interface{}{
		"userId":                                 userID,
		"count":                                  maxTweetsNbr,
		"includePromotedContent":                 false,
		"withQuickPromoteEligibilityTweetFields": false,
		"withVoice":                              true,
		"withV2Timeline":                         true,
	}
	features := map[string]interface{}{
		"rweb_lists_timeline_redesign_enabled":                              true,
		"responsive_web_graphql_exclude_directive_enabled":                  true,
		"verified_phone_label_enabled":                                      false,
		"creator_subscriptions_tweet_preview_api_enabled":                   true,
		"responsive_web_graphql_timeline_navigation_enabled":                true,
		"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
		"tweetypie_unmention_optimization_enabled":                          true,
		"vibe_api_enabled":                                                        true,
		"responsive_web_edit_tweet_api_enabled":                                   true,
		"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
		"view_counts_everywhere_api_enabled":                                      true,
		"longform_notetweets_consumption_enabled":                                 true,
		"tweet_awards_web_tipping_enabled":                                        false,
		"freedom_of_speech_not_reach_fetch_enabled":                               true,
		"standardized_nudges_misinfo":                                             true,
		"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": false,
		"interactive_text_enabled":                                                true,
		"responsive_web_text_conversations_enabled":                               false,
		"longform_notetweets_rich_text_read_enabled":                              true,
		"longform_notetweets_inline_media_enabled":                                false,
		"responsive_web_enhance_cards_enabled":                                    false,
	}

	if cursor != "" {
		variables["cursor"] = cursor
	}

	query := url.Values{}
	query.Set("variables", mapToJSONString(variables))
	query.Set("features", mapToJSONString(features))
	req.URL.RawQuery = query.Encode()

	var timeline twitterscraper.TimelineV2
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return twitterscraper.TimelineV2{}, err
	}

	// save to cache

	// return value
	return twitterscraper.TimelineV2{}, nil
}
