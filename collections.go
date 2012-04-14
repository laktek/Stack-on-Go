package stackongo

type AccessTokens struct {
	Items         []AccessToken
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Answers struct {
	Items         []Answer
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Badges struct {
	Items         []Badge
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Comments struct {
	Items         []Comment
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Errors struct {
	Items         []Error
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Events struct {
	Items         []Event
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type infoCollection struct {
	Items         []Info
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type InboxItems struct {
	Items         []InboxItem
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Filters struct {
	Items         []Filter
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type NetworkUsers struct {
	Items         []NetworkUser
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Posts struct {
	Items []Post
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Privileges struct {
	Items []Privilege
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Questions struct {
	Items []Question
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type QuestionTimelines struct {
	Items []QuestionTimeline
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Reputations struct {
	Items []Reputation
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Revisions struct {
	Items []Revision
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Sites struct {
	Items []Site
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type SuggestedEdits struct {
	Items []SuggestedEdit
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Tags struct {
	Items []Tag
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type TagScores struct {
	Items []TagScore
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type TagSynonyms struct {
	Items []TagSynonym
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type TagWikis struct {
	Items []TagWiki
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type TopTags struct {
	Items []TopTag
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Users struct {
	Items []User
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type UserTimelines struct {
	Items []UserTimeline
	Error_id      int
	Error_name    string
	Error_message string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}
