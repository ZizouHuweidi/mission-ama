package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/context"
	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	employee struct {
		controller.Controller
	}

	employeeForm struct {
		Name       string `form:"name" validate:"required"`
		Phone      int    `form:"phone" validate:"required"`
		CSP        bool   `form:"csp"`
		Occupation string `form:"occupation"`
		Submission controller.FormSubmission
	}
)

func (c *employee) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "employee"
	page.Title = "New Employee"
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

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	e, err := c.Container.ORM.Employee.
		Create().
		SetName(form.Name).
		SetPhone(form.Phone).
		SetCSP(form.CSP).
		SetOccupation(form.Occupation).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		ctx.Logger().Infof("Employee created: %s", e.Name)
	default:
		return c.Fail(err, "unable to create employee")
	}

	return c.Get(ctx)
}
