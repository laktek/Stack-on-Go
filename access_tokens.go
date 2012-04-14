package stackongo

import (
	"fmt"
	"strings"
)

// InspectAccessTokens returns the properties for a set of access tokens. 
func InspectAccessTokens(access_tokens []string, params map[string]string) (output *AccessTokens, error error) {
	request_path := fmt.Sprintf("access-tokens/%v", strings.Join(access_tokens, ";"))

	output = new(AccessTokens)
	error = get(request_path, params, output)
	return
}

// DeauthenticateAccessTokens de-authorizes the app for the users with given access tokens. 
func DeauthenticateAccessTokens(access_tokens []string, params map[string]string) (output *AccessTokens, error error) {
	request_path := fmt.Sprintf("apps/%v/de-authenticate", strings.Join(access_tokens, ";"))

	output = new(AccessTokens)
	error = get(request_path, params, output)
	return
}

// InvalidateAccessTokens invalidates the given access tokens. 
func InvalidateAccessTokens(access_tokens []string, params map[string]string) (output *AccessTokens, error error) {
	request_path := fmt.Sprintf("access-tokens/%v/invalidate", strings.Join(access_tokens, ";"))

	output = new(AccessTokens)
	error = get(request_path, params, output)
	return
}
