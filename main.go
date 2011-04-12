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
	"flag"
	"fmt"
)

//flag parse
var port *int = flag.Int("port", 8081, "http port for server")
var initdb *bool = flag.Bool("initdb", false, "create initial record in mongodb")
var database *string = flag.String("database","public_web","mongo database name")

func loadNewsItem(path string) (*NewsItem, os.Error) {
	var item *NewsItem

	cachedItem, found := cachedNewsItems[path]
	if found {
		t := time.Seconds()
		log.Println("Cached at", cachedItem.CachedAt)
		log.Println("Now it is", t)
		if (t - cachedItem.CachedAt) < (60 * 10) { // cache for 10 minutes
			log.Println("Returning page from cache")
			return cachedItem.NewsItem, nil
		}
	}
	mongo, err := mgo.Mongo("localhost")
	if err != nil {
		return nil, err
	}
	defer mongo.Close()

	c := mongo.DB(*database).C("newsitems")
	item = &NewsItem{}

	err = c.Find(bson.M{"page.permalink": path}).One(item)
	log.Println("Retrieving Page from Mongo")
	go cacheNewsItem(item)
	return item, err
}


func loadNewsItems() ([]*NewsItem, os.Error) {
	var item *NewsItem

	mongo, err := mgo.Mongo("localhost")
	if err != nil {
		return nil, err
	}
	defer mongo.Close()
	
	items := make([]*NewsItem, 25)
	
	c := mongo.DB(*database).C("newsitems")
	item = &NewsItem{}

	iter, err := c.Find(bson.M{}).Iter()
	i := 0

	for iter.Next(&item) != mgo.NotFound {
		items[i] = item
		i++
	}

	return items, err
}

func viewHandler(req *web.Request) {
	path := req.Param.Get("path")
	log.Println("Path:", path)
	parm := req.Param.Get("invalidate")
	if parm == "true" {
		removeCachedNewsItem(path)
	}
	log.Println("Parm:", parm)
	p, err := loadNewsItem(path)
	list, err := loadNewsItems()
	if err != nil {
		renderTemplate(req, web.StatusNotFound, "404", p, list)
	} else {
		var templateName string
		if len(p.Template) != 0 {
			templateName = p.Template
		} else {
			templateName = "index"
		}
		renderTemplate(req, web.StatusOK, templateName, p, list)
	}
}
func homeHandler(req *web.Request) {

	p, err := loadNewsItems()

	if err != nil {
		renderTemplate(req, web.StatusNotFound, "404", p[0], p)
	} else {
		renderTemplate(req, web.StatusOK, "index", p[0], p)
	}
}


var cachedNewsItems = make(map[string]*CachedNewsItem)
var templates = make(map[string]*template.Template)
//var mongo *mgo.Session

func init() {
	for _, tmpl := range []string{"index", "404"} {
		templates[tmpl] = template.MustParseFile("templates/"+tmpl+".html", nil)
	}

}

func cacheNewsItem(n *NewsItem) {
	cachedNewsItems[n.Page.Permalink] = &CachedNewsItem{NewsItem: n, CachedAt: time.Seconds()}
}

func removeCachedNewsItem(permalink string) {
	cachedNewsItems[permalink] = nil, false
}

func renderTemplate(req *web.Request, status int, tmpl string, n *NewsItem, items []*NewsItem) {
	err := templates[tmpl].Execute(
		req.Respond(status),
		map[string]interface{}{
			"newsItem": n,
			"newsItems": items,
			"xsrf":     req.Param.Get("xsrf"),
		})
	if err != nil {
		log.Println("error rendering", tmpl, err)
	}
}

func loadFirstRecord(){
	//open mongo
	mongo, err := mgo.Mongo("127.0.0.1")
	if err != nil {
	panic(err)
	}

	c := mongo.DB(*database).C("newsitems")

	err = c.Insert(&NewsItem{Page:Page{Permalink: "/news/gophertimes-born", Title:"Gopher Times", Description:"Gopher Times is born.",	Keywords:"Go, Golang, Go News,Golang news",PageTitle:"Gopher Times",Content:"",Template:"index"}, Tags:[]string{"golang","gophertimes"},ContributedBy: "Brian Ketelsen", Byline: "Brian Ketelsen", PostedTime: time.Seconds(), Blurb: "Gopher Times is Born!", FullDescription: "I'm hoping that Gopher Times will serve as a source of quality news for the Go community"})
	if err != nil {
	log.Println(err)
	}
}

func main() {
	flag.Parse()
		if *initdb {
			loadFirstRecord()
			return
		}
		var portString = fmt.Sprintf(":%d", *port)


	h := web.ProcessForm(10000, true, // limit size of form to 10k, enable xsrf
		web.NewRouter().
			Register("/static/<path:.*>", "GET", web.DirectoryHandler("static/")).
			//			Register("/favicon.ico", "GET", web.FileHandler("static/favicon.ico")).
			Register("/", "GET", homeHandler).
			Register("/<path:(.*)>", "GET", viewHandler))
	server.Run(portString, h)

}
