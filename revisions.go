package stackongo

import (
	"fmt"
	"strings"
)

// Revisions returns edit revisions identified by given ids
func (session Session) Revisions(ids []int, params map[string]string) (output *Revisions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"revisions", strings.Join(string_ids, ";")}, "/")

	output = new(Revisions)
	error = session.get(request_path, params, output)
	return
}

// RevisionsForPosts returns the revisions for the posts identified with given ids
func (session Session) RevisionsForPosts(ids []int, params map[string]string) (output *Revisions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";"), "revisions"}, "/")

	output = new(Revisions)
	error = session.get(request_path, params, output)
	return
}
