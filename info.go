package stackongo

import (
	"os"
)

// Info returns a collection of statistics about the site 
func (session Session) Info() (output Info, error os.Error) {
  collection := new(infoCollection)
	error = session.get("info", map[string]string{}, collection)
	return collection.Items[0], error
}
