package stackongo

import (
	"fmt"
	"strings"
)

// AllQuestions returns all questions
func (session Session) AllQuestions(params map[string]string) (output *Questions, error error) {
	output = new(Questions)
	error = session.get("questions", params, output)
	return
}

// Questions returns the questions for given ids
func (session Session) GetQuestions(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";")}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// UnansweredQuestions returns all unanswered questions
func (session Session) UnansweredQuestions(params map[string]string) (output *Questions, error error) {
	output = new(Questions)
	error = session.get("questions/unanswered", params, output)
	return
}

// QuestionsWithNoAnswers returns questions with no answers
func (session Session) QuestionsWithNoAnswers(params map[string]string) (output *Questions, error error) {
	output = new(Questions)
	error = session.get("questions/no-answers", params, output)
	return
}

// ReleatedQuestions returns the questions releated to the questions identified with given ids
func (session Session) RelatedQuestions(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "related"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// LinkedQuestions returns the questions linked to the questions identified with given ids
func (session Session) LinkedQuestions(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "linked"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// QuestionsFromUsers returns the questions asked by the users with given ids
func (session Session) QuestionsFromUsers(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// QuestionsWithNoAnswersFromUsers returns the questions without answers asked by the users with given ids
func (session Session) QuestionsWithNoAnswersFromUsers(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions", "no-answers"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// UnacceptedQuestionsFromUsers returns the unaccepted questions asked by the users with given ids
func (session Session) UnacceptedQuestionsFromUsers(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions", "unaccepted"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// UnansweredQuestionsFromUsers returns the unanswered questions asked by the users with given ids
func (session Session) UnansweredQuestionsFromUsers(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions", "unanswered"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// FavoriteQuestionsFromUsers returns the favorite questions by users with given ids
func (session Session) FavoriteQuestionsFromUsers(ids []int, params map[string]string) (output *Questions, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "favorites"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// TopQuestionsFromUsers returns the top questions from the users identified with given ids for the questions with given tags
func (session Session) TopQuestionsFromUsers(ids []int, tags []string, params map[string]string) (output *Questions, error error) {

	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}

	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "tags", strings.Join(tags, ";"), "top-questions"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// FAQForTags returns the frequently asked questions for the given tags
func (session Session) FAQForTags(tags []string, params map[string]string) (output *Questions, error error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "faq"}, "/")

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// Search queries the post titles with the given query and returns the matching questions. You can used params to set other criteria such as `tagged`
func (session Session) Search(query string, params map[string]string) (output *Questions, error error) {
	request_path := "search"
	// set query as a param
	params["intitle"] = query

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}

// Similar returns questions similar to the given query
func (session Session) Similar(query string, params map[string]string) (output *Questions, error error) {
	request_path := "similar"
	// set query as a param
	params["title"] = query

	output = new(Questions)
	error = session.get(request_path, params, output)
	return
}
