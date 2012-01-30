package stackongo

import (
	"testing"
)

func TestRevisions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/revisions/1;2;3", dummyRevisionsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	revisions, err := session.Revisions([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

	if len(revisions) != 3 {
		t.Error("Number of items are wrong.")
	}

	if revisions[0].Revision_guid != "C8AF433A-8DFE-4906-9418-EB7D5B4522EA" {
		t.Error("guid is invalid.")
	}

	if revisions[0].Creation_date != 1220598880 {
		t.Error("Creation date is invalid.")
	}

	if revisions[0].User.Display_name != "JasonMichael" {
		t.Error("User is invalid.")
	}

}

func TestRevisionsForPosts(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/posts/1;2;3/revisions", dummyRevisionsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.RevisionsForPosts([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyRevisionsResponse string = `
{
  "items": [
    {
      "revision_guid": "C8AF433A-8DFE-4906-9418-EB7D5B4522EA",
      "revision_number": 0,
      "revision_type": "vote_based",
      "post_type": "answer",
      "post_id": 16321,
      "comment": "&lt;b&gt;Post Deleted&lt;/b&gt;  by &lt;a href=&quot;/users/1935/jasonmichael&quot;&gt;JasonMichael&lt;/a&gt;",
      "creation_date": 1220598880,
      "is_rollback": false,
      "set_community_wiki": false,
      "user": {
        "user_id": 1935,
        "display_name": "JasonMichael",
        "reputation": 645,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/16fdddee3335cb859fe7bb300cf02cdb?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/1935/jasonmichael"
      }
    },
    {
      "revision_guid": "F7658FFB-6F1C-412F-A6D9-18B17B560381",
      "revision_number": 0,
      "revision_type": "vote_based",
      "post_type": "answer",
      "post_id": 16445,
      "comment": "&lt;b&gt;Post Deleted&lt;/b&gt;  by &lt;a href=&quot;/users/1444/humpton&quot;&gt;Humpton&lt;/a&gt;",
      "creation_date": 1219590393,
      "is_rollback": false,
      "set_community_wiki": false,
      "user": {
        "user_id": 1444,
        "display_name": "Humpton",
        "reputation": 327,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/15c189676f875b5a245f3bd6e87744f1?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/1444/humpton"
      }
    },
    {
      "revision_guid": "CA1797BD-E4C4-4649-AC1D-EB2C5B69B4BF",
      "revision_number": 1,
      "revision_type": "single_user",
      "post_type": "answer",
      "post_id": 16545,
      "creation_date": 1219165179,
      "is_rollback": false,
      "set_community_wiki": false,
      "user": {
        "display_name": "mauro",
        "user_type": "does_not_exist"
      }
    }
  ]
}
`
