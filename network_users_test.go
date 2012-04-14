package stackongo

import (
	"testing"
)

func TestAssociatedAccounts(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/associated", dummyNetworkUsersResponse, t)
	defer closeDummyServer(dummy_server)

	users, err := AssociatedAccounts([]int{1, 2, 3}, map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.Error())
	}

	if len(users.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if users.Items[0].Site_name != "Stack Overflow" {
		t.Error("Site name invalid.")
	}

	if users.Items[0].User_id != 1 {
		t.Error("user id invalid.")
	}

	if users.Items[0].Badge_counts.Gold != 25 {
		t.Error("badge count invalid.")
	}

	if users.Items[0].Last_access_date != 1328143328 {
		t.Error("last access date invalid.")
	}

}

//Test Data

var dummyNetworkUsersResponse string = `
{
  "items": [
    {
      "site_name": "Stack Overflow",
      "site_url": "http://stackoverflow.com",
      "user_id": 1,
      "reputation": 17614,
      "account_id": 1,
      "creation_date": 1217514151,
      "badge_counts": {
        "gold": 25,
        "silver": 84,
        "bronze": 97
      },
      "last_access_date": 1328143328,
      "answer_count": 148,
      "question_count": 15
    },
    {
      "site_name": "Server Fault",
      "site_url": "http://serverfault.com",
      "user_id": 1,
      "reputation": 4524,
      "account_id": 1,
      "creation_date": 1241075307,
      "badge_counts": {
        "gold": 4,
        "silver": 33,
        "bronze": 58
      },
      "last_access_date": 1328142890,
      "answer_count": 42,
      "question_count": 19
    },
    {
      "site_name": "Super User",
      "site_url": "http://superuser.com",
      "user_id": 1,
      "reputation": 8083,
      "account_id": 1,
      "creation_date": 1247439102,
      "badge_counts": {
        "gold": 9,
        "silver": 47,
        "bronze": 86
      },
      "last_access_date": 1328145149,
      "answer_count": 149,
      "question_count": 14
    }
  ]
}
`
