package stackongo

import (
	"testing"
	"url"
	"reflect"
)

func TestAuthURL(t *testing.T) {
	given_uri, _ := url.Parse("https://stackexchange.com/oauth?client_id=abc,redirect_url=www.my_app.com&state=test&scope=read_inbox%2Cno_expiry")
	result_uri, _ := url.Parse(AuthURL("abc", "www.my_app.com", map[string]string{"state": "test", "scope": "read_inbox,no_expiry"}))

	if reflect.DeepEqual(result_uri, given_uri) {
		t.Error("URL doesn't match")
	}
}

func TestGenerateAccessTokenValid(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/oauth/access_token", dummyAccessTokenResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	auth_url = dummy_server.URL + "/oauth/access_token"

	token_output, _ := GenerateAccessToken("abc", "secret", "def", "www.my_app.com")

	if token_output["access_token"] != "my_token" {
		t.Error("invalid access token.")
	}

	if token_output["expires"] != "1234" {
		t.Error("invalid expiry.")
	}

}

func TestGenerateAccessTokenInvalid(t *testing.T) {
	dummy_server := returnDummyErrorResponseForPath("/oauth/access_token", dummyAccessTokenErrorResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	auth_url = dummy_server.URL + "/oauth/access_token"

	_, error := GenerateAccessToken("abc", "secret", "def", "www.my_app.com")

	if error.String() != "invalid_request: some reason" {
		t.Error("invalid error message.")
	}

}

//test data
var dummyAccessTokenResponse string = "access_token=my_token&expires=1234"

var dummyAccessTokenErrorResponse string = `
{ "error": { "type": "invalid_request", "message": "some reason" } }
`
