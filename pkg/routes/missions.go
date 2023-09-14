package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	missions struct {
		controller.Controller
	}

	missionTable struct {
		Title string
		Body  string
	}
)

func (c *missions) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "missions"
	page.Title = "AMA Missions"
	page.Data = c.fetchMissions(ctx)

	return c.RenderPage(ctx, page)
}

// fetchPosts is an mock example of fetching posts to illustrate how paging works
func (c *missions) fetchMissions(ctx echo.Context) error {
	m, err := c.Container.ORM.Mission.Query().All(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "unable to fetch missions")
	}

	fmt.Print(m)

	return c.Get(ctx)
}
