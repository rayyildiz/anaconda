package anaconda

import (
	"net/url"
)

// Verify the credentials by making a very small request
func (api TwitterApi) VerifyCredentials() (ok bool, err error) {
	v := cleanValues(nil)
	v.Set("include_entities", "false")
	v.Set("skip_status", "true")

	_, err = api.GetSelf(v)
	return err == nil, err
}

// Get the user object for the authenticated user. Requests /account/verify_credentials
func (api TwitterApi) GetSelf(v url.Values) (u User, err error) {
	responseCh := make(chan response)
	api.queryQueue <- query{api.baseUrl + "/account/verify_credentials.json", v, &u, _GET, responseCh}
	return u, (<-responseCh).err
}
