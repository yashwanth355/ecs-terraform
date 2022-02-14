package porepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	
	// "database/sql"
	// "reflect"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"
)

type PoRepo interface {
	GetPoCreationInfo(ctx context.Context, req *model.Input) (interface{}, error)
	GetPOFormInfo(ctx context.Context, req *model.GetPoFormInfoRequestBody) (interface{}, error)
	ViewPoDetails(ctx context.Context, req *model.PurchaseOrderDetails) (interface{}, error)
	// ListPurchaseOrders(ctx context.Context, req *model.ListPurchaseOrderRequest) (interface{}, error)
	GetPortandOrigin(ctx context.Context, req *model.Input) (interface{}, error)
	GetBalQuoteQtyForPoOrder(ctx context.Context, req *model.Input) (interface{}, error)
}

type poRepo struct {
	db *gorm.DB
}

func NewPoRepo() PoRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}
	newDb.AutoMigrate(&entity.User{})
	return &poRepo{
		db: newDb,
	}
}
func (po poRepo) GetBalQuoteQtyForPoOrder(ctx context.Context, req *model.Input) (interface{}, error) {
	
	logger := logger2.GetLoggerWithContext(ctx)
	if req.Type == "getBalqtyforPo" {
		sqlGetPOBalQty := `select sum(total_quantity) from dbo.pur_gc_po_con_master_newpg where quote_no=$1`
		rows, err := po.db.Raw(sqlGetPOBalQty, req.QuotationId).Rows()
		if err != nil {
			logger.Error("error while fetching records from dbo.pur_gc_po_con_master_newpg ", err.Error())
		}
		var oQPO []model.QtyforPo
		defer rows.Close()
		for rows.Next() {
			var oQ model.QtyforPo
			err = rows.Scan(&oQ.OrderQty)						
		}
		return oQPO, err
		
	}
	return nil,nil
}
func (po poRepo) GetPoCreationInfo(ctx context.Context, req *model.Input) (interface{}, error) {
	const (
		CONTAINERTYPES = "containerTypes"
		)
	logger := logger2.GetLoggerWithContext(ctx)
	if req.Type == CONTAINERTYPES {
		rows, err := po.db.Raw("select conttypeid, conttypename from dbo.sales_container_types").Rows()
		if err != nil {
			logger.Error("error while fetching records from dbo.sales_container_types ", err.Error())
		}
		var allContainerTypes []model.ContainerTypesList
		defer rows.Close()
		for rows.Next() {
			var ct model.ContainerTypesList
			err = rows.Scan(&ct.ConttypeId, &ct.ConttypeName)
			allContainerTypes = append(allContainerTypes, ct)
			
		}
		return allContainerTypes, err
		
	}
	return nil,nil
}

