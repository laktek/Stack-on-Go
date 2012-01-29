package stackongo

type questionsCollection struct {
	Items []Question
	Error
}

type commentsCollection struct {
	Items []Comment
	Error
}

type answersCollection struct {
	Items []Answer
	Error
}

type postsCollection struct {
	Items []Post
	Error
}

type infoCollection struct {
	Items []Info
	Error
}

type privilegesCollection struct {
	Items []Privilege
	Error
}
