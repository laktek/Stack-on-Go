package stackongo

type AccessTokens struct {
	Items []AccessToken
	Error
	MetaInfo
}

type Answers struct {
	Items []Answer
	Error
	MetaInfo
}

type Badges struct {
	Items []Badge
	Error
	MetaInfo
}

type Comments struct {
	Items []Comment
	Error
	MetaInfo
}

type Errors struct {
	Items []Error
	Error
	MetaInfo
}

type Events struct {
	Items []Event
	Error
	MetaInfo
}

type infoCollection struct {
	Items []Info
	Error
	MetaInfo
}

type InboxItems struct {
	Items []InboxItem
	Error
	MetaInfo
}

type Filters struct {
	Items []Filter
	Error
	MetaInfo
}

type NetworkUsers struct {
	Items []NetworkUser
	Error
	MetaInfo
}

type Posts struct {
	Items []Post
	Error
	MetaInfo
}

type Privileges struct {
	Items []Privilege
	Error
	MetaInfo
}

type Questions struct {
	Items []Question
	Error
	MetaInfo
}

type QuestionTimelines struct {
	Items []QuestionTimeline
	Error
	MetaInfo
}

type Reputations struct {
	Items []Reputation
	Error
	MetaInfo
}

type Revisions struct {
	Items []Revision
	Error
	MetaInfo
}

type Sites struct {
	Items []Site
	Error
	MetaInfo
}

type SuggestedEdits struct {
	Items []SuggestedEdit
	Error
	MetaInfo
}

type Tags struct {
	Items []Tag
	Error
	MetaInfo
}

type TagScores struct {
	Items []TagScore
	Error
	MetaInfo
}

type TagSynonyms struct {
	Items []TagSynonym
	Error
	MetaInfo
}

type TagWikis struct {
	Items []TagWiki
	Error
	MetaInfo
}

type TopTags struct {
	Items []TopTag
	Error
	MetaInfo
}

type Users struct {
	Items []User
	Error
	MetaInfo
}

type UserTimelines struct {
	Items []UserTimeline
	Error
	MetaInfo
}