func (po poRepo) GetPOFormInfo(ctx context.Context, req *model.GetPoFormInfoRequestBody) (interface{}, error) {
	const (
		POSUBCATEGORY = "posubcategory"
		SUPPLIERINFO  = "supplierinfo"
		BILLINGINFO   = "billinginfo"
		DELIVERYINFO  = "deliveryinfo"
		ALLSUPPLIERS  = "allsuppliers"
		GREENCOFFEE   = "greencoffee"
		GCCOMPOSITION = "gccomposition"
	)
	logger := logger2.GetLoggerWithContext(ctx)
	if req.Type == POSUBCATEGORY {
		rows, err := po.db.Raw("select vendortypeid,initcap(vendortypename) from dbo.pur_vendor_types").Rows()
		if err != nil {
			logger.Error("error while getting records from dbo.pur_vendor_types ", err.Error())
		}
		var allPOSubs []model.GetPoFormInfoRequestBody
		defer rows.Close()
		for rows.Next() {
			var pos model.GetPoFormInfoRequestBody
			err = rows.Scan(&pos.SupplierTypeID, &pos.SupplierTypeName)
			allPOSubs = append(allPOSubs, pos)
			return allPOSubs, nil
		}
	}

	if req.Type == SUPPLIERINFO {
		sqlStatement2 := `select country,vendorid,initcap(vendorname),initcap(address1)||','||initcap(address2)||','||initcap(city)||','||pincode||','||initcap(state)||','||'Phone:'||phone||','||'Mobile:'||mobile||','||'GST NO:'||gstin address 
							from dbo.pur_vendor_master where vendorid=$1`

		rows, err := po.db.Raw(sqlStatement2, req.SupplierID).Rows()
		if err != nil {
			logger.Error("error while getting records from dbo.pur_vendor_master  ", err.Error())
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&req.SupplierCoun, &req.SupplierID, &req.SupplierName, &req.SupplierAddress)
		}

		if req.SupplierCoun != "INDIA" {
			req.SupplierType = "International"
		} else {
			req.SupplierType = "Domestic"
			return req, nil
		}
	}
	if req.Type == BILLINGINFO {
		logger.Info("getting billing info details")
		sqlStatement := `select potypeid,initcap(potypename),initcap(potypefullname)||','||initcap(address) as fulladdress from dbo.pur_po_types`

		rows, err := po.db.Raw(sqlStatement).Rows()
		if err != nil {
			logger.Error("error while getting records from dbo.pur_vendor_master  ", err.Error())
			return nil, err
		}
		defer rows.Close()
		var allPCA []model.POCreatedAt
		defer rows.Close()
		for rows.Next() {
			var pca model.POCreatedAt
			err = rows.Scan(&pca.POTypeID, &pca.POTypeName, &pca.POAddress)
			allPCA = append(allPCA, pca)
		}
		return allPCA, nil
	}
	if req.Type == DELIVERYINFO {
		logger.Info("get deliveryinfo details:")
		sqlStatement := `select potypeid,initcap(potypename),initcap(potypefullname)||','||initcap(address) as fulladdress from dbo.pur_po_types`
		rows, err := po.db.Raw(sqlStatement).Rows()
		if err != nil {
			logger.Error("error while getting  deliveryinfo records from dbo.dbo.pur_po_types  ", err.Error())
			return nil, err
		}
		logger.Info("Query executed")
		var allPCF []model.POCreatedFor
		defer rows.Close()
		for rows.Next() {
			var pcf model.POCreatedFor
			err = rows.Scan(&pcf.POTypeID, &pcf.POTypeName, &pcf.POAddress)
			allPCF = append(allPCF, pcf)
		}
		return allPCF, nil
	}
	if req.Type == ALLSUPPLIERS {
		logger.Info("get vendors", req.Type)
		logger.Info("get vendors based on supplierID", req.SupplierTypeID)
		sqlStatementSup := `select vendorid,initcap(vendorname) from dbo.pur_vendor_master 
							where ((groupid='FAC-2') or (groupid='FAC-3') or (groupid='FAC-4') or (groupid='FAC-9')) and vendortypeid=$1`
		rows, err := po.db.Raw(sqlStatementSup, &req.SupplierTypeID).Rows()
		if err != nil {
			logger.Error("error while getting  deliveryinfo records from dbo.dbo.pur_po_types  ", err.Error())
			return nil, err
		}
		logger.Info("Query executed")
		var allSuppliers []model.GetPoFormInfoRequestBody
		defer rows.Close()
		for rows.Next() {
			var a model.GetPoFormInfoRequestBody
			err = rows.Scan(&a.SupplierID, &a.SupplierName)
			allSuppliers = append(allSuppliers, a)
		}
		return allSuppliers, nil
	}
	if req.Type == GREENCOFFEE {
		logger.Info("get GC", req.Type)
		sqlStatement4 := `select itemid,initcap(itemname),cat_type from dbo.inv_gc_item_master`
		rows, err := po.db.Raw(sqlStatement4).Rows()
		if err != nil {
			if err != nil {
				logger.Error("error while getting  GC records from dbo.dbo.pur_po_types  ", err.Error())
				return nil, err
			}
		}
		logger.Info("GC Query executed")
		var allGc []model.GreenCoffee
		defer rows.Close()
		for rows.Next() {
			var gc model.GreenCoffee
			err = rows.Scan(&gc.ItemID, &gc.ItemName, &gc.GCCoffeeType)
			allGc = append(allGc, gc)
		}

		return allGc, err
	}

	if req.Type == GCCOMPOSITION {
		logger.Info("get GC new composition based on the GD ID", req.Type)
		logger.Info("Entered Item id is:", &req.ItemID)
		sqlStatement5 := `select itemid,density,moisture,browns,blacks,brokenbits,insectedbeans,bleached,husk,sticks,stones,beansretained from dbo.pur_gc_po_composition_master_newpg
							where itemid=$1`
		rows, err := po.db.Raw(sqlStatement5, &req.ItemID).Rows()
		if err != nil {
			if err != nil {
				logger.Error("error while getting  getting   GC new composition based on the GD I ", err.Error())
				return nil, err
			}
		}
		logger.Info("GC Query executed")
		var allGcComp []model.GreenCoffee
		defer rows.Close()
		for rows.Next() {
			var gc model.GreenCoffee
			err = rows.Scan(&gc.ItemID, &gc.Density, &gc.Moisture, &gc.Browns, &gc.Blacks, &gc.BrokenBits, &gc.InsectedBeans,
				&gc.Bleached, &gc.Husk, &gc.Sticks, &gc.Stones, &gc.BeansRetained)
			allGcComp = append(allGcComp, gc)
		}
		logger.Info(allGcComp)
		return allGcComp, nil
	}
	errInfo := fmt.Sprintf("req type should be either of  %s %s %s %s %s %s  %s", POSUBCATEGORY, SUPPLIERINFO, BILLINGINFO, DELIVERYINFO, ALLSUPPLIERS, GREENCOFFEE, GCCOMPOSITION)
	return nil, errors.New(errInfo)
}

