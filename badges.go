package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getBadges(path string, params map[string]string) (output *Badges, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(Badges))
	output = parsed_response.(*Badges)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// AllBadges returns all badges in site 
func (session Session) AllBadges(params map[string]string) (output *Badges, error os.Error) {
	return session.getBadges("badges", params)
}

// Badges returns the badges with the given ids
func (session Session) GetBadges(ids []int, params map[string]string) (output *Badges, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"badges", strings.Join(string_ids, ";")}, "/")
	return session.getBadges(request_path, params)
}

// NamedBadges returns all explicitly named badges 
func (session Session) NamedBadges(params map[string]string) (output *Badges, error os.Error) {
	return session.getBadges("badges/name", params)
}

// TagBadges returns the badges that are awarded for participation in specific tags 
func (session Session) TagBadges(params map[string]string) (output *Badges, error os.Error) {
	return session.getBadges("badges/tags", params)
}

// RecentBadgeRecipients returns the recent recipients of the given badges 
func (session Session) RecentBadgeRecipients(ids []int, params map[string]string) (output *Badges, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"badges", strings.Join(string_ids, ";"), "recipients"}, "/")
	return session.getBadges(request_path, params)
}

// RecentAllBadgeRecipients returns badges recently awarded
func (session Session) RecentAllBadgeRecipients(params map[string]string) (output *Badges, error os.Error) {
	return session.getBadges("badges/recipients", params)
}

// BadgesOfUsers returns the badges earned by the users identified by a set of ids 
func (session Session) BadgesOfUsers(ids []int, params map[string]string) (output *Badges, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "badges"}, "/")
	return session.getBadges(request_path, params)
}
