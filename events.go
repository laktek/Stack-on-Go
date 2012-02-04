package stackongo

import (
	"os"
)

func (session Session) getEvents(path string, params map[string]string) (output *Events, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(Events))
	output = parsed_response.(*Events)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return
}

// Events returns a stream of events that have occurred on the site 
// This method requires an access_token.
func (session Session) Events(params map[string]string, auth map[string]string) (output *Events, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	return session.getEvents("events", params)
}
