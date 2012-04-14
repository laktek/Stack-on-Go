package stackongo

type AccessToken struct {
	Access_token    string
	Expires_on_date int64
	Account_id      int
	Scope           []string
}

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
	Owner                ShallowUser
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
	User        ShallowUser
	Link        string
}

type BadgeCount struct {
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
	Owner         ShallowUser
	Reply_to_user ShallowUser
	Link          string
}

type Error struct {
	Error_id      int
	Error_name    string
	Error_message string
}

type Event struct {
	Event_type    string //one of question_posted, answer_posted, comment_posted, post_edited, or user_created
	Event_id      int    //refers to an event
	Creation_date int64
	Link          string
	Excerpt       string
}

type Filter struct {
	Filter          string
	Included_fields []string
	Filter_type     string //one of safe, unsafe, or invalid
}

type InboxItem struct {
	Item_type     string //one of comment, chat_message, new_answer, careers_message, careers_invitations, or meta_question
	Question_id   int    // refers to a question
	Answer_id     int    // refers to an answer
	Comment_id    int    //refers to a comment
	Title         string
	Creation_date int64
	Is_unread     bool
	Site          Site
	Body          string
	Link          string
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

type NetworkUser struct {
	Site_name        string
	Site_url         string
	User_id          int
	Reputation       int
	Account_id       int
	Creation_date    int64
	User_type        string //one of unregistered, registered, moderator, or does_not_exist
	Badge_counts     BadgeCount
	Last_access_date int64
	Answer_count     int
	Question_count   int
}

type MetaInfo struct {
  Backoff int 
  Has_more bool
  Page int
  Page_size int 
  Quota_max int 
  Quota_remaining int 
  Total int
  Type string
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
	Owner              ShallowUser
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
	Owner                ShallowUser
	Comments             []Comment
	Answers              []Answer
	Link                 string
	Is_answered          bool
}

type QuestionTimeline struct {
	Timeline_type   string //one of question, answer, comment, unaccepted_answer, accepted_answer, vote_aggregate, revision, or post_state_changed
	Question_id     int
	Post_id         int
	Comment_id      int
	Revision_guid   string
	Up_vote_count   int
	Down_vote_count int
	Creation_date   int64
	User            ShallowUser
	Owner           ShallowUser
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
	User               ShallowUser
}

type RelatedSite struct {
	Name               string
	Site_url           string
	Relation           string //one of parent, meta, chat, or other
	Api_site_parameter string
}

type ShallowUser struct {
	User_id       int
	Display_name  string
	Reputation    int
	User_type     string //one of unregistered, registered, moderator, or does_not_exist
	Profile_image string
	Link          string
}

type Site struct {
	Site_type           string
	Name                string
	Logo_url            string
	Api_site_parameter  string
	Site_url            string
	Audience            string
	Icon_url            string
	Aliases             []string
	Site_state          string //one of normal, closed_beta, open_beta, or linked_meta
	Styling             Styling
	Closed_beta_date    int64
	Open_beta_date      int64
	Launch_date         int64
	Favicon_url         string
	Related_sites       []RelatedSite
	Twitter_account     string
	Markdown_extensions []string
}

type Styling struct {
	Link_color           string
	Tag_foreground_color string
	Tag_background_color string
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
	Proposing_user    ShallowUser
}

type Tag struct {
	Name               string
	Count              int
	Is_required        bool
	Is_moderator_only  bool
	User_id            int
	Has_synonyms       bool
	Last_activity_date int64
}

type TagScore struct {
	User       ShallowUser
	Score      int
	Post_count int
}

type TagSynonym struct {
	From_tag          string
	To_tag            string
	Applied_count     int
	Last_applied_date int64
	Creation_date     int64
}

type TagWiki struct {
	Tag_name               string
	Body                   string
	Excerpt                string
	Body_last_edit_date    int64
	Excerpt_last_edit_date int64
	Last_body_editor       ShallowUser
	Last_excerpt_editor    ShallowUser
}

type TopTag struct {
	Tag_name       string
	Question_score int
	Question_count int
	Answer_score   int
	Answer_count   int
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
	Badge_counts              BadgeCount
	Question_count            int
	Answer_count              int
	Up_vote_count             int
	Down_vote_count           int
	About_me                  string
	View_count                int
	Accept_rate               int
}

type UserTimeline struct {
	Creation_date     int64
	Post_type         string //one of question, or answer
	Timeline_type     string //one of commented, asked, answered, badge, revision, accepted, reviewed, or suggested
	User_id           int
	Post_id           int
	Comment_id        int
	Suggested_edit_id int
	Badge_id          int
	Title             string
	Detail            string
	Link              string
}
