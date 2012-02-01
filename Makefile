include $(GOROOT)/src/Make.inc

TARG=stackongo
GOFILES=\
	session.go\
	answers.go\
	badges.go\
	comments.go\
	collections.go\
	info.go\
	posts.go\
	privileges.go\
	questions.go\
	reputations.go\
	resource_types.go\
	revisions.go\
	suggested_edits.go\
	users.go\
	user_timelines.go\

include $(GOROOT)/src/Make.pkg

