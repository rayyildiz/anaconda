package anaconda

import (
	"net/url"
	"strconv"
)

type Cursor struct {
	Previous_cursor     int64
	Previous_cursor_str string

	Ids []int64

	Next_cursor     int64
	Next_cursor_str string
}

type UserCursor struct {
	Previous_cursor     int64
	Previous_cursor_str string
	Next_cursor         int64
	Next_cursor_str     string
	Users               []User
}

type FriendsIdsCursor struct {
	Previous_cursor     int64
	Previous_cursor_str string
	Next_cursor         int64
	Next_cursor_str     string
	Ids                 []int64
}

type FriendsIdsPage struct {
	Ids   []int64
	Error error
}

type Friendship struct {
	Name        string
	Id_str      string
	Id          int64
	Connections []string
	Screen_name string
}

type FollowersPage struct {
	Followers []User
	Error     error
}

type FriendsPage struct {
	Friends []User
	Error   error
}

// FIXME: Might want to consolidate this with FriendsIdsPage and just
//		  have "UserIdsPage".
type FollowersIdsPage struct {
	Ids   []int64
	Error error
}

// GetFriendshipsNoRetweets returns a collection of user_ids that the currently authenticated user does not want to receive retweets from.
// It does not currently support the stringify_ids parameter.
func (api TwitterApi) GetFriendshipsNoRetweets() (ids []int64, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/no_retweets/ids.json", nil, &ids, _GET, responseCh}
	return ids, (<-responseCh).err
}

func (api TwitterApi) GetFollowersIds(v url.Values) (c Cursor, err error) {
	err = api.apiGet(api.baseUrl+"/followers/ids.json", v, &c)
	return
}

// Like GetFollowersIds, but returns a channel instead of a cursor and pre-fetches the remaining results
// This channel is closed once all values have been fetched
func (api TwitterApi) GetFollowersIdsAll(v url.Values) (result chan FollowersIdsPage) {
	result = make(chan FollowersIdsPage)

	v = cleanValues(v)
	go func(a TwitterApi, v url.Values, result chan FollowersIdsPage) {
		// Cursor defaults to the first page ("-1")
		nextCursor := "-1"
		for {
			v.Set("cursor", nextCursor)
			c, err := a.GetFollowersIds(v)

			// throttledQuery() handles all rate-limiting errors
			// if GetFollowersList() returns an error, it must be api different kind of error

			result <- FollowersIdsPage{c.Ids, err}

			nextCursor = c.Next_cursor_str
			if err != nil || nextCursor == "0" {
				close(result)
				break
			}
		}
	}(api, v, result)
	return result
}

func (api TwitterApi) GetFriendsIds(v url.Values) (c Cursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friends/ids.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) GetFriendshipsLookup(v url.Values) (friendships []Friendship, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/lookup.json", v, &friendships, _GET, responseCh}
	return friendships, (<-responseCh).err
}

