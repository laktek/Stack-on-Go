package stackongo

import (
	"os"
)

func (session Session) getInfo(path string, params map[string]string) (output Info, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(infoCollection))
	collection := parsed_response.(*infoCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items[0]
	}

	return output, error

}

// Info returns a collection of statistics about the site 
func (session Session) Info() (output Info, error os.Error) {
	return session.getInfo("info", map[string]string{})
}
