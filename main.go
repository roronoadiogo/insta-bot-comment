package main

import (
	"github.com/gocolly/colly"
	"github.com/roronoadiogo/insta-bot-comment/infra/services"
	"github.com/roronoadiogo/insta-bot-comment/usecases"
)

func main() {
	println("Start the project")

	a := services.NewWebsiteAccessor()

	a.AccessSite(usecases.WebParams{
		Url:      "www.instagram.com",
		Instance: *colly.NewCollector(),
	})

}
