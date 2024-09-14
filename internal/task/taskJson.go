package task

import "time"

var TimeNow = time.Now().Format("02 Jan 2006 15:04:05")

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	Id          int    `json:"Id"`
	Description string `json:"Description"`
	Status      Status `json:"Status"`
	CreatedAt   string `json:"Created_at"`
	UpdatedAt   string `json:"Updated_at"`
}
