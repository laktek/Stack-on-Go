package stackongo

import (
	"os"
	"fmt"
)

func (session Session) getPrivileges(path string, params map[string]string) (output *Privileges, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(Privileges))
	output = parsed_response.(*Privileges)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// AllPrivileges returns all privileges available in site 
func (session Session) AllPrivileges(params map[string]string) (output *Privileges, error os.Error) {
	return session.getPrivileges("privileges", params)
}

// PrivilegesForUser returns the privileges a user has 
func (session Session) PrivilegesForUser(id int, params map[string]string) (output *Privileges, error os.Error) {
	return session.getPrivileges(fmt.Sprintf("users/%v/privileges", id), params)
}
