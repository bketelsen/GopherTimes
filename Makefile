include $(GOROOT)/src/Make.inc

TARG=clarityservicesdotcom
GOFILES=\
	main.go \
	core.go \
	schema.go \

include $(GOROOT)/src/Make.cmd
