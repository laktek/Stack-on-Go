package stackongo

type Answer struct {
	Question_id          int
	Answer_id            int
	Locked_date          int64
	Creation_date        int64
	Last_edit_date       int64
	Last_activity_date   int64
	Score                int
	Community_owned_date int64
	Is_accepted          bool
	Body                 string
	Owner                Shallow_user
	Title                string
	Up_vote_count        int
	Down_vote_count      int
	Comments             []Comment
	Link                 string
}

type Badge struct {
	Badge_id    int
	Rank        string
	Name        string
	Description string
	Award_count int
	Badge_type  string
	User        Shallow_user
	Link        string
}

type Badge_count struct {
	Gold   int
	Silver int
	Bronze int
}

type Comment struct {
	Comment_id    int
	Post_id       int
	Creation_date int64
	Post_type     string //one of question, or answer
	Score         int
	Edited        bool
	Body          string
	Owner         Shallow_user
	Reply_to_user Shallow_user
	Link          string
}

type Error struct {
	Error_id      int
	Error_name    string
	Error_message string
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

type Migration_info struct {
	Question_id int
	Other_site  Site
	On_date     int64
}

type Post struct {
	Post_id            int
	Post_type          string
	Body               string
	Owner              Shallow_user
	Creation_date      int64
	Last_activity_date int64
	Last_edit_date     int64
	Score              int
	Up_vote_count      int
	Down_vote_count    int
	Comments           []Comment
}

type Privilege struct {
	Short_description string
	Description       string
	Reputation        int
}

type Question struct {
	Question_id          int
	Last_edit_date       int64
	Creation_date        int64
	Last_activity_date   int64
	Locked_date          int64
	Community_owned_date int64
	Score                int
	Answer_count         int
	Accepted_answer_id   int
	Migrated_to          Migration_info
	Migrated_from        Migration_info
	Bounty_closes_date   int64
	Bounty_amount        int
	Closed_date          int64
	Protected_date       int64
	Body                 string
	Title                string
	Tags                 []string
	Closed_reason        string
	Up_vote_count        int
	Down_vote_count      int
	Favorite_count       int
	View_count           int
	Owner                Shallow_user
	Comments             []Comment
	Answers              []Answer
	Link                 string
	Is_answered          bool
}

type Reputation struct {
	User_id           int
	Post_id           int
	Post_type         string //one of question, or answer
	Vote_type         string //one of accepts, up_votes, down_votes, bounties_offered, bounties_won, spam, or suggested_edits
	Title             string
	Link              string
	Reputation_change int
	On_date           int64
}

type Revision struct {
	Revision_guid      string
	Revision_number    int
	Revision_type      string //one of single_user, or vote_based
	Post_type          string //one of question, or answer
	Post_id            int
	Comment            string
	Creation_date      int64
	Is_rollback        bool
	Last_body          string
	Last_title         string
	Last_tags          []string
	Body               string
	Title              string
	Tags               []string
	Set_community_wiki bool
	User               Shallow_user
}

type Shallow_user struct {
	User_id       int
	Display_name  string
	Reputation    int
	User_type     string //one of unregistered, registered, moderator, or does_not_exist
	Profile_image string
	Link          string
}

type Site struct {
	Site_type          string
	Name               string
	Logo_url           string
	Api_site_parameter string
	Site_url           string
	Audience           string
	Icon_url           string
	Aliases            []string
	Site_state         string //one of normal, closed_beta, open_beta, or linked_meta
	//styling styling
	Closed_beta_date int64
	Open_beta_date   int64
	Launch_date      int64
	Favicon_url      string
	//related_sites an array of related_site
	Twitter_account     string
	Markdown_extensions []string
}

type SuggestedEdit struct {
	Suggested_edit_id int
	Post_id           int
	Post_type         string //one of question, or answer
	Body              string
	Title             string
	Tags              []string
	Comment           string
	Creation_date     int64
	Approval_date     int64
	Rejection_date    int64
	Proposing_user    Shallow_user
}

type User struct {
	User_id                   int
	User_type                 string //one of unregistered, registered, moderator, or does_not_exist
	Creation_date             int64
	Display_name              string
	Profile_image             string
	Reputation                int
	Reputation_change_day     int
	Reputation_change_week    int
	Reputation_change_month   int
	Reputation_change_quarter int
	Reputation_change_year    int
	Age                       int
	Last_access_date          int64
	Last_modified_date        int64
	Is_employee               bool
	Link                      string
	Website_url               string
	Location                  string
	Account_id                int
	Timed_penalty_date        int64
	Badge_counts              Badge_count
	Question_count            int
	Answer_count              int
	Up_vote_count             int
	Down_vote_count           int
	About_me                  string
	View_count                int
	Accept_rate               int
}

type UserTimeline struct {
Creation_date int64
Post_type string //one of question, or answer
Timeline_type string //one of commented, asked, answered, badge, revision, accepted, reviewed, or suggested
User_id int
Post_id int
Comment_id int
Suggested_edit_id int
Badge_id int
Title string
Detail string
Link string
}
