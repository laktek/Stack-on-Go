package stackongo

import (
	"testing"
)

func TestAllComments(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/comments", dummyCommentsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	comments, err := session.AllComments(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(comments) != 3 {
		t.Error("Number of items wrong.")
	}

	if comments[0].Comment_id != 11354978 {
		t.Error("ID invalid.")
	}

	if comments[0].Owner.Display_name != "mynameisneo" {
		t.Error("Owner invalid.")
	}

	if comments[1].Reply_to_user.Display_name != "user1056824" {
		t.Error("Owner invalid.")
	}

	if comments[0].Creation_date != 1327798867 {
		t.Error("Date invalid.")
	}

}

func TestComments(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/comments/1;2;3", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.Comments([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestCommentsForQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/1;2;3/comments", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.CommentsForQuestions([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestCommentsForAnswers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/answers/1;2;3/comments", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.CommentsForAnswers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestCommentsForPosts(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/posts/1;2;3/comments", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.CommentsForPosts([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestCommentsFromUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/comments", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.CommentsFromUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestCommentsMentionedUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/mentioned", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.CommentsMentionedUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestCommentsFromUsersTo(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/comments/4", dummyCommentsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.CommentsFromUsersTo([]int{1, 2, 3}, 4, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyCommentsResponse string = `
{
  "items": [
    {
      "comment_id": 11354978,
      "post_id": 9050020,
      "creation_date": 1327798867,
      "edited": false,
      "owner": {
        "user_id": 1127391,
        "display_name": "mynameisneo",
        "reputation": 57,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/002fa792e5474d3c531da293c6d1a924?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/1127391/mynameisneo"
      }
    },
    {
      "comment_id": 11354977,
      "post_id": 9039268,
      "creation_date": 1327798866,
      "edited": false,
      "owner": {
        "user_id": 133414,
        "display_name": "mvanveen",
        "reputation": 266,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/6695f249160eba47fed3e5fcefd2c942?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/133414/mvanveen"
      },
      "reply_to_user": {
        "user_id": 1056824,
        "display_name": "user1056824",
        "reputation": 11,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/573fc5433388d9245b0e9f1c059e2bb1?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/1056824/user1056824"
      }
    },
    {
      "comment_id": 11354976,
      "post_id": 9049774,
      "creation_date": 1327798862,
      "edited": false,
      "owner": {
        "user_id": 771771,
        "display_name": "Marty Griffin",
        "reputation": 39,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/6e7c1d04d78c984607e864b93227f0c8?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/771771/marty-griffin"
      }
    }
  ]
}
`