func (po poRepo) GetPortandOrigin(ctx context.Context, req *model.Input) (interface{}, error) {
	const (
		ORIGIN   = "originDetails"
		PORTLOAD = "portLoadingDetails"
	)
	logger := logger2.GetLoggerWithContext(ctx)
	if req.Type == ORIGIN {
		var origins []model.Origin
		rows, err := po.db.Raw("select distinct initcap(origin) from dbo.pur_gc_contract_master where origin != '';").Rows()
		if err != nil {
			logger.Error("error while getting record", err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var origin model.Origin
			err = rows.Scan(&origin.Origin)
			origins = append(origins, origin)
		}
		return origins, nil
	}
	if req.Type == PORTLOAD {
		var items []model.PortLoading
		rows, err := po.db.Raw("select distinct initcap(poloading) from dbo.pur_gc_contract_master where poloading != '';").Rows()
		if err != nil {
			logger.Error("error while getting record", err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var portsload model.PortLoading
			err = rows.Scan(&portsload.Port)
			items = append(items, portsload)
		}
		return items, nil
	}
	return nil, nil
}

func (po poRepo) ViewPoDetails(ctx context.Context, req *model.PurchaseOrderDetails) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	var viewPOResp []model.PurchaseOrderDetails
	if req.PoNO != "" {
		logger.Error("Entered PO View Module")
		logger.Error("selected PO NO:", req.PoNO)
		//check if po is import or domestic
		sqlStatementIDC1 := `SELECT posubcat FROM dbo.pur_gc_po_con_master_newpg
						where pono=$1`
		rows, err := po.db.Raw(sqlStatementIDC1, req.PoNO).Rows()
		
		if err != nil {
			logger.Error("Fetching PO Details from DB failed")
			logger.Error(err.Error())
		}
		// defer rows.Close()
		for rows.Next() {
			var cat model.PurchaseOrderDetails
			err = rows.Scan(&cat.POSubCategory)
			logger.Error("Scanned pocat is",cat)
			viewPOResp=append(viewPOResp,cat)
			logger.Error(viewPOResp)
		}
		if req.POSubCategory == "Import" {
			req.SupplierType = "Import"
			sqlStatementPOV1 := `SELECT total_quantity,cid,poid, podate, pocat,vendorid,itemid,
						billing_at_id, delivery_at_id,currencyid,status,dispatchterms, origin,
						poloading, insurance, destination, forwarding,nocontainers,container_type,
						payment_terms,remarks,taxes_duties, transport_mode, transit_insurence, packing_forward,
						othercharges,rate,noofbags,netweight,
						purchase_type, terminal_month, booked_term_rate,booked_differential, fixed_term_rate, fixed_differential,
						purchase_price, market_price, po_margin, total_price,gross_price,fixationdate,quantity_mt
						FROM dbo.pur_gc_po_con_master_newpg
						where pono=$1`
			rows, err := po.db.Raw(sqlStatementPOV1, req.PoNO).Rows()
			logger.Error("PO Master Query Executed")
			if err!= nil {
				logger.Error("Fetching PO Details from DB failed")
				logger.Error(err.Error())
			}
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&req.TotalQuantity, &req.Contract, &req.PoId, &req.PoDate, &req.POCategory,
					&req.SupplierID, &req.ItemID, &req.POBillTypeID, &req.PODelTypeID, &req.CurrencyID, &req.Status,
					&req.IncoTermsID, &req.Origin, &req.PortOfLoad, &req.Insurance, &req.PlaceOfDestination, &req.Forwarding, &req.NoOfContainers,
					&req.ContainerType, &req.PaymentTerms, &req.Comments, &req.TaxDuties, &req.ModeOfTransport, &req.TransitInsurance,
					&req.PackForward, &req.OtherCharges, &req.Rate, &req.NoOfBags, &req.NetWt,
					&req.PurchaseType, &req.TerminalMonth, &req.BookedTerminalRate, &req.BookedDifferential, &req.FixedTerminalRate, &req.FixedDifferential, &req.PurchasePrice, &req.MarketPrice,
					&req.POMargin, &req.TotalPrice, &req.GrossPrice, &req.FixationDate, &req.MTQuantity)

			}	
				
			//Fetch Incoterms details:
			if req.IncoTermsID != "" {
				logger.Error("get incoterms for id :", req.IncoTermsID)
				sqlStatementIT1 := `SELECT incoterms FROM dbo.cms_incoterms_master where incotermsid=$1`
				rows, err := po.db.Raw(sqlStatementIT1, req.IncoTermsID).Rows()
				if err != nil {
					logger.Error("Fetching Incoterms Details from DB failed")

				}
				defer rows.Close()
				for rows.Next() {
					var inc model.PurchaseOrderDetails
					err = rows.Scan(&inc.IncoTerms)
					viewPOResp=append(viewPOResp,inc)
				}
			}

		} else {
			//DOMESTIC PO VIEW
			req.SupplierType = "Domestic"

			sqlStatementDPOV1 := `SELECT poid, podate, pocat, posubcat, 
							vendorid,itemid, billing_at_id, delivery_at_id,
							currencyid,status,
							advancetype, advance, payment_terms_days, 
							taxes_duties, transport_mode, transit_insurence, 
							packing_forward,othercharges,rate,remarks,
							purchase_type,terminal_month,terminal_price,
							purchase_price,market_price,total_price,gross_price,total_quantity,fixationdate
							FROM dbo.pur_gc_po_con_master_newpg
							where pono=$1`
			rows, err := po.db.Raw(sqlStatementDPOV1, req.PoNO).Rows()
			logger.Error("PO Master Query Executed")
			if err != nil {
				logger.Error("Fetching PO Details from DB failed")
				logger.Error(err.Error())
				// return events.APIGatewayProxyResponse{500, headers, nil, errd1.Error(), false}, nil
			}
			// defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&req.PoId, &req.PoDate, &req.POCategory, &req.POSubCategory,
					&req.SupplierID, &req.ItemID, &req.POBillTypeID, &req.PODelTypeID, &req.CurrencyID,
					&req.Status, &req.AdvanceType, &req.Advance,
					&req.PaymentTermsDays, &req.TaxDuties, &req.ModeOfTransport, &req.TransitInsurance,
					&req.DPackForward, &req.OtherCharges, &req.Rate, &req.Comments, &req.PurchaseType,
					&req.TerminalMonth, &req.DTerminalPrice, &req.PurchasePriceInr,
					&req.MarketPriceInr, &req.TotalPrice, &req.GrossPrice, &req.TotalQuantity, &req.FixationDate)
			}
		}
		// ------COMMON to IMPORT && DOMESTIC-------//
		
		// ---------------_Fetch Billing Address Info------------------------
		logger.Error("Entered Billing Module")
		sqlStatementPOVB2 := `SELECT 
						 potypeid,
						 initcap(bdi.potypename),
						 initcap(bdi.potypefullname)||','||initcap(bdi.address) as billingaddress
						 from dbo.pur_po_types bdi
						 where 
						 bdi.potypeid=(select pom.billing_at_id from dbo.pur_gc_po_con_master_newpg pom where pom.pono=$1)`
		rows, err = po.db.Raw(sqlStatementPOVB2, req.PoNO).Rows()
		logger.Error("PO Types Query Executed")
		if err != nil {
			logger.Error("Issue in fetching billing address from DB failed")
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, errb2.Error(), false}, nil
		}

		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&req.POBillTypeID, &req.POBillTypeName, &req.POBillAddress)
			logger.Error(req.POBillAddress)			
		}
		//---------------_Fetch Delivery Address Info------------------------
		logger.Error("Entered PO Delivery Module")
		sqlStatementPOVD2 := `SELECT 
						  initcap(bdi.potypename),
						 initcap(bdi.potypefullname)||','||initcap(bdi.address) as billingaddress
						 from dbo.pur_po_types bdi
						 where 
						 bdi.potypeid=(select pom.delivery_at_id from dbo.pur_gc_po_con_master_newpg pom where pom.pono=$1)`
		rows, err = po.db.Raw(sqlStatementPOVD2, req.PoNO).Rows()
		logger.Error("PO Delivery Address Query Executed")
		if err != nil {
			logger.Error("Fetching PO Delivery Details from DB failed")
			logger.Error(err.Error())			
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&req.PODelTypeName, &req.PODelAddress)
		}
		//-------__Fetch Vendor Information---------------------
		logger.Error("Entered PO Vendor Module")
		sqlStatementPOV3 := `SELECT				
						vm.vendortypeid,
						vm.country,
						initcap(vm.vendorname),
						initcap(vm.address1)||','||initcap(vm.address2)||','||initcap(vm.city)||','||pincode||','||initcap(vm.state)||' -'||SUBSTRING (vm.gstin, 1 , 2)||','||'Phone:'||vm.phone||','||'Mobile:'||vm.mobile||','||'GST NO:'||vm.gstin||','||'PAN NO:'||vm.panno,
						vm.email
						from 
						dbo.pur_vendor_master_newpg vm
						where vm.vendorid=(select pom.vendorid from dbo.pur_gc_po_con_master_newpg pom where pom.pono=$1)`
		
		rows, err  = po.db.Raw(sqlStatementPOV3, req.PoNO).Rows()
		logger.Error("Vendor Details fetch Query Executed")
		if err != nil {
			logger.Error("Fetching Vendor Details from DB failed")
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&req.SupplierTypeID, &req.SupplierCountry, &req.SupplierName, &req.SupplierAddress, &req.SupplierEmail)
		}
	
		//-------------Fetch Currencuy Info----------------------------
		logger.Error("Entered Currency Fetch Module")
		sqlStatementPOV4 := `SELECT currencyname,currencycode
							from dbo.project_currency_master 
							where currencyid=$1`
		rows, err  = po.db.Raw(sqlStatementPOV4, req.CurrencyID).Rows()
		logger.Error("Currency Details fetch Query Executed")
		if err != nil {
			logger.Error("Fetching Currency Details from DB failed")
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, err4.Error(), false}, nil
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&req.CurrencyName, &req.CurrencyCode)
		}
		if req.AdvanceType == "101" {
			req.AdvanceType = "Percentage"
		} else {
			req.AdvanceType = "Amount"
		}
		logger.Error("Currency Name & Code are: ", req.CurrencyName, req.CurrencyCode)

		//----------_Fetch Green Coffee Item Information--------------------
		if req.ItemID != "" {
			logger.Error("Entered GC Item Fetch Module")
			sqlStatementPOV5 := `SELECT im.itemid,initcap(im.itemname),im.cat_type
							from dbo.inv_gc_item_master_newpg im
							where
							im.itemid=$1`
			rows, err  = po.db.Raw(sqlStatementPOV5, req.ItemID).Rows()
			logger.Error("GC Details fetch Query Executed")
			if err != nil {
				logger.Error("Fetching GC Details from DB failed")
				logger.Error(err.Error())
				// return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
			}
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&req.ItemID, &req.ItemName, &req.GCCoffeeType)
			}
		}

		// ---------------------Fetch GC Composition Details--------------------------------------//
		logger.Error("The GC Composition for the Item #", req.ItemID)
		sqlStatementPOGC1 := `SELECT density, moisture, browns, blacks, brokenbits, insectedbeans, bleached, husk, sticks, stones, beansretained
						FROM dbo.pur_gc_po_composition_master_newpg where itemid=$1`
		rows, err  = po.db.Raw(sqlStatementPOGC1, req.ItemID).Rows()
		logger.Error("GC Fetch Query Executed")
		if err != nil {
			logger.Error("Fetching GC Composition Details from DB failed")
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}

		for rows.Next() {
			err = rows.Scan(&req.Density, &req.Moisture, &req.Browns, &req.Blacks, &req.BrokenBits, &req.InsectedBeans, &req.Bleached, &req.Husk, &req.Sticks,
				&req.Stones, &req.BeansRetained)

		}

		// ---------------------Fetch Multiple Dispatch Info-------------------------------------//
		logger.Error("Fetching Single/Multiple Dispatch Information the Contract #")
		sqlStatementMDInfo1 := `select d.detid,d.dispatch_date,d.quantity, d.dispatch_type,d.dispatch_count,
							m.delivered_quantity, (m.expected_quantity-m.delivered_quantity) as balance_quantity
							from dbo.pur_gc_po_dispatch_master_newpg d
							left join dbo.inv_gc_po_mrin_master_newpg as m on m.detid=d.detid
							where d.pono=$1`
		rows, err  = po.db.Raw(sqlStatementMDInfo1, req.PoNO).Rows()
		
		logger.Error("Multi Dispatch Info Fetch Query Executed")
		if err != nil {
			logger.Error("Multi Dispatch Info Fetch Query failed")
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		var mid model.ItemDispatch
		
		
		for rows.Next() {
			
			err = rows.Scan(&mid.DispatchID, &mid.DispatchDate, &mid.DispatchQuantity, &req.DispatchType, &req.DispatchCount, &mid.DeliveredQuantity, &mid.BalanceQuantity)
			// itemDisp = append(itemDisp, mid)
			gcMultiDispatch := append(req.ItemDispatchDetails, mid)
			req.ItemDispatchDetails = gcMultiDispatch
		}
		if err != nil {
			
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		// logger.Error("Multi Dispatch Details:", req.ItemDispatchDetails)

		//---------------Fetch Domestic Tax info for Domestic PO-------------------

		if req.POSubCategory == "Domestic" {
			logger.Error("Selected supplier type Domestic Code:", req.POSubCategory)
			sqlStatementDTax1 := `SELECT sgst, cgst, igst,pack_forward, installation,
							 freight, handling, misc, hamali, mandifee, full_tax,
							  insurance FROM dbo.pur_gc_po_details_taxes_newpg 
							  where pono=$1`
			rows, err  = po.db.Raw(sqlStatementDTax1, req.PoNO).Rows()
			logger.Error("Domestic Tax Info Fetch Query Executed")
			if err != nil {
				logger.Error("Domestic Tax Info Fetch Query failed")
				logger.Error(err.Error())
				// return events.APIGatewayProxyResponse{500, headers, nil, errDTax1.Error(), false}, nil
			}

			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&req.SGST, &req.CGST, &req.IGST, &req.DPackForward, &req.DInstallation, &req.DFreight,
					&req.DHandling, &req.DMisc, &req.DHamali, &req.DMandiFee, &req.DFullTax, &req.DInsurance)
			}	
			if err != nil {
			
				logger.Error(err.Error())
			}
		}
		//----------Quote Info for Speciality Green Coffee Item Information--------------------
		if req.GCCoffeeType != "regular" {
			logger.Error("Entered Quote date & Quote Info Fetch Module for speciaity Coffee")
			sqlStatementSPQ := `SELECT 
							 pom.quote_no,
							 pom.quote_date,
							 pom.quote_price
							 from dbo.pur_gc_po_con_master_newpg pom
							 where pom.pono=$1`
			rows, err  = po.db.Raw(sqlStatementSPQ, req.PoNO).Rows()
			logger.Error("Quote Info Fetch Module for speciaity Coffee Query Executed")
			if err != nil {
				logger.Error("Quote Info Fetch Module for speciaity Coffee from DB failed")
				logger.Error(err.Error())
				// return events.APIGatewayProxyResponse{500, headers, nil, errSPQ.Error(), false}, nil
			}
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&req.QuotNo, &req.QuotDate, &req.QuotPrice)
			}
			if err != nil {
			
				logger.Error(err.Error())
			}
			logger.Error(req.QuotNo, req.QuotDate)
		}
		//------Consolidated Finance Status------------------//
		sqlStatementCFS := `SELECT accpay_status,qc_status,payable_amount
						FROM dbo.pur_gc_po_con_master_newpg
						where pono=$1`
		rows, err  = po.db.Raw(sqlStatementCFS, req.PoNO).Rows()
		logger.Error("Consolidated Finance Status Query Executed")
		if err != nil {
			logger.Error("Fetching Consolidated Finance Status from DB failed")
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&req.QCStatus, &req.APStatus, &req.PayableAmount)
		}
		if err != nil {
			
			logger.Error(err.Error())
		}
		//---------------------Fetch Audit Log Info-------------------------------------//
		logger.Error("Fetching Audit Log Info #")
		sqlStatementAI := `select u.username as createduser, gc.created_date,
			gc.description, v.username as modifieduser, gc.modified_date
   			from dbo.auditlog_pur_gc_master_newpg gc
   			inner join dbo.users_master_newpg u on gc.createdby=u.userid
  			left join dbo.users_master_newpg v on gc.modifiedby=v.userid
   			where gc.pono=$1 order by logid desc limit 1`
		rows, err  = po.db.Raw(sqlStatementAI, req.PoNO).Rows()
		logger.Error("Audit Info Fetch Query Executed")
		if err != nil {
			logger.Error("Audit Info Fetch Query failed")
			logger.Error(err.Error())
			// return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}

		for rows.Next() {
			var al model.AuditLogGCPO
			err = rows.Scan(&al.CreatedUserid, &al.CreatedDate, &al.Description, &al.ModifiedUserid, &al.ModifiedDate)
			auditDetails := append(req.AuditLogDetails, al)
			req.AuditLogDetails = auditDetails
			logger.Error("added one")
		}
		if err != nil {
			
			logger.Error(err.Error())
		}
		logger.Error("Audit Details:", req.AuditLogDetails)
		return req,err
		
	} else {
		logger.Error("Couldnt find po")
		// return events.APIGatewayProxyResponse{200, headers, nil, string("Couldn't find PO Details"), false}, nil
	}
	// return events.APIGatewayProxyResponse{200, headers, nil, string("success"), false}, nil
	return nil,nil
}

