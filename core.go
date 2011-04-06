package main

import (
	"github.com/garyburd/twister/web"
	"template"
	"os"
)

func coreErrorHandler(req *web.Request, status int, reason os.Error, header web.HeaderMap) {

	coreTempl.Execute(
		req.Responder.Respond(status, header),
		map[string]interface{}{
			"req":     req,
			"status":  status,
			"message": reason,
			"xsrf":    req.Param.Get(web.XSRFParamName),
		})
}

func coreHandler(req *web.Request) {
	coreTempl.Execute(
		req.Respond(web.StatusOK, web.HeaderContentType, "text/html"),
		map[string]interface{}{
			"req":     req,
			"status":  web.StatusOK,
			"message": "ok",
			"xsrf":    req.Param.Get(web.XSRFParamName),
		})
}

var coreTempl = template.MustParse(coreStr, template.FormatterMap{"": template.HTMLFormatter})

const coreStr = `
<html>
<head>
<title>Clarity Services, Inc.</title>
</head>
<body>
<hr>
Status: {status} {message}
Oh No!  We broke something here.  Promise we'll fix it later.
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
</html> `