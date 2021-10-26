package main

import (
	"fmt"

	"github.com/ferjozsot23/webScraperGO/csv"
	"github.com/gocolly/colly"
)

var protocol = ""
var domain = ""
var path = ""
var URL = fmt.Sprintf("%s://%s/%s", protocol, domain, path)

func main() {
	// Creating the colly Collector
	collector := colly.NewCollector(
		colly.AllowedDomains(domain),
	)
	var data [][]string
	// Añadiendo la primera fila
	data = append(data, [][]string{{"", "", ""}}...)

	htmlFunctionExtracting := func(htmlElement *colly.HTMLElement) {

		field1 := htmlElement.ChildText("") // extracting text from html element
		field2 := htmlElement.ChildText("")
		field3 := htmlElement.ChildAttrs("", "") // extracting value from attr of a html element

		data = append(data, [][]string{{field1, field2, field3[0]}}...)

	}

	// Scraping all the div's elements with ".elementor-column-wrap"
	collector.OnHTML(".", htmlFunctionExtracting)

	collector.Visit(URL)
	csv.SaveDataOnCSVFormat(data, "data")

}
