package leadrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"
	apputils "usermvc/utility/apputils"
	logger2 "usermvc/utility/logger"

	"github.com/jinzhu/gorm"
)

type LeadRepo interface {
	GetLeadDetails(ctx context.Context, req model.GetLeadDetailsRequestBody) (interface{}, error)
	GetAllLeadDetails(ctx context.Context) ([]*model.GetAlleadsResonseBody, error)
	GetCmsLeads(ctx context.Context, req model.GetLeadDetailsRequestBody) (entity.CmsLeadsMaster, error)

	GetSalutation(ctx context.Context, salutationid int64) (*entity.CmsSalutationMaster, error)
	GetCmsLeadsShippingAddress(ctx context.Context, leadId string) (*entity.CmsLeadsShippingAddressMaster, error)
	GetCmsAccountProductSegment(ctx context.Context, productsegmentid int) (*entity.CmsAccountProductSegmentMaster, error)
	GetCmsPhonecodes(ctx context.Context, contact_extid int64) (*entity.CmsPhonecodesMaster, error)
	GetCmsAccountType(ctx context.Context, accountID int64) (*entity.CmsAccountTypeMaster, error)
	GetcmsCoffeetype(ctx context.Context, Id int64) (*entity.CmsCoffeetypeMaster, error)
	GetCmsLeadsBillingAddress(ctx context.Context, leadID string) (*entity.CmsLeadsBillingAddressMaster, error)
	GetLeadCreationInfo(ctx context.Context, req *model.GetLeadCreationInfoRequest) (interface{}, error)
	GetProdPackcategoryName(ctx context.Context, categoryId int) (*string, error)
	GetProdPackCategoryWeight(ctx context.Context, weightId int) (*string, error)
	//GetLeadsInfo(ctx context.Context, req model.GetLeadInfoReq) ([]*model.LeadInfo, error)
	GetQuoatotionCreateInfoReq(ctx context.Context, req model.GetQuoatotionCreateInfoReq) (interface{}, error)

	LeadExists(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest) (bool, error)

	CreateNewLead(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest) error

	InsertLeadRecord(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (string, error)

	AddBillingAddressOnNewLead(leadId string, ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	AddShippingAddressOnNewLead(leadId string, ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	LogToAuditLogOnNewLead(leadId string, ctx context.Context,
		equestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	LogNotificationOnLeadCreateUpdate(leadId string, status string, ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	UpdateLead(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest) error

	UpdateLeadsShippingaddress(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	UpdateLeadsBillingaddress(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	EditAuditLogEntryOnLeadAmend(ctx context.Context,
		requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error)

	ProvideLeadsData(ctx context.Context,
		reqParams model.ProvideLeadsInfoReqContext) ([]model.LeadInfo, error)
}

type leadRepo struct {
	db *gorm.DB
}

func NewLeadRepo() LeadRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}

	newDb.AutoMigrate(&entity.CmsLeadsMaster{})

	return &leadRepo{
		db: newDb,
	}
}
func (lr leadRepo) GetLeadDetails(ctx context.Context, req model.GetLeadDetailsRequestBody) (interface{}, error) {

	var cmsLeadsMaster []*entity.CmsLeadsMaster
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	if err := lr.db.Table("cms_leads_master").Where(&entity.CmsLeadsMaster{Leadid: "12"}).Model(&entity.CmsLeadsMaster{}).Find(&cmsLeadsMaster); err != nil {
		logger.Error("error while getting record from cms_leads_master where req is ", req)
	}

	return cmsLeadsMaster, nil
}

func (lr leadRepo) GetCmsLeads(ctx context.Context, req model.GetLeadDetailsRequestBody) (entity.CmsLeadsMaster, error) {
	var cmsLeadsMaster entity.CmsLeadsMaster
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	if err := lr.db.Table("dbo.cms_leads_master").Where(&entity.CmsLeadsMaster{Leadid: req.Id}).Model(&entity.CmsLeadsMaster{}).Find(&cmsLeadsMaster); err != nil {
		logger.Error("error while getting record from cms_leads_master where req is ", req)
	}

	return cmsLeadsMaster, nil
}
func (lr leadRepo) GetAllLeadDetails(ctx context.Context) ([]*model.GetAlleadsResonseBody, error) {
	var result []*model.GetAlleadsResonseBody
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	rows, err := lr.db.Raw(`SELECT "accountname","aliases","contactfirstname","contactlastname","phone","email","approvalstatus" FROM "cms_leads_master"`).Rows()

	if err != nil {
		logger.Error(err.Error())
	}
	for rows.Next() {
		var account model.GetAlleadsResonseBody
		err = rows.Scan(&account.AccountName, &account.Aliases, &account.Contactfirstname, &account.Contactlastname, &account.Phone, &account.Email, &account.ApprovalStatus)
		result = append(result, &account)
	}
	res, _ := json.Marshal(result)
	logger.Info("response from cms_leads_master", string(res))
	return result, nil
}

func (lr leadRepo) GetSalutation(ctx context.Context, salutationid int64) (*entity.CmsSalutationMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_salutation_master where SalutationId  ", salutationid)
	var salutation entity.CmsSalutationMaster
	if err := lr.db.Table("dbo.cms_salutation_master").Where("salutationid=?", salutationid).Find(&salutation).Error; err != nil {
		logger.Error("error while getting data from cms_salutation_master")
		return nil, err
	}
	return &salutation, nil
}

func (lr leadRepo) GetCmsLeadsShippingAddress(ctx context.Context, leadId string) (*entity.CmsLeadsShippingAddressMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_account_product_segment_master where leadId  ", leadId)
	var cmsLeadsShippingAddressMaster entity.CmsLeadsShippingAddressMaster
	if err := lr.db.Table("dbo.cms_leads_shipping_address_master").Model(&entity.CmsLeadsShippingAddressMaster{}).Where("leadid=?", leadId).Find(&cmsLeadsShippingAddressMaster).Error; err != nil {
		logger.Error("error while getting data from cms_salutation_master")
		return nil, err
	}
	return &cmsLeadsShippingAddressMaster, nil
}

func (lr leadRepo) GetCmsAccountProductSegment(ctx context.Context, productsegmentid int) (*entity.CmsAccountProductSegmentMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_account_product_segment_master where leadId  ", productsegmentid)
	var cmsAccountProductSegmentMaster entity.CmsAccountProductSegmentMaster
	if err := lr.db.Table("dbo.cms_account_product_segment_master").Where("productsegmentid=?", productsegmentid).Find(&cmsAccountProductSegmentMaster).Error; err != nil {
		logger.Error("error while getting data from cms_salutation_master")
		return nil, err
	}
	return &cmsAccountProductSegmentMaster, nil
}

func (lr leadRepo) GetCmsPhonecodes(ctx context.Context, contact_extid int64) (*entity.CmsPhonecodesMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_phonecodes_master where contact_extid ", contact_extid)
	var cmsPhonecodesMaster entity.CmsPhonecodesMaster
	//if err := lr.db.Table("cms_phonecodes_master").Model(&entity.CmsPhonecodesMaster{}).Where("id=?", contact_extid).Find(&cmsPhonecodesMaster).Error; err != nil {
	//	logger.Error("error while getting data from cms_phonecodes_master")
	//	return nil, err
	//}

	sqlStatement := `SELECT * FROM "dbo.cms_phonecodes_master" where id=$1`
	rows, err := lr.db.Raw(sqlStatement, contact_extid).Rows()
	if err != nil {
		logger.Info(err, "unable to add contact ext code", contact_extid)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&cmsPhonecodesMaster.Id, &cmsPhonecodesMaster.Countryname, &cmsPhonecodesMaster.Dial)
		logger.Info("added dial code")
	}
	return &cmsPhonecodesMaster, nil
}

func (lr leadRepo) GetCmsAccountType(ctx context.Context, accountTypeID int64) (*entity.CmsAccountTypeMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_account_type_master where accountId ", accountTypeID)
	var cmsAccountTypeMaster entity.CmsAccountTypeMaster
	if err := lr.db.Table("dbo.cms_account_type_master").Model(&entity.CmsAccountTypeMaster{}).Where("accounttypeid=?", accountTypeID).Find(&cmsAccountTypeMaster).Error; err != nil {
		logger.Error("error while getting data from cms_phonecodes_master")
		return nil, err
	}
	return &cmsAccountTypeMaster, nil
}

func (lr leadRepo) GetcmsCoffeetype(ctx context.Context, Id int64) (*entity.CmsCoffeetypeMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_coffeetype_master where leadId  ", Id)
	var cmsCoffeetypeMaster entity.CmsCoffeetypeMaster
	if err := lr.db.Table("dbo.cms_coffeetype_master").Model(&entity.CmsCoffeetypeMaster{}).Where("id=?", Id).Find(&cmsCoffeetypeMaster).Error; err != nil {
		logger.Error("error while getting data from cms_coffeetype_master")
		return nil, err
	}
	return &cmsCoffeetypeMaster, nil
}

func (lr leadRepo) GetCmsLeadsBillingAddress(ctx context.Context, leadID string) (*entity.CmsLeadsBillingAddressMaster, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("getting record from cms_leads_billing_address_master where leadId  ", leadID)
	var cmsLeadsBillingAddressMaster entity.CmsLeadsBillingAddressMaster
	if err := lr.db.Table("dbo.cms_leads_billing_address_master").Model(&entity.CmsLeadsBillingAddressMaster{}).Where("leadid=?", leadID).Find(&cmsLeadsBillingAddressMaster).Error; err != nil {
		logger.Error("error while getting data from cms_coffeetype_master")
		return nil, err
	}
	return &cmsLeadsBillingAddressMaster, nil
}

func (lr leadRepo) GetProdPackcategoryName(ctx context.Context, categoryId int) (*string, error) {
	type categoryName struct {
		Categorytypename string
	}
	var result categoryName
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	sqlStatement := `select categorytypename from cms_prod_pack_category_type where categorytypeid=$1`
	if err := lr.db.Raw(sqlStatement, categoryId).Scan(&result).Error; err != nil {
		logger.Error("error while get all account details from Accont master ", err.Error())
		return nil, err
	}

	return &result.Categorytypename, nil
}

func (lr leadRepo) GetProdPackCategoryWeight(ctx context.Context, weightId int) (*string, error) {
	type Weight struct {
		Weightname string
	}
	var result Weight
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	sqlStatement := `select weightname from cms_prod_pack_category_weight where weightid=$1`
	if err := lr.db.Raw(sqlStatement, weightId).Scan(&result).Error; err != nil {
		logger.Error("error while get all account details from Accont master ", err.Error())
		return nil, err
	}
	return &result.Weightname, nil
}

func (lr leadRepo) GetQuoatotionCreateInfoReq(ctx context.Context, req model.GetQuoatotionCreateInfoReq) (interface{}, error) {
	const (
		INCOTERMS        = "incoterms"
		CURRENCIES       = "currencies"
		LOADINGPORTS     = "loadingports"
		DESTINATIONPORTS = "destinationports"
		ACCOUNTDETAILS   = "accountdetails"
		VIEWQUOTE        = "viewquote"
		REQUESTPRICE     = "requestprice"
	)
	logger := logger2.GetLoggerWithContext(ctx)
	if req.Type == INCOTERMS {
		logger.Info("going to fetch all records from cms_incoterms_master")
		var response entity.CmsIncotermsMaster
		err := lr.db.Table("cms_incoterms_master").Model(&entity.CmsIncotermsMaster{}).Find(&response).Error
		if err != nil {
			logger.Error("error while get all account details from cms_leads_master ", err.Error())
			return nil, err
		}
		return response, err
	}
	if req.Type == CURRENCIES {
		sqlStatement := `SELECT "currencyid", "currencyname", "currencycode" FROM "project_currency_master"`
		rows, err := lr.db.Table("project_currency_master").Raw(sqlStatement).Rows()
		if err != nil {
			logger.Error("error while get all account details from project_currency_master ", err.Error())
			return nil, err
		}
		var allCurrencies []model.Currencies
		defer rows.Close()
		for rows.Next() {
			var currency model.Currencies
			err = rows.Scan(&currency.Currencyid, &currency.Currencyname, &currency.Currencycode)
			allCurrencies = append(allCurrencies, currency)
		}
		return allCurrencies, nil
	}
	if req.Type == LOADINGPORTS {
		logger.Info("going to fetch all records from cms_portloading_master")
		var response model.Loadingports
		err := lr.db.Table("cms_portloading_master").Model(&entity.CmsPortloadingMaster{}).Find(&response).Error
		if err != nil {
			logger.Error("error while get all account details from cms_portloading_master ", err.Error())
			return nil, err
		}
		return response, err
	}
	if req.Type == DESTINATIONPORTS {
		logger.Info("going to fetch all records from cms_destination_master")
		var response model.Destinationports
		err := lr.db.Table("cms_destination_master").Model(&entity.CmsDestinationMaster{}).Find(&response).Error
		if err != nil {
			logger.Error("error while get all account details from cms_destination_master ", err.Error())
			return nil, err
		}
		return response, err
	}
	if req.Type == ACCOUNTDETAILS {
		var allAccounts []model.AccountDetails
		sqlStatement := `select a.accountid, a.accountname, a.accounttypeid, concat(b.street,' ', b.city,' ', b.stateprovince, ' ', b.postalcode) as address
		from 
	   accounts_master a
		INNER JOIN accounts_billing_address_master b on b.accountid =a.accountid
		INNER Join cms_leads_master l on l.accountid = a.accountid`

		rows, err := lr.db.Raw(sqlStatement).Rows()

		defer rows.Close()
		for rows.Next() {
			var account model.AccountDetails
			err = rows.Scan(&account.Accountid, &account.Accountname, &account.Accounttypeid, &account.Address)

			var accounttypes []string
			if account.Accounttypeid != "" {
				z := strings.Split(account.Accounttypeid, ",")
				for i, z := range z {
					logger.Info("get account name", i, z)
					sqlStatement := `SELECT accounttype FROM "cms_account_type_master" where accounttypeid=$1`
					rows1, err1 := lr.db.Raw(sqlStatement, z).Rows()

					if err1 != nil {
						logger.Error(err, "unable to add account names")
					}

					for rows1.Next() {
						var accounttype string
						err = rows1.Scan(&accounttype)
						accounttypes = append(accounttypes, accounttype)
					}
				}
			}
			account.Accounttypename = strings.Join(accounttypes, ",")
			if account.Accountid != 0 {
				sqlStatement := `SELECT contactid, contactfirstname FROM "contacts_master" where accountid=$1`
				rows2, err2 := lr.db.Raw(sqlStatement, account.Accountid).Rows()

				if err2 != nil {
					logger.Error(err)
					logger.Info("unable to add account type", account.Accountid)
				}

				for rows2.Next() {
					var contact model.ContactDetails
					err = rows2.Scan(&contact.Contactid, &contact.Contactname)
					allAccounts := append(account.Contacts, contact)
					account.Contacts = allAccounts
					logger.Info("added one account", allAccounts)
				}
			}
			allAccounts = append(allAccounts, account)
		}
		return allAccounts, nil
	}
	if req.Type == VIEWQUOTE {
		logger.Info("get account details", req.Type)
		sqlStatement := `SELECT q.accountid,a.accountname, q.accounttypename, q.contactid, t.contactfirstname, q.createddate, u.firstname as createdby , 
        r.currencyname,
        q.currencycode,
        q.currencyid,
        q.fromdate,
		q.todate,
		q.paymentterm,
		q.otherspecification,
		q.remarks,
		q.destinationcountryid,
		q.destination,
		q.finalaccountid,
		concat(b.street,' ', b.city,' ', b.stateprovince, ' ', b.postalcode) as address,
		c.incoterms,
		q.incotermsid,
		s.status,
		q.portloading,
        q.portloadingid,
		q.destinationid,
		q.remarksfromgmc from crm_quote_master q
        INNER JOIN accounts_master a on q.accountid = a.accountid 
		INNER JOIN accounts_billing_address_master b on q.accountid = b.accountid
		INNER JOIN cms_incoterms_master c on q.incotermsid = c.incotermsid
        INNER JOIN cms_allstatus_master s ON q.statusid = s.id
		INNER JOIN userdetails_master u ON q.createdby = u.userid
        INNER JOIN contacts_master t ON q.contactid = t.contactid
		INNER JOIN project_currency_master r ON q.currencyid = r.currencyid
		where q.quoteid =$1`
		rows, err := lr.db.Raw(sqlStatement, req.QuoteId).Rows()
		if err != nil {
			logger.Error("error while getting data ", err.Error())
			return nil, err
		}
		var quote model.GetQuoteDetails
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&quote.Accountid,
				&quote.Accountname,
				&quote.Accounttypename,
				&quote.Contactid,
				&quote.Contactname,
				&quote.CreatedDate,
				&quote.Createdby,
				&quote.Currencyname,
				&quote.Currencycode,
				&quote.Currencyid,
				&quote.Fromdate,
				&quote.Todate,
				&quote.Paymentterms,
				&quote.Otherspecifications,
				&quote.Remarksfrommarketing,
				&quote.Destinationcountryid,
				&quote.Destination,
				&quote.Finalclientaccountid,
				&quote.Billingaddress,
				&quote.Incoterms,
				&quote.Incotermsid,
				&quote.Status,
				&quote.PortLoading,
				&quote.Portloadingid,
				&quote.Portdestinationid,
				&quote.Remarksfromgmc)

		}

		if quote.Finalclientaccountid != "" {

			sqlStatement := `select accountname as address from accounts_master b where accountid=$1`

			rows, err := lr.db.Raw(sqlStatement, quote.Finalclientaccountid).Rows()

			if err != nil {
				logger.Info(err, "unable to add final account name", quote.Finalclientaccountid)
				return nil, err
			}

			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&quote.Finalclientaccountname)
				logger.Info("added final account name")
			}
		}
		return quote, nil
	}

	if req.Type == REQUESTPRICE {
		logger.Info("update request price status", req.Type)
		sqlStatement := `UPDATE crm_quote_master SET statusid=2 where quoteid=$1`
		rows, err := lr.db.Raw(sqlStatement, req.QuoteId).Rows()
		if err != nil {
			logger.Info("error while getting data ", err.Error())
			return err, nil
		}
		return rows, nil
	}
	return nil, nil
}

