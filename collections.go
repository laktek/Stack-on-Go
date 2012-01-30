package stackongo

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

type revisionsCollection struct {
	Items []Revision
	Error
}
