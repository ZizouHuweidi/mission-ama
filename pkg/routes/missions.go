package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	missions struct {
		controller.Controller
	}

	missionTable struct {
		Name        string
		Purpose     string
		Destination string
		StartDate   string
		EndDate     string
		Transport   string
		Created_at  string
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

func (c *missions) fetchMissions(ctx echo.Context) []missionTable {
	mis, err := c.Container.ORM.Mission.Query().All(ctx.Request().Context())
	if err != nil {
		ctx.Logger().Errorf("failed querying missions: %v", err)
		return nil
	}

	missionData := make([]missionTable, len(mis))
	for i, mi := range mis {

		missionData[i] = missionTable{
			Name:        mi.Name,
			Purpose:     mi.Purpose,
			Destination: mi.Destination,
			StartDate:   mi.StartDate.String(),
			EndDate:     mi.EndDate.String(),
			Transport:   mi.Transport,
			Created_at:  mi.CreatedAt.String(),
		}
		ctx.Logger().Infof("Mission found: %s", mi.Name)
	}

	return missionData

}
