package model

type GCDetails struct {
	Create bool `json:"create"`

	Update                bool   `json:"update"`
	GroupId               string `json:"group_id"`
	ItemCode              string `json:"item_code"`
	SCode                 string `json:"s_code"`
	ItemName              string `json:"item_name"`
	ItemDesc              string `json:"item_desc"`
	HsnCode               string `json:"hsn_code"`
	ConvertionRatio       string `json:"convertion_ratio"`
	ItemCatId             string `json:"item_catid"`
	Uom                   string `json:"uom"`
	ShowStock             bool   `json:"show_stock"`
	EnableStatus          bool   `json:"enable_status"`
	IsRawMaterial         bool   `json:"is_rawmeterial"`
	DisplayInPo           bool   `json:"display_inpo"`
	DisplayInDailyUpdates bool   `json:"display_in_dailyupdates"`
	IsSpecialCoffee       bool   `json:"is_specialcoffee"`
	CoffeeType            string `json:"coffee_type"`
	CreatedOn             string `json:"created_on"`
	CreatedBy             string `json:"created_by"`
	UpdatedOn             string `json:"updated_on"`
	UpdatedBy             string `json:"updated_by"`
	LCode                 string `json:"lcode"`
	LName                 string `json:"lname"`
	LGroupCode            string `json:"lgroupcode"`
	Itemid                string `json:"item_id"`
	Itemidsno             int    `json:"itemidsno"`
	CategoryType          string `json:"cat_type"`
	//Special composition Info Section---------------------------
	Density       int `json:"density"`
	Moisture      int `json:"moisture"`
	Browns        int `json:"browns"`
	Blacks        int `json:"blacks"`
	BrokenBits    int `json:"broken_bits"`
	InsectedBeans int `json:"insected_beans"`
	Bleached      int `json:"bleached"`
	Husk          int `json:"husk"`
	Sticks        int `json:"sticks"`
	Stones        int `json:"stones"`
	BeansRetained int `json:"beans_retained"`
}
