include $(GOROOT)/src/Make.inc

TARG=gophertimes
GOFILES=\
	main.go \
	core.go \
	schema.go \

include $(GOROOT)/src/Make.cmd
