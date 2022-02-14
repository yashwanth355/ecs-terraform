package model

type DesignationName struct {
	DesgName string `json:"desgName"`
}

type GetAllDesignationNameResponse struct {
	Status  int
	Payload []*DesignationName
}
