// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/zizouhuweidi/mission-ama/ent/employee"
	"github.com/zizouhuweidi/mission-ama/ent/mission"
	"github.com/zizouhuweidi/mission-ama/ent/project"
)

// Mission is the model entity for the Mission schema.
type Mission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Purpose holds the value of the "purpose" field.
	Purpose string `json:"purpose,omitempty"`
	// Destination holds the value of the "destination" field.
	Destination string `json:"destination,omitempty"`
	// StartDate holds the value of the "startDate" field.
	StartDate time.Time `json:"startDate,omitempty"`
	// EndDate holds the value of the "endDate" field.
	EndDate time.Time `json:"endDate,omitempty"`
	// Transport holds the value of the "transport" field.
	Transport string `json:"transport,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionQuery when eager-loading is set.
	Edges             MissionEdges `json:"edges"`
	employee_missions *int
	project_missions  *int
	selectValues      sql.SelectValues
}

// MissionEdges holds the relations/edges for other nodes in the graph.
type MissionEdges struct {
	// Employee holds the value of the employee edge.
	Employee *Employee `json:"employee,omitempty"`
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[0] {
		if e.Employee == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[1] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			values[i] = new(sql.NullInt64)
		case mission.FieldName, mission.FieldPurpose, mission.FieldDestination, mission.FieldTransport:
			values[i] = new(sql.NullString)
		case mission.FieldStartDate, mission.FieldEndDate, mission.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case mission.ForeignKeys[0]: // employee_missions
			values[i] = new(sql.NullInt64)
		case mission.ForeignKeys[1]: // project_missions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mission fields.
func (m *Mission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case mission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case mission.FieldPurpose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purpose", values[i])
			} else if value.Valid {
				m.Purpose = value.String
			}
		case mission.FieldDestination:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination", values[i])
			} else if value.Valid {
				m.Destination = value.String
			}
		case mission.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startDate", values[i])
			} else if value.Valid {
				m.StartDate = value.Time
			}
		case mission.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endDate", values[i])
			} else if value.Valid {
				m.EndDate = value.Time
			}
		case mission.FieldTransport:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transport", values[i])
			} else if value.Valid {
				m.Transport = value.String
			}
		case mission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case mission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employee_missions", value)
			} else if value.Valid {
				m.employee_missions = new(int)
				*m.employee_missions = int(value.Int64)
			}
		case mission.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_missions", value)
			} else if value.Valid {
				m.project_missions = new(int)
				*m.project_missions = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Mission.
// This includes values selected through modifiers, order, etc.
func (m *Mission) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryEmployee queries the "employee" edge of the Mission entity.
func (m *Mission) QueryEmployee() *EmployeeQuery {
	return NewMissionClient(m.config).QueryEmployee(m)
}

// QueryProject queries the "project" edge of the Mission entity.
func (m *Mission) QueryProject() *ProjectQuery {
	return NewMissionClient(m.config).QueryProject(m)
}

// Update returns a builder for updating this Mission.
// Note that you need to call Mission.Unwrap() before calling this method if this Mission
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mission) Update() *MissionUpdateOne {
	return NewMissionClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Mission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mission) Unwrap() *Mission {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mission is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mission) String() string {
	var builder strings.Builder
	builder.WriteString("Mission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("purpose=")
	builder.WriteString(m.Purpose)
	builder.WriteString(", ")
	builder.WriteString("destination=")
	builder.WriteString(m.Destination)
	builder.WriteString(", ")
	builder.WriteString("startDate=")
	builder.WriteString(m.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("endDate=")
	builder.WriteString(m.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("transport=")
	builder.WriteString(m.Transport)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Missions is a parsable slice of Mission.
type Missions []*Mission
