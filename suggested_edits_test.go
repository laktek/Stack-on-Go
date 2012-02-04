package stackongo

import (
	"testing"
  "strings"
)

func TestAllSuggestedEdits(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/suggested-edits", dummySuggestedEditsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	suggested_edits, err := session.AllSuggestedEdits(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(suggested_edits.Items) != 2 {
		t.Error("Number of items wrong.")
	}

	if suggested_edits.Items[0].Suggested_edit_id != 190741 {
		t.Error("ID invalid.")
	}

	if suggested_edits.Items[0].Proposing_user.Display_name != "Natali" {
		t.Error("Owner invalid.")
	}

	if suggested_edits.Items[0].Creation_date != 1327922327 {
		t.Error("Date invalid.")
	}

	if strings.Join(suggested_edits.Items[1].Tags, ",") != "c#,jquery,asp.net,html" {
		t.Error("Tags invalid.")
	}

}

func TestSuggestedEdits(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/suggested-edits/1;2;3", dummySuggestedEditsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.SuggestedEdits([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestSuggestedEditsForPosts(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/posts/1;2;3/suggested-edits", dummySuggestedEditsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.SuggestedEditsForPosts([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}
}

func TestSuggestedEditsFromUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/suggested-edits", dummySuggestedEditsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.SuggestedEditsFromUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummySuggestedEditsResponse string = `
{
  "items": [
    {
      "suggested_edit_id": 190741,
      "post_id": 9062968,
      "post_type": "question",
      "comment": "improveing formating",
      "creation_date": 1327922327,
      "proposing_user": {
        "user_id": 601868,
        "display_name": "Natali",
        "reputation": 405,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/b15f339505879a64a8ca7b7eaa6f8585?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/601868/natali"
      }
    },
    {
      "suggested_edit_id": 190740,
      "post_id": 9052990,
      "post_type": "question",
      "title": "How to get values in hidden fields without clicking the submit button",
      "tags": [
        "c#",
        "jquery",
        "asp.net",
        "html"
      ],
      "comment": "changed tags, spelling correction, changed title to be more relevent and improved formating",
      "creation_date": 1327922303,
      "proposing_user": {
        "user_id": 27922,
        "display_name": "TheAlbear",
        "reputation": 1265,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/e46f50e7b35efdcf2945f6e40b10df7f?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/27922/thealbear"
      }
    }
  ]
}
`
