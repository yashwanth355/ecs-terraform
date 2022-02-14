package model

type GetPoFormInfoRequestBody struct {
	Type             string         `json:"type"`
	SupplierName     string         `json:"supplier_name"`
	SupplierID       string         `json:"supplier_id"`
	SupplierType     string         `json:"supplier_type"`
	SupplierCoun     string         `json:"supplier_country"`
	SupplierAddress  string         `json:"supplier_address"`
	POCreatedAt      []POCreatedAt  `json:"po_created_at"`
	POCreatedFor     []POCreatedFor `json:"po_created_for"`
	GreenCoffee      []GreenCoffee  `json:"green_coffee_types"`
	SupplierTypeID   string         `json:"supplier_type_id"`
	SupplierTypeName string         `json:"supplier_type_name"`
	ItemID           string         `json:"item_id"`
}

type POCreatedAt struct {
	POTypeID   int    `json:"billing_at_id"`
	POTypeName string `json:"billing_at_name"`
	POAddress  string `json:"billing_at_address"`
}
type POCreatedFor struct {
	POTypeID   int    `json:"delivery_at_id"`
	POTypeName string `json:"delivery_at_name"`
	POAddress  string `json:"delivery_at_address"`
}
type GreenCoffee struct {
	ItemID       string `json:"item_id"`
	ItemName     string `json:"item_name"`
	GCCoffeeType string `json:"gc_type"`
	// GreenCoffee 		[]GreenCoffee `json:"green_coffee_types"`

	Density       int `json:"density"`
	Moisture      int `json:"moisture"`
	Browns        int `json:"browns"`
	Blacks        int `json:"blacks"`
	BrokenBits    int `json:"brokenbits"`
	InsectedBeans int `json:"insectedbeans"`
	Bleached      int `json:"bleached"`
	Husk          int `json:"husk"`
	Sticks        int `json:"sticks"`
	Stones        int `json:"stones"`
	BeansRetained int `json:"beansretained"`
}
type ContainerTypesList struct {
	ConttypeId   string `json:"conttype_id"`
	ConttypeName string `json:"conttype_name"`
}