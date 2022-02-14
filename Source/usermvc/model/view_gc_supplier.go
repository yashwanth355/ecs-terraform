package model

type GCViewDetails struct {
	GroupId               string `json:"group_id"`
	ItemCode              string `json:"item_code"`
	SCode                 string `json:"s_code"`
	ItemName              string `json:"item_name"`
	ItemDesc              string `json:"item_desc"`
	HsnCode               string `json:"hsn_code"`
	ConvertionRatio       string `json:"convertion_ratio"`
	ItemCatId             string `json:"item_catid"`
	ItemCatName           string `json:"item_catname"`
	Uom                   string `json:"uom"`
	UomName               string `json:"uom_name"`
	DisplayInPo           bool   `json:"display_inpo"`
	DisplayInDailyUpdates bool   `json:"display_in_dailyupdates"`
	IsSpecialCoffee       bool   `json:"is_specialcoffee"`
	CategoryType          string `json:"cat_type"`
	CoffeeType            string `json:"coffee_type"`
	LName                 string `json:"lname"`
	LGroupCode            string `json:"lgroupcode"`
	//AuditLogDetails       []AuditLogGC        `json:"audit_log_gc"`
	StockLocation []ItemStockLocation `json:"item_stock_location"`
	VendorList    []VendorList        `json:"vendor_list"`

	//Special composition Info Section---------------------------
	Density       int    `json:"density"`
	Moisture      int    `json:"moisture"`
	Browns        int    `json:"browns"`
	Blacks        int    `json:"blacks"`
	BrokenBits    int    `json:"broken_bits"`
	InsectedBeans int    `json:"insected_beans"`
	Bleached      int    `json:"bleached"`
	Husk          int    `json:"husk"`
	Sticks        int    `json:"sticks"`
	Stones        int    `json:"stones"`
	BeansRetained int    `json:"beans_retained"`
	Itemid        string `json:"item_id"`
}

type ItemStockLocation struct {
	Entity    string `json:"entity"`
	Name      string `json:"name"`
	Quantity  string `json:"quantity"`
	Value     string `json:"value"`
	UnitPrice string `json:"unit_price"`
}

type VendorList struct {
	VendorName  string `json:"vendor_name"`
	ContactName string `json:"contact_name"`
	State       string `json:"state"`
	Country     string `json:"country"`
	City        string `json:"city"`
}
