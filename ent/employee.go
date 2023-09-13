// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/zizouhuweidi/mission-ama/ent/employee"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone int `json:"phone,omitempty"`
	// CSP holds the value of the "CSP" field.
	CSP bool `json:"CSP,omitempty"`
	// Occupation holds the value of the "occupation" field.
	Occupation string `json:"occupation,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeQuery when eager-loading is set.
	Edges               EmployeeEdges `json:"edges"`
	employee_supervisor *int
	selectValues        sql.SelectValues
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	// Missions holds the value of the missions edge.
	Missions []*Mission `json:"missions,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Supervisee holds the value of the supervisee edge.
	Supervisee *Employee `json:"supervisee,omitempty"`
	// Supervisor holds the value of the supervisor edge.
	Supervisor []*Employee `json:"supervisor,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MissionsOrErr returns the Missions value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) MissionsOrErr() ([]*Mission, error) {
	if e.loadedTypes[0] {
		return e.Missions, nil
	}
	return nil, &NotLoadedError{edge: "missions"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// SuperviseeOrErr returns the Supervisee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeEdges) SuperviseeOrErr() (*Employee, error) {
	if e.loadedTypes[2] {
		if e.Supervisee == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Supervisee, nil
	}
	return nil, &NotLoadedError{edge: "supervisee"}
}

// SupervisorOrErr returns the Supervisor value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) SupervisorOrErr() ([]*Employee, error) {
	if e.loadedTypes[3] {
		return e.Supervisor, nil
	}
	return nil, &NotLoadedError{edge: "supervisor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case employee.FieldCSP:
			values[i] = new(sql.NullBool)
		case employee.FieldID, employee.FieldPhone:
			values[i] = new(sql.NullInt64)
		case employee.FieldName, employee.FieldOccupation:
			values[i] = new(sql.NullString)
		case employee.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case employee.ForeignKeys[0]: // employee_supervisor
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case employee.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case employee.FieldPhone:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				e.Phone = int(value.Int64)
			}
		case employee.FieldCSP:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field CSP", values[i])
			} else if value.Valid {
				e.CSP = value.Bool
			}
		case employee.FieldOccupation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field occupation", values[i])
			} else if value.Valid {
				e.Occupation = value.String
			}
		case employee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case employee.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employee_supervisor", value)
			} else if value.Valid {
				e.employee_supervisor = new(int)
				*e.employee_supervisor = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Employee.
// This includes values selected through modifiers, order, etc.
func (e *Employee) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryMissions queries the "missions" edge of the Employee entity.
func (e *Employee) QueryMissions() *MissionQuery {
	return NewEmployeeClient(e.config).QueryMissions(e)
}

// QueryProjects queries the "projects" edge of the Employee entity.
func (e *Employee) QueryProjects() *ProjectQuery {
	return NewEmployeeClient(e.config).QueryProjects(e)
}

// QuerySupervisee queries the "supervisee" edge of the Employee entity.
func (e *Employee) QuerySupervisee() *EmployeeQuery {
	return NewEmployeeClient(e.config).QuerySupervisee(e)
}

// QuerySupervisor queries the "supervisor" edge of the Employee entity.
func (e *Employee) QuerySupervisor() *EmployeeQuery {
	return NewEmployeeClient(e.config).QuerySupervisor(e)
}

// Update returns a builder for updating this Employee.
// Note that you need to call Employee.Unwrap() before calling this method if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return NewEmployeeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Employee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(fmt.Sprintf("%v", e.Phone))
	builder.WriteString(", ")
	builder.WriteString("CSP=")
	builder.WriteString(fmt.Sprintf("%v", e.CSP))
	builder.WriteString(", ")
	builder.WriteString("occupation=")
	builder.WriteString(e.Occupation)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee
