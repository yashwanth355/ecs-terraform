package model

import "time"

type ProjectDetailRequest struct {
	ProjectId          string    `json:"ProjectId"`
	ProjectIdsNo       int       `json:"ProjectIdsNo"`
	ProjectName        string    `json:"ProjectName"`
	ProjectOwner       string    `json:"ProjectOwner"`
	ProjectDate        time.Time `json:"ProjectDate"`
	DueDate            time.Time `json:"DueDate"`
	ProjectDescription string    `json:"ProjectDescription"`
	DepId              string    `json:"DepId"`
	FinancialYearId    string    `json:"FinancialYearId"`
	ProjectType        string    `json:"ProjectType"`
}
