package anaconda

import (
	"net/url"
	"strconv"
)

type Location struct {
	Name  string `json:"name"`
	Woeid int    `json:"woeid"`
}

type Trend struct {
	Name            string `json:"name"`
	Query           string `json:"query"`
	Url             string `json:"url"`
	PromotedContent string `json:"promoted_content"`
}

type TrendResponse struct {
	Trends    []Trend    `json:"trends"`
	AsOf      string     `json:"as_of"`
	CreatedAt string     `json:"created_at"`
	Locations []Location `json:"locations"`
}

type TrendLocation struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	ParentId    int    `json:"parentid"`
	PlaceType   struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"placeType"`
	Url   string `json:"url"`
	Woeid int32  `json:"woeid"`
}

// https://developer.twitter.com/en/docs/trends/trends-for-location/api-reference/get-trends-place
func (api TwitterApi) GetTrendsByPlace(id int64, v url.Values) (trendResp TrendResponse, err error) {
	responseCh := make(chan response)
	v = cleanValues(v)
	v.Set("id", strconv.FormatInt(id, 10))
	api.queryQueue <- query{api.baseUrl + "/trends/place.json", v, &[]interface{}{&trendResp}, _GET, responseCh}
	return trendResp, (<-responseCh).err
}

// https://developer.twitter.com/en/docs/trends/locations-with-trending-topics/api-reference/get-trends-available
func (api TwitterApi) GetTrendsAvailableLocations(v url.Values) (locations []TrendLocation, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/trends/available.json", v, &locations, _GET, responseCh}
	return locations, (<-responseCh).err
}

// https://developer.twitter.com/en/docs/trends/locations-with-trending-topics/api-reference/get-trends-closest
func (api TwitterApi) GetTrendsClosestLocations(lat float64, long float64, v url.Values) (locations []TrendLocation, err error) {
	responseCh := make(chan response)
	v = cleanValues(v)
	v.Set("lat", strconv.FormatFloat(lat, 'f', 6, 64))
	v.Set("long", strconv.FormatFloat(long, 'f', 6, 64))
	api.queryQueue <- query{api.baseUrl + "/trends/closest.json", v, &locations, _GET, responseCh}
	return locations, (<-responseCh).err
}
