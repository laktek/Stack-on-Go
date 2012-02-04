package stackongo

import (
	"testing"
)

func TestTimelineForQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/1;2;3/timeline", dummyQuestionTimelinesResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	question_timelines, err := session.TimelineForQuestions([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(question_timelines.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if question_timelines.Items[0].Timeline_type != "comment" {
		t.Error("Timeline type invalid.")
	}

	if question_timelines.Items[0].Creation_date != 1328052381 {
		t.Error("Date invalid.")
	}

	if question_timelines.Items[0].User.Display_name != "eouw0o83hf" {
		t.Error("User invalid.")
	}

}

//Test Data

var dummyQuestionTimelinesResponse string = `
{
  "items": [
    {
      "timeline_type": "comment",
      "question_id": 9087138,
      "post_id": 9087152,
      "comment_id": 11413013,
      "creation_date": 1328052381,
      "user": {
        "user_id": 570190,
        "display_name": "eouw0o83hf",
        "reputation": 242,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/e0ff68bf6b06206afd062af7ed0d640a?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/570190/eouw0o83hf",
        "accept_rate": 75
      }
    },
    {
      "timeline_type": "comment",
      "question_id": 9087138,
      "post_id": 9087152,
      "comment_id": 11412816,
      "creation_date": 1328051616,
      "user": {
        "user_id": 264918,
        "display_name": "EB.",
        "reputation": 60,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/21ddc54277d0234a30a70b3460270dd8?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/264918/eb",
        "accept_rate": 83
      }
    },
    {
      "timeline_type": "comment",
      "question_id": 9087138,
      "post_id": 9087152,
      "comment_id": 11412653,
      "creation_date": 1328050976,
      "user": {
        "user_id": 570190,
        "display_name": "eouw0o83hf",
        "reputation": 242,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/e0ff68bf6b06206afd062af7ed0d640a?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/570190/eouw0o83hf",
        "accept_rate": 75
      }
    }
  ]
}
`
