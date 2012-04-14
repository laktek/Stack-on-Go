package stackongo

import "fmt"

// AllErrors returns the various error codes that can be produced by the API. 
func AllErrors(params map[string]string) (output *Errors, error error) {
	output = new(Errors)
	error = get("errors", params, output)
	return
}

// SimulateError allows you to simulate an error response
func SimulateError(id int) (output *Error, error error) {
	request_path := fmt.Sprintf("errors/%v", id)

	output = new(Error)
	error = get(request_path, map[string]string{}, output)
	return
}
