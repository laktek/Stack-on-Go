package stackongo

import (
	"os"
  "strings"
	"fmt"
)

func (session Session) getRevisions(path string, params map[string]string) (output []Revision, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(revisionsCollection))
	collection := parsed_response.(*revisionsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// Revisions returns edit revisions identified by given ids
func (session Session) Revisions(ids []int, params map[string]string) (output []Revision, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"revisions", strings.Join(string_ids, ";")}, "/")
	return session.getRevisions(request_path, params)
}

// RevisionsForPosts returns the revisions for the posts identified with given ids
func (session Session) RevisionsForPosts(ids []int, params map[string]string) (output []Revision, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";"), "revisions"}, "/")
	return session.getRevisions(request_path, params)
}
