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
	resource_types.go\

include $(GOROOT)/src/Make.pkg

