package stackongo

import (
	"fmt"
	"strings"
)

// ReputationChangesForUsers returns a subset of the reputation changes for users with given ids. 
func (session Session) ReputationChangesForUsers(ids []int, params map[string]string) (output *Reputations, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "reputation"}, "/")

	output = new(Reputations)
	error = session.get(request_path, params, output)
	return
}
