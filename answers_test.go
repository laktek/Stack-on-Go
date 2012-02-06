package stackongo

import (
	"testing"
)

func TestAllAnswers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/answers", dummyAnswersResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	answers, err := session.AllAnswers(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(answers.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if answers.Items[0].Answer_id != 9051117 {
		t.Error("ID invalid.")
	}

	if answers.Items[0].Owner.Display_name != "Dalar" {
		t.Error("Owner invalid.")
	}

	if answers.Items[0].Creation_date != 1327814223 {
		t.Error("Date invalid.")
	}

}

func TestGetAnswers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/answers/1;2;3", dummyAnswersResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.GetAnswers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestAnswersForQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/1;2;3/answers", dummyAnswersResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.AnswersForQuestions([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}
}

func TestAnswersFromUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/answers", dummyAnswersResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.AnswersFromUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}
}

func TestTopAnswersFromUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/tags/hello;world/top-answers", dummyAnswersResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.TopAnswersFromUsers([]int{1, 2, 3}, []string{"hello", "world"}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}
}

//Test Data

var dummyAnswersResponse string = `
{
  "items": [
    {
      "question_id": 9023805,
      "answer_id": 9051117,
      "creation_date": 1327814223,
      "last_activity_date": 1327814223,
      "score": 0,
      "is_accepted": false,
      "owner": {
        "user_id": 325674,
        "display_name": "Dalar",
        "reputation": 176,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/d0c5cda979ec453235876282edbe1188?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/325674/dalar"
      }
    },
    {
      "question_id": 9051070,
      "answer_id": 9051116,
      "creation_date": 1327814175,
      "last_activity_date": 1327814175,
      "score": 0,
      "is_accepted": false,
      "owner": {
        "user_id": 197913,
        "display_name": "Damir Arh",
        "reputation": 1019,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/c85e6426e70c4ecf205a6187e7dbf174?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/197913/damir-arh"
      }
    },
    {
      "question_id": 8415119,
      "answer_id": 9051112,
      "creation_date": 1327814056,
      "last_activity_date": 1327814056,
      "score": 0,
      "is_accepted": false,
      "owner": {
        "user_id": 1094837,
        "display_name": "Joseph Torraca",
        "reputation": 78,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/292351179bec9b5119d0e3a7e743537a?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/1094837/joseph-torraca"
      }
    }
  ]
}
`
