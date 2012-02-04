package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getUserTimelines(path string, params map[string]string) (output *UserTimelines, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(UserTimelines))
	output = parsed_response.(*UserTimelines)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// TimelineForUsers returns a subset of the actions the users with given ids have taken on the site. 
func (session Session) TimelineForUsers(ids []int, params map[string]string) (output *UserTimelines, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "timeline"}, "/")
	return session.getUserTimelines(request_path, params)
}
