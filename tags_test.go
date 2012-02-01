package stackongo

import (
	"testing"
)

func TestAllTags(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/tags", dummyTagsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	tags, err := session.AllTags(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(tags) != 3 {
		t.Error("Number of items wrong.")
	}

	if tags[0].Name != "c#" {
		t.Error("Name invalid.")
	}

	if tags[0].Count != 261768 {
		t.Error("Tag count invalid.")
	}

	if tags[0].Has_synonyms != true {
		t.Error("boolean invalid.")
	}

}

func TestTagsForUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/tags", dummyTagsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.TagsForUsers([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestRelatedTags(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/tags/tag1;tag2;tag3/tags", dummyTagsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.RelatedTags([]string{"tag1", "tag2", "tag3"}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyTagsResponse string = `
{
  "items": [
    {
      "name": "c#",
      "count": 261768,
      "is_required": false,
      "is_moderator_only": false,
      "has_synonyms": true
    },
    {
      "name": "java",
      "count": 202323,
      "is_required": false,
      "is_moderator_only": false,
      "has_synonyms": true
    },
    {
      "name": "php",
      "count": 187267,
      "is_required": false,
      "is_moderator_only": false,
      "has_synonyms": true
    }
  ]
}
`
