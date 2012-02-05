package stackongo

import (
	"os"
)

// Inbox returns authenticated user's inbox. 
// This method requires an access_token, with a scope containing "read_inbox".
func Inbox(params map[string]string, auth map[string]string) (output *InboxItems, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	output = new(InboxItems)
	error = get("inbox", params, output)
	return
}

// UnreadInbox returns unread items in an authenticated user's inbox. 
// This method requires an access_token, with a scope containing "read_inbox".
func UnreadInbox(params map[string]string, auth map[string]string) (output *InboxItems, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	output = new(InboxItems)
	error = get("inbox/unread", params, output)
	return
}
