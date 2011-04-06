package main

import (
	"github.com/garyburd/twister/web"
	"github.com/garyburd/twister/server"
	"template"
	"os"
	"log"
	
)

func loadPage(path string)(page *Page, err os.Error){
	//will load a mongodb document
	//deserialized into a page struct
	log.Println(path)
	return &Page{Path:path,Title:"Clarity Services Home Page", Content:"Your body goes here"}, nil
}

func viewHandler(req *web.Request) {
	path := req.Param.Get("path")
	p, err := loadPage(path)
	if err != nil {
		// respond with a simple error message for now
		w := req.Respond(web.StatusInternalServerError)
		println(w, "500 Internal Server Error")
	}

	renderTemplate(req, "public_base", p)
	
}

var templates = make(map[string]*template.Template)

func init() {
	for _, tmpl := range []string{"public_base"} {
		templates[tmpl] = template.MustParseFile("templates/"+tmpl+".html", nil)
	}
}

func renderTemplate(req *web.Request, tmpl string, p *Page) {
	err := templates[tmpl].Execute(
		req.Respond(web.StatusOK),
		map[string]interface{}{
			"page": p,
			"xsrf": req.Param.Get("xsrf"),
		})
	if err != nil {
		log.Println("error rendering", tmpl, err)
	}
}

func main() {

	// this is a bad regex, will only handle single word, not full path
	const pathParam = "<path:.*>"

	h := web.ProcessForm(10000, true, // limit size of form to 10k, enable xsrf
		web.NewRouter().
			Register("/static/<path:.*>", "GET", web.DirectoryHandler("static/")).
			Register("/<path:(.*)>", "GET", viewHandler))
	server.Run(":8080", h)
}

