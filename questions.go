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

// QuestionsFromUsers returns the questions asked by the users with given ids
func (session Session) QuestionsFromUsers(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions"}, "/")
	return session.getQuestions(request_path, params)
}

// QuestionsWithNoAnswersFromUsers returns the questions without answers asked by the users with given ids
func (session Session) QuestionsWithNoAnswersFromUsers(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions", "no-answers"}, "/")
	return session.getQuestions(request_path, params)
}

// UnacceptedQuestionsFromUsers returns the unaccepted questions asked by the users with given ids
func (session Session) UnacceptedQuestionsFromUsers(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions", "unaccepted"}, "/")
	return session.getQuestions(request_path, params)
}

// UnansweredQuestionsFromUsers returns the unanswered questions asked by the users with given ids
func (session Session) UnansweredQuestionsFromUsers(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "questions", "unanswered"}, "/")
	return session.getQuestions(request_path, params)
}

// FavoriteQuestionsFromUsers returns the favorite questions by users with given ids
func (session Session) FavoriteQuestionsFromUsers(ids []int, params map[string]string) (output []Question, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "favorites"}, "/")
	return session.getQuestions(request_path, params)
}

// TopQuestionsFromUsers returns the top questions from the users identified with given ids for the questions with given tags
func (session Session) TopQuestionsFromUsers(ids []int, tags []string, params map[string]string) (output []Question, error os.Error) {

	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}

	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "tags", strings.Join(tags, ";"), "top-questions"}, "/")
	return session.getQuestions(request_path, params)
}

// FAQForTags returns the frequently asked questions for the given tags
func (session Session) FAQForTags(tags []string, params map[string]string) (output []Question, error os.Error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "faq"}, "/")
	return session.getQuestions(request_path, params)
}

// Search queries the post titles with the given query and returns the matching questions. You can used params to set other criteria such as `tagged`
func (session Session) Search(query string, params map[string]string) (output []Question, error os.Error) {
	request_path := "search"
	// set query as a param
	params["intitle"] = query

	return session.getQuestions(request_path, params)
}

// Similar returns questions similar to the given query
func (session Session) Similar(query string, params map[string]string) (output []Question, error os.Error) {
	request_path := "similar"
	// set query as a param
	params["title"] = query
	return session.getQuestions(request_path, params)
}
