package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Link    string `json:"link,omitempty"`
}

func main() {
	ScrapeHTML()
}

//ScrapeHTML is function to scrape
func ScrapeHTML() {
	res, err := http.Get("https://www.devdungeon.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	rows := make([]Article, 0)

	doc.Find(".view-content").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(Article)
		row.Title = sel.Find("article a").Text()
		row.Content = sel.Find("article .field-item").Text()
		rows = append(rows, *row)
	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bts))
}
