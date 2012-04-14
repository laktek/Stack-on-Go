package stackongo

// Events returns a stream of events that have occurred on the site 
// This method requires an access_token.
func (session Session) Events(params map[string]string, auth map[string]string) (output *Events, error error) {
	//add auth params
	for key, value := range auth {
		params[key] = value
	}

	output = new(Events)
	error = session.get("events", params, output)
	return
}
