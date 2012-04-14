package stackongo

import (
	"testing"
)

func TestEvents(t *testing.T) {
	dummy_server := returnDummyResponseForPathAndParams("/2.0/events", map[string]string{"key": "app123", "access_token": "abc"}, dummyEventsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	events, err := session.Events(map[string]string{"page": "1"}, map[string]string{"key": "app123", "access_token": "abc"})

	if err != nil {
		t.Error(err.Error())
	}

	if len(events.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if events.Items[0].Event_type != "comment_posted" {
		t.Error("Event type invalid.")
	}

	if events.Items[0].Event_id != 11462515 {
		t.Error("Event id invalid.")
	}

	if events.Items[0].Creation_date != 1328226264 {
		t.Error("Date invalid.")
	}

}

//Test Data

var dummyEventsResponse string = `
{
  "items": [
    {
      "event_type": "comment_posted",
      "event_id": 11462515,
      "creation_date": 1328226264
    },
    {
      "event_type": "answer_posted",
      "event_id": 9121759,
      "creation_date": 1328226257
    },
    {
      "event_type": "question_posted",
      "event_id": 9121758,
      "creation_date": 1328226255
    }
  ]
}
`
