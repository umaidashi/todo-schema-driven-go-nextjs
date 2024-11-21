package model

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/samber/lo"

	"entgo.io/ent/schema/field"
)

type Priority struct {
	Name    string `json:"name"` // DB に保存される値
	Label   string `json:"label"`
	Color   string `json:"color"`
	BgColor string `json:"bgColor"`
}

var PriorityHigh = Priority{
	Name:    "HIGH",
	Label:   "High",
	Color:   "#FF0000",
	BgColor: "#FF0000",
}

var PriorityMiddle = Priority{
	Name:    "MIDDLE",
	Label:   "Middle",
	Color:   "#FF0000",
	BgColor: "#FF0000",
}

var PriorityLow = Priority{
	Name:    "LOW",
	Label:   "Low",
	Color:   "#FF0000",
	BgColor: "#FF0000",
}

var PRIORITIES = []Priority{
	PriorityHigh,
}

func (p Priority) Value(p2 Priority) (driver.Value, error) {
	return p2.Name, nil
}

func (Priority) ScanValue() field.ValueScanner {
	return &sql.NullString{}
}

func (t Priority) FromValue(v driver.Value) (Priority, error) {
	if str, ok := v.(string); ok {
		return PriorityOf(str), nil
	}
	return Priority{}, fmt.Errorf("unexpected type %T", v)
}

var priorityMap = lo.KeyBy(PRIORITIES, func(p Priority) string {
	return p.Name
})

func PriorityOf(name string) Priority {
	v, ok := priorityMap[name]
	if ok {
		return v
	}
	panic("priority is invalid. name = " + name)
}
