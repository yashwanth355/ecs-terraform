package supplierrepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"
)

type SupplierRepo interface {
	InsertSupplierDetails(ctx context.Context, supplierdetails entity.VendorDetails) (interface{}, error)
	ViewSupplierDetails(ctx context.Context, req *model.VendorDetails) (interface{}, error)
	ListSupplierDetails(ctx context.Context, req *model.ListSupplier) (interface{}, error)
	//	UpdateSupplierDetails(ctx context.Context, supplieretails entity.VendorDetails) (interface{}, error)
}

type supplierRepo struct {
	db *gorm.DB
}

func NewsupplierRepo() SupplierRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}

	newDb.AutoMigrate(&entity.CmsLeadsMaster{})

	return &supplierRepo{
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

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (sc supplierRepo) InsertSupplierDetails(ctx context.Context, vendor entity.VendorDetails) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	fmt.Println(logger)
	sqlStatementImp1 := `INSERT INTO dbo.pur_vendor_master_newpg(
		vendorid,vendoridsno,vendorname, vendortypeid,vendorcatid,groupid,contactname,address1,address2,
		country,state,city,pincode,phone,mobile,email,web,panno,gstin,msme,bankname,branch,accounttype,accountno,ifscode,micrcode,
		auto_gen_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27)`

	_, err := sc.db.Raw(sqlStatementImp1,
		vendor.VendorId, vendor.VendorIdSno, vendor.VendorName, NewNullString(vendor.VendorTypeid), NewNullString(vendor.VendorCategoryid),
		vendor.VendorGroupid, vendor.ContactName, vendor.Address1, vendor.Address2, vendor.Country, vendor.State,
		vendor.City, vendor.Pincode, vendor.Phone, vendor.Mobile, vendor.Email, vendor.Website, vendor.PanNo,
		vendor.GSTIdentificationNo, vendor.MSMESSI, vendor.BankName, vendor.Branch, vendor.AccountType,
		vendor.AccountNumber, vendor.IfscCode, vendor.MicrCode, vendor.AutoGenID).Rows()
	if err != nil {
		logger.Error("error")
	}

	if vendor.Update && vendor.VendorId != "" {
		sqlStatementImp1 := `update dbo.pur_vendor_master_newpg set
	vendorname = $1, vendortypeid = $2,	vendorcatid = $3,groupid = $4,contactname = $5,address1 = $6,
	address2 = $7,country = $8,state = $9,city = $10,pincode = $11, phone = $12,mobile = $13,email = $14,
	web = $15,panno  =$16,gstin = $17,msme = $18,bankname = $19,branch = $20,accounttype = $21,
	accountno = $22,ifscode = $23,micrcode = $24 where vendorid = $25`

		_, err = sc.db.Raw(sqlStatementImp1,
			vendor.VendorName, NewNullString(vendor.VendorTypeid), NewNullString(vendor.VendorCategoryid),
			vendor.VendorGroupid, vendor.ContactName, vendor.Address1, vendor.Address2, vendor.Country,
			vendor.State, vendor.City, vendor.Pincode, vendor.Phone, vendor.Mobile, vendor.Email, vendor.Website,
			vendor.PanNo, vendor.GSTIdentificationNo, vendor.MSMESSI, vendor.BankName, vendor.Branch, vendor.AccountType,
			vendor.AccountNumber, vendor.IfscCode, vendor.MicrCode, vendor.VendorId).Rows()
		if err != nil {
			logger.Error("error")
		}
	}

	return vendor, nil
}
func (sc supplierRepo) ViewSupplierDetails(ctx context.Context, vendor *model.VendorDetails) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting supplier details")
	if vendor.View && vendor.VendorId != "" {
		sqlStatementMDInfo1 := `select d.vendortypeid,d.vendorname, d.groupid, d.contactname,d.address1,d.address2,
							d.city, d.pincode, d.country,d.phone,d.mobile, d.web, d.accountno, d.ifscode,d.bankname,
							d.panno,d.branch, d.gstin,d.msme,d.micrcode,d.accounttype,d.state,m.vendorcatid,m.vendorcatname,n.groupname,d.email
							from dbo.pur_vendor_master_newpg d
							inner join dbo.pur_vendor_category as m 
							on d.vendorcatid=m.vendorcatid
							inner join dbo.pur_vendor_groups as n 
							on d.groupid=n.groupid
							where d.vendorid=$1`
		rows1, err1 := sc.db.Raw(sqlStatementMDInfo1, vendor.VendorId).Rows()
		log.Println("fetch query executed")
		if err1 != nil {
			log.Println("Query failed")
			log.Println(err1.Error())
			return nil, nil
		}

		for rows1.Next() {

			err1 = rows1.Scan(&vendor.VendorTypeid, &vendor.VendorName, &vendor.VendorGroupid, &vendor.ContactName, &vendor.Address1, &vendor.Address2,
				&vendor.City, &vendor.Pincode, &vendor.Country, &vendor.Phone, &vendor.Mobile, &vendor.Website, &vendor.AccountNumber,
				&vendor.IfscCode, &vendor.BankName, &vendor.PanNo, &vendor.Branch, &vendor.GSTIdentificationNo, &vendor.MSMESSI,
				&vendor.MicrCode, &vendor.AccountType, &vendor.State, &vendor.VendorCategoryid, &vendor.VendorCategoryName, &vendor.VendorGroupName, &vendor.Email)

		}
		if err1 != nil {
			logger.Error("error while viewing")
		}
	}
	return vendor, nil
	//Working fine end of function
}
func (sc supplierRepo) ListSupplierDetails(ctx context.Context, req *model.ListSupplier) (interface{}, error) {
	const (
		Instantcoffeesupplier = "Instant Coffee Suppliers"
		Indigeneous           = "Indigeneous Green Coffee Suppliers"
		ImportedGreencoffee   = "Imported Green Coffee Suppliers"
	)
	logger := logger2.GetLoggerWithContext(ctx)
	var allsupplierdetails []model.ListSupplierDetails
	if req.Type == Instantcoffeesupplier {
		sqlstmt1 := `SELECT m.vendorname,m.contactname, m.city, m.state,m.country,g.groupname,m.phone
			 from dbo.pur_vendor_master_newpg m
			 INNER JOIN dbo.pur_vendor_groups g  ON g.groupid = m.groupid
		WHERE g.groupid='FAC-9'	order by vendoridsno desc`
		rows, err := sc.db.Raw(sqlstmt1).Rows()
		if err != nil {
			logger.Error("Error while getting")
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var list1 model.ListSupplierDetails
			err = rows.Scan(&list1.Name, &list1.ContactPerson, &list1.City, &list1.State, &list1.Country, &list1.Group, &list1.Phone)

			allsupplierdetails = append(allsupplierdetails, list1)
		}

	} else if req.Type == Indigeneous {
		sqlstmt1 := `SELECT m.vendorname,m.contactname,m.city,m.state,m.country,g.groupname,m.phone
				 from dbo.pur_vendor_master_newpg m
				 INNER JOIN dbo.pur_vendor_groups g  ON g.groupid = m.groupid
	WHERE g.groupid='FAC-3' order by vendoridsno desc`
		rows, err := sc.db.Raw(sqlstmt1).Rows()
		if err != nil {
			logger.Error("Error while getting")
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var suppli model.ListSupplierDetails
			err = rows.Scan(&suppli.Name, &suppli.ContactPerson, &suppli.City, &suppli.State,
				&suppli.Country, &suppli.Group, &suppli.Phone)
			allsupplierdetails = append(allsupplierdetails, suppli)
		}
	} else if req.Type == ImportedGreencoffee {

		sqlstmt1 := `SELECT m.vendorname,m.contactname,m.city,m.state,m.country,g.groupname,m.phone
				 from dbo.pur_vendor_master_newpg m
				 INNER JOIN dbo.pur_vendor_groups g  ON g.groupid = m.groupid WHERE g.groupid='FAC-2' order by vendoridsno desc`
		rows, err := sc.db.Raw(sqlstmt1).Rows()
		if err != nil {
			logger.Error("Error while getting")
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var suppli1 model.ListSupplierDetails
			err = rows.Scan(&suppli1.Name, &suppli1.ContactPerson, &suppli1.City, &suppli1.State,
				&suppli1.Country, &suppli1.Group, &suppli1.Phone)
			allsupplierdetails = append(allsupplierdetails, suppli1)
		}
	} else {
		sqlstmt := `SELECT m.vendorname,m.contactname,m.city,m.state,m.country,g.groupname,m.phone
			  from dbo.pur_vendor_master_newpg m
			INNER JOIN dbo.pur_vendor_groups g  ON g.groupid = m.groupid`
		//WHERE g.groupid='FAC-9'	order by vendoridsno desc`
		rows, err := sc.db.Raw(sqlstmt).Rows()
		if err != nil {
			logger.Error("Error while getting")
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var suppli2 model.ListSupplierDetails
			err = rows.Scan(&suppli2.Name, &suppli2.ContactPerson, &suppli2.City, &suppli2.State,
				&suppli2.Country, &suppli2.Group, &suppli2.Phone)
			allsupplierdetails = append(allsupplierdetails, suppli2)
		}
	}
	return allsupplierdetails, nil
}
