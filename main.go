package main

import (
	"github.com/garyburd/twister/web"
	"github.com/garyburd/twister/server"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
	"template"
	"os"
	"log"
)

func loadPage(path string) (*Page, os.Error) {

	mongo, err := mgo.Mongo("msg2")
	if err != nil {
		return nil, err
	}

	c := mongo.DB("public_web").C("page")
	page := &Page{}

	err = c.Find(bson.M{"path": path}).One(page)

	return page, err
}

func viewHandler(req *web.Request) {
	path := req.Param.Get("path")
	p, err := loadPage(path)
	if err != nil {
		renderTemplate(req, web.StatusNotFound, "404", p)
	} else {
		renderTemplate(req, web.StatusOK, "public_base", p)
	}
}

var templates = make(map[string]*template.Template)
var mongo *mgo.Session

func init() {
	for _, tmpl := range []string{"public_base", "404"} {
		templates[tmpl] = template.MustParseFile("templates/"+tmpl+".html", nil)
	}

}

func renderTemplate(req *web.Request, status int, tmpl string, p *Page) {
	err := templates[tmpl].Execute(
		req.Respond(status),
		map[string]interface{}{
			"page": p,
			"xsrf": req.Param.Get("xsrf"),
		})
	if err != nil {
		log.Println("error rendering", tmpl, err)
	}
}

func main() {

	mongo, err := mgo.Mongo("msg2")
	if err != nil {
		panic(err)
	}

	c := mongo.DB("public_web").C("page")

	// this is a bad regex, will only handle single word, not full path
	const pathParam = "<path:.*>"

	h := web.ProcessForm(10000, true, // limit size of form to 10k, enable xsrf
		web.NewRouter().
			Register("/static/<path:.*>", "GET", web.DirectoryHandler("static/")).
			Register("/favicon.ico", "GET", web.FileHandler("static/favicon.ico")).
			Register("/<path:(.*)>", "GET", viewHandler))
	server.Run(":8081", h)

	defer mongo.Close()
}
