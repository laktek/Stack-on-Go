package stackongo

import (
	"testing"
)

func TestInbox(t *testing.T) {
	dummy_server := returnDummyResponseForPathAndParams("/2.0/inbox", map[string]string{"key": "app123", "access_token": "abc"}, dummyInboxResponse, t)
	defer closeDummyServer(dummy_server)

	inbox_items, err := Inbox(map[string]string{"page": "1"}, map[string]string{"key": "app123", "access_token": "abc"})

	if err != nil {
		t.Error(err.Error())
	}

	if len(inbox_items.Items) != 1 {
		t.Error("Number of items wrong.")
	}

	if inbox_items.Items[0].Item_type != "new_answer" {
		t.Error("Type invalid.")
	}

	if inbox_items.Items[0].Question_id != 1202 {
		t.Error("Question id invalid.")
	}

	if inbox_items.Items[0].Title != "Sample question" {
		t.Error("Title invalid.")
	}

	if inbox_items.Items[0].Creation_date != 1328148124 {
		t.Error("Creation date invalid.")
	}

	if inbox_items.Items[0].Site.Name != "Stack Apps" {
		t.Error("Creation date invalid.")
	}

}

func TestUnreadInbox(t *testing.T) {
	dummy_server := returnDummyResponseForPathAndParams("/2.0/inbox/unread", map[string]string{"key": "app123", "access_token": "abc"}, dummyInboxResponse, t)
	defer closeDummyServer(dummy_server)

	_, err := UnreadInbox(map[string]string{"page": "1"}, map[string]string{"key": "app123", "access_token": "abc"})

	if err != nil {
		t.Error(err.Error())
	}

}

//Test Data

var dummyInboxResponse string = `
{
  "items": [
    {
      "item_type": "new_answer",
      "question_id": 1202,
      "answer_id": 1,
      "title": "Sample question",
      "creation_date": 1328148124,
      "is_unread": false,
      "site": {
        "site_type": "main_site",
        "name": "Stack Apps",
        "logo_url": "http://sstatic.net/stackapps/img/logo.png",
        "api_site_parameter": "stackapps",
        "site_url": "http://stackapps.com",
        "audience": "apps, scripts, and development with the Stack Exchange API",
        "icon_url": "http://sstatic.net/stackapps/img/apple-touch-icon.png",
        "site_state": "normal",
        "styling": {
          "link_color": "#0077DD",
          "tag_foreground_color": "#555555",
          "tag_background_color": "#E7ECEC"
        },
        "launch_date": 1274313600,
        "favicon_url": "http://sstatic.net/stackapps/img/favicon.ico",
        "related_sites": [
          {
            "name": "Chat Stack Exchange",
            "site_url": "http://chat.stackexchange.com",
            "relation": "chat"
          }
        ]
      },
      "link": "http://stackapps.com/questions/1220/sample-question/1#1"
    }
  ]
}
`
