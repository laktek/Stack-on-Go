package stackongo

import (
	"fmt"
	"strings"
)

// AllSuggestedEdits returns all the suggested edits in the systems. 
func (session Session) AllSuggestedEdits(params map[string]string) (output *SuggestedEdits, error error) {
	output = new(SuggestedEdits)
	error = session.get("suggested-edits", params, output)
	return
}

// SuggestedEdits returns suggested edits identified in ids. 
func (session Session) GetSuggestedEdits(ids []int, params map[string]string) (output *SuggestedEdits, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"suggested-edits", strings.Join(string_ids, ";")}, "/")

	output = new(SuggestedEdits)
	error = session.get(request_path, params, output)
	return
}

// SuggestedEditsForPosts returns the suggested edits for the posts identified with given ids
func (session Session) SuggestedEditsForPosts(ids []int, params map[string]string) (output *SuggestedEdits, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";"), "suggested-edits"}, "/")

	output = new(SuggestedEdits)
	error = session.get(request_path, params, output)
	return
}

// SuggestedEditsFromUsers returns the suggested edits submitted by users with given ids. 
func (session Session) SuggestedEditsFromUsers(ids []int, params map[string]string) (output *SuggestedEdits, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "suggested-edits"}, "/")

	output = new(SuggestedEdits)
	error = session.get(request_path, params, output)
	return
}
