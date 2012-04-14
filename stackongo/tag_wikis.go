package stackongo

import "strings"

// WikisForTags returns the wikis that go with the given set of tags 
func (session Session) WikisForTags(tags []string, params map[string]string) (output *TagWikis, error error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "wikis"}, "/")

	output = new(TagWikis)
	error = session.get(request_path, params, output)
	return
}
