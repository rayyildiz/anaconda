package anaconda

import (
	"fmt"
	"net/url"
	"strconv"
)

func (api TwitterApi) GetTweet(id int64, v url.Values) (tweet Tweet, err error) {
	v = cleanValues(v)
	v.Set("id", strconv.FormatInt(id, 10))

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/show.json", v, &tweet, _GET, responseCh}
	return tweet, (<-responseCh).err
}

func (api TwitterApi) GetTweetsLookupByIds(ids []int64, v url.Values) (tweet []Tweet, err error) {
	var pids string
	for w, i := range ids {
		pids += strconv.FormatInt(i, 10)
		if w != len(ids)-1 {
			pids += ","
		}
	}
	v = cleanValues(v)
	v.Set("id", pids)
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/lookup.json", v, &tweet, _GET, responseCh}
	return tweet, (<-responseCh).err
}

func (api TwitterApi) GetRetweets(id int64, v url.Values) (tweets []Tweet, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + fmt.Sprintf("/statuses/retweets/%d.json", id), v, &tweets, _GET, responseCh}
	return tweets, (<-responseCh).err
}

//PostTweet will create a tweet with the specified status message
func (api TwitterApi) PostTweet(status string, v url.Values) (tweet Tweet, err error) {
	v = cleanValues(v)
	v.Set("status", status)
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/statuses/update.json", v, &tweet, _POST, responseCh}
	return tweet, (<-responseCh).err
}

//DeleteTweet will destroy (delete) the status (tweet) with the specified ID, assuming that the authenticated user is the author of the status (tweet).
//If trimUser is set to true, only the user's Id will be provided in the user object returned.
func (api TwitterApi) DeleteTweet(id int64, trimUser bool) (tweet Tweet, err error) {
	v := url.Values{}
	if trimUser {
		v.Set("trim_user", "t")
	}
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + fmt.Sprintf("/statuses/destroy/%d.json", id), v, &tweet, _POST, responseCh}
	return tweet, (<-responseCh).err
}

//Retweet will retweet the status (tweet) with the specified ID.
//trimUser functions as in DeleteTweet
func (api TwitterApi) Retweet(id int64, trimUser bool) (rt Tweet, err error) {
	v := url.Values{}
	if trimUser {
		v.Set("trim_user", "t")
	}
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + fmt.Sprintf("/statuses/retweet/%d.json", id), v, &rt, _POST, responseCh}
	return rt, (<-responseCh).err
}

//UnRetweet will renove retweet Untweets a retweeted status.
//Returns the original Tweet with retweet details embedded.
//
//https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-statuses-unretweet-id
//trim_user: tweet returned in a timeline will include a user object
//including only the status authors numerical ID.
func (api TwitterApi) UnRetweet(id int64, trimUser bool) (rt Tweet, err error) {
	v := url.Values{}
	if trimUser {
		v.Set("trim_user", "t")
	}
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + fmt.Sprintf("/statuses/unretweet/%d.json", id), v, &rt, _POST, responseCh}
	return rt, (<-responseCh).err
}

// Favorite will favorite the status (tweet) with the specified ID.
// https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-favorites-create
func (api TwitterApi) Favorite(id int64) (rt Tweet, err error) {
	v := url.Values{}
	v.Set("id", fmt.Sprint(id))
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + fmt.Sprintf("/favorites/create.json"), v, &rt, _POST, responseCh}
	return rt, (<-responseCh).err
}

// Un-favorites the status specified in the ID parameter as the authenticating user.
// Returns the un-favorited status in the requested format when successful.
// https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-favorites-destroy
func (api TwitterApi) Unfavorite(id int64) (rt Tweet, err error) {
	v := url.Values{}
	v.Set("id", fmt.Sprint(id))
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + fmt.Sprintf("/favorites/destroy.json"), v, &rt, _POST, responseCh}
	return rt, (<-responseCh).err
}
