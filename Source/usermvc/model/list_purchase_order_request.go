package model

import (
	"database/sql"
)

type ListPurchaseOrderRequest struct {
	Type string `json:"type"`
}

type PurchaseOrderDetailsT struct {
	PoNo           string         `json:"po_no"`
	PoDate         string         `json:"po_date"`
	ApprovalStatus string         `json:"approval_status"`
	Vendor         string         `json:"vendor_name"`
	VendorTypeId   string         `json:"vendor_typeid"`
	QuotNo         sql.NullString `json:"quot_no"`
	Category       string         `json:"category"`
	Currency       string         `json:"currency"`
	PoValue        sql.NullString `json:"po_value"`
	TaxValue       string         `json:"tax_value"`
	Advance        string         `json:"advance"`
	PoQty          string         `json:"po_qty"`
	MrinQty        string         `json:"mrin_qty"`
	Balance        string         `json:"balance"`
}
