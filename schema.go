package main

type Page struct {
	Path    string
	Title   string
	Content string
	Products		[]Product
	PressReleases	[]PressRelease
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
