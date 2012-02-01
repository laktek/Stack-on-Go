package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getTags(path string, params map[string]string) (output []Tag, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(tagsCollection))
	collection := parsed_response.(*tagsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// AllTags returns all tags in site 
func (session Session) AllTags(params map[string]string) (output []Tag, error os.Error) {
	return session.getTags("tags", params)
}

// TagsForUsers returns the tags the users identified with given ids have been active in. 
func (session Session) TagsForUsers(ids []int, params map[string]string) (output []Tag, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "tags"}, "/")
	return session.getTags(request_path, params)
}

// RelatedTags returns the tags that are most related to those in given tags. 
func (session Session) RelatedTags(tags []string, params map[string]string) (output []Tag, error os.Error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "tags"}, "/")
	return session.getTags(request_path, params)
}