/*
*
 */
func (leadRepoRef leadRepo) UpdateLead(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest) error {

	db := leadRepoRef.db
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	var opError error
	var success bool

	success, opError = doUpdateLead(ctx, requestPayload, db)

	if opError == nil && success {

		success, opError = leadRepoRef.UpdateLeadsBillingaddress(ctx,
			requestPayload, db)
	}
	if success && opError == nil {

		success, opError = leadRepoRef.UpdateLeadsShippingaddress(ctx,
			requestPayload, db)
	}
	if success && opError == nil {

		success, opError = leadRepoRef.EditAuditLogEntryOnLeadAmend(ctx,
			requestPayload, db)
	}
	if success && opError == nil {

		success, opError = leadRepoRef.LogNotificationOnLeadCreateUpdate(requestPayload.LeadId,
			`Lead Updated`, ctx, requestPayload, db)
	}
	if success && opError == nil {
		return tx.Commit().Error
	}
	if !success && opError != nil {
		tx.Rollback()
	}
	return opError
}

/*
*
 */
func (leadRepoRef leadRepo) CreateNewLead(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest) error {

	db := leadRepoRef.db
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	var opError error
	newLeadid, insertLeadErr := leadRepoRef.InsertLeadRecord(ctx, requestPayload, db)
	if insertLeadErr == nil {

		var success bool
		success, opError = leadRepoRef.LogNotificationOnLeadCreateUpdate(newLeadid, `Lead Created`, ctx,
			requestPayload, db)
		if success && opError == nil {

			success, opError = leadRepoRef.AddBillingAddressOnNewLead(newLeadid, ctx,
				requestPayload, db)
		}
		if success && opError == nil {

			success, opError = leadRepoRef.AddShippingAddressOnNewLead(newLeadid, ctx,
				requestPayload, db)
		}
		if success && opError == nil {

			success, opError = leadRepoRef.LogToAuditLogOnNewLead(newLeadid, ctx,
				requestPayload, db)
		}
		if success && opError == nil {

			return tx.Commit().Error
		}
		if !success || opError != nil {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
		opError = insertLeadErr
	}
	return opError
}

/*
*
 */
func (leadRepoRef leadRepo) LeadExists(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest) (bool, error) {

	query := `select accountname from dbo.cms_leads_master where accountname = $1`
	rows, err := leadRepoRef.db.Raw(query, requestPayload.Accountname).Rows()
	defer rows.Close()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while executing query to check if record exists in dbo.cms_leads_master with given Lead Name ", err.Error())
		return true, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

/*
*
 */
func (leadRepoRef leadRepo) InsertLeadRecord(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (string, error) {

	newLeadNum, err := getNextLeadNumber(ctx, db)
	if err == nil && newLeadNum != 0 {
		newLeadId, err := doInsertLead(newLeadNum, ctx, requestPayload, db)
		if err == nil && newLeadId != "" {
			return newLeadId, nil
		}
	}
	return "", err
}

/*
*
 */
func (leadRepoRef leadRepo) AddBillingAddressOnNewLead(leadId string, ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	insertBAddressRecordSQL := `INSERT INTO dbo.cms_leads_billing_address_master 
		(leadid, billingid, street, city, stateprovince, postalcode, country) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := leadRepoRef.db.Raw(insertBAddressRecordSQL, leadId, leadId,
		requestPayload.BillingStreetAddress, requestPayload.BillingCity,
		requestPayload.BillingState, requestPayload.BillingPostalCode,
		requestPayload.BillingCountry).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while creating Billing Address on New Lead Creation in dbo.cms_leads_billing_address_master ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) AddShippingAddressOnNewLead(leadId string, ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	insertShipAddressRecordSQL := `INSERT INTO dbo.cms_leads_shipping_address_master 
		(leadid, shippingid, street, city, stateprovince, postalcode, country) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := leadRepoRef.db.Raw(insertShipAddressRecordSQL, leadId, leadId,
		requestPayload.ContactStreetAddress, apputils.NullColumnValue(requestPayload.ContactCity),
		requestPayload.ContactState, requestPayload.ContactPostalCode,
		requestPayload.ContactCountrycode).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while creating Shipping Address on New Lead Creation in dbo.cms_leads_shipping_address_master ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) LogToAuditLogOnNewLead(leadId string, ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	insertAuditLogRecordSQL := `INSERT INTO dbo.auditlog_cms_leads_master_newpg(
		leadid, createdby, created_date, description)
		VALUES($1, $2, $3, $4)`

	_, err := db.Raw(insertAuditLogRecordSQL, leadId,
		requestPayload.CreatedUserid, requestPayload.CreatedDate,
		`Lead Created`).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while adding AUDIT LOG record on New Lead Creation in dbo.auditlog_cms_leads_master_newpg ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) UpdateLeadsBillingaddress(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	updateLeadBARecordSQL := `UPDATE dbo.cms_leads_billing_address_master 
		SET street=$1, city=$2, stateprovince=$3, 
		postalcode=$4, country=$5 where billingid=$6`

	_, err := db.Raw(updateLeadBARecordSQL, requestPayload.BillingStreetAddress,
		requestPayload.BillingCity, requestPayload.BillingState, requestPayload.BillingPostalCode,
		requestPayload.BillingCountry, requestPayload.LeadId).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while updating Leads Billing Address in dbo.cms_leads_billing_address_master ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) UpdateLeadsShippingaddress(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	updateLeadShpAdrsRecordSQL := `UPDATE dbo.cms_leads_shipping_address_master 
			SET street=$1, city=$2, stateprovince=$3, 
			postalcode=$4, country=$5 where shippingid=$6`

	_, err := db.Raw(updateLeadShpAdrsRecordSQL, requestPayload.ContactStreetAddress,
		apputils.NullColumnValue(requestPayload.ContactCity), requestPayload.ContactState, requestPayload.ContactPostalCode,
		requestPayload.ContactCountrycode, requestPayload.LeadId).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while updating Leads Shipping Address in dbo.cms_leads_billing_address_master ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) EditAuditLogEntryOnLeadAmend(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	updateAuditLogEntrySQL := `update dbo.auditlog_cms_leads_master_newpg
		set	description=$1, modifiedby=$2, modified_date=$3 where leadid=$4`

	_, err := db.Raw(updateAuditLogEntrySQL,
		"Lead Details Modified",
		requestPayload.ModifiedUserid,
		time.Now().Format("2006-01-02"),
		requestPayload.LeadId).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while updating AUDIT LOG entry in dbo.auditlog_cms_leads_master_newpg on Lead Details Update ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) LogNotificationOnLeadCreateUpdate(leadId string, status string, ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	insertNotifRecordSQL := `insert into dbo.notifications_master_newpg 
	(userid, objid, status, feature_category) 
	values($1, $2, $3, 'Lead')`

	_, err := leadRepoRef.db.Raw(insertNotifRecordSQL,
		requestPayload.CreatedUserid, leadId, status).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while Loggin Notification Record on New Lead Creation in dbo.notifications_master_newpg ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func getNextLeadNumber(ctx context.Context, db *gorm.DB) (int64, error) {
	type LeadId struct{ LeadNumber int64 }
	var leadIdStruct LeadId
	query := `SELECT idsno as LeadNumber FROM dbo.cms_leads_master where idsno is not null ORDER BY idsno DESC LIMIT 1`
	rows, err := db.Raw(query).Rows()
	defer rows.Close()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("error reading last lead number from cms_leads_master", err.Error())
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&leadIdStruct.LeadNumber)
	}
	return (leadIdStruct.LeadNumber + 1), nil
}

/*
*
 */
func makeLeadId(newLeadNumber int64) string {
	return "Lead-" + strconv.FormatInt(newLeadNumber, 10)
}

/*
*
 */
func doInsertLead(newLeadNum int64, ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (string, error) {

	newLeadId := makeLeadId(newLeadNum)

	insertLeadRecordSQL := `INSERT INTO dbo.cms_leads_master ( leadid, autogencode,
		legacyid, accountname, accounttypeid, phone, email, createddate, createduserid,
		shipping_continentid, countryid, approxannualrev, website, productsegmentid,
		leadscore, masterstatus,contactfirstname, contactlastname, manfacunit, instcoffee,
		price, approvalstatus, contact_salutationid, contact_position, contact_mobile,
		shipping_continent, shipping_country, coffeetypeid, aliases, isactive,
		otherinformation, contact_ext_id ) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, 
		$15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, 
		$28, $29, $30, $31, $32)`

	_, err := db.Raw(insertLeadRecordSQL,
		newLeadId, newLeadId, newLeadNum,
		requestPayload.Accountname, requestPayload.Accounttypeid,
		requestPayload.ContactPhone, requestPayload.ContactEmail, requestPayload.CreatedDate,
		requestPayload.CreatedUserid, requestPayload.ShippingContinentid,
		requestPayload.ShippingCountryid, requestPayload.Approximativeannualrevenue,
		requestPayload.Website, requestPayload.Productsegmentid,
		requestPayload.Leadscore, requestPayload.Masterstatus,
		requestPayload.Contactfirstname, requestPayload.Contactlastname,
		requestPayload.Manfacunit, requestPayload.Instcoffee,
		requestPayload.Price, requestPayload.Approvalstatus,
		requestPayload.ContactSalutationid, requestPayload.ContactPosition,
		requestPayload.ContactMobile, requestPayload.ShippingContinent,
		requestPayload.ShippingCountry,
		apputils.NullColumnValue(requestPayload.Coffeetypeid),
		requestPayload.Aliases, requestPayload.Isactive,
		apputils.NullColumnValue(requestPayload.OtherInformation),
		requestPayload.Contact_ext).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while inserting New Lead Record to dbo.cms_leads_master ", err.Error())
		return "", err
	}
	return newLeadId, nil
}

/*
*
 */
func doUpdateLead(ctx context.Context,
	requestPayload model.InsertLeadDetailsRequest, db *gorm.DB) (bool, error) {

	updateLeadRecordSQL := `UPDATE dbo.cms_leads_master SET 
	accountname=$1, accounttypeid=$2, contact_mobile=$3,
	email=$4, phone=$5, modifieddate=$6, modifieduserid=$7,
	shipping_continentid=$8, countryid=$9, approxannualrev=$10,
	website=$11, productsegmentid=$12, leadscore=$13,
	contactfirstname=$14, contactlastname=$15, manfacunit=$16,
	instcoffee=$17, price=$18, contact_salutationid=$19, contact_position=$20,
	shipping_continent=$21, shipping_country=$22,
	coffeetypeid=$23, aliases=$24, otherinformation=$25, contact_ext_id=$26
	where leadid=$27`

	_, err := db.Raw(updateLeadRecordSQL,
		requestPayload.Accountname, requestPayload.Accounttypeid,
		requestPayload.ContactMobile, requestPayload.ContactEmail,
		requestPayload.ContactPhone, requestPayload.ModifiedDate,
		requestPayload.ModifiedUserid, requestPayload.ShippingContinentid,
		requestPayload.ShippingCountryid, requestPayload.Approximativeannualrevenue,
		requestPayload.Website, requestPayload.Productsegmentid,
		requestPayload.Leadscore, requestPayload.Contactfirstname,
		requestPayload.Contactlastname, requestPayload.Manfacunit,
		requestPayload.Instcoffee, requestPayload.Price,
		requestPayload.ContactSalutationid, requestPayload.ContactPosition,
		requestPayload.ShippingContinent, requestPayload.ShippingCountry,
		requestPayload.Coffeetypeid, requestPayload.Aliases,
		requestPayload.OtherInformation, requestPayload.Contact_ext,
		requestPayload.LeadId).Rows()

	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while updating Lead Record in dbo.cms_leads_master ", err.Error())
		return false, err
	}
	return true, nil
}

/*
*
 */
func (leadRepoRef leadRepo) GetLeadCreationInfo(ctx context.Context,
	reqParams *model.GetLeadCreationInfoRequest) (interface{}, error) {

	var resultObj interface{}
	var opError error
	resultObj = nil
	opError = nil

	requestedInfoType := reqParams.Type
	const (
		ACCOUNTDETAILS  = "accountDetails"
		PRODUCTSEGMENTS = "productsegments"
		PHONECODES      = "phonecodes"
		COUNTRIES       = "countries"
		COFFEETYPES     = "coffeetypes"
		SALUTATIONS     = "salutations"
		CONTINENTS      = "continents"
	)
	switch {

	case requestedInfoType == ACCOUNTDETAILS:
		resultObj, opError = getAccountTypeInfo(ctx, leadRepoRef)
	case requestedInfoType == PRODUCTSEGMENTS:
		resultObj, opError = getAccountProductSegments(ctx, leadRepoRef)
	case requestedInfoType == PHONECODES:
		resultObj, opError = getPhoneCodes(ctx, leadRepoRef)
	case requestedInfoType == COUNTRIES:
		resultObj, opError = getContinentCountries(reqParams.ContinentName, ctx, leadRepoRef)
	case requestedInfoType == COFFEETYPES:
		resultObj, opError = getCoffeeTypes(ctx, leadRepoRef)
	case requestedInfoType == SALUTATIONS:
		resultObj, opError = getSalutations(ctx, leadRepoRef)
	case requestedInfoType == CONTINENTS:
		resultObj, opError = getContinentsInfo(ctx, leadRepoRef)
	default:
		errMsg := fmt.Sprintf("Requested Information Type should be either of %s, %s, %s ,%s, %s, %s, %s", ACCOUNTDETAILS, PRODUCTSEGMENTS, PHONECODES, COUNTRIES, COFFEETYPES, SALUTATIONS, CONTINENTS)
		return nil, errors.New(errMsg)
	}
	return resultObj, opError
}

/*
*
 */
func getAccountTypeInfo(ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var accountTypeInfo []model.AccountsInformation
	query := `SELECT accounttypeid, accounttype FROM dbo.cms_account_type_master`
	db := leadRepoRef.db
	err := db.Raw(query).Scan(&accountTypeInfo).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getAccountTypeInfo in querying dbo.cms_account_type_master ", err.Error())
		return nil, err
	}
	return accountTypeInfo, err
}

