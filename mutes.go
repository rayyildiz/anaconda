package anaconda

import (
	"net/url"
	"strconv"
)

func (api TwitterApi) GetMutedUsersList(v url.Values) (c UserCursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/mutes/users/list.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) GetMutedUsersIds(v url.Values) (c Cursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/mutes/users/ids.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) MuteUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return api.Mute(v)
}

func (api TwitterApi) MuteUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return api.Mute(v)
}

func (api TwitterApi) Mute(v url.Values) (user User, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/mutes/users/create.json", v, &user, _POST, responseCh}
	return user, (<-responseCh).err
}

func (api TwitterApi) UnmuteUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return api.Unmute(v)
}

func (api TwitterApi) UnmuteUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return api.Unmute(v)
}

func (api TwitterApi) Unmute(v url.Values) (user User, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/mutes/users/destroy.json", v, &user, _POST, responseCh}
	return user, (<-responseCh).err
}
