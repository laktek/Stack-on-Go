package stackongo

import (
	"fmt"
	"strings"
)

// AllTags returns all tags in site 
func (session Session) AllTags(params map[string]string) (output *Tags, error error) {
	output = new(Tags)
	error = session.get("tags", params, output)
	return
}

// TagsForUsers returns the tags the users identified with given ids have been active in. 
func (session Session) TagsForUsers(ids []int, params map[string]string) (output *Tags, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "tags"}, "/")

	output = new(Tags)
	error = session.get(request_path, params, output)
	return
}

// RelatedTags returns the tags that are most related to those in given tags. 
func (session Session) RelatedTags(tags []string, params map[string]string) (output *Tags, error error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "tags"}, "/")

	output = new(Tags)
	error = session.get(request_path, params, output)
	return
}
