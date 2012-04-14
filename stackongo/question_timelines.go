package stackongo

import (
	"fmt"
	"strings"
)

// TimelineForQuestions returns a subset of the events that have happened to the questions identified with ids. 
func (session Session) TimelineForQuestions(ids []int, params map[string]string) (output *QuestionTimelines, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "timeline"}, "/")

	output = new(QuestionTimelines)
	error = session.get(request_path, params, output)
	return
}
