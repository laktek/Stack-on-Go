package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getSuggestedEdits(path string, params map[string]string) (output *SuggestedEdits, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(SuggestedEdits))
	output = parsed_response.(*SuggestedEdits)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// AllSuggestedEdits returns all the suggested edits in the systems. 
func (session Session) AllSuggestedEdits(params map[string]string) (output *SuggestedEdits, error os.Error) {
	return session.getSuggestedEdits("suggested-edits", params)
}

// SuggestedEdits returns suggested edits identified in ids. 
func (session Session) SuggestedEdits(ids []int, params map[string]string) (output *SuggestedEdits, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"suggested-edits", strings.Join(string_ids, ";")}, "/")
	return session.getSuggestedEdits(request_path, params)
}

// SuggestedEditsForPosts returns the suggested edits for the posts identified with given ids
func (session Session) SuggestedEditsForPosts(ids []int, params map[string]string) (output *SuggestedEdits, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";"), "suggested-edits"}, "/")
	return session.getSuggestedEdits(request_path, params)
}

// SuggestedEditsFromUsers returns the suggested edits submitted by users with given ids. 
func (session Session) SuggestedEditsFromUsers(ids []int, params map[string]string) (output *SuggestedEdits, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "suggested-edits"}, "/")
	return session.getSuggestedEdits(request_path, params)
}
