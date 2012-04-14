package stackongo

import "errors"

// Info returns a collection of statistics about the site 
func (session Session) Info() (output Info, error error) {
	collection := new(infoCollection)
	error = session.get("info", map[string]string{}, collection)

	if len(collection.Items) > 0 {
		output = collection.Items[0]
	} else {
		error = errors.New("Site not found")
	}

	return
}
