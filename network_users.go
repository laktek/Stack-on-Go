package stackongo

import (
	"os"
	"http"
	"fmt"
	"strings"
)

// AssociatedAccounts returns all associated accounts for the given user ids.
func AssociatedAccounts(ids []int, params map[string]string) (output []NetworkUser, error os.Error) {

	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "associated"}, "/")

	// make the request
	client := new(http.Client)
	response, err := client.Get(setupEndpoint(request_path, params).String())

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(networkUsersCollection))
	collection := parsed_response.(*networkUsersCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}
