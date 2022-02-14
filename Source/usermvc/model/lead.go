package model

type Salutations struct {
	Salutationid int64  `json:"id"`
	Salutation   string `json:"salutation"`
}

type AccountsInformation struct {
	Accounttypeid string `json:"id"`
	Accounttype   string `json:"accounttype"`
}

type CoffeeTypes struct {
	CoffeeType   string `json:"coffeetype"`
	CoffeeTypeId string `json:"id"`
}

type ProductSegments struct {
	Productsegmentid int    `json:"id"`
	Productsegment   string `json:"productsegment"`
}