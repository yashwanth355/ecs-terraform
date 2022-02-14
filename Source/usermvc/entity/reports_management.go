package entity

type ConfirmedOrderReport struct {
	SerialNo     int    `gorm:"column:serialno"`
	ContactNo    string `gorm:"column:contactno"`
	CustomerName string `gorm:"column:customername"`
	MarketingRep string `gorm:"column:marketingrep"`
	ProductCode  string `gorm:"column:productcode"`
	SampleCode   string `gorm:"column:samplecode"`
	PackingType  string `gorm:"column:packingtype"`
	BrandName    string `gorm:"column:brandname"`
	Destination  string `gorm:"column:destination"`
	InCoterms    string `gorm:"column:incoterms"`
	Quantity     string `gorm:"column:quantity"`
	Price        string `gorm:"column:price"`
	Value        string `gorm:"column:value"`
}
