package model

import (
	"github.com/samber/lo"
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
	PriorityMiddle,
	PriorityLow,
}

var PRIORITY_NAMES = lo.Map(PRIORITIES, func(p Priority, _ int) string {
	return p.Name
})

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
