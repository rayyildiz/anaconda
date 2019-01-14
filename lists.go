package anaconda

import (
	"net/url"
	"strconv"
	"strings"
)

// CreateList implements /lists/create.json
func (api TwitterApi) CreateList(name, description string, v url.Values) (list List, err error) {
	v = cleanValues(v)
	v.Set("name", name)
	v.Set("description", description)

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/create.json", v, &list, _POST, responseCh}
	return list, (<-responseCh).err
}

// AddUserToList implements /lists/members/create.json
func (api TwitterApi) AddUserToList(screenName string, listID int64, v url.Values) (users []User, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("screen_name", screenName)

	var addUserToListResponse AddUserToListResponse

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/members/create.json", v, &addUserToListResponse, _POST, responseCh}
	return addUserToListResponse.Users, (<-responseCh).err
}

// AddMultipleUsersToList implements /lists/members/create_all.json
func (api TwitterApi) AddMultipleUsersToList(screenNames []string, listID int64, v url.Values) (list List, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("screen_name", strings.Join(screenNames, ","))

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/members/create_all.json", v, &list, _POST, responseCh}
	r := <-responseCh
	return list, r.err
}

// RemoveUserFromList implements /lists/members/destroy.json
func (a TwitterApi) RemoveUserFromList(screenName string, listID int64, v url.Values) (list List, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("screen_name", screenName)

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/members/destroy.json", v, &list, _POST, response_ch}
	r := <-response_ch
	return list, r.err
}

// RemoveMultipleUsersFromList implements /lists/members/destroy_all.json
func (a TwitterApi) RemoveMultipleUsersFromList(screenNames []string, listID int64, v url.Values) (list List, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("screen_name", strings.Join(screenNames, ","))

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/members/destroy_all.json", v, &list, _POST, response_ch}
	r := <-response_ch
	return list, r.err
}

// GetListsOwnedBy implements /lists/ownerships.json
// screen_name, count, and cursor are all optional values
func (api TwitterApi) GetListsOwnedBy(userID int64, v url.Values) (lists []List, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(userID, 10))

	var listResponse ListResponse

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/ownerships.json", v, &listResponse, _GET, responseCh}
	return listResponse.Lists, (<-responseCh).err
}

func (api TwitterApi) GetListTweets(listID int64, includeRTs bool, v url.Values) (tweets []Tweet, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("include_rts", strconv.FormatBool(includeRTs))

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/statuses.json", v, &tweets, _GET, responseCh}
	return tweets, (<-responseCh).err
}

// GetList implements /lists/show.json
func (api TwitterApi) GetList(listID int64, v url.Values) (list List, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/show.json", v, &list, _GET, responseCh}
	return list, (<-responseCh).err
}

func (api TwitterApi) GetListTweetsBySlug(slug string, ownerScreenName string, includeRTs bool, v url.Values) (tweets []Tweet, err error) {
	v = cleanValues(v)
	v.Set("slug", slug)
	v.Set("owner_screen_name", ownerScreenName)
	v.Set("include_rts", strconv.FormatBool(includeRTs))

	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/lists/statuses.json", v, &tweets, _GET, responseCh}
	return tweets, (<-responseCh).err
}
