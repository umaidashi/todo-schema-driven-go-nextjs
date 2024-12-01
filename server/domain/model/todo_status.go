package model

import (
	"github.com/samber/lo"
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

var TodoStatusInProgress = TodoStatus{
	Name:    "IN_PROGRESS",
	Label:   "進行中",
	Color:   "#ff0000",
	BgColor: "#ffcccc",
}

var TodoStatusDone = TodoStatus{
	Name:    "DONE",
	Label:   "完了",
	Color:   "#ff0000",
	BgColor: "#ffcccc",
}

var TODO_STATUSES = []TodoStatus{
	TodoStatusPending,
	TodoStatusInProgress,
	TodoStatusDone,
}

var TODO_STATUS_NAMES = lo.Map(TODO_STATUSES, func(s TodoStatus, _ int) string {
	return s.Name
})

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
