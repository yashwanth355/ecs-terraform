package model

type GetLeadInfoReq struct {
	GetInfo string `json:"getinfo"`
	UserId  string `json:"userid"`
}

type ProvideLeadsInfoReqContext struct {
	Filter string `json:"filter"`
	UserId string `json:"loggedinuserid"`
}
