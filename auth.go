package stackongo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type authError struct {
	Error map[string]string
}

var auth_url string = "https://stackexchange.com/oauth/access_token"

// AuthURL returns the URL to redirect the user for authentication 
// It accepts the following arguments
// client_id - Your App's registered ID
// redirect_uri - URI to redirect after authentication
// options - a map containing the following:
// scope - set of previlages you need to access - can be empty, "read_inbox", "no_expiry" or "read_inbox,no_expiry"
// state - optional string which will be returned as it is
func AuthURL(client_id, redirect_uri string, options map[string]string) (output string) {
	auth_url, _ := url.Parse("https://stackexchange.com/oauth")
	auth_query := auth_url.Query()
	auth_query.Add("client_id", client_id)
	auth_query.Add("redirect_uri", redirect_uri)

	for key, value := range options {
		auth_query.Add(key, value)
	}

	auth_url.RawQuery = auth_query.Encode()
	return auth_url.String()
}

func ObtainAccessToken(client_id, client_secret, code, redirect_uri string) (output map[string]string, error error) {
	client := &http.Client{Transport: getTransport()}

	parsed_auth_url, _ := url.Parse(auth_url)
	auth_query := parsed_auth_url.Query()
	auth_query.Add("client_id", client_id)
	auth_query.Add("client_secret", client_secret)
	auth_query.Add("code", code)
	auth_query.Add("redirect_uri", redirect_uri)

	// make the request
	response, error := client.PostForm(auth_url, auth_query)

	if error != nil {
		return
	}

	//check whether the response is a bad request
	if response.StatusCode != 200 {
		collection := new(authError)
		error = parseResponse(response, collection)

		error = errors.New(collection.Error["type"] + ": " + collection.Error["message"])
	} else {
		// if not process the output 
		bytes, err2 := ioutil.ReadAll(response.Body)

		if err2 != nil {
			return output, err2
		}

		collection, err3 := url.ParseQuery(string(bytes))
		output = map[string]string{"access_token": collection.Get("access_token"), "expires": collection.Get("expires")}
		error = err3
	}

	return
}
