include $(GOROOT)/src/Make.inc

TARG=stackongo
GOFILES=\
	session.go\
	resource_types.go\
	collections.go\
	questions.go\
	comments.go\

include $(GOROOT)/src/Make.pkg

