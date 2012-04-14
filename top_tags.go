package stackongo

import (
	"fmt"
	"strings"
)

// TopTagsByAnswerForUser returns a single user's top tags by answer score. 
func (session Session) TopTagsByAnswerForUser(id int, params map[string]string) (output *TopTags, error error) {
	request_path := strings.Join([]string{"users", fmt.Sprintf("%v", id), "top-answer-tags"}, "/")

	output = new(TopTags)
	error = session.get(request_path, params, output)
	return
}

// TopTagsByQuestionForUser returns a single user's top tags by question score. 
func (session Session) TopTagsByQuestionForUser(id int, params map[string]string) (output *TopTags, error error) {
	request_path := strings.Join([]string{"users", fmt.Sprintf("%v", id), "top-question-tags"}, "/")

	output = new(TopTags)
	error = session.get(request_path, params, output)
	return
}
