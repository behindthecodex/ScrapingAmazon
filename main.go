package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("body", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		//fmt.Println(e.Attr("href"))

		e.ForEach("a.a-link-normal.s-access-detail-page.s-color-twister-title-link.a-text-normal", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Attr("href"))
			fmt.Println(el.Attr("title"))
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.amazon.com/s?i=computers-intl-ship&bbn=16225007011&rh=n%3A16225007011&dc&fst=as%3Aoff&qid=1587960854&ref=sr_ex_n_1")
}
