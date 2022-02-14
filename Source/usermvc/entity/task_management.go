package entity

import "time"

type TaskMangementMaster struct {
	TaskId            string    `gorm:"column:taskid"`
	TaskIdsNo         int       `gorm:"column:taskidsno"`
	ProjectId         string    `gorm:"column:projectid"`
	TaskName          string    `gorm:"column:taskname"`
	TaskStart         time.Time `gorm:"column:taskstart"`
	CloseStatus       bool      `gorm:"column:closestatus"`
	TaskDescription   string    `gorm:"column:taskdesc"`
	AssignedTo        string    `gorm:"column:assignedto"`
	Status            string    `gorm:"column:status"`
	CancelStatus      bool      `gorm:"column:cancelstatus"`
	StartingDate      time.Time `gorm:"column:sdate"`
	EndingDate        time.Time `gorm:"column:edate"`
	TaskStartDatetime time.Time `gorm:"column:taskstartdatetime"`
	CustomerId        string    `gorm:"column:custid"`
	OriginId          string    `gorm:"column:originid"`
}

type ListTaskDetails struct {
	TaskId       string    `gorm:"column:taskid"`
	OriginId     string    `gorm:"column:originid"`
	TaskName     string    `gorm:"column:taskname"`
	StartingDate time.Time `gorm:"column:sdate"`
	Status       string    `gorm:"column:status"`
	AssignedTo   string    `gorm:"column:assignedto"`
}
