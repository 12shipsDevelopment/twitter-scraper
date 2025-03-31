package services

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func newRequest(method string, url string, withRelies bool) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("include_profile_interstitial_type", "1")
	q.Add("include_blocking", "1")
	q.Add("include_blocked_by", "1")
	q.Add("include_followed_by", "1")
	q.Add("include_want_retweets", "1")
	q.Add("include_mute_edge", "1")
	q.Add("include_can_dm", "1")
	q.Add("include_can_media_tag", "1")
	q.Add("include_ext_has_nft_avatar", "1")
	q.Add("include_ext_is_blue_verified", "1")
	q.Add("include_ext_verified_type", "1")
	q.Add("skip_status", "1")
	q.Add("cards_platform", "Web-12")
	q.Add("include_cards", "1")
	q.Add("include_ext_alt_text", "true")
	q.Add("include_ext_limited_action_results", "false")
	q.Add("include_quote_count", "true")
	q.Add("include_reply_count", "1")
	q.Add("tweet_mode", "extended")
	q.Add("include_ext_collab_control", "true")
	q.Add("include_ext_views", "true")
	q.Add("include_entities", "true")
	q.Add("include_user_entities", "true")
	q.Add("include_ext_media_color", "true")
	q.Add("include_ext_media_availability", "true")
	q.Add("include_ext_sensitive_media_warning", "true")
	q.Add("include_ext_trusted_friends_metadata", "true")
	q.Add("send_error_codes", "true")
	q.Add("simple_quoted_tweet", "true")
	q.Add("include_tweet_replies", strconv.FormatBool(withRelies))
	q.Add("ext", "mediaStats,highlightedLabel,hasNftAvatar,voiceInfo,birdwatchPivot,enrichments,superFollowMetadata,unmentionInfo,editControl,collab_control,vibe")
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func mapToJSONString(data map[string]interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
