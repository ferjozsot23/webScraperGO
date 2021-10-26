package main

import (
	"fmt"

	"github.com/ferjozsot23/webScraperGO/csv"
	"github.com/gocolly/colly"
)

var protocol = "https"
var domain = "deceac-el.espe.edu.ec"
var path = "personal-docente-e-investigadores/"
var URL = fmt.Sprintf("%s://%s/%s", protocol, domain, path)

func main() {
	// Creating the colly Collector
	collector := colly.NewCollector(
		colly.AllowedDomains(domain),
	)
	var data [][]string
	// Añadiendo la primera fila
	data = append(data, [][]string{{"Professor", "Email", "ImageURL"}}...)

	htmlFunctionExtracting := func(htmlElement *colly.HTMLElement) {

		professor := htmlElement.ChildText("h5") // extracting text from html element
		email := htmlElement.ChildText("span[style='color: #008000; font-family: Agency FB,serif;']")
		image := htmlElement.ChildAttrs("img", "src") // extracting value from attr of a html element

		if professor != "" {
			if email == "" {
				email = "email empty"
			}
			if image[0] == "" {
				image[0] = "image source empty"
			}

			data = append(data, [][]string{{professor, email, image[0]}}...)
		}
	}

	// Scraping all the div's elements with ".elementor-column-wrap"
	collector.OnHTML(".elementor-column-wrap", htmlFunctionExtracting)

	collector.Visit(URL)
	csv.SaveDataOnCSVFormat(data, "data-espeec")

}
