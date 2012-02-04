include $(GOROOT)/src/Make.inc

TARG=stackongo
GOFILES=\
	session.go\
	access_tokens.go\
	answers.go\
	auth.go\
	badges.go\
	comments.go\
	collections.go\
	errors.go\
	events.go\
	filters.go\
	info.go\
	inbox.go\
	network_users.go\
	posts.go\
	privileges.go\
	questions.go\
	question_timelines.go\
	reputations.go\
	resource_types.go\
	revisions.go\
	sites.go\
	suggested_edits.go\
	tags.go\
	tag_scores.go\
	tag_synonyms.go\
	tag_wikis.go\
	top_tags.go\
	users.go\
	user_timelines.go\

include $(GOROOT)/src/Make.pkg

