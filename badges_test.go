package stackongo

import (
	"testing"
)

func TestAllBadges(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/badges", dummyBadgesResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	badges, err := session.AllBadges(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

	if len(badges.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if badges.Items[0].Badge_id != 455 {
		t.Error("ID invalid.")
	}

	if badges.Items[0].Rank != "bronze" {
		t.Error("Rank invalid.")
	}

	if badges.Items[0].Badge_type != "tag_based" {
		t.Error("Badge type invalid.")
	}

}

func TestGetBadges(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/badges/1;2;3", dummyBadgesResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.GetBadges([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

func TestNamedBadges(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/badges/name", dummyBadgesResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.NamedBadges(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

func TestTagBadges(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/badges/tags", dummyBadgesResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	_, err := session.TagBadges(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

func TestRecentAllBadgeRecipients(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/badges/recipients", dummyBadgesResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	badges, err := session.RecentAllBadgeRecipients(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

	if badges.Items[0].User.Display_name != "Joel Martinez" {
		t.Error("User invalid.")
	}

}

func TestRecentBadgeRecipients(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/badges/1;2;3/recipients", dummyBadgesResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	badges, err := session.RecentBadgeRecipients([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

	if badges.Items[0].User.Display_name != "Joel Martinez" {
		t.Error("User invalid.")
	}

}

func TestBadgeOfUsers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/1;2;3/badges", dummyBadgesResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("stackoverflow")
	badges, err := session.BadgesOfUsers([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

	if badges.Items[0].User.Display_name != "Joel Martinez" {
		t.Error("User invalid.")
	}

}

//Test Data

var dummyBadgesResponse string = `
{
  "items": [
    {
      "badge_id": 455,
      "rank": "bronze",
      "name": "xna",
      "award_count": 2,
      "badge_type": "tag_based",
      "user": {
        "user_id": 5416,
        "display_name": "Joel Martinez",
        "reputation": 15030,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/c7fbf557a6c37f0ae2b300f6799c767f?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/5416/joel-martinez"
      },
      "link": "http://stackoverflow.com/badges/455/xna"
    },
    {
      "badge_id": 456,
      "rank": "bronze",
      "name": "elisp",
      "award_count": 7,
      "badge_type": "tag_based",
      "link": "http://stackoverflow.com/badges/456/elisp"
    },
    {
      "badge_id": 457,
      "rank": "bronze",
      "name": "latex",
      "award_count": 10,
      "badge_type": "tag_based",
      "link": "http://stackoverflow.com/badges/457/latex"
    }
  ]
}
`
