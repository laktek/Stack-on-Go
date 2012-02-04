package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getQuestionTimelines(path string, params map[string]string) (output *QuestionTimelines, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(QuestionTimelines))
	output = parsed_response.(*QuestionTimelines)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// TimelineForQuestions returns a subset of the events that have happened to the questions identified with ids. 
func (session Session) TimelineForQuestions(ids []int, params map[string]string) (output *QuestionTimelines, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "timeline"}, "/")
	return session.getQuestionTimelines(request_path, params)
}
