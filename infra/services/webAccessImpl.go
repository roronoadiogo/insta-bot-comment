package services

import (
	"github.com/gocolly/colly"
	"github.com/roronoadiogo/insta-bot-comment/usecases"
)

type websiteAccessor struct{}

func NewWebsiteAccessor() usecases.AccessWebSite {
	return &websiteAccessor{}
}

func (web *websiteAccessor) AccessSite(w usecases.WebParams) {

	c := w.Instance

	c.OnHTML("label._aa48", func(e *colly.HTMLElement) {
		x := e.DOM.HasClass("_aa48")
		println(x)

	})

	c.Visit("http://www.instagram.com")
	c.Wait()

}

func (web *websiteAccessor) ClickAction() {}
