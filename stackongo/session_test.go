package stackongo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func closeDummyServer(dummy_server *httptest.Server) {
	transport = nil
	dummy_server.Close()
}

func createDummyServer(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	dummy_server := httptest.NewServer(http.HandlerFunc(handler))

	//change the host to use the test server
	SetTransport(&http.Transport{Proxy: func(*http.Request) (*url.URL, error) { return url.Parse(dummy_server.URL) }})

	//turn off SSL
	UseSSL = false

	return dummy_server
}

func returnDummyResponseForPath(path string, dummy_response string, t *testing.T) *httptest.Server {
	//serve dummy responses
	dummy_data := []byte(dummy_response)

	return createDummyServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Error("Path doesn't match")
		}
		w.Write(dummy_data)
	})
}

func returnDummyResponseForPathAndParams(path string, params map[string]string, dummy_response string, t *testing.T) *httptest.Server {
	//serve dummy responses
	dummy_data := []byte(dummy_response)

	return createDummyServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Error("Path doesn't match")
		}

		for key, value := range params {
			if r.URL.Query().Get(key) != value {
				t.Error("Expected " + key + " to equal " + value + ". Got " + r.URL.Query().Get(key))
			}
		}
		w.Write(dummy_data)
	})
}

func returnDummyErrorResponseForPath(path string, dummy_response string, t *testing.T) *httptest.Server {
	//serve dummy responses
	dummy_data := []byte(dummy_response)

	return createDummyServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Error("Path doesn't match")
		}
		w.WriteHeader(400)
		w.Write(dummy_data)
	})
}

func TestMetaInfo(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions", dummyMetaInfoResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	results, err := session.AllQuestions(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.Error())
	}

	if results.Has_more != true {
		t.Error("Has more field invalid.")
	}

	if results.Quota_remaining != 9989 {
		t.Error("Quota remaining invalid.")
	}

	if results.Quota_max != 10000 {
		t.Error("Quota max invalid.")
	}

}

func TestRequestError(t *testing.T) {
	dummy_server := returnDummyErrorResponseForPath("/2.0/questions", dummyErrorResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	result, err := session.AllQuestions(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil && result.Error_name != "no_method" && result.Error_message != "simulated" {
		t.Error("Y U GAVE NO ERROR?")
	}

}

//test data

var dummyMetaInfoResponse string = `
{
  "items": [],
  "quota_remaining": 9989,
  "quota_max": 10000,
  "has_more": true
}
`
