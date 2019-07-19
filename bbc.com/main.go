/*
 * this package for get data from scrapping BBC without unit testing , because , this art just for improve and learn
 */
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery" // library for scrapping
)

// struct for representatif for data when we put
type Article struct {
	Title   string `json:"title,omitempty"`   // title struct for add.in tittle in this scrapping
	Content string `json:"content,omitempty"` // content struct for add.in content in this scrapping
}

// function main for run this file
func main() {
	ScrapeHTML() // call function in this file
}

//ScrapeHTML is function to scrape
func ScrapeHTML() {
	// setting your target website get scrapping
	website := "https://www.bbc.com/indonesia/indonesia?xtor=SEC-3-[GOO]-[70466280620]-[346260440805]-S-[berita%20internasional]"
	// connect the website
	res, err := http.Get(website)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// call function to get all HTML in this website
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// make array struct for saving data
	rows := make([]Article, 0)
	// call class including all data tittle and content in this website
	doc.Find(".eagle").Children().Each(func(i int, sel *goquery.Selection) {
		// call this struct and representatif
		row := new(Article)
		row.Title = sel.Find(".title-link__title-text").Text()
		row.Content = sel.Find(".eagle-item__summary").Text()
		rows = append(rows, *row)
	})
	// change data struct in json
	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	// print to json
	log.Println(string(bts))
}
