package stackongo

import (
	"os"
	"http"
)

// AllSites returns all sites available in StackExchange network
func AllSites(params map[string]string) (output *Sites, error os.Error) {
	client := new(http.Client)

	// make the request
	response, err := client.Get(setupEndpoint("sites", params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(Sites))
	output = parsed_response.(*Sites)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}
