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
	question_timelines.go\
	reputations.go\
	resource_types.go\
	revisions.go\
	suggested_edits.go\
	tags.go\
	tag_synonyms.go\
	tag_wikis.go\
	users.go\
	user_timelines.go\

include $(GOROOT)/src/Make.pkg

