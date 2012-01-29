package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getAnswers(path string, params map[string]string) (output []Answer, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(answersCollection))
	collection := parsed_response.(*answersCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// AllAnswers returns all answers in site 
func (session Session) AllAnswers(params map[string]string) (output []Answer, error os.Error) {
	return session.getAnswers("answers", params)
}

// Answers returns the answers with the given ids
func (session Session) Answers(ids []int, params map[string]string) (output []Answer, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"answers", strings.Join(string_ids, ";")}, "/")
	return session.getAnswers(request_path, params)
}

// AnswersForQuestions returns the answers for the questions identified with given ids
func (session Session) AnswersForQuestions(ids []int, params map[string]string) (output []Answer, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "answers"}, "/")
	return session.getAnswers(request_path, params)
}

// AnswersFromUsers returns the answers from the users identified with given ids
func (session Session) AnswersFromUsers(ids []int, params map[string]string) (output []Answer, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "answers"}, "/")
	return session.getAnswers(request_path, params)
}

// TopAnswersFromUsers returns the top answers from the users identified with given ids for the questions with given tags
func (session Session) TopAnswersFromUsers(ids []int, tags []string, params map[string]string) (output []Answer, error os.Error) {

	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}

	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "tags", strings.Join(tags, ";"), "top-answers"}, "/")
	return session.getAnswers(request_path, params)
}
