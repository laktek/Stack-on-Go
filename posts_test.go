package stackongo

import (
	"testing"
)

func TestAllPosts(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/posts", dummyPostsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	posts, err := session.AllPosts(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(posts.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if posts.Items[0].Post_id != 9051104 {
		t.Error("ID invalid.")
	}

	if posts.Items[0].Post_type != "question" {
		t.Error("Post type invalid.")
	}

	if posts.Items[0].Owner.Display_name != "atbebtg" {
		t.Error("Owner invalid.")
	}

	if posts.Items[0].Creation_date != 1327813841 {
		t.Error("Date invalid.")
	}

}

func TestGetPosts(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/posts/1;2;3", dummyPostsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.GetPosts([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyPostsResponse string = `
{
  "items": [
    {
      "post_id": 9051104,
      "post_type": "question",
      "owner": {
        "user_id": 502574,
        "display_name": "atbebtg",
        "reputation": 614,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/3d796e73a7bf18f9338290dbca0d6717?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/502574/atbebtg"
      },
      "creation_date": 1327813841,
      "last_activity_date": 1327813841,
      "score": 0
    },
    {
      "post_id": 9051103,
      "post_type": "question",
      "owner": {
        "user_id": 460920,
        "display_name": "user460920",
        "reputation": 16,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/10ea58fd1c8456a24d9f671d446935df?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/460920/user460920"
      },
      "creation_date": 1327813831,
      "last_activity_date": 1327813831,
      "score": 0
    },
    {
      "post_id": 9051102,
      "post_type": "question",
      "owner": {
        "user_id": 483619,
        "display_name": "Dave Ferguson",
        "reputation": 554,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/a7bdbd77eff0feb358f106588a83b149?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/483619/dave-ferguson"
      },
      "creation_date": 1327813751,
      "last_activity_date": 1327813751,
      "score": 0
    }
  ]
}
`
