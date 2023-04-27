package usecases

import (
	"github.com/gocolly/colly"
)

type WebParams struct {
	Url      string
	Instance colly.Collector
}

type AccessWebSite interface {
	AccessSite(w WebParams)
	ClickAction()
}
