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
var database *string = flag.String("database", "public_web", "mongo database name")

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

func tagList() []interface{} {

    mongo, err := mgo.Mongo("localhost")
    if err != nil {
        return nil
    }
    defer mongo.Close()

    var tags map[string]interface{}

    db := mongo.DB(*database)
    err = db.Run(bson.D{{"distinct", "newsitems"}, {"key", "tags"}}, &tags)
    if tags["values"] != nil {
        return tags["values"].([]interface{})
    }
    return nil
}


func loadNewsItems(search bson.M) ([]*NewsItem, os.Error) {

    mongo, err := mgo.Mongo("localhost")
    if err != nil {
        return nil, err
    }
    defer mongo.Close()

    items := make([]*NewsItem, 0)

    c := mongo.DB(*database).C("newsitems")

    iter, err := c.Find(search).Sort(bson.M{"postedtime": -1}).Iter()
    var i int = 0

    for {
        item := &NewsItem{}
        err = iter.Next(&item)
        if err != nil {
            break
        }
        items = append(items, item)

        i++
    }

    if err != mgo.NotFound {
        return items[0:i], err
    }

    return items, nil
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
    list, err := loadNewsItems(bson.M{})

    if err != nil {
        renderSingleTemplate(req, web.StatusNotFound, "404", p, list)
    } else {
        /*		var templateName string
        		if len(p.Template) != 0 {
        			templateName = p.Template
        		} else {
        			templateName = "article"
        		}
        */
        renderSingleTemplate(req, web.StatusOK, "article", p, list)
    }
}

func tagsHandler(req *web.Request) {
    tag := req.Param.Get("tag")

    results, err := loadNewsItems(bson.M{"tags":tag})
    p, err := loadNewsItems(bson.M{})
    if err != nil {
        renderListTemplate(req, web.StatusNotFound, "404", results, p)
    } else {
        renderListTemplate(req, web.StatusOK, "index", results, p)
    }
}

func categoryHandler(req *web.Request) {
    category := req.Param.Get("category")

    //p, err := loadNewsItemsByTag(tag)
    results, err := loadNewsItems(bson.M{"newscategory":category})
    p, err := loadNewsItems(bson.M{})
    if err != nil {
        renderListTemplate(req, web.StatusNotFound, "404", results, p)
    } else {
        renderListTemplate(req, web.StatusOK, "index", results, p)
    }
}


func editHandler(req *web.Request) {
    path := req.Param.Get("path")
    log.Println("Path:", path)
    n, err := loadNewsItem(path)
    if err != nil {
        n = &NewsItem{Page: Page{Permalink: path, Title: "Title", Description: "Description", Keywords: "Go, Golang, Go News,Golang news", PageTitle: "Page Title", Content: "Content", Template: "article"}, Tags: []string{"golang", "gophertimes"}, ContributedBy: "", Byline: "Brian Ketelsen", PostedTime: time.Seconds(), Blurb: "Article Blurb", FullDescription: "Article Full"}
    }
    renderEditTemplate(req, "edit", n)

}

func saveHandler(req *web.Request) {
    permalink := req.Param.Get("page-permalink")
    title := req.Param.Get("page-title")
    description := req.Param.Get("page-description")
    pageTitle := req.Param.Get("page-page-title")
    keywords := req.Param.Get("page-keywords")
    content := req.Param.Get("page-content")
    template := req.Param.Get("page-template")

    byline := req.Param.Get("newsitem-byline")
    blurb := req.Param.Get("newsitem-blurb")
    fulldescription := req.Param.Get("newsitem-fulldescription")
    imagepath := req.Param.Get("newsitem-imagepath")
    externallink := req.Param.Get("newsitem-externallink")
    newscategory := req.Param.Get("newsitem-newscategory")
    contributedby := req.Param.Get("newsitem-contributedby")

    n := &NewsItem{Page: Page{Permalink: permalink,
        Title:       title,
        Description: description,
        Keywords:    keywords,
        PageTitle:   pageTitle,
        Content:     content,
        Template:    template},
        Tags:            []string{"golang", "gophertimes"},
        ContributedBy:   contributedby,
        Byline:          byline,
        PostedTime:      time.Seconds(),
        Blurb:           blurb,
        ImagePath:       imagepath,
        ExternalLink:    externallink,
        NewsCategory:    newscategory,
        FullDescription: fulldescription}

    mongo, err := mgo.Mongo("127.0.0.1")
    if err != nil {
        panic(err)
    }

    c := mongo.DB(*database).C("newsitems")

    err = c.Insert(n)
    if err != nil {
        log.Println(err)
    }

    req.Redirect(permalink, false)

}

