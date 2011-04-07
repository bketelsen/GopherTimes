package main

import (
	"github.com/garyburd/twister/web"
	"github.com/garyburd/twister/server"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
	"template"
	"os"
	"log"
	"time"
)

func loadPage(path string) (*Page, os.Error) {
	var page *Page

	cachedPage, found := cachedPages[path]
	if found {
		t := time.Seconds()
		log.Println("Cached at", cachedPage.CachedAt)
		log.Println("Now it is", t)
		if (t - cachedPage.CachedAt) < (60 * 10) { // cache for 10 minutes
			log.Println("Returning page from cache")
			return cachedPage.CachedPage, nil
		}
	}
	mongo, err := mgo.Mongo("localhost")
	if err != nil {
		return nil, err
	}
	defer mongo.Close()

	c := mongo.DB("public_web").C("page")
	page = &Page{}

	err = c.Find(bson.M{"path": path}).One(page)
	log.Println("Retrieving Page from Mongo")
	go cachePage(page)
	return page, err
}

func viewHandler(req *web.Request) {
	path := req.Param.Get("path")
	p, err := loadPage(path)
	if err != nil {
		renderTemplate(req, web.StatusNotFound, "404", p)
	} else {
		var templateName string
		if len(p.Template) != 0 {
			templateName = p.Template
		} else {
			templateName = "public_base"
		}
		renderTemplate(req, web.StatusOK, templateName, p)
	}
}

var cachedPages = make(map[string]*CachedPage)
var templates = make(map[string]*template.Template)
//var mongo *mgo.Session

func init() {
	for _, tmpl := range []string{"index", "404"} {
		templates[tmpl] = template.MustParseFile("templates/"+tmpl+".html", nil)
	}

}

func cachePage(p *Page) {
	cachedPages[p.Path] = &CachedPage{CachedPage: p, CachedAt: time.Seconds()}
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



	h := web.ProcessForm(10000, true, // limit size of form to 10k, enable xsrf
		web.NewRouter().
			Register("/static/<path:.*>", "GET", web.DirectoryHandler("static/")).
			Register("/favicon.ico", "GET", web.FileHandler("static/favicon.ico")).
			Register("/<path:(.*)>", "GET", viewHandler))
	server.Run(":8081", h)

}
