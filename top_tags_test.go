package stackongo

import (
	"testing"
)

func TestTopTagsByAnswerForUser(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1/top-answer-tags", dummyTopTagsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	tags, err := session.TopTagsByAnswerForUser(1, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(tags.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if tags.Items[0].Tag_name != "sql-server" {
		t.Error("Name invalid.")
	}

	if tags.Items[0].Answer_score != 89 {
		t.Error("Answer score invalid.")
	}

	if tags.Items[0].Answer_count != 8 {
		t.Error("Answer count invalid.")
	}

}

func TestTopTagsByQuestionForUser(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1/top-question-tags", dummyTopTagsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	tags, err := session.TopTagsByQuestionForUser(1, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(tags.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if tags.Items[0].Tag_name != "sql-server" {
		t.Error("Name invalid.")
	}

	if tags.Items[0].Question_score != 369 {
		t.Error("Question score invalid.")
	}

	if tags.Items[0].Question_count != 4 {
		t.Error("Question count invalid.")
	}

}

//Test Data

var dummyTopTagsResponse string = `
{
  "items": [
    {
      "tag_name": "sql-server",
      "question_score": 369,
      "question_count": 4,
      "answer_score": 89,
      "answer_count": 8
    },
    {
      "tag_name": "sql",
      "question_score": 292,
      "question_count": 2,
      "answer_score": 60,
      "answer_count": 4
    },
    {
      "tag_name": "parameters",
      "question_score": 196,
      "question_count": 1,
      "answer_score": 0,
      "answer_count": 0
    }
  ]
}
`
