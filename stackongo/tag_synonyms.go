package stackongo

import "strings"

// AllTagSynonyms returns all tag synonyms in site 
func (session Session) AllTagSynonyms(params map[string]string) (output *TagSynonyms, error error) {
	output = new(TagSynonyms)
	error = session.get("tags/synonyms", params, output)
	return
}

// SynonymsForTags returns all the synonyms that point to the given tags 
func (session Session) SynonymsForTags(tags []string, params map[string]string) (output *TagSynonyms, error error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "synonyms"}, "/")

	output = new(TagSynonyms)
	error = session.get(request_path, params, output)
	return
}
