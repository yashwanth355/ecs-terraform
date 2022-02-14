package model

type UserRole struct {
	userName string `json:"user_name"`
}
type GetAllUser struct {
	Status  int
	Payload []*UserRole
}
