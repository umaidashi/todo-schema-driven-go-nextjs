package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
  "entgo.io/ent/schema/edge"
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
		field.Enum("priority").Values("low", "medium", "high"),
		field.String("status").Default("pending"), // TODO: GoType(model.TodoStatus)
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
    edge.To("labels", Label.Type),
  }
}
