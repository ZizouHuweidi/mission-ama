package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/context"
	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	project struct {
		controller.Controller
	}

	ProjectForm struct {
		Name        string `form:"name" validate:"required"`
		Description string `form:"name" validate:"required"`
		Submission  controller.FormSubmission
	}
)

func (c *project) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "project"
	page.Title = "New Project"
	page.Form = ProjectForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*ProjectForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *project) Post(ctx echo.Context) error {
	var form ProjectForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to bind form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if !form.Submission.HasErrors() {
		e, err := c.Container.ORM.Project.
			Create().
			SetName(form.Name).
			SetDescription(form.Description).
			Save(ctx.Request().Context())

		switch err.(type) {
		case nil:
			ctx.Logger().Infof("Project created: %s", e.Name)
		default:
			return c.Fail(err, "unable to create project")
		}
	}
	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	return c.Get(ctx)
}
