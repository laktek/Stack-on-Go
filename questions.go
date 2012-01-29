package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getQuestions(path string, params map[string]string) (output []Question, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(questionsCollection))
	collection := parsed_response.(*questionsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// AllQuestions returns all questions
func (session Session) AllQuestions(params map[string]string) (output []Question, error os.Error) {
	return session.getQuestions("questions", params)
}

// Questions returns the questions for given ids
func (session Session) Questions(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";")}, "/")
	return session.getQuestions(request_path, params)
}

// UnansweredQuestions returns all unanswered questions
func (session Session) UnansweredQuestions(params map[string]string) (output []Question, error os.Error) {
	return session.getQuestions("questions/unanswered", params)
}

// QuestionsWithNoAnswers returns questions with no answers
func (session Session) QuestionsWithNoAnswers(params map[string]string) (output []Question, error os.Error) {
	return session.getQuestions("questions/no-answers", params)
}

// ReleatedQuestions returns the questions releated to the questions identified with given ids
func (session Session) RelatedQuestions(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "related"}, "/")
	return session.getQuestions(request_path, params)
}

// LinkedQuestions returns the questions linked to the questions identified with given ids
func (session Session) LinkedQuestions(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "linked"}, "/")
	return session.getQuestions(request_path, params)
}
