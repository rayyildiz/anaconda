package anaconda

import (
	"net/url"
)

// GetHomeTimeline returns the most recent tweets and retweets posted by the user
// and the users that they follow.
// https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-home_timeline
// By default, include_entities is set to "true"
func (api TwitterApi) GetHomeTimeline(v url.Values) (timeline []Tweet, err error) {
	v = cleanValues(v)
	if val := v.Get("include_entities"); val == "" {
		v.Set("include_entities", "true")
	}

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/home_timeline.json", v, &timeline, _GET, responseCh}
	return timeline, (<-responseCh).err
}

// GetUserTimeline returns a collection of the most recent Tweets posted by the user indicated by the screen_name or user_id parameters.
// https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline
func (api TwitterApi) GetUserTimeline(v url.Values) (timeline []Tweet, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/user_timeline.json", v, &timeline, _GET, responseCh}
	return timeline, (<-responseCh).err
}

// GetMentionsTimeline returns the most recent mentions (Tweets containing a users’s @screen_name) for the authenticating user.
// The timeline returned is the equivalent of the one seen when you view your mentions on twitter.com.
// https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-mentions_timeline
func (api TwitterApi) GetMentionsTimeline(v url.Values) (timeline []Tweet, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/mentions_timeline.json", v, &timeline, _GET, responseCh}
	return timeline, (<-responseCh).err
}

// GetRetweetsOfMe returns the most recent Tweets authored by the authenticating user that have been retweeted by others.
// https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-retweets_of_me
func (api TwitterApi) GetRetweetsOfMe(v url.Values) (tweets []Tweet, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/retweets_of_me.json", v, &tweets, _GET, responseCh}
	return tweets, (<-responseCh).err
}
