package stackongo

import (
	"os"
	"http"
  "strings"
)

// CreateFilter creates a new filter given a list of includes, excludes, a base filter, and whether or not this filter should be "unsafe". 
func CreateFilter(params map[string]string) (output []Filter, error os.Error) {
	client := new(http.Client)

	// make the request
	response, err := client.Get(setupEndpoint("filters/create", params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(filtersCollection))
	collection := parsed_response.(*filtersCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error
}

// InspectFilters returns the fields included by the given filters, and the "safeness" of those filters.
func InspectFilters(filters []string, params map[string]string) (output []Filter, error os.Error) {

	request_path := strings.Join([]string{"filters", strings.Join(filters, ";")}, "/")

	// make the request
	client := new(http.Client)
	response, err := client.Get(setupEndpoint(request_path, params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(filtersCollection))
	collection := parsed_response.(*filtersCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}
