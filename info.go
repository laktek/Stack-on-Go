package stackongo

import (
	"os"
)

// Info returns a collection of statistics about the site 
func (session Session) Info() (output Info, error os.Error) {
  collection := new(infoCollection)
	error = session.get("info", map[string]string{}, collection)

  if len(collection.Items) > 0 {
		output = collection.Items[0]
	} else {
		error = os.NewError("Site not found")
	}

	return
}
