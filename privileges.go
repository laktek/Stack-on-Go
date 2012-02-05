package stackongo

import (
	"os"
	"fmt"
)

// AllPrivileges returns all privileges available in site 
func (session Session) AllPrivileges(params map[string]string) (output *Privileges, error os.Error) {
	output = new(Privileges)
	error = session.get("privileges", params, output)
	return
}

// PrivilegesForUser returns the privileges a user has 
func (session Session) PrivilegesForUser(id int, params map[string]string) (output *Privileges, error os.Error) {
	request_path := fmt.Sprintf("users/%v/privileges", id)

	output = new(Privileges)
	error = session.get(request_path, params, output)
	return
}
