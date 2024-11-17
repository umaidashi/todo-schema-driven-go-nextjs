// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/pkg/ent/label"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Label is the model entity for the Label schema.
type Label struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LabelQuery when eager-loading is set.
	Edges        LabelEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LabelEdges holds the relations/edges for other nodes in the graph.
type LabelEdges struct {
	// Todos holds the value of the todos edge.
	Todos []*Todo `json:"todos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodosOrErr returns the Todos value or an error if the edge
// was not loaded in eager-loading.
func (e LabelEdges) TodosOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todos, nil
	}
	return nil, &NotLoadedError{edge: "todos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Label) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case label.FieldID:
			values[i] = new(sql.NullInt64)
		case label.FieldTitle:
			values[i] = new(sql.NullString)
		case label.FieldCreatedAt, label.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Label fields.
func (l *Label) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case label.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case label.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				l.Title = value.String
			}
		case label.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case label.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Label.
// This includes values selected through modifiers, order, etc.
func (l *Label) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryTodos queries the "todos" edge of the Label entity.
func (l *Label) QueryTodos() *TodoQuery {
	return NewLabelClient(l.config).QueryTodos(l)
}

// Update returns a builder for updating this Label.
// Note that you need to call Label.Unwrap() before calling this method if this Label
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Label) Update() *LabelUpdateOne {
	return NewLabelClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Label entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Label) Unwrap() *Label {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Label is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Label) String() string {
	var builder strings.Builder
	builder.WriteString("Label(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("title=")
	builder.WriteString(l.Title)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Labels is a parsable slice of Label.
type Labels []*Label
