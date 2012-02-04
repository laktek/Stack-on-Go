package stackongo

import (
	"testing"
)

func TestReputationChangesForUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/reputation", dummyReputationsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	reputations, err := session.ReputationChangesForUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(reputations.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if reputations.Items[0].User_id != 1 {
		t.Error("User ID is invalid.")
	}

	if reputations.Items[0].Post_type != "answer" {
		t.Error("Post type is invalid.")
	}

	if reputations.Items[0].On_date != 1326828986 {
		t.Error("On date is invalid.")
	}

}

//Test Data

var dummyReputationsResponse string = `
{
  "items": [
    {
      "user_id": 1,
      "post_id": 225592,
      "post_type": "answer",
      "vote_type": "up_votes",
      "reputation_change": 160,
      "on_date": 1326828986
    },
    {
      "user_id": 2,
      "post_id": 1722923,
      "post_type": "answer",
      "vote_type": "up_votes",
      "reputation_change": 10,
      "on_date": 1326366463
    },
    {
      "user_id": 3,
      "post_id": 7038708,
      "post_type": "question",
      "vote_type": "up_votes",
      "reputation_change": 10,
      "on_date": 1325566470
    }
  ]
}
`
