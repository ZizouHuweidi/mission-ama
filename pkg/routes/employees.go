package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	employees struct {
		controller.Controller
	}

	employeeTable struct {
		Title string
		Body  string
	}
)

func (c *employees) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "employees"
	page.Title = "AMA Employees"
	page.Data = c.fetchEmployees(ctx)

	return c.RenderPage(ctx, page)
}

// fetchPosts is an mock example of fetching posts to illustrate how paging works
func (c *employees) fetchEmployees(ctx echo.Context) error {
	m, err := c.Container.ORM.Employee.Query().All(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "unable to fetch employees")
	}

	fmt.Print(m)

	return c.Get(ctx)
}
