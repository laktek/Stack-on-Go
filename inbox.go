package stackongo

import (
	"os"
	"http"
)

func getInbox(path string, params map[string]string) (output *InboxItems, error os.Error) {

	client := new(http.Client)

	// make the reques
	response, err := client.Get(setupEndpoint(path, params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(InboxItems))
	output = parsed_response.(*InboxItems)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return
}

// Inbox returns authenticated user's inbox. 
// This method requires an access_token, with a scope containing "read_inbox".
func Inbox(params map[string]string, auth map[string]string) (output *InboxItems, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	return getInbox("inbox", params)
}

// UnreadInbox returns unread items in an authenticated user's inbox. 
// This method requires an access_token, with a scope containing "read_inbox".
func UnreadInbox(params map[string]string, auth map[string]string) (output *InboxItems, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	return getInbox("inbox/unread", params)
}
