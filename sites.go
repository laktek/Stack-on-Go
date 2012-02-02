package stackongo

import (
	"os"
	"http"
)

// Sites returns all sites available in StackExchange network
func Sites(params map[string]string) (output []Site, error os.Error) {
	client := new(http.Client)

	// make the request
	response, err := client.Get(setupEndpoint("sites", params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(sitesCollection))
	collection := parsed_response.(*sitesCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}
