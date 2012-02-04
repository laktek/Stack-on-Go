package stackongo

import (
	"testing"
	"url"
)

func TestAuthURL(t *testing.T) {
	result_uri, _ := url.Parse(AuthURL("abc", "www.my_app.com", map[string]string{"state": "test", "scope": "read_inbox,no_expiry"}))

	if result_uri.Query().Encode() != "state=test&scope=read_inbox%2Cno_expiry&redirect_uri=www.my_app.com&client_id=abc" {
		print(result_uri.Query().Encode())
		t.Error("URL doesn't match")
	}
}

func TestObtainAccessTokenValid(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/oauth/access_token", dummyAccessTokenResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	auth_url = dummy_server.URL + "/oauth/access_token"

	token_output, _ := ObtainAccessToken("abc", "secret", "def", "www.my_app.com")

	if token_output["access_token"] != "my_token" {
		t.Error("invalid access token.")
	}

	if token_output["expires"] != "1234" {
		t.Error("invalid expiry.")
	}

}

func TestObtainAccessTokenInvalid(t *testing.T) {
	dummy_server := returnDummyErrorResponseForPath("/oauth/access_token", dummyAccessTokenErrorResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	auth_url = dummy_server.URL + "/oauth/access_token"

	_, error := ObtainAccessToken("abc", "secret", "def", "www.my_app.com")

	if error.String() != "invalid_request: some reason" {
		t.Error("invalid error message.")
	}

}

//test data
var dummyAccessTokenResponse string = "access_token=my_token&expires=1234"

var dummyAccessTokenErrorResponse string = `
{ "error": { "type": "invalid_request", "message": "some reason" } }
`
