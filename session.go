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
func setupEndpoint(path string, params map[string]string) *url.URL {
	base_url, _ := url.Parse(host)
	endpoint, _ := base_url.Parse("/2.0/" + path)

	query := endpoint.Query()
	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint
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

// make the request
func (session Session) get(section string, params map[string]string) (*http.Response, os.Error) {
	client := new(http.Client)

	//set parameters for querystring
	params["site"] = session.Site

	return client.Get(setupEndpoint(section, params).String())
}
