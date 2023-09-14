package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/context"
	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	project struct {
		controller.Controller
	}

	projectForm struct {
		Email      string `form:"email" validate:"required,email"`
		Message    string `form:"message" validate:"required"`
		Submission controller.FormSubmission
	}
)

func (c *project) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "project"
	page.Title = "new project"
	page.Form = projectForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*projectForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *project) Post(ctx echo.Context) error {
	var form projectForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to bind form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if !form.Submission.HasErrors() {
		err := c.Container.Mail.
			Compose().
			To(form.Email).
			Subject("project form submitted").
			Body(fmt.Sprintf("The message is: %s", form.Message)).
			Send(ctx)
		if err != nil {
			return c.Fail(err, "unable to send email")
		}
	}

	return c.Get(ctx)
}
