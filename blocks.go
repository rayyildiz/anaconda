package anaconda

import (
	"net/url"
	"strconv"
)

func (api TwitterApi) GetBlocksList(v url.Values) (c UserCursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/blocks/list.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) GetBlocksIds(v url.Values) (c Cursor, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/blocks/ids.json", v, &c, _GET, responseCh}
	return c, (<-responseCh).err
}

func (api TwitterApi) BlockUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return api.Block(v)
}

func (api TwitterApi) BlockUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return api.Block(v)
}

func (api TwitterApi) Block(v url.Values) (user User, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/blocks/create.json", v, &user, _POST, responseCh}
	return user, (<-responseCh).err
}

func (api TwitterApi) UnblockUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return api.Unblock(v)
}

func (api TwitterApi) UnblockUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return api.Unblock(v)
}

func (api TwitterApi) Unblock(v url.Values) (user User, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/blocks/destroy.json", v, &user, _POST, responseCh}
	return user, (<-responseCh).err
}
