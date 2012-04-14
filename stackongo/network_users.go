package stackongo

import (
	"fmt"
	"strings"
)

// AssociatedAccounts returns all associated accounts for the given user ids.
func AssociatedAccounts(ids []int, params map[string]string) (output *NetworkUsers, error error) {

	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "associated"}, "/")

	output = new(NetworkUsers)
	error = get(request_path, params, output)
	return

}