/*
*
 */
func getAccountProductSegments(ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var productSegments []model.ProductSegments
	query := `SELECT productsegmentid,productsegment FROM dbo.cms_account_product_segment_master`
	db := leadRepoRef.db
	err := db.Raw(query).Scan(&productSegments).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getAccountProductSegments in querying dbo.cms_account_product_segment_master ", err.Error())
		return nil, err
	}
	return productSegments, err
}

/*
*
 */
func getPhoneCodes(ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var allPhoneCodes []model.PhoneCodes
	db := leadRepoRef.db
	query := `SELECT id, Country_Name, Dial FROM dbo.cms_phonecodes_master`
	rows, err := db.Raw(query).Rows()
	defer rows.Close()
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getPhoneCodes in querying dbo.cms_phonecodes_master ", err.Error())
		return nil, err
	}
	var onePhoneCodeEntry model.PhoneCodes
	for rows.Next() {
		err = rows.Scan(&onePhoneCodeEntry.Id, &onePhoneCodeEntry.Countryname, &onePhoneCodeEntry.Dialcode)
		allPhoneCodes = append(allPhoneCodes, onePhoneCodeEntry)
	}
	return allPhoneCodes, err
}

/*
*
 */
func getContinentCountries(continent string, ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var countriesInfo []model.AccountsInformation
	query := `select countryname from dbo.continents_countries_master where continentname=$1`
	db := leadRepoRef.db
	err := db.Raw(query, continent).Scan(&countriesInfo).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getContinentCountries in querying dbo.continents_countries_master ", err.Error())
		return nil, err
	}
	return countriesInfo, err
}

