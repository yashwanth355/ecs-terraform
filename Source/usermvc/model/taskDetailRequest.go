package model

import "time"

type TaskDetailRequest struct {
	TaskId            string    `json:"TaskId"`
	TaskIdsNo         int       `json:"TaskIdsNo"`
	ProjectId         string    `json:"ProjectId"`
	TaskName          string    `json:"TaskName"`
	TaskStart         time.Time `json:"TaskStart"`
	CloseStatus       bool      `json:"CloseStatus"`
	TaskDescription   string    `json:"TaskDescription"`
	AssignedTo        string    `json:"AssignedTo"`
	Status            string    `json:"Status"`
	CancelStatus      bool      `json:"CancelStatus"`
	StartingDate      time.Time `json:"StartingDate"`
	EndingDate        time.Time `json:"EndingDate"`
	TaskStartDatetime time.Time `json:"TaskStartDatetime"`
	CustomerId        string    `json:"CustomerId"`
	OriginId          string    `json:"OriginId"`
}
