package stackongo

import (
	"testing"
)

func TestTopAnswerers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/tags/test/top-answerers/all_time", dummyTagScoresResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	tag_scores, err := session.TopAnswerers("test", "all_time", map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(tag_scores.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if tag_scores.Items[0].Score != 45 {
		t.Error("Score invalid.")
	}

	if tag_scores.Items[0].Post_count != 1 {
		t.Error("Post count invalid.")
	}

	if tag_scores.Items[0].User.Display_name != "user208987" {
		t.Error("User invalid.")
	}

}

func TestTopAskers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/tags/test/top-askers/all_time", dummyTagScoresResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.TopAskers("test", "all_time", map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyTagScoresResponse string = `
{
  "items": [
    {
      "user": {
        "user_id": 208987,
        "display_name": "user208987",
        "reputation": 234,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/449b9a01c4cba54b275435cfabc54e30?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/208987/user208987"
      },
      "score": 45,
      "post_count": 1
    },
    {
      "user": {
        "user_id": 7743,
        "display_name": "Jeroen Dirks",
        "reputation": 1661,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/2d3c00ac7272c33d84be8e310b95e93c?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/7743/jeroen-dirks",
        "accept_rate": 97
      },
      "score": 41,
      "post_count": 6
    },
    {
      "user": {
        "user_id": 211210,
        "display_name": "Malte Schledjewski",
        "reputation": 330,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/9c9a69a4563e030a3a0cbed8c9ca19ee?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/211210/malte-schledjewski"
      },
      "score": 41,
      "post_count": 1
    }
  ]
}
`
