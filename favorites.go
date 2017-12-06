package anaconda

import (
	"net/url"
)

func (api TwitterApi) GetFavorites(v url.Values) (favorites []Tweet, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/favorites/list.json", v, &favorites, _GET, responseCh}
	return favorites, (<-responseCh).err
}
