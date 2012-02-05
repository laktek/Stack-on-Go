package stackongo

import (
	"os"
	"strings"
	"fmt"
)

// AllPosts returns all Posts in site 
func (session Session) AllPosts(params map[string]string) (output *Posts, error os.Error) {
	output = new(Posts)
	error = session.get("posts", params, output)
	return
}

// Posts returns the posts with the given ids
func (session Session) GetPosts(ids []int, params map[string]string) (output *Posts, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";")}, "/")

	output = new(Posts)
	error = session.get(request_path, params, output)
	return

}