func (api TwitterApi) GetFriendshipsIncoming(v url.Values) (c Cursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/incoming.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) GetFriendshipsOutgoing(v url.Values) (c Cursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/outgoing.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) GetFollowersList(v url.Values) (c UserCursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/followers/list.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) GetFriendsList(v url.Values) (c UserCursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friends/list.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

// Like GetFriendsList, but returns a channel instead of a cursor and pre-fetches the remaining results
// This channel is closed once all values have been fetched
func (api TwitterApi) GetFriendsListAll(v url.Values) (result chan FriendsPage) {
	result = make(chan FriendsPage)

	v = cleanValues(v)
	go func(a TwitterApi, v url.Values, result chan FriendsPage) {
		// Cursor defaults to the first page ("-1")
		nextCursor := "-1"
		for {
			v.Set("cursor", nextCursor)
			c, err := a.GetFriendsList(v)

			// throttledQuery() handles all rate-limiting errors
			// if GetFriendsListAll() returns an error, it must be api different kind of error

			result <- FriendsPage{c.Users, err}

			nextCursor = c.Next_cursor_str
			if err != nil || nextCursor == "0" {
				close(result)
				break
			}
		}
	}(api, v, result)
	return result
}

// Like GetFollowersList, but returns a channel instead of a cursor and pre-fetches the remaining results
// This channel is closed once all values have been fetched
func (api TwitterApi) GetFollowersListAll(v url.Values) (result chan FollowersPage) {
	result = make(chan FollowersPage)

	v = cleanValues(v)
	go func(a TwitterApi, v url.Values, result chan FollowersPage) {
		// Cursor defaults to the first page ("-1")
		nextCursor := "-1"
		for {
			v.Set("cursor", nextCursor)
			c, err := a.GetFollowersList(v)

			// throttledQuery() handles all rate-limiting errors
			// if GetFollowersList() returns an error, it must be api different kind of error

			result <- FollowersPage{c.Users, err}

			nextCursor = c.Next_cursor_str
			if err != nil || nextCursor == "0" {
				close(result)
				break
			}
		}
	}(api, v, result)
	return result
}

func (api TwitterApi) GetFollowersUser(id int64, v url.Values) (c Cursor, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/followers/ids.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

// Like GetFriendsIds, but returns a channel instead of a cursor and pre-fetches the remaining results
// This channel is closed once all values have been fetched
func (api TwitterApi) GetFriendsIdsAll(v url.Values) (result chan FriendsIdsPage) {
	result = make(chan FriendsIdsPage)

	v = cleanValues(v)
	go func(a TwitterApi, v url.Values, result chan FriendsIdsPage) {
		// Cursor defaults to the first page ("-1")
		nextCursor := "-1"
		for {
			v.Set("cursor", nextCursor)
			c, err := a.GetFriendsIds(v)

			// throttledQuery() handles all rate-limiting errors
			// if GetFollowersList() returns an error, it must be api different kind of error

			result <- FriendsIdsPage{c.Ids, err}

			nextCursor = c.Next_cursor_str
			if err != nil || nextCursor == "0" {
				close(result)
				break
			}
		}
	}(api, v, result)
	return result
}

func (api TwitterApi) GetFriendsUser(id int64, v url.Values) (c Cursor, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friends/ids.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

// FollowUserId follows the user with the specified userId.
// This implements the /friendships/create endpoint, though the function name
// uses the terminology 'follow' as this is most consistent with colloquial Twitter terminology.
func (api TwitterApi) FollowUserId(userId int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(userId, 10))
	return api.postFriendshipsCreateImpl(v)
}

// FollowUserId follows the user with the specified screenname (username).
// This implements the /friendships/create endpoint, though the function name
// uses the terminology 'follow' as this is most consistent with colloquial Twitter terminology.
func (api TwitterApi) FollowUser(screenName string) (user User, err error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	return api.postFriendshipsCreateImpl(v)
}

func (api TwitterApi) postFriendshipsCreateImpl(v url.Values) (user User, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/create.json", v, &user, _POST, responseCh}
	return user, (<-responseCh).err
}

// UnfollowUserId unfollows the user with the specified userId.
// This implements the /friendships/destroy endpoint, though the function name
// uses the terminology 'unfollow' as this is most consistent with colloquial Twitter terminology.
func (api TwitterApi) UnfollowUserId(userId int64) (u User, err error) {
	v := url.Values{}
	v.Set("user_id", strconv.FormatInt(userId, 10))
	// Set other values before calling this method:
	// page, count, include_entities
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/destroy.json", v, &u, _POST, responseCh}
	return u, (<-responseCh).err
}

// UnfollowUser unfollows the user with the specified screenname (username)
// This implements the /friendships/destroy endpoint, though the function name
// uses the terminology 'unfollow' as this is most consistent with colloquial Twitter terminology.
func (api TwitterApi) UnfollowUser(screenname string) (u User, err error) {
	v := url.Values{}
	v.Set("screen_name", screenname)
	// Set other values before calling this method:
	// page, count, include_entities
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/friendships/destroy.json", v, &u, _POST, responseCh}
	return u, (<-responseCh).err
}
