package model

type ListGCDetails struct {
	ItemId       string `json:"item_id"`
	ItemCode     string `json:"item_code"`
	SCode        string `json:"s_code"`
	ItemName     string `json:"item_name"`
	GroupId      string `json:"group_id"`
	Groupname    string `json:"group_name"`
	Uom          string `json:"uom"`
	CategoryName string `json:"category_name"`
	DisplayInPo  bool   `json:"display_inpo"`
	CoffeeType   string `json:"coffee_type"`
}

type InputG struct {
	Type           string `json:"type"`
	GroupId        string `json:"group_id"`
	AdvancedFilter bool   `json:"deep_filter"`
	FilterParam    string `json:"deep_filter_args"`
}
