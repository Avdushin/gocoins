package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	prices()
}

func prices() {
	url := "https://mainfin.ru/currency/cny" // change cny to other currency (Example: /usd)
	res, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	price := doc.Find(".mark-text").Text()
	fmt.Printf(" ¥%s", price)
}
