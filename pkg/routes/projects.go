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
		Created_at  string
	}
)

func (c *projects) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "projects"
	page.Title = "AMA Projects"

	page.Data = c.fetchProjects(ctx)

	return c.RenderPage(ctx, page)
}

func (c *projects) fetchProjects(ctx echo.Context) []projectTable {
	pjs, err := c.Container.ORM.Project.Query().All(ctx.Request().Context())
	if err != nil {
		ctx.Logger().Errorf("failed querying projects: %v", err)
		return nil
	}

	projectData := make([]projectTable, len(pjs))
	for i, pj := range pjs {
		projectData[i] = projectTable{
			Name:        pj.Name,
			Description: pj.Description,
			Created_at:  pj.CreatedAt.String(),
		}
		ctx.Logger().Infof("Project found: %s", pj.Name)
	}

	return projectData

}
