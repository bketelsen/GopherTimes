package main

import "time"


type Page struct {
    Permalink   string
    Title       string
    Description string
    Keywords    string
    PageTitle   string
    Content     string
    Template    string
}

type CachedNewsItem struct {
    NewsItem *NewsItem
    CachedAt int64
}

type CachedNewsItemArray struct {
	NewsItems []*NewsItem
	CachedAt int64
}

type NewsItem struct {
    Page
    Byline          string
    PostedTime      int64
    Blurb           string
    FullDescription string
    ImagePath       string
    ExternalLink    string
    NewsCategory    string
    ContributedBy   string
    Tags            []string
}

func (n *NewsItem) PostedTimeEnglish() string {
    localTime := time.SecondsToLocalTime(n.PostedTime)
    return localTime.Format("_2 January 2006")
}
