package main

import (
	"github.com/garyburd/twister/web"
	"github.com/garyburd/twister/server"
	"template"
	"os"
	
)

func loadPage(path string)(content string, err os.Error){
	//will return a mongodb document
	content = "Hello " + path
	return
}

func viewHandler(req *web.Request) {
	path := req.Param.Get("path")
	p, err := loadPage(path)
	if err != nil {
		// do something!
	}
	println(p)
	
	
}

func main() {
	// this is a bad regex, will only handle single word, not full path
	const pathParam = "<path:.*>"
	h := web.ProcessForm(10000, true, // limit size of form to 10k, enable xsrf
		web.NewRouter().
			Register("/"+pathParam, "GET", viewHandler).
			Register("/static/<path:.*>", "GET", web.DirectoryHandler("static/")))
	server.Run(":8080", h)
}

var homeTempl = template.MustParse(homeStr, template.FormatterMap{"": template.HTMLFormatter})

const homeStr = `
<html>
<head>
</head>
<body>
{.section req}
<table>
<tr><th align="left" valign="top">RemoteAddr</th><td>{RemoteAddr}</td></tr>
<tr><th align="left" valign="top">Method</th><td>{Method}</td></tr>
<tr><th align="left" valign="top">URL</th><td>{URL}</td></tr>
<tr><th align="left" valign="top">ProtocolVersion</th><td>{ProtocolVersion}</td></tr>
<tr><th align="left" valign="top">Param</th><td>{Param}</td></tr>
<tr><th align="left" valign="top">ContentType</th><td>{ContentType}</td></tr>
<tr><th align="left" valign="top">ContentLength</th><td>{ContentLength}</td></tr>
<tr><th align="left" valign="top">Header</th><td>{Header}</td></tr>
</table>
{.end}
</body>
</html>`