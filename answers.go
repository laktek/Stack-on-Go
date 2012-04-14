package stackongo

import (
	"fmt"
	"strings"
)

// AllAnswers returns all answers in site 
func (session Session) AllAnswers(params map[string]string) (output *Answers, error error) {
	output = new(Answers)
	error = session.get("answers", params, output)
	return
}

// Answers returns the answers with the given ids
func (session Session) GetAnswers(ids []int, params map[string]string) (output *Answers, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"answers", strings.Join(string_ids, ";")}, "/")

	output = new(Answers)
	error = session.get(request_path, params, output)
	return
}

// AnswersForQuestions returns the answers for the questions identified with given ids
func (session Session) AnswersForQuestions(ids []int, params map[string]string) (output *Answers, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "answers"}, "/")

	output = new(Answers)
	error = session.get(request_path, params, output)
	return
}

// AnswersFromUsers returns the answers from the users identified with given ids
func (session Session) AnswersFromUsers(ids []int, params map[string]string) (output *Answers, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "answers"}, "/")

	output = new(Answers)
	error = session.get(request_path, params, output)
	return
}

// TopAnswersFromUsers returns the top answers from the users identified with given ids for the questions with given tags
func (session Session) TopAnswersFromUsers(ids []int, tags []string, params map[string]string) (output *Answers, error error) {

	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}

	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "tags", strings.Join(tags, ";"), "top-answers"}, "/")

	output = new(Answers)
	error = session.get(request_path, params, output)
	return
}
