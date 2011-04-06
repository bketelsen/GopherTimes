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
	//will load a mongodb document
	//deserialized into a page struct
	log.Println(path)
	
	mongo, err := mgo.Mongo("127.0.0.1")
	if err != nil {
		panic(err)
	}

	c := mongo.DB("public_web").C("page")

	log.Println(c)
	page := &Page{}
	
    err = c.Find(bson.M{"path": path}).One(page)
	log.Println(page)
	log.Println(err)
	
	return page, err
}

func viewHandler(req *web.Request) {
	path := req.Param.Get("path")
	p, err := loadPage(path)
	if err != nil {
		// respond with a simple error message for now
		w := req.Respond(web.StatusNotFound)
		println(w, "404 Not Found")
		return
	}

	renderTemplate(req, "public_base", p)

}

var templates = make(map[string]*template.Template)
var mongo *mgo.Session

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

	mongo, err := mgo.Mongo("127.0.0.1")
	if err != nil {
		panic(err)
	}

	c := mongo.DB("public_web").C("page")
	
	log.Println(c)
	err = c.Insert(&Page{Path: "about/us", Title: "Clarity Services Home Page", Content: "About Clarity Services Here"})
	if err != nil {
		log.Println(err)
	}


	// this is a bad regex, will only handle single word, not full path
	const pathParam = "<path:.*>"

	h := web.ProcessForm(10000, true, // limit size of form to 10k, enable xsrf
		web.NewRouter().
			Register("/static/<path:.*>", "GET", web.DirectoryHandler("static/")).
			Register("/favicon.ico", "GET", web.FileHandler("static/favicon.ico")).
			Register("/<path:(.*)>", "GET", viewHandler))
	server.Run(":8080", h)

	defer mongo.Close()
}
