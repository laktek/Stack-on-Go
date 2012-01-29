package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getComments(path string, params map[string]string) (output []Comment, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(commentsCollection))
	collection := parsed_response.(*commentsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// AllComments returns all comments in site 
func (session Session) AllComments(params map[string]string) (output []Comment, error os.Error) {
	return session.getComments("comments", params)
}

// Comments returns the comments with the given ids
func (session Session) Comments(ids []int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"comments", strings.Join(string_ids, ";")}, "/")
	return session.getComments(request_path, params)
}

// CommentsForQuestions returns the comments for the questions identified with given ids
func (session Session) CommentsForQuestions(ids []int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "comments"}, "/")
	return session.getComments(request_path, params)
}

// CommentsForAnswers returns the comments for the answers identified with given ids
func (session Session) CommentsForAnswers(ids []int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"answers", strings.Join(string_ids, ";"), "comments"}, "/")
	return session.getComments(request_path, params)
}

// CommentsForPosts returns the comments for the posts identified with given ids
func (session Session) CommentsForPosts(ids []int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";"), "comments"}, "/")
	return session.getComments(request_path, params)
}

// CommentsFromUsers returns the comments from the users identified with given ids
func (session Session) CommentsFromUsers(ids []int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "comments"}, "/")
	return session.getComments(request_path, params)
}

// CommentsMentionedUsers returns the comments mentioning the users identified with given ids
func (session Session) CommentsMentionedUsers(ids []int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "mentioned"}, "/")
	return session.getComments(request_path, params)
}

// CommentsFromUsersTo returns the comments to a user from the users identified with given ids
func (session Session) CommentsFromUsersTo(ids []int, to int, params map[string]string) (output []Comment, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "comments", fmt.Sprintf("%v", to)}, "/")
	return session.getComments(request_path, params)
}
