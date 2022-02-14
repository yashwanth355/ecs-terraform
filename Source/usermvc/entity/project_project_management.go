package entity

import "time"

type ProjectMangementMaster struct {
	ProjectId          string    `gorm:"column:projectid"`
	ProjectIdsNo       int       `gorm:"column:projectidsno"`
	ProjectName        string    `gorm:"column:projectname"`
	ProjectOwner       string    `gorm:"column:project_owner"`
	ProjectDate        time.Time `gorm:"column:projectdate"`
	DueDate            time.Time `gorm:"column:duedate"`
	ProjectDescription string    `gorm:"column:projectdesc"`
	DepId              string    `gorm:"column:depid"`
	FinancialYearId    string    `gorm:"column:financialyearid"`
	ProjectType        string    `gorm:"column:project_type"`
}

type ProjectManagementDetails struct {
	ProjectId   string `gorm:"column:projectid"`
	ProjectName string `gorm:"column:projectname"`
}
