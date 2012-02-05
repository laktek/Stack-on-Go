package stackongo

import (
	"os"
)

// AllSites returns all sites available in StackExchange network
func AllSites(params map[string]string) (output *Sites, error os.Error) {
	output = new(Sites)
	error = get("sites", params, output)
	return
}
