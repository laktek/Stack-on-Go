package stackongo

import (
	"fmt"
	"strings"
)

// AllBadges returns all badges in site 
func (session Session) AllBadges(params map[string]string) (output *Badges, error error) {
	output = new(Badges)
	error = session.get("badges", params, output)
	return
}

// Badges returns the badges with the given ids
func (session Session) GetBadges(ids []int, params map[string]string) (output *Badges, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"badges", strings.Join(string_ids, ";")}, "/")

	output = new(Badges)
	error = session.get(request_path, params, output)
	return
}

// NamedBadges returns all explicitly named badges 
func (session Session) NamedBadges(params map[string]string) (output *Badges, error error) {
	output = new(Badges)
	error = session.get("badges/name", params, output)
	return
}

// TagBadges returns the badges that are awarded for participation in specific tags 
func (session Session) TagBadges(params map[string]string) (output *Badges, error error) {
	output = new(Badges)
	error = session.get("badges/tags", params, output)
	return
}

// RecentBadgeRecipients returns the recent recipients of the given badges 
func (session Session) RecentBadgeRecipients(ids []int, params map[string]string) (output *Badges, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"badges", strings.Join(string_ids, ";"), "recipients"}, "/")

	output = new(Badges)
	error = session.get(request_path, params, output)
	return

}

// RecentAllBadgeRecipients returns badges recently awarded
func (session Session) RecentAllBadgeRecipients(params map[string]string) (output *Badges, error error) {
	output = new(Badges)
	error = session.get("badges/recipients", params, output)
	return
}

// BadgesOfUsers returns the badges earned by the users identified by a set of ids 
func (session Session) BadgesOfUsers(ids []int, params map[string]string) (output *Badges, error error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "badges"}, "/")

	output = new(Badges)
	error = session.get(request_path, params, output)
	return
}
