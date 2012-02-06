package stackongo

import (
	"testing"
)

func TestAllTagSynonyms(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/tags/synonyms", dummyTagSynonymsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	tag_synonyms, err := session.AllTagSynonyms(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(tag_synonyms.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if tag_synonyms.Items[0].From_tag != "acoustic-echo-cancellatio" {
		t.Error("From tag invalid.")
	}

	if tag_synonyms.Items[0].To_tag != "aec" {
		t.Error("To tag invalid.")
	}

	if tag_synonyms.Items[0].Applied_count != 0 {
		t.Error("Applied count invalid.")
	}

	if tag_synonyms.Items[0].Creation_date != 1327953917 {
		t.Error("Creation date invalid.")
	}

}

func TestSynonymsForTags(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/tags/tag1;tag2;tag3/synonyms", dummyTagSynonymsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.SynonymsForTags([]string{"tag1", "tag2", "tag3"}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyTagSynonymsResponse string = `
{
  "items": [
    {
      "from_tag": "acoustic-echo-cancellatio",
      "to_tag": "aec",
      "applied_count": 0,
      "creation_date": 1327953917
    },
    {
      "from_tag": "validate",
      "to_tag": "validation",
      "applied_count": 3,
      "last_applied_date": 1328085114,
      "creation_date": 1327411904
    },
    {
      "from_tag": "assets-pipeline",
      "to_tag": "asset-pipeline",
      "applied_count": 0,
      "creation_date": 1327698162
    }
  ]
}
`