/*
*
 */
func getCoffeeTypes(ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var allCoffeeTypes []model.CoffeeTypes
	query := `SELECT id, coffeetype FROM dbo.cms_coffeetype_master`
	db := leadRepoRef.db
	err := db.Raw(query).Scan(&allCoffeeTypes).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getCoffeeTypes in querying dbo.cms_coffeetype_master ", err.Error())
		return nil, err
	}
	return allCoffeeTypes, err
}

/*
*
 */
func getSalutations(ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var salutionsInfo []model.Salutations
	query := `SELECT salutationid, salutation FROM dbo.cms_salutation_master where isactive=$1`
	db := leadRepoRef.db
	err := db.Raw(query, 1).Scan(&salutionsInfo).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getSalutations in querying dbo.cms_salutation_master ", err.Error())
		return nil, err
	}
	return salutionsInfo, err
}

/*
*
 */
func getContinentsInfo(ctx context.Context, leadRepoRef leadRepo) (interface{}, error) {

	var continentsInfo []model.Continents
	query := `SELECT distinct "continent_name", "continent_code" FROM "continents"`
	db := leadRepoRef.db
	err := db.Raw(query).Scan(&continentsInfo).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: getContinentsInfo in querying public.continents ", err.Error())
		return nil, err
	}
	return continentsInfo, err
}

/*
*
 */
func (leadRepoRef leadRepo) ProvideLeadsData(ctx context.Context,
	reqParams model.ProvideLeadsInfoReqContext) ([]model.LeadInfo, error) {

	var leads []model.LeadInfo
	queryPart := `ORDER BY createddate desc`
	if reqParams.Filter != "" {
		queryPart = "where " + reqParams.Filter + " ORDER BY createddate desc"
	}
	query := fmt.Sprintf(`SELECT leadid, accountname, aliases, contactfirstname, 
		contactlastname, contact_mobile, email, leadscore,masterstatus 
		FROM dbo.LeadsGrid %s`, queryPart)

	db := leadRepoRef.db
	err := db.Raw(query).Scan(&leads).Error
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: Lead Repository -> ProvideLeadsData in querying dbo.LeadsGrid ", err.Error())
		return nil, err
	}
	return leads, nil
}
