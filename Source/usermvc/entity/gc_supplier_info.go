package entity

type GCDetails struct {
	Create                bool   `gorm:"column:create"`
	Update                bool   `gorm:"column:update"`
	GroupId               string `gorm:"column:group_id"`
	ItemCode              string `gorm:"column:item_code"`
	SCode                 string `gorm:"column:s_code"`
	ItemName              string `gorm:"column:item_name"`
	ItemDesc              string `gorm:"column:item_desc"`
	HsnCode               string `gorm:"column:hsn_code"`
	ConvertionRatio       string `gorm:"column:convertion_ratio"`
	ItemCatId             string `gorm:"column:item_catid"`
	Uom                   string `gorm:"column:uom"`
	ShowStock             bool   `gorm:"column:show_stock"`
	EnableStatus          bool   `gorm:"column:enable_status"`
	IsRawMaterial         bool   `gorm:"column:is_rawmeterial"`
	DisplayInPo           bool   `gorm:"column:display_inpo"`
	DisplayInDailyUpdates bool   `gorm:"column:display_in_dailyupdates"`
	IsSpecialCoffee       bool   `gorm:"column:is_specialcoffee"`
	CoffeeType            string `gorm:"column:coffee_type"`
	CreatedOn             string `gorm:"column:created_on"`
	CreatedBy             string `gorm:"column:created_by"`
	UpdatedOn             string `gorm:"column:updated_on"`
	UpdatedBy             string `gorm:"column:updated_by"`
	LCode                 string `gorm:"column:lcode"`
	LName                 string `gorm:"column:lname"`
	LGroupCode            string `gorm:"column:lgroupcode"`
	Itemid                string `gorm:"column:item_id"`
	Itemidsno             int    `gorm:"column:itemidsno"`
	CategoryType          string `gorm:"column:cat_type"`
	//Special composition Info Section---------------------------
	Density       int `gorm:"column:density"`
	Moisture      int `gorm:"column:moisture"`
	Browns        int `gorm:"column:browns"`
	Blacks        int `gorm:"column:blacks"`
	BrokenBits    int `gorm:"column:broken_bits"`
	InsectedBeans int `gorm:"column:insected_beans"`
	Bleached      int `gorm:"column:bleached"`
	Husk          int `gorm:"column:husk"`
	Sticks        int `gorm:"column:sticks"`
	Stones        int `gorm:"column:stones"`
	BeansRetained int `gorm:"column:beans_retained"`
}
