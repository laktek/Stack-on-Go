package stackongo

//All collection types are private

type answersCollection struct {
	Items []Answer
	Error
}

type badgesCollection struct {
	Items []Badge
	Error
}

type commentsCollection struct {
	Items []Comment
	Error
}

type infoCollection struct {
	Items []Info
	Error
}

type postsCollection struct {
	Items []Post
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

type questionTimelinesCollection struct {
	Items []QuestionTimeline
	Error
}

type reputationsCollection struct {
	Items []Reputation
	Error
}

type revisionsCollection struct {
	Items []Revision
	Error
}

type suggestedEditsCollection struct {
	Items []SuggestedEdit
	Error
}

type tagsCollection struct {
	Items []Tag
	Error
}

type tagSynonymsCollection struct {
	Items []TagSynonym
	Error
}

type tagWikisCollection struct {
	Items []TagWiki
	Error
}

type usersCollection struct {
	Items []User
	Error
}

type userTimelinesCollection struct {
	Items []UserTimeline
	Error
}