func homeHandler(req *web.Request) {

    p, err := loadNewsItems(bson.M{})

    if err != nil {
        log.Println(err.String())
        renderListTemplate(req, web.StatusNotFound, "404", p[0:1], p)
    } else {
        renderListTemplate(req, web.StatusOK, "index", p[0:1], p[1:])
    }
}


var cachedNewsItems = make(map[string]*CachedNewsItem)
var templates = make(map[string]*template.Template)
//var mongo *mgo.Session

func init() {
    for _, tmpl := range []string{"index", "404", "article", "edit"} {
        templates[tmpl] = template.MustParseFile("templates/"+tmpl+".html", nil)
    }

}

func cacheNewsItem(n *NewsItem) {
    cachedNewsItems[n.Page.Permalink] = &CachedNewsItem{NewsItem: n, CachedAt: time.Seconds()}
}

func removeCachedNewsItem(permalink string) {
    cachedNewsItems[permalink] = nil, false
}


func renderEditTemplate(req *web.Request, tmpl string, n *NewsItem) {

    err := templates[tmpl].Execute(
        req.Respond(web.StatusOK),
        map[string]interface{}{
            "item": n,
            "xsrf": req.Param.Get("xsrf"),
        })
    if err != nil {
        log.Println("error rendering", tmpl, err)
    }
}


func renderSingleTemplate(req *web.Request, status int, tmpl string, n *NewsItem, items []*NewsItem) {

    err := templates[tmpl].Execute(
        req.Respond(status),
        map[string]interface{}{
            "item":      n,
            "newsItems": items,
            "tags":      tagList(),
            "xsrf":      req.Param.Get("xsrf"),
        })
    if err != nil {
        log.Println("error rendering", tmpl, err)
    }
}

func renderListTemplate(req *web.Request, status int, tmpl string, results []*NewsItem, items []*NewsItem) {

    err := templates[tmpl].Execute(
        req.Respond(status),
        map[string]interface{}{
            "results":   results,
            "newsItems": items,
            "tags":      tagList(),
            "xsrf":      req.Param.Get("xsrf"),
        })
    if err != nil {
        log.Println("error rendering", tmpl, err)
    }
}


func loadFirstRecord() {
    //open mongo
    mongo, err := mgo.Mongo("127.0.0.1")
    if err != nil {
        panic(err)
    }

    c := mongo.DB(*database).C("newsitems")

    err = c.Insert(&NewsItem{Page: Page{Permalink: "news/gophertimes-born", Title: "Gopher Times", Description: "Gopher Times is born.", Keywords: "Go, Golang, Go News,Golang news", PageTitle: "Gopher Times", Content: "", Template: "index"}, NewsCategory:"news", Tags: []string{"golang", "gophertimes"}, ContributedBy: "Brian Ketelsen", Byline: "Brian Ketelsen", PostedTime: time.Seconds(), Blurb: "Gopher Times is Born!", FullDescription: "I'm hoping that Gopher Times will serve as a source of quality news for the Go community"})
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
          	Register("/category/<category:(.*)>", "GET", categoryHandler). 
            Register("/tags/<tag:(.*)>", "GET", tagsHandler).
            Register("/edit/<path:(.*)>", "GET", editHandler, "POST", saveHandler).
            Register("/<path:(.*)>", "GET", viewHandler))
    server.Run(portString, h)

}
