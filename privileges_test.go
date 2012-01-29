package stackongo

import (
	"testing"
)

func TestAllPrivileges(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/privileges", dummyPrivilegesResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	privileges, err := session.AllPrivileges(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

	if len(privileges) != 3 {
		t.Error("Number of items wrong.")
	}

	if privileges[0].Short_description != "create posts" {
		t.Error("Short description is invalid.")
	}

	if privileges[0].Reputation != 1 {
		t.Error("Reputation is invalid.")
	}

}

func TestPrivilegesForUser(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1/privileges", dummyPrivilegesResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.PrivilegesForUser(1, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyPrivilegesResponse string = `
{
  "items": [
    {
      "short_description": "create posts",
      "description": "Ask and answer questions",
      "reputation": 1
    },
    {
      "short_description": "participate in meta",
      "description": "Participate in per-site meta",
      "reputation": 5
    },
    {
      "short_description": "remove new user restrictions",
      "description": "Remove new user restrictions",
      "reputation": 10
    }
  ]
}
`
