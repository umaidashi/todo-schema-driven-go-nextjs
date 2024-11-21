package model

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/samber/lo"

	"entgo.io/ent/schema/field"
)

type TodoStatus struct {
	Name    string `json:"name"` // DB に保存される値
	Label   string `json:"label"`
	Color   string `json:"color"`
	BgColor string `json:"bgColor"`
}

var TodoStatusPending = TodoStatus{
	Name:    "PENDING",
	Label:   "未着手",
	Color:   "#ff0000",
	BgColor: "#ffcccc",
}

var TODO_STATUSES = []TodoStatus{
	TodoStatusPending,
}

func (t TodoStatus) Value(s TodoStatus) (driver.Value, error) {
	return s.Name, nil
}

func (TodoStatus) ScanValue() field.ValueScanner {
	return &sql.NullString{}
}

func (t TodoStatus) FromValue(v driver.Value) (TodoStatus, error) {
	if str, ok := v.(string); ok {
		return TodoStatusOf(str), nil
	}
	return TodoStatus{}, fmt.Errorf("unexpected type %T", v)
}

var todoStatusMap = lo.KeyBy(TODO_STATUSES, func(s TodoStatus) string {
	return s.Name
})

func TodoStatusOf(name string) TodoStatus {
	v, ok := todoStatusMap[name]
	if ok {
		return v
	}
	panic("status is invalid. name = " + name)
}
