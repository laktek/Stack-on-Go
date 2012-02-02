package stackongo

import (
	"os"
	"http"
	"fmt"
)

// AllErrors returns the various error codes that can be produced by the API. 
func AllErrors(params map[string]string) (output []Error, error os.Error) {
	client := new(http.Client)

	// make the request
	response, err := client.Get(setupEndpoint("errors", params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(errorsCollection))
	collection := parsed_response.(*errorsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error
}

// SimulateError allows you to simulate an error response
func SimulateError(id int) (error os.Error) {
	client := new(http.Client)

	request_path := fmt.Sprintf("errors/%v", id)

	// make the request
	response, err := client.Get(setupEndpoint(request_path, map[string]string{}).String())

	if err != nil {
		return err
	}

	parsed_response, error := parseResponse(response, new(Error))
	collection := parsed_response.(*Error)

	//overload the generic error with details
	error = os.NewError(collection.Error_name + ": " + collection.Error_message)

	return error

}
