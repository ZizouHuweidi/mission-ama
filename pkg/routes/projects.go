package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	projects struct {
		controller.Controller
	}

	projectTable struct {
		Name        string
		Description string
	}
)

func (c *projects) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "projects"
	page.Title = "AMA Projects"

	// page.Data

	return c.RenderPage(ctx, page)
}
