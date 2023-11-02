package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/zizouhuweidi/mission-ama/ent"

	"github.com/zizouhuweidi/mission-ama/pkg/context"
	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	mission struct {
		controller.Controller
	}

	missionForm struct {
		Name        string `form:"name" validate:"required"`
		Purpose     string `form:"purpose" validate:"required"`
		Destination string `form:"destination" validate:"required"`
		startDate   string `form:"startDate" validate:"required"`
		endDate     string `form:"endDate" validate:"required"`
		Transport   string `form:"transport" validate:"required"`
		EmployeeID  int    `form:"projectID" validate:"required"`
		ProjectID   int    `form:"projectID"`
		Submission  controller.FormSubmission
	}

	pageData struct {
		Projects  []ent.Project
		Employees []ent.Employee
	}
)

func (c *mission) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "mission"
	page.Title = "New Mission"
	// page.Data = c.fetch(ctx)
	page.Form = missionForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*missionForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *mission) fetch(ctx echo.Context) {

}

func (c *mission) Post(ctx echo.Context) error {
	var form missionForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to bind form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	e, err := c.Container.ORM.Mission.
		Create().
		SetName(form.Name).
		SetProjectID(form.ProjectID).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		ctx.Logger().Infof("Mission created: %s", e.Name)
	default:
		return c.Fail(err, "unable to create mission")
	}

	return c.Get(ctx)
}
