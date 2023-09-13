package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("purpose").
			Optional(),
		field.String("destination").
			NotEmpty(),
		field.Time("startDate").
			Optional(),
		field.Time("endDate").
			Optional(),
		field.String("transport"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).
			Ref("missions").
			Unique(),
		edge.From("project", Project.Type).
			Ref("missions").
			Unique(),
	}
}
