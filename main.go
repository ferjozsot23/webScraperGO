package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ferjozsot23/webScraperGO/csv"
	"github.com/gocolly/colly"
)



var data = [][]string{{"Item", "Category", "Price", "ImagenURL"}}

func main() {
	if len(os.Args) != 4 {
		log.Println("Put domain, path and number of pages")
		os.Exit(1)
	}
	var domain = os.Args[1]
	var path = os.Args[2]
	var pages, err = strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("Put a correct number of pages")
	}

	for i := 1; i <= pages; i++ {
		collector := colly.NewCollector(
			colly.AllowedDomains(domain),
		)
		collector.OnHTML(".product-inner", func(htmlElement *colly.HTMLElement) {

			item := htmlElement.ChildText(".product-loop-header h2")       // extracting text from html element
			category := htmlElement.ChildText(".product-loop-header span") // extracting text from html element
			price := htmlElement.ChildText(".product-loop-footer bdi")
			image := htmlElement.ChildAttr("img", "src") // extracting value from attr of a html element

			data = append(data, [][]string{{item, category, price, image}}...)
		})

		collector.Visit(fmt.Sprintf("%s%d", path, i))
	}
	csv.SaveDataOnCSVFormat(data, "data-timecenterec")


}
