package stackongo

import (
	"testing"
)

func TestAllErrors(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/errors", dummyErrorsResponse, t)
	defer closeDummyServer(dummy_server)

	errors, err := AllErrors(map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(errors.Items) != 3 {
		t.Error("Number of items wrong.")
	}

	if errors.Items[0].Error_id != 400 {
		t.Error("Error id invalid.")
	}

	if errors.Items[0].Error_name != "bad_parameter" {
		t.Error("error name invalid.")
	}

}

func TestSimulateError(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/errors/404", dummyErrorResponse, t)
	defer closeDummyServer(dummy_server)

	error, err := SimulateError(404)

	if err != nil {
		t.Error(err.String())
	}

	if error.Error_name != "no_method" {
		t.Error("Error name invalid.")
	}

	if error.Error_message != "simulated" {
		t.Error("Error message invalid.")
	}

}

//Test Data

var dummyErrorResponse string = `
{
  "error_id": 404,
  "error_name": "no_method",
  "error_message": "simulated"
}
`

var dummyErrorsResponse string = `
{
  "items": [
    {
      "error_id": 400,
      "error_name": "bad_parameter",
      "description": "An malformed parameter was passed"
    },
    {
      "error_id": 401,
      "error_name": "access_token_required",
      "description": "No access_token was passed"
    },
    {
      "error_id": 402,
      "error_name": "invalid_access_token",
      "description": "An access_token that is malformed, expired, or otherwise incorrect was passed"
    }
  ]
}
`
