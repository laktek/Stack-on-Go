package stackongo

import (
	"os"
	"strings"
	"fmt"
)

// AllUsers returns all users in site 
func (session Session) AllUsers(params map[string]string) (output *Users, error os.Error) {
	output = new(Users)
	error = session.get("users", params, output)
	return
}

// Users returns the users with the given ids
func (session Session) GetUsers(ids []int, params map[string]string) (output *Users, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";")}, "/")

	output = new(Users)
	error = session.get(request_path, params, output)
	return
}

// AuthenticatedUser returns the user associated with the passed auth_token.
func (session Session) AuthenticatedUser(params map[string]string, auth map[string]string) (output User, error os.Error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	collection := new(Users)
	error = session.get("me", params, collection)

	if len(collection.Items) > 0 {
		output = collection.Items[0]
	} else {
		error = os.NewError("User not found")
	}

	return

}

// Moderators returns those users on a site who can exercise moderation powers. 
func (session Session) Moderators(params map[string]string) (output *Users, error os.Error) {
	output = new(Users)
	error = session.get("users/moderators", params, output)
	return
}

// ElectedModerators returns those users on a site who both have moderator powers, and were actually elected. 
func (session Session) ElectedModerators(params map[string]string) (output *Users, error os.Error) {
	output = new(Users)
	error = session.get("users/moderators/elected", params, output)
	return
}
