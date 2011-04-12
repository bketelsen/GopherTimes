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
<title>Gopher Times</title>
</head>
<body>
<hr>
Status: {status} {message}
We're sorry, but there was a problem fulfilling your request.
</body>
</html> `