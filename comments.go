package stackongo

import (
	"fmt"
	"strings"
)

// AllComments returns all comments in site 
func (session Session) AllComments(params map[string]string) (output *Comments, error error) {
	output = new(Comments)
	error = session.get("comments", params, output)
	return
}

// Comments returns the comments with the given ids
func (session Session) GetComments(ids []int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"comments", strings.Join(string_ids, ";")}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}

// CommentsForQuestions returns the comments for the questions identified with given ids
func (session Session) CommentsForQuestions(ids []int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"questions", strings.Join(string_ids, ";"), "comments"}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}

// CommentsForAnswers returns the comments for the answers identified with given ids
func (session Session) CommentsForAnswers(ids []int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"answers", strings.Join(string_ids, ";"), "comments"}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}

// CommentsForPosts returns the comments for the posts identified with given ids
func (session Session) CommentsForPosts(ids []int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"posts", strings.Join(string_ids, ";"), "comments"}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}

// CommentsFromUsers returns the comments from the users identified with given ids
func (session Session) CommentsFromUsers(ids []int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "comments"}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}

// CommentsMentionedUsers returns the comments mentioning the users identified with given ids
func (session Session) CommentsMentionedUsers(ids []int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "mentioned"}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}

// CommentsFromUsersTo returns the comments to a user from the users identified with given ids
func (session Session) CommentsFromUsersTo(ids []int, to int, params map[string]string) (output *Comments, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "comments", fmt.Sprintf("%v", to)}, "/")

	output = new(Comments)
	error = session.get(request_path, params, output)
	return
}
