package stackongo

import (
	"testing"
)

func TestAllUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users", dummyUsersResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	users, err := session.AllUsers(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(users) != 1 {
		t.Error("Number of items wrong.")
	}

	if users[0].User_id != 22656 {
		t.Error("ID invalid.")
	}

	if users[0].User_type != "registered" {
		t.Error("Post type invalid.")
	}

	if users[0].Creation_date != 1222430705 {
		t.Error("Date invalid.")
	}

	if users[0].Is_employee != false {
		t.Error("Boolean doesn't match.")
	}

	if users[0].Badge_counts.Gold != 105 {
		t.Error("Badge count is invalid.")
	}

}

func TestUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3", dummyUsersResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.Users([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestModerators(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/moderators", dummyUsersResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.Moderators(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestElectedModerators(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/moderators/elected", dummyUsersResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.ElectedModerators(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyUsersResponse string = `
{
  "items": [
    {
      "user_id": 22656,
      "user_type": "registered",
      "creation_date": 1222430705,
      "display_name": "Jon Skeet",
      "profile_image": "http://www.gravatar.com/avatar/6d8ebb117e8d83d74ea95fbdd0f87e13?d=identicon&r=PG",
      "reputation": 397366,
      "reputation_change_day": 30,
      "reputation_change_week": 1135,
      "reputation_change_month": 30,
      "reputation_change_quarter": 10890,
      "reputation_change_year": 10890,
      "age": 35,
      "last_access_date": 1328051866,
      "last_modified_date": 1328017043,
      "is_employee": false,
      "link": "http://stackoverflow.com/users/22656/jon-skeet",
      "website_url": "http://csharpindepth.com",
      "location": "Reading, United Kingdom",
      "account_id": 11683,
      "badge_counts": {
        "gold": 105,
        "silver": 1672,
        "bronze": 2946
      },
      "accept_rate": 95
    }
  ]
}
`
