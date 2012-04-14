package stackongo

import "strings"

// TopAnswerers returns the top 30 answerers active in a single tag, for the given period 
// period should be either `all_time` or `month` 
func (session Session) TopAnswerers(tag string, period string, params map[string]string) (output *TagScores, error error) {
	request_path := strings.Join([]string{"tags", tag, "top-answerers", period}, "/")

	output = new(TagScores)
	error = session.get(request_path, params, output)
	return
}

// TopAskers returns the top 30 askers active in a single tag, for the given period 
// period should be either `all_time` or `month` 
func (session Session) TopAskers(tag string, period string, params map[string]string) (output *TagScores, error error) {
	request_path := strings.Join([]string{"tags", tag, "top-askers", period}, "/")

	output = new(TagScores)
	error = session.get(request_path, params, output)
	return
}
