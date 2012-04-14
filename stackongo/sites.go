package stackongo

// AllSites returns all sites available in StackExchange network
func AllSites(params map[string]string) (output *Sites, error error) {
	output = new(Sites)
	error = get("sites", params, output)
	return
}
