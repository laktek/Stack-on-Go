package gostack

import (
	"http"
	"json"
	"io/ioutil"
	"url"
	"os"
)

const host = "http://api.stackexchange.com"

type Site struct {
	Param string
}

type Info struct {
	Total_questions      int
	Total_unanswered     int
	Total_accepted       int
	Total_answers        int
	Questions_per_minute float32
	Answers_per_minute   float32
	Total_comments       int
	Total_votes          int
	Total_badges         int
	Badges_per_minute    float32
	Total_users          int
	New_active_users     int
	Api_revision         string
}

type Privilege struct {
	Short_description string
	Description       string
	Reputation        int
}

type Question struct {
Question_id int
Last_edit_date date
Creation_date date
✔ last_activity_date
date
✔ locked_date
date
✔ score
integer
✔ community_owned_date
date
✔ answer_count
integer
✔ accepted_answer_id
integer
✔ migrated_to
migration_info
✔ migrated_from
migration_info
✔ bounty_closes_date
date
✔ bounty_amount
integer
✔ closed_date
date
✔ protected_date
date
✘ body
string
unchanged in unsafe filters
✔ title
string
✔ tags
an array of string
✔ closed_reason
string
✘ up_vote_count
integer
✘ down_vote_count
integer
✘ favorite_count
integer
✔ view_count
integer
✔ owner
shallow_user
✘ comments
an array of comment
✘ answers
an array of answer
✔ link
string
unchanged in unsafe filters
✔ is_answered
boolean

}

type Error struct {
	Error_id      int
	Error_name    string
	Error_message string
}

type infoCollection struct {
	Items []Info
	Error
}

type privilegesCollection struct {
	Items []Privilege
	Error
}

type questionsCollection struct {
	Items []Question
	Error
}

func NewSite(param string) *Site {
	return &Site{Param: param}
}

// construct the endpoint URL
func (site Site) setupEndpoint(path string) *url.URL {
	base_url, _ := url.Parse(host)
	endpoint, _ := base_url.Parse("/2.0/" + path)

	query := endpoint.Query()
	query.Set("site", site.Param)
	endpoint.RawQuery = query.Encode()

	return endpoint
}

// make the request
func (site Site) get(section string) (*http.Response, os.Error) {
	client := new(http.Client)
	return client.Get(site.setupEndpoint(section).String())
}

// parse the response
func parseResponse(response *http.Response, result interface{}) (interface{}, os.Error) {

	//read the response
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return result, err
	}

	//parse JSON
	err2 := json.Unmarshal(bytes, result)

	//check whether the response is a bad request
	if response.StatusCode == 400 {
		err2 = os.NewError("Bad Request")
	}

	return result, err2
}

func (site Site) GetInfo() (output Info, error os.Error) {
	// make the request
	response, err := site.get("info")

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(infoCollection))
	collection := parsed_response.(*infoCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items[0]
	}

	return output, error
}

func (site Site) GetPrivileges() (output []Privilege, error os.Error) {
	// make the request
	response, err := site.get("privileges")

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(privilegesCollection))
	collection := parsed_response.(*privilegesCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error
}

func (site Site) GetAllQuestions() (output []Question, error os.error) {
	// make the request
	response, err := site.get("questions")

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(questionsCollection))
	collection := parsed_response.(*questionsCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error
}

func (site Site) GetQuestions() (output []Question, error os.error) {

}
