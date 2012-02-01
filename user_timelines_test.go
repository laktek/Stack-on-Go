package stackongo

import (
	"testing"
)

func TestTimelineForUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/timeline", dummyUserTimelinesResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	user_timelines, err := session.TimelineForUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(user_timelines) != 3 {
		t.Error("Number of items wrong.")
	}

	if user_timelines[0].User_id != 22656 {
		t.Error("ID invalid.")
	}

	if user_timelines[0].Post_type != "answer" {
		t.Error("Post type invalid.")
	}

	if user_timelines[0].Creation_date != 1328047513 {
		t.Error("Date invalid.")
	}

	if user_timelines[0].Title != "Thread concurrency issue even within one single command?" {
		t.Error("Title invalid.")
	}

}

//Test Data

var dummyUserTimelinesResponse string = `
{
  "items": [
    {
      "creation_date": 1328047513,
      "post_type": "answer",
      "timeline_type": "answered",
      "user_id": 22656,
      "post_id": 9087675,
      "title": "Thread concurrency issue even within one single command?"
    },
    {
      "creation_date": 1328044996,
      "post_type": "question",
      "timeline_type": "commented",
      "user_id": 22656,
      "post_id": 9087138,
      "comment_id": 11410707,
      "title": "Display varchar Date as D/M/Y?",
      "detail": "Why is your date field in the database as a varchar in the first place? If you can possibly fix the schema, that would be the best approach."
    },
    {
      "creation_date": 1328044954,
      "post_type": "answer",
      "timeline_type": "commented",
      "user_id": 22656,
      "post_id": 9086766,
      "comment_id": 11410691,
      "title": "Joining two tables and returning multiple records as one row using LINQ",
      "detail": "@Zajn: You might want to use officeNames = offices.Select(o => o.OfficeName).ToList() or something like that - it's hard to say without knowing more about what you're doing. Iterating over the office names should work anyway though."
    }
  ]
}
`
