package stackongo

import (
	"os"
	"http"
	"strings"
	"fmt"
)

func getAccessTokens(path string, params map[string]string) (output []AccessToken, error os.Error) {

	client := new(http.Client)

	// make the reques
	response, err := client.Get(setupEndpoint(path, params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(accessTokensCollection))
	collection := parsed_response.(*accessTokensCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error
}

// InspectAccessTokens returns the properties for a set of access tokens. 
func InspectAccessTokens(access_tokens []string, params map[string]string) (output []AccessToken, error os.Error) {
	request_path := fmt.Sprintf("access-tokens/%v", strings.Join(access_tokens, ";"))
	return getAccessTokens(request_path, params)
}

// DeauthenticateAccessTokens de-authorizes the app for the users with given access tokens. 
func DeauthenticateAccessTokens(access_tokens []string, params map[string]string) (output []AccessToken, error os.Error) {
	request_path := fmt.Sprintf("apps/%v/de-authenticate", strings.Join(access_tokens, ";"))
	return getAccessTokens(request_path, params)
}

// InvalidateAccessTokens invalidates the given access tokens. 
func InvalidateAccessTokens(access_tokens []string, params map[string]string) (output []AccessToken, error os.Error) {
	request_path := fmt.Sprintf("access-tokens/%v/invalidate", strings.Join(access_tokens, ";"))
	return getAccessTokens(request_path, params)
}
