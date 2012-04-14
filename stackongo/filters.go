package stackongo

import "strings"

// CreateFilter creates a new filter given a list of includes, excludes, a base filter, and whether or not this filter should be "unsafe". 
func CreateFilter(params map[string]string) (output *Filters, error error) {
	output = new(Filters)
	error = get("filters/create", params, output)
	return
}

// InspectFilters returns the fields included by the given filters, and the "safeness" of those filters.
func InspectFilters(filters []string, params map[string]string) (output *Filters, error error) {
	request_path := strings.Join([]string{"filters", strings.Join(filters, ";")}, "/")

	output = new(Filters)
	error = get(request_path, params, output)
	return

}
