package task

import "time"

type (
	Task struct {
		Name      string        `json:"name"`
		Status    string        `json:"status"`
		Value     int           `json:"value"`
		Duration  time.Duration `json:"duration"`
		CreatedAt time.Time     `json:"createdAt"`
	} // @name Task

	Create struct {
		Name  string `json:"name" validate:"required"`
		Value int    `json:"value" validate:"required"`
	} // @name Create
)
