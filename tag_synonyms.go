package stackongo

import (
	"os"
	"strings"
)

func (session Session) getTagSynonyms(path string, params map[string]string) (output *TagSynonyms, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(TagSynonyms))
	output = parsed_response.(*TagSynonyms)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// AllTagSynonyms returns all tag synonyms in site 
func (session Session) AllTagSynonyms(params map[string]string) (output *TagSynonyms, error os.Error) {
	return session.getTagSynonyms("tags/synonyms", params)
}

// SynonymsForTags returns all the synonyms that point to the given tags 
func (session Session) SynonymsForTags(tags []string, params map[string]string) (output *TagSynonyms, error os.Error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "synonyms"}, "/")
	return session.getTagSynonyms(request_path, params)
}