//  func (po poRepo)	ListPurchaseOrders(ctx context.Context, req *model.ListPurchaseOrderRequest) (interface{}, error) {
// 	const (
// 		APPROVEDPOS ="approvedpos"
// 		PENDINGPOS ="pendingpos"
// 	)

// 	logger := logger2.GetLoggerWithContext(ctx)
// 	var allPurchaseOrderDetails []model.PurchaseOrderDetails
// 	if req.Type == PENDINGPOS {
// 		sqlStatement := `SELECT q.pono,
// 		TO_CHAR(q.podate,'DD-MON-YY') as podate,q.quotno,q.approvalstatus,initcap(a.vendorname), q.pocat, initcap(c.vendortypename), b.currencycode,
// 		d.final_price,q.taxes_duties,q.advance
// 		  FROM dbo.pur_gc_po_con_master_newpg as q
// 		  INNER JOIN dbo.pur_vendor_master as a ON q.vendorid = a.vendorid
// 		  INNER JOIN dbo.pur_vendor_types as c ON a.vendortypeid = c.vendortypeid
// 		  INNER JOIN dbo.project_currency_master as b ON q.currencyid = b.currencyid
// 		  Left JOIN dbo.pur_gc_price_details_newpg as d ON q.pono = d.pono
// 		  where q.approvalstatus=false order by q.poidsno DESC`
// 		rows, err  := po.db.Raw(sqlStatement).Rows()
//         if err!= nil {
//         	logger.Error("error while getting data from database err: ",err.Error())
//         	return nil, err
// 		}
// 		defer rows.Close()
// 		for rows.Next() {
// 			var po model.PurchaseOrderDetails
// 			err = rows.Scan(&po.PoNo, &po.PoDate, &po.QuotNo, &po.ApprovalStatus, &po.Vendor, &po.Category, &po.VendorTypeId, &po.Currency, &po.PoValue, &po.TaxValue, &po.Advance)
// 			allPurchaseOrderDetails = append(allPurchaseOrderDetails, po)
// 		}

