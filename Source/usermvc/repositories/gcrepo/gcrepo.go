package gcrepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"

	logger2 "usermvc/utility/logger"

	"github.com/jinzhu/gorm"
)

type GcRepo interface {
	InsertGcDetails(ctx context.Context, gcs entity.GCDetails) (interface{}, error)
	ViewGcDetails(ctx context.Context, gcDetails *model.GCViewDetails) (interface{}, error)
	ListGcDetails(ctx context.Context, req *model.ListGCDetails) (interface{}, error)
}

type gcRepo struct {
	db *gorm.DB
}

func NewgcRepo() GcRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}
	newDb.AutoMigrate(&entity.GCDetails{})
	return &gcRepo{
		db: newDb,
	}
}

type NullString struct {
	sql.NullString
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (gc gcRepo) InsertGcDetails(ctx context.Context, gcDetails entity.GCDetails) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	fmt.Println(logger)

	sqlStatementDT1 := `select itemidsno from dbo.inv_gc_item_master_newpg order by itemidsno desc limit 1`

	rows1, err := gc.db.Raw(sqlStatementDT1).Rows()

	for rows1.Next() {
		err = rows1.Scan(&gcDetails.Itemidsno)
	}

	gcDetails.Itemidsno = gcDetails.Itemidsno + 1
	gcDetails.Itemid = "FAC-" + strconv.Itoa(gcDetails.Itemidsno)
	gcDetails.LCode = "GCITEM-FAC-" + strconv.Itoa(gcDetails.Itemidsno)
	gcDetails.LGroupCode = "GCITEM-" + gcDetails.GroupId

	if gcDetails.IsSpecialCoffee {
		gcDetails.CategoryType = "speciality"
	} else {
		gcDetails.CategoryType = "regular"
	}

	sqlStatement2 := `INSERT INTO dbo.inv_gc_item_master_newpg (
		groupid, itemcode,
		s_code, itemname,itemdesc, hsncode, convertionratio,
		itemcatid,uom,coffee_type,display_inpo,dailyprice_enable,
		createdon, createdby, lcode,lname,lgroupcode, itemid, itemidsno, cat_type)
		 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,$15,$16,$17,$18,$19,$20)`
	_, err = gc.db.Raw(sqlStatement2,
		gcDetails.GroupId, gcDetails.ItemCode, gcDetails.SCode,
		gcDetails.ItemName, gcDetails.ItemDesc, gcDetails.HsnCode,
		gcDetails.ConvertionRatio, gcDetails.ItemCatId,
		gcDetails.Uom, gcDetails.CoffeeType, gcDetails.DisplayInPo,
		gcDetails.DisplayInDailyUpdates, time.Now().Format("2006-01-02"),
		gcDetails.CreatedBy, gcDetails.LCode, gcDetails.LName, gcDetails.LGroupCode,
		gcDetails.Itemid, gcDetails.Itemidsno, gcDetails.CategoryType).Rows()
	//log.Println(gcDetails)

	if err != nil {
		logger.Error("error")
	}

	sqlStatement3 := `INSERT INTO dbo.pur_gc_po_composition_master_newpg (density,
		moisture, browns, blacks,
		brokenbits, insectedbeans,
		bleached, husk, sticks, stones, beansretained, itemid) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err = gc.db.Raw(sqlStatement3,
		gcDetails.Density, gcDetails.Moisture, gcDetails.Browns, gcDetails.Blacks,
		gcDetails.BrokenBits, gcDetails.InsectedBeans, gcDetails.Bleached, gcDetails.Husk,
		gcDetails.Sticks, gcDetails.Stones, gcDetails.BeansRetained,
		gcDetails.Itemid).Rows()
	//defer rows3.Close()
	if err != nil {
		logger.Error("error")
	}
	if gcDetails.Update && gcDetails.Itemid != "" {

		var categoryType string
		if gcDetails.IsSpecialCoffee {
			categoryType = "speciality"
		} else {
			categoryType = "regular"
		}

		sqlStatement1 := `UPDATE dbo.inv_gc_item_master_newpg SET 
			groupid=$1,itemcode=$2,s_code=$3,itemname=$4,itemdesc=$5,hsncode=$6, 
			convertionratio=$7,itemcatid=$8,uom=$9,coffee_type=$10,
			display_inpo=$11,dailyprice_enable=$12,updatedon=$13,
			updatedby=$14, lname=$15,lgroupcode=$16, cat_type=$17 where itemid=$18`

		_, err = gc.db.Raw(sqlStatement1,
			gcDetails.GroupId, gcDetails.ItemCode, gcDetails.SCode, gcDetails.ItemName,
			gcDetails.ItemDesc, gcDetails.HsnCode, gcDetails.ConvertionRatio, gcDetails.ItemCatId,
			gcDetails.Uom, gcDetails.CoffeeType, gcDetails.DisplayInPo, gcDetails.DisplayInDailyUpdates,
			gcDetails.UpdatedOn, gcDetails.UpdatedBy, gcDetails.LName, gcDetails.LGroupCode, categoryType, gcDetails.Itemid).Rows()

		sqlStatement2 := `UPDATE dbo.pur_gc_po_composition_master_newpg SET 
			density=$1,
			moisture=$2,
			browns=$3,
			blacks=$4,
			brokenbits=$5,
			insectedbeans=$6,
			bleached=$7, 
			husk=$8,
			sticks=$9, 
			stones=$10, 
			beansretained=$11 where itemid=$12`

		_, err = gc.db.Raw(sqlStatement2,
			gcDetails.Density,
			gcDetails.Moisture,
			gcDetails.Browns,
			gcDetails.Blacks,
			gcDetails.BrokenBits,
			gcDetails.InsectedBeans,
			gcDetails.Bleached,
			gcDetails.Husk,
			gcDetails.Sticks,
			gcDetails.Stones,
			gcDetails.BeansRetained,
			gcDetails.Itemid).Rows()

		if err != nil {
			log.Println("unable to update gc details", err.Error())
			return nil, nil
		}
	}
	return gcDetails, nil
}

func (gc gcRepo) ViewGcDetails(ctx context.Context, gCDetails *model.GCViewDetails) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("viewing gc details")

	if gCDetails.Itemid != " " {
		sqlStatement :=
			`SELECT 
			gc.groupid, gc.itemcode, gc.s_code, gc.itemname, gc.itemdesc,
			gc.hsncode, gc.convertionratio,gc.itemcatid,
			initcap(category.itemcatname) as catname,gc.uom, initcap(u.uomname) as uomname, coffee_type, display_inpo,
			dailyprice_enable, initcap(gc.lname), gc.lgroupcode, gc.cat_type
			from dbo.inv_gc_item_master_newpg gc
			INNER JOIN dbo.INV_ITEM_CATEGORY as category ON gc.itemcatid = category.itemcatid
            INNER JOIN dbo.PROJECT_UOM as u ON gc.uom = u.uom
			where gc.Itemid=$1`

		rows, err := gc.db.Raw(sqlStatement, gCDetails.Itemid).Rows()

		if err != nil {
			log.Println(err)
			return nil, nil
		}

		for rows.Next() {
			err = rows.Scan(&gCDetails.GroupId, &gCDetails.ItemCode, &gCDetails.SCode,
				&gCDetails.ItemName,
				&gCDetails.ItemDesc, &gCDetails.HsnCode, &gCDetails.ConvertionRatio, &gCDetails.DisplayInPo,
				&gCDetails.DisplayInDailyUpdates, &gCDetails.LName, &gCDetails.LGroupCode, &gCDetails.CategoryType, &gCDetails.ItemCatId,
				&gCDetails.ItemCatName, &gCDetails.Uom, &gCDetails.UomName, &gCDetails.CoffeeType)
		}
		if err != nil {
			logger.Error("error")
		}
		if gCDetails.CategoryType == "speciality" {
			gCDetails.IsSpecialCoffee = true
		}
		sqlStatementPOGC1 := `SELECT density, moisture, browns, blacks, brokenbits, insectedbeans, bleached, husk, sticks, stones, beansretained
					FROM dbo.pur_gc_po_composition_master_newpg where itemid=$1`
		rows7, err7 := gc.db.Raw(sqlStatementPOGC1, gCDetails.Itemid).Rows()

		if err7 != nil {
			log.Println("Fetching GC Composition Details from DB failed")
			log.Println(err7.Error())
			return nil, nil
		}

		for rows7.Next() {
			err7 = rows7.Scan(&gCDetails.Density, &gCDetails.Moisture, &gCDetails.Browns, &gCDetails.Blacks,
				&gCDetails.BrokenBits, &gCDetails.InsectedBeans, &gCDetails.Bleached,
				&gCDetails.Husk, &gCDetails.Sticks, &gCDetails.Stones, &gCDetails.BeansRetained)
		}
		if err7 != nil {
			logger.Error("error")
		}
		log.Println("Fetching stock location Info #")
		sqlStatementSL := `select initcap(entity.entityname) as entityname, initcap(master.name) as name, loc.quantity,loc.value,loc.unitprice
	from dbo.inv_gc_item_master_newpg gc
	  INNER JOIN dbo.INV_GC_ITEM_LOCATION_OPENING as loc ON gc.lcode = loc.lcode
	  INNER JOIN dbo.inv_gc_location_master as master ON loc.id = master.locationid
	  INNER JOIN dbo.project_entity_master as entity ON master.entityid = entity.entityid
	  where gc.itemid=$1`
		rowsSL, errSL := gc.db.Raw(sqlStatementSL, gCDetails.Itemid).Rows()
		if errSL != nil {
			log.Println("Fetching stock location Info failed")
			log.Println(errSL.Error())
		}

		for rowsSL.Next() {
			var sl model.ItemStockLocation
			errSL = rowsSL.Scan(&sl.Entity, &sl.Name, &sl.Quantity, &sl.Value, &sl.UnitPrice)
			stockDetails := append(gCDetails.StockLocation, sl)
			gCDetails.StockLocation = stockDetails
			log.Println("added one")
		}
		log.Println("Fetching vendor info #")
		sqlStatementVL := `SELECT distinct initcap(master.vendorname) as vendorname,initcap(master.contactname) as contactname,initcap(master.state) as state,
	initcap(master.country) as country,initcap(master.city) as city
				from dbo.pur_gc_po_con_master_newpg po
				inner join dbo.pur_vendor_master as master on po.vendorid = master.vendorid
				where po.Itemid=$1`
		rowsVL, errVL := gc.db.Raw(sqlStatementVL, gCDetails.Itemid).Rows()
		if errVL != nil {
			log.Println("Fetching stock location Info failed")
			log.Println(errSL.Error())
		}

		for rowsVL.Next() {
			var vl model.VendorList
			errVL = rowsVL.Scan(&vl.VendorName, &vl.ContactName, &vl.State, &vl.Country, &vl.City)
			vendorDetails := append(gCDetails.VendorList, vl)
			gCDetails.VendorList = vendorDetails
			log.Println("added one")
		}

		if errVL != nil {
			logger.Error("error")
		}
	}
	return gCDetails, nil
}

/*func (gc gcRepo) ListGcDetails(ctx context.Context, req *model.InputG) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	const(

	)
	var listgc []model.ListGCDetails
	if req.Type ==
	  sqlstmt := SELECT
	  gc.itemid,gc.itemcode as item_code, gc.s_code as s_code, gc.itemname as item_name,
	  category.itemcatname as category_name, u.uomname as uom,
	  gc.groupid as group_id, grp.groupname as group_name, gc.display_inpo as display_inpo,
	  gc.coffee_type as coffee_type from dbo.inv_gc_item_master_newpg gc
		INNER JOIN dbo.INV_GC_ITEMGROUPS as grp ON gc.groupid = grp.groupid
		INNER JOIN dbo.INV_ITEM_CATEGORY as category ON gc.itemcatid = category.itemcatid
		INNER JOIN dbo.PROJECT_UOM as u ON gc.uom = u.uom)`

rows, err = gc.db.Raw(sqlStatement1).Row()
if err != nil {
//logger.Error("error")
return nil

}*/

func (gc gcRepo) ListGcDetails(ctx context.Context, req *model.ListGCDetails) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)

	var input model.InputG
	var filter1, filter2 string
	if input.Type == "filterGcs" {
		filter1 = "where gc.groupid=" + input.GroupId + " order by gc.itemidsno DESC"
		filter2 = ""
		if input.AdvancedFilter {
			logger.Info("Filter Gcs and Advanced filters selected")
			filter1 = "where gc.groupid=" + input.GroupId + "order by gc.itemidsno DESC"
			filter2 = "where " + input.FilterParam
		} else if input.Type == "allGcs" {
			filter1 = "order by gc.itemidsno DESC"
			filter2 = ""
			if input.AdvancedFilter {
				logger.Info("All Gcs and Advanced filter selected")
				filter1 = "order by gc.itemidsno DESC"
				filter2 = "where" + input.FilterParam
			}
		} else if input.Type == "specialGcs" {
			filter1 = "where gc.cat_type='speciality' order by gc.itemidsno DESC"
			filter2 = "where " + input.FilterParam
		}
	}
	res := dynamicQuery(filter1, filter2)
	return res, nil
}
func dynamicQuery(param1, param2 string) []model.ListGCDetails {
	var all []model.ListGCDetails
	var db *sql.DB
	sqlStatement1 := `select * from (SELECT
						gc.itemid,gc.itemcode as item_code, gc.s_code as s_code, gc.itemname as item_name,
						category.itemcatname as category_name, u.uomname as uom,
						gc.groupid as group_id, grp.groupname as group_name, gc.display_inpo as display_inpo,
						gc.coffee_type as coffee_type from dbo.inv_gc_item_master_newpg gc
			  			INNER JOIN dbo.INV_GC_ITEMGROUPS as grp ON gc.groupid = grp.groupid
			  			INNER JOIN dbo.INV_ITEM_CATEGORY as category ON gc.itemcatid = category.itemcatid
	          			INNER JOIN dbo.PROJECT_UOM as u ON gc.uom = u.uom %s) a %s`
	rows, err := db.Query(sqlStatement1)
	//rows, err = gc.db.Raw(sqlStatement1).Row()
	if err != nil {
		//logger.Error("error")
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var gcd model.ListGCDetails
		err = rows.Scan(&gcd.ItemId, &gcd.ItemCode, &gcd.SCode, &gcd.ItemName, &gcd.CategoryName, &gcd.Uom, &gcd.GroupId, &gcd.Groupname,
			&gcd.DisplayInPo, &gcd.CoffeeType)
		all = append(all, gcd)
	}
	return all
}
