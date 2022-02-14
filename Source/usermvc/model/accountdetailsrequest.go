package model

type AccountDetailsRequest struct {
	LeadId               int    `json:"leadid"`
	Role                 string `json:"role"`
	ConvertLeadToAccount bool   `json:"convertleadtoaccount"`
	Approve              bool   `json:"approve"`
	Reject               bool   `json:"reject"`
	Comments             string `json:"comments"`
}
