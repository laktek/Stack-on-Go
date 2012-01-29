package stackongo

import (
	"http"
	"json"
	"io/ioutil"
	"url"
	"os"
)

var host string = "http://api.stackexchange.com"

type Session struct {
	Site string
}

func NewSession(site string) *Session {
	return &Session{Site: site}
}

func setHost(url string) {
	host = url
}

// construct the endpoint URL
func (session Session) setupEndpoint(path string, params map[string]string) *url.URL {
	base_url, _ := url.Parse(host)
	endpoint, _ := base_url.Parse("/2.0/" + path)

	//set parameters for querystring
	query := endpoint.Query()
	query.Set("site", session.Site)

	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint
}

// make the request
func (session Session) get(section string, params map[string]string) (*http.Response, os.Error) {
	client := new(http.Client)
	return client.Get(session.setupEndpoint(section, params).String())
}

// parse the response
func parseResponse(response *http.Response, result interface{}) (interface{}, os.Error) {

	//read the response
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return result, err
	}

	//parse JSON
	err2 := json.Unmarshal(bytes, result)

	if err2 != nil {
		print(err2.String())
	}

	//check whether the response is a bad request
	if response.StatusCode == 400 {
		err2 = os.NewError("Bad Request")
	}

	return result, err2
}

/*
func (site Site) GetInfo() (output Info, error os.Error) {
	// make the request
	response, err := site.get("info")

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(infoCollection))
	collection := parsed_response.(*infoCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items[0]
	}

	return output, error
}

func (site Site) GetPrivileges() (output []Privilege, error os.Error) {
	// make the request
	response, err := site.get("privileges")

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(privilegesCollection))
	collection := parsed_response.(*privilegesCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error
}
*/
