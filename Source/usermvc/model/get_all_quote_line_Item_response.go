package model

type GetAllQouteLineItemsResponseBody struct {
	Quoteitemid        string `json:"quotelineitem_id"`
	Quoteid            int    `json:"quote_id"`
	Sampleid           string `json:"sample_code"`
	Expectedorderqty   int    `json:"expectedorder_kgs"`
	Categoryname       string `json:"category"`
	Packcategorytypeid *int   `json:"categorytype_id"`
	Packweightid       *int   `json:"weight_id"`
	CategoryType       string `json:"categorytype"`
	Weight             string `json:"weight"`
	Packupcid          int    `json:"upc_id"`
}

type GetAllQouteLineItemsResponse struct {
	Status  int
	Payload []GetAllQouteLineItemsResponseBody
}

type Categorytypename struct {
	Categorytypename string `json:"categorytypename"`
}
