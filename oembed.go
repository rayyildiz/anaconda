package anaconda

import (
	"net/http"
	"net/url"
	"strconv"
)

type OEmbed struct {
	Type          string
	Width         int
	Cache_age     string
	Height        int
	Author_url    string
	Html          string
	Version       string
	Provider_name string
	Provider_url  string
	Url           string
	Author_name   string
}

// No authorization on this endpoint. Its the only one.
func (api TwitterApi) GetOEmbed(v url.Values) (o OEmbed, err error) {
	resp, err := http.Get(api.baseUrlV1() + "/statuses/oembed.json?" + v.Encode())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = decodeResponse(resp, &o)
	return
}

// Calls GetOEmbed with the corresponding id. Convenience wrapper for GetOEmbed()
func (api TwitterApi) GetOEmbedId(id int64, v url.Values) (o OEmbed, err error) {
	v = cleanValues(v)
	v.Set("id", strconv.FormatInt(id, 10))
	resp, err := http.Get(api.baseUrlV1() + "/statuses/oembed.json?" + v.Encode())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = decodeResponse(resp, &o)
	return
}

func (api TwitterApi) baseUrlV1() string {
	if api.baseUrl == BaseUrl {
		return BaseUrlV1
	}

	if api.baseUrl == "" {
		return BaseUrlV1
	}

	return api.baseUrl
}
