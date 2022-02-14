package model

type DepartmentNames struct {
	Deptname string `json:"deptname"`
}
type GetAllDepartmentResponse struct {
	Status  int
	Payload []*DepartmentNames
}
