package stackongo

import (
	"fmt"
	"strings"
)

// TimelineForUsers returns a subset of the actions the users with given ids have taken on the site. 
func (session Session) TimelineForUsers(ids []int, params map[string]string) (output *UserTimelines, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "timeline"}, "/")

	output = new(UserTimelines)
	error = session.get(request_path, params, output)
	return
}
