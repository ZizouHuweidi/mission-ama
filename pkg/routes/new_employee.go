package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/context"
	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	employee struct {
		controller.Controller
	}

	employeeForm struct {
		Email       string `form:"email" validate:"required,email"`
		Message     string `form:"message" validate:"required"`
		Subemployee controller.FormSubmission
	}
)

func (c *employee) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "employee"
	page.Title = "employee us"
	page.Form = employeeForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*employeeForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *employee) Post(ctx echo.Context) error {
	var form employeeForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to bind form")
	}

	if err := form.Subemployee.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form subemployee")
	}

	if !form.Subemployee.HasErrors() {
		err := c.Container.Mail.
			Compose().
			To(form.Email).
			Subject("employee form submitted").
			Body(fmt.Sprintf("The message is: %s", form.Message)).
			Send(ctx)
		if err != nil {
			return c.Fail(err, "unable to send email")
		}
	}

	return c.Get(ctx)
}
