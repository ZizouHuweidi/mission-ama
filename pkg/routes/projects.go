package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	projects struct {
		controller.Controller
	}

	projectTable struct {
		Title string
		Body  string
	}
)

func (c *projects) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "projects"
	page.Title = "AMA Projects"
	// page.Data = c.fetchProjects(ctx)

	return c.RenderPage(ctx, page)
}

// fetchPosts is an mock example of fetching posts to illustrate how paging works
func (c *projects) fetchProjects(ctx echo.Context) error {
	p, err := c.Container.ORM.Project.Query().All(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "unable to fetch projects")
	}

	fmt.Print(p)

	return c.Get(ctx)
}
