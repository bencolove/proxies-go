package scrape

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

// From https://www.freeproxylists.net/?c=&pt=&pr=&a%5B%5D=1&a%5B%5D=2&u=70

const URL = "https://www.freeproxylists.net/?c=&pt=&pr=&a%5B%5D=1&a%5B%5D=2&u=70"

func CollectProxy() error {
	c := colly.NewCollector()

	c.OnResponse(func(resp *colly.Response) {
		fmt.Printf("statusCode: %d\n", resp.StatusCode)
		fmt.Println(string(resp.Body))
	})

	c.OnHTML("tbody", func(elem *colly.HTMLElement) {
		elem.ForEach("tr.Odd", func(i int, s *colly.HTMLElement) {
			fmt.Printf("index: %d, class: %v\n", i, s.Attr("class"))
		})
	})

	return c.Visit(URL)
}
