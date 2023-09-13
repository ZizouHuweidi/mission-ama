package routes

import (
	"html/template"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	missions struct {
		controller.Controller
	}

	missionsData struct {
		ShowCacheWarning bool
		FrontendTabs     []missionsTab
		BackendTabs      []missionsTab
	}

	missionsTab struct {
		Title string
		Body  template.HTML
	}
)

func (c *missions) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "missions"
	page.Title = "missions"

	// This page will be cached!
	page.Cache.Enabled = true
	page.Cache.Tags = []string{"page_missions", "page:list"}

	// A simple example of how the Data field can contain anything you want to send to the templates
	// even though you wouldn't normally send markup like this
	page.Data = missionsData{
		ShowCacheWarning: true,
		FrontendTabs: []missionsTab{
			{
				Title: "HTMX",
				Body:  template.HTML(`Completes HTML as a hypertext by providing attributes to AJAXify anything and much more. Visit <a href="https://htmx.org/">htmx.org</a> to learn more.`),
			},
			{
				Title: "Alpine.js",
				Body:  template.HTML(`Drop-in, Vue-like functionality written directly in your markup. Visit <a href="https://alpinejs.dev/">alpinejs.dev</a> to learn more.`),
			},
			{
				Title: "Bulma",
				Body:  template.HTML(`Ready-to-use frontend components that you can easily combine to build responsive web interfaces with no JavaScript requirements. Visit <a href="https://bulma.io/">bulma.io</a> to learn more.`),
			},
		},
		BackendTabs: []missionsTab{
			{
				Title: "Echo",
				Body:  template.HTML(`High performance, extensible, minimalist Go web framework. Visit <a href="https://echo.labstack.com/">echo.labstack.com</a> to learn more.`),
			},
			{
				Title: "Ent",
				Body:  template.HTML(`Simple, yet powerful ORM for modeling and querying data. Visit <a href="https://entgo.io/">entgo.io</a> to learn more.`),
			},
		},
	}

	return c.RenderPage(ctx, page)
}
