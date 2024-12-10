package model

import "time"

type Todo struct {
	ID          TodoId     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartAt     time.Time  `json:"startAt"`
	EndAt       *time.Time `json:"endAt"`
	Priority    Priority   `json:"priority"`
	Status      TodoStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TodoId int
