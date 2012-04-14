package stackongo

import (
	"testing"
)

func TestAllSites(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/sites", dummySitesResponse, t)
	defer closeDummyServer(dummy_server)

	sites, err := AllSites(map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.Error())
	}

	if len(sites.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if sites.Items[0].Site_type != "main_site" {
		t.Error("Site type invalid.")
	}

	if sites.Items[0].Aliases[0] != "http://www.stackoverflow.com" {
		t.Error("alias invalid.")
	}

	if sites.Items[0].Related_sites[0].Name != "Stack Overflow Chat" {
		t.Error("related site invalid.")
	}

}

//Test Data

var dummySitesResponse string = `
{
  "items": [
    {
      "site_type": "main_site",
      "name": "Stack Overflow",
      "logo_url": "http://sstatic.net/stackoverflow/img/logo.png",
      "api_site_parameter": "stackoverflow",
      "site_url": "http://stackoverflow.com",
      "audience": "professional and enthusiast programmers",
      "icon_url": "http://sstatic.net/stackoverflow/img/apple-touch-icon.png",
      "aliases": [
        "http://www.stackoverflow.com"
      ],
      "site_state": "normal",
      "styling": {
        "link_color": "#0077CC",
        "tag_foreground_color": "#3E6D8E",
        "tag_background_color": "#E0EAF1"
      },
      "launch_date": 1221436800,
      "favicon_url": "http://sstatic.net/stackoverflow/img/favicon.ico",
      "related_sites": [
        {
          "name": "Stack Overflow Chat",
          "site_url": "http://chat.stackoverflow.com",
          "relation": "chat"
        }
      ],
      "markdown_extensions": [
        "Prettify"
      ]
    },
    {
      "site_type": "main_site",
      "name": "Server Fault",
      "logo_url": "http://sstatic.net/serverfault/img/logo.png",
      "api_site_parameter": "serverfault",
      "site_url": "http://serverfault.com",
      "audience": "system administrators and desktop support professionals",
      "icon_url": "http://sstatic.net/serverfault/img/apple-touch-icon.png",
      "site_state": "normal",
      "styling": {
        "link_color": "#10456A",
        "tag_foreground_color": "#444444",
        "tag_background_color": "#F3F1D9"
      },
      "launch_date": 1243296000,
      "favicon_url": "http://sstatic.net/serverfault/img/favicon.ico",
      "related_sites": [
        {
          "name": "Meta Server Fault",
          "site_url": "http://meta.serverfault.com",
          "relation": "meta",
          "api_site_parameter": "meta.serverfault"
        },
        {
          "name": "Chat Stack Exchange",
          "site_url": "http://chat.stackexchange.com",
          "relation": "chat"
        }
      ],
      "twitter_account": "ServerFault"
    },
    {
      "site_type": "main_site",
      "name": "Super User",
      "logo_url": "http://sstatic.net/superuser/img/logo.png",
      "api_site_parameter": "superuser",
      "site_url": "http://superuser.com",
      "audience": "computer enthusiasts and power users",
      "icon_url": "http://sstatic.net/superuser/img/apple-touch-icon.png",
      "site_state": "normal",
      "styling": {
        "link_color": "#1086A4",
        "tag_foreground_color": "#1087A4",
        "tag_background_color": "#FFFFFF"
      },
      "launch_date": 1250553600,
      "favicon_url": "http://sstatic.net/superuser/img/favicon.ico",
      "related_sites": [
        {
          "name": "Meta Super User",
          "site_url": "http://meta.superuser.com",
          "relation": "meta",
          "api_site_parameter": "meta.superuser"
        },
        {
          "name": "Chat Stack Exchange",
          "site_url": "http://chat.stackexchange.com",
          "relation": "chat"
        }
      ],
      "twitter_account": "StackSuper_User"
    }
  ]
}
`
