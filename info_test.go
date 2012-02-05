package stackongo

import (
	"testing"
)

func TestInfo(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/info", dummyInfoResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	info, err := session.Info()

	if err != nil {
		t.Error(err.String())
	}

	if info.Total_badges != 2575799 {
		t.Error("Total badges invalid.")
	}

	if info.Badges_per_minute != 1.41 {
		t.Error("Badges per minute invalid.")
	}

	if info.Api_revision != "2012.1.27.662" {
		t.Error("API revision invalid.")
	}

}

func TestNoInfo(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/info", dummyMetaInfoResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	_, err := session.Info()

	if err.String() != "Site not found" {
		t.Error("Error didn't match")
	}
}

//Test Data

var dummyInfoResponse string = `
{
  "items": [
    {
      "total_questions": 2578583,
      "total_unanswered": 2061766,
      "total_accepted": 1615008,
      "total_answers": 5429736,
      "questions_per_minute": 1.4,
      "answers_per_minute": 2.95,
      "total_comments": 10606280,
      "total_votes": 16305351,
      "total_badges": 2575799,
      "badges_per_minute": 1.41,
      "total_users": 969429,
      "new_active_users": 49,
      "api_revision": "2012.1.27.662"
    }
  ],
  "quota_remaining": 9974,
  "quota_max": 10000,
  "backoff": 10,
  "has_more": false
}
`
