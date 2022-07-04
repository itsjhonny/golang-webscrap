package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type NoticeModule struct {
	Title   string
	Summary string
	Date    string
	Link    string
	Body    string
}

func main() {
	c := colly.NewCollector()

	notice_array := make([]NoticeModule, 0, 1)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s\n", r.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})
	c.OnHTML("section.container", func(h *colly.HTMLElement) {

		h.ForEach("div.widget--info__text-container", func(_ int, elem *colly.HTMLElement) {

			if coded_link, ok := elem.DOM.Find("a").Attr("href"); ok {

				decoded_link, err := url.QueryUnescape(coded_link)
				if err != nil {
					log.Fatal(err)
					return
				}

				cleaned_link := <-clearUrl(decoded_link)

				new := <-getNew(cleaned_link)
				notice_array = append(notice_array, new)
			}

		})

	})

	c.OnScraped(func(r *colly.Response) {

		fmt.Println("\nfinalizou")
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		fmt.Printf("%v", enc.Encode(notice_array))

	})

	c.Visit("https://g1.globo.com/busca/?q=noticias&order=recent&species=not%C3%ADcias")

}

func clearUrl(url string) <-chan string {
	d := make(chan string)
	go func() {
		split_u := strings.Split(url, "&u=")
		split_syn := strings.Split(split_u[1], "&syn=")
		d <- split_syn[0]
	}()

	return d
}

func getNew(url string) <-chan NoticeModule {
	d := make(chan NoticeModule)
	notice := NoticeModule{}
	go func() {
		c := colly.NewCollector()

		notice.Link = url

		c.OnHTML("main", func(h *colly.HTMLElement) {

			title := h.ChildText("h1.content-head__title")
			notice.Title = strings.TrimSpace(title)

			summary := h.ChildText("h2.content-head__subtitle")
			notice.Summary = strings.TrimSpace(summary)

			date := h.ChildAttr("p.content-publication-data__updated > time", "datetime")
			notice.Date = strings.TrimSpace(date)

			body := h.ChildText("p.content-text__container")
			notice.Body = strings.TrimSpace(body)

		})

		c.OnScraped(func(r *colly.Response) {

			d <- notice

		})

		c.Visit(url)

	}()

	return d

}
