package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
  "entgo.io/ent/schema/edge"
)

type Label struct {
	ent.Schema
}

func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (Label) Edges() []ent.Edge {
	return []ent.Edge{
    edge.From("todos", Todo.Type).
      Ref("labels"),
  }
}
