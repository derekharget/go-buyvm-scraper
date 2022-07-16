package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	pid      int
	quantity int
	name     string
}

func main() {

	var productData = []Product{}

	c := colly.NewCollector(
		colly.AllowedDomains("my.frantech.ca"),
	)

	for i := 1; i <= 12; i++ {

		c.OnHTML("div#product"+strconv.Itoa(i), func(e *colly.HTMLElement) {

			pidLink := e.ChildAttrs("a", "href")[0]
			pid := strings.Split(pidLink, "&pid=")[1]

			pidQuantity := e.DOM.Find(".package-qty").Text()

			var quantity string

			if len(pidQuantity) > 1 {
				quantity = strings.Replace(strings.Split(pidQuantity, " Available")[0], " ", "", -1)
			} else {
				quantity = "999"
			}

			// Package name

			pidName := e.DOM.Find(".package-name").Text()

			var p = new(Product)
			p.name = pidName
			p.quantity, _ = strconv.Atoi(quantity)
			p.pid, _ = strconv.Atoi(pid)

			productData = append(productData, *p)

			fmt.Println(productData)
		})
	}

	c.Visit("https://my.frantech.ca/cart.php?gid=37")

	// for i := 1; i < 12; i++ {

	// 	c.OnHTML("div#product"+strconv.Itoa(i+1), func(e *colly.HTMLElement) {
	// 		data := e.ChildAttrs("a", "href")[0]
	// 		links := strings.Split(data, "&pid=")[1]
	// 		fmt.Println(links)
	// 	})
	// }

	// c.Visit("https://my.frantech.ca/cart.php?gid=39")
}
