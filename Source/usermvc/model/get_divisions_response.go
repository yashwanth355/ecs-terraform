package model

type Divisions struct {
	Divmaster string `json:"division"`
}

type GetDivisionsResponse struct {
	Status  int
	Payload []*Divisions
}
