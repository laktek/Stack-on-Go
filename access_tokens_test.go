package stackongo

import (
	"testing"
)

func TestInspectAccessTokens(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/access-tokens/abc;def;ghi", dummyAccessTokensResponse, t)
	defer closeDummyServer(dummy_server)

	access_tokens, err := InspectAccessTokens([]string{"abc", "def", "ghi"}, map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(access_tokens.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if access_tokens.Items[0].Access_token != "abc" {
		t.Error("Access token invalid.")
	}

	if access_tokens.Items[0].Expires_on_date != 1328148124 {
		t.Error("Expires date invalid.")
	}

	if access_tokens.Items[0].Account_id != 1001 {
		t.Error("Account id invalid.")
	}

	if access_tokens.Items[0].Scope[0] != "read_inbox" {
		t.Error("Scope invalid.")
	}

}

func TestDeauthenticateAccessTokens(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/apps/abc;def;ghi/de-authenticate", dummyAccessTokensResponse, t)
	defer closeDummyServer(dummy_server)

	_, err := DeauthenticateAccessTokens([]string{"abc", "def", "ghi"}, map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TesInvalidateAccessTokens(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/access-tokens/abc;def;ghi/invalidate", dummyAccessTokensResponse, t)
	defer closeDummyServer(dummy_server)

	_, err := InvalidateAccessTokens([]string{"abc", "def", "ghi"}, map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyAccessTokensResponse string = `
{
  "items": [
    {
      "access_token": "abc",
      "expires_on_date": 1328148124,
      "account_id": 1001,
      "scope": ["read_inbox", "no_expiry"]
    },
    {
      "access_token": "def",
      "expires_on_date": 1328148124,
      "account_id": 1002,
      "scope": ["read_inbox"]
    },
    {
      "access_token": "ghi",
      "expires_on_date": 1328148124,
      "account_id": 1003,
      "scope": ["read_inbox"]
    }
  ]
}
`