// 	} else if req.Type ==  APPROVEDPOS {
//     sqlStatement := `SELECT q.pono,
// 		TO_CHAR(q.podate,'DD-MON-YY') as podate,q.quotno,q.approvalstatus,initcap(a.vendorname), q.pocat, initcap(c.vendortypename), b.currencycode,
// 			d.final_price,q.taxes_duties,q.advance
// 		FROM dbo.pur_gc_po_con_master_newpg as q
// 		INNER JOIN dbo.pur_vendor_master as a ON q.vendorid = a.vendorid
// 		INNER JOIN dbo.pur_vendor_types as c ON a.vendortypeid = c.vendortypeid
// 		INNER JOIN dbo.project_currency_master as b ON q.currencyid = b.currencyid
// 		Left JOIN dbo.pur_gc_price_details_newpg as d ON q.pono = d.pono
// 		where q.approvalstatus=true order by q.poidsno DESC`
//     rows,err := po.db.Raw(sqlStatement).Rows()
//     if err != nil {
//     	logger.Error("error while getting  date from database err: ",err.Error() )
//     	return nil, err
// 	}
// 		defer rows.Close()
// 		for rows.Next() {
// 			var po model.PurchaseOrderDetails
// 			err = rows.Scan(&po.PoNo, &po.PoDate, &po.QuotNo, &po.ApprovalStatus, &po.Vendor, &po.Category, &po.VendorTypeId, &po.Currency, &po.PoValue, &po.TaxValue, &po.Advance)
// 			allPurchaseOrderDetails = append(allPurchaseOrderDetails, po)
// 		}

