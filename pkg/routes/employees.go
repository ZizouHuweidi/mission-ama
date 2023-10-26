package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	employees struct {
		controller.Controller
	}

	employeeTable struct {
		Name       string
		Phone      int
		CSP        bool
		Occupation string
		Created_at string
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

func (c *employees) fetchEmployees(ctx echo.Context) []employeeTable {
	emps, err := c.Container.ORM.Employee.Query().All(ctx.Request().Context())
	if err != nil {
		ctx.Logger().Errorf("failed querying employees: %v", err)
		return nil
	}

	employeeData := make([]employeeTable, len(emps))
	for i, em := range emps {
		employeeData[i] = employeeTable{
			Name:       em.Name,
			Phone:      em.Phone,
			CSP:        em.CSP,
			Occupation: em.Occupation,
			Created_at: em.CreatedAt.String(),
		}
		ctx.Logger().Infof("Employee found: %s", em.Name)
	}

	return employeeData

}
