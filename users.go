package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getUsers(path string, params map[string]string) (output []User, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(usersCollection))
	collection := parsed_response.(*usersCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// AllUsers returns all users in site 
func (session Session) AllUsers(params map[string]string) (output []User, error os.Error) {
	return session.getUsers("users", params)
}

// Users returns the users with the given ids
func (session Session) Users(ids []int, params map[string]string) (output []User, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";")}, "/")
	return session.getUsers(request_path, params)
}

// AuthenticatedUser returns the user associated with the passed auth_token.
func (session Session) AuthenticatedUser(params map[string]string, auth map[string]string) (output User, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

  results, error := session.getUsers("me", params)
  return results[0], error
}

// Moderators returns those users on a site who can exercise moderation powers. 
func (session Session) Moderators(params map[string]string) (output []User, error os.Error) {
	return session.getUsers("users/moderators", params)
}

// ElectedModerators returns those users on a site who both have moderator powers, and were actually elected. 
func (session Session) ElectedModerators(params map[string]string) (output []User, error os.Error) {
	return session.getUsers("users/moderators/elected", params)
}