// 	} else {
// 		rows, err  = po.db.Raw(`SELECT q.pono,
// 		TO_CHAR(q.podate,'DD-MON-YY') as podate,q.quotno,q.approvalstatus,initcap(a.vendorname), q.pocat, initcap(c.vendortypename), b.currencycode,
// 		d.final_price,q.taxes_duties,q.advance
// 		  FROM dbo.pur_gc_po_con_master_newpg as q
// 		  INNER JOIN dbo.pur_vendor_master as a ON q.vendorid = a.vendorid
// 		  INNER JOIN dbo.pur_vendor_types as c ON a.vendortypeid = c.vendortypeid
// 		  INNER JOIN dbo.project_currency_master as b ON q.currencyid = b.currencyid
// 		  Left JOIN dbo.pur_gc_price_details_newpg as d ON q.pono = d.pono
// 		  order by q.poidsno DESC`).Rows()
// 		  if err!= nil {
// 		  	logger.Error("error while getting data from databases err: ",err.Error())
// 		  	return nil, err
// 		  }
// 		defer rows.Close()
// 		for rows.Next() {
// 			var po model.PurchaseOrderDetails
// 			err = rows.Scan(&po.PoNo, &po.PoDate, &po.QuotNo, &po.ApprovalStatus, &po.Vendor, &po.Category, &po.VendorTypeId, &po.Currency, &po.PoValue, &po.TaxValue, &po.Advance)
// 			allPurchaseOrderDetails = append(allPurchaseOrderDetails, po)
// 		}
// 	}
// 	return nil, nil
//  }
