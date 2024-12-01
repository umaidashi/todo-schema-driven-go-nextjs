package schema

import (
	"server/domain/model"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description"),
		field.Time("start_at"),
		field.Time("end_at"),
		field.Enum("priority").Values(model.PRIORITY_NAMES...),
		field.Enum("status").Values(model.TODO_STATUS_NAMES...),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("labels", Label.Type),
	}
}
