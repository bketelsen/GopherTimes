include $(GOROOT)/src/Make.inc

TARG=gophertimes
GOFILES=\
	main.go \
	core.go \
	schema.go \

include $(GOROOT)/src/Make.cmd


format:
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w main.go
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w core.go
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w schema.go