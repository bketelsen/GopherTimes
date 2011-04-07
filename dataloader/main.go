package main

import (
	"launchpad.net/mgo"
		"log"

)

type Page struct {
	Path    string
	Title   string
	Description string
	Keywords		string
	PageTitle	string
	Content string
	Product	Product
	PressRelease PressRelease
}

type Product struct {

	Name            string
	Blurb           string
	FullDescription string
	ImagePath       string
}

type PressRelease struct {
	
	Date      string
	Title     string
	PathToPdf string
}

func main() {

	mongo, err := mgo.Mongo("localhost")
	if err != nil {
		panic(err)
	}

	c := mongo.DB("public_web").C("page")
	
	p := &Page{Path:"products/clear_idfraud",Title:"Clarity Services - Clear ID Fraud", Content:"About Clear ID Fraud Here", Product: Product{Name:"Clear ID Fraud", Blurb:"Clear ID Fraud Blurb"}}

	err = c.Insert(p)
	if err != nil {
		log.Println(err)
	}

	defer mongo.Close()
}
