package accountrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"
)

type AccountRepo interface {
	//Create(context context.Context, accountDetailsRequest entity.AccountDetails) ()
	Insert(context context.Context, accountDetailsRequest entity.AccountDetails) (*model.AccountDetailsResponse, error)
	GetAllAccountDetails(ctx context.Context) ([]*model.GetAllAccountsResponseBody, error)
	GetAllLeadAccounts(ctx context.Context) ([]*model.GetAlleadsResonseBody, error)
	GetAllQuoteLineItems(ctx context.Context, req *model.GetAllQuoteLineRequest) ([]model.GetAllQouteLineItemsResponseBody, error)
	GetAllQoutes(ctx context.Context, req *model.GetAllQoutesRequestBody) ([]*model.QuoteDetails, error)
	//GetLeadCreationInfo(ctx context.Context, req *model.GetLeadCreationInfoRequest) (interface{}, error)
	//GetQuotationCreationInfo(ctx context.Context, req *model.GetQuoatotionCreateInfoReq) (interface{}, error)

}

const (
	MarketingMExecutive = "Marketing Executive"
	ManagingDirector    = "Managing Director"
	PENDINGQUOTES       = "pendingquotes"
	MYQOUTES            = "myquotes"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo() AccountRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}
	//Need to check with this part
	//newDb.AutoMigrate(&entity.AccountMaster{})
	return &accountRepo{
		db: newDb,
	}
}

func (r accountRepo) Create(context context.Context, account entity.AccountDetails) error {
	if err := r.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r accountRepo) Insert(context context.Context, accountDetailsRequest entity.AccountDetails) (*model.AccountDetailsResponse, error) {
	logger := logger2.GetLoggerWithContext(context)
	lead := accountDetailsRequest
	if lead.Role == MarketingMExecutive && lead.ConvertLeadToAccount {
		sqlQuery := `UPDATE CMS_LEADS_MASTER
						  	  SET 
						  	  masterstatus='Pending Approval'
						      WHERE 
					    	  leadid=$1`
		err := r.db.Update(sqlQuery, lead.LeadId).Error
		if err != nil {
			fmt.Println("error")
		}
	}
	if (lead.Role == MarketingMExecutive) && (lead.ConvertLeadToAccount) {
		//sqlQuery := `UPDATE CMS_LEADS_MASTER
		//				  	  SET
		//				  	  masterstatus='Pending Approval'
		//				      WHERE
		//			    	  leadid=$1`
		rows, err := r.db.DB().Query(repositories.UpdateQuery, lead.LeadId)
		logger.Info("Updated Status to Pending Approval")

		if err != nil {
			logger.Info(err.Error())
			return nil, errors.New(err.Error())
			return nil, err
			//return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()
		res, _ := json.Marshal(rows)
		logger.Info("response from database", string(res))
		return &model.AccountDetailsResponse{
			StatusCode: 200,
			Payload:    string(res),
		}, nil

		// LEAD REJECTION MODULE
	} else if (lead.Role == ManagingDirector) && (lead.Reject) {

		sqlQuery := `UPDATE CMS_LEADS_MASTER
							SET 
							masterstatus='Rejected',
							comments=$1
	 						WHERE 
	  						leadid=$2`
		logger.Info("Updating Status to Rejected")
		rows, err := r.db.DB().Query(sqlQuery, lead.Comments, lead.LeadId)

		if err != nil {
			logger.Info(err.Error())
			return nil, err
			//return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		fmt.Println(res)
		return &model.AccountDetailsResponse{
			StatusCode: 200,
			Payload:    string(res),
		}, nil
		//return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil
		// APPROVAL BY MD
	} else if (lead.Role == ManagingDirector) && (lead.ConvertLeadToAccount || lead.Approve) {
		sqlQuery := `UPDATE CMS_LEADS_MASTER
							  SET 
							  masterstatus='account Created',
					  		  accountid=(select floor(100 + random() * 899)::numeric)
					  		  WHERE 
					  		  leadid=$1`
		rows, err := r.db.DB().Query(sqlQuery, lead.LeadId)
		logger.Info("account ID assigned")
		sqlStatementina1 := `INSERT INTO Accounts_master (
			accountid,
			accountname,
			accounttypeid,
			phone,
			email,
			createddate,
			createduserid,
			approxannualrev,
			website,
			productsegmentid,
			masterstatus,
			recordtypeid,
			shipping_continent,
			shipping_country,
			comments,
			aliases,
			isactive,
			otherinformation)
			SELECT
			accountid,
			accountname,
			accounttypeid,
			phone,
			email,
			createddate,
			createduserid ,
			approxannualrev,
			website,
			productsegmentid,
			masterstatus,
			recordtypeid,
			shipping_continent,
			shipping_country,
			comments,
			aliases,
			isactive,
			otherinformation
			FROM
			cms_leads_master
			WHERE leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementina1, lead.LeadId)
		// Get Accountid from Lead Record
		// Set account status to Prospect in accounts_master
		sqlStatementstat1 := `UPDATE accounts_master
					 		 SET 
					 		 account_owner=u.username,
					 		 masterstatus='Prospect',
							 comments=$1
					 		 FROM accounts_master acc
					 		 INNER JOIN
					 		 CMS_LEADS_MASTER ld
					 		 ON ld.accountid = acc.accountid
					 		 INNER JOIN
					 		 userdetails_master u on u.userid=ld.createduserid
					 		 where ld.leadid=$2`

		rows, err = r.db.DB().Query(sqlStatementstat1, lead.Comments, lead.LeadId)
		sqlStatementapp1 := `UPDATE CMS_LEADS_MASTER
					  	  SET 
					  	  masterstatus='Appoved'
						  comments=$1	
					      WHERE 
					      leadid=$2`

		rows, err = r.db.DB().Query(sqlStatementapp1, lead.Comments, lead.LeadId)
		fmt.Println("Lead Status is set to Approved")
		sqlStatementcon1 := `insert into contacts_master(
			contactfirstname,
			contactlastname,
			contactemail,
			contactphone,
			contactmobilenumber,
			accountid,
			position,
			salutationid) 
			select
			contactfirstname,
			contactlastname,
			email,
			phone,
			contact_mobile,
			accountid,
			contact_position,
			contact_salutationid
			from
			cms_leads_master where leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementcon1, lead.LeadId)
		fmt.Println("Lead Contact data is inserted into Contacts_Master Successfully")

		//Insert into accounts_billing_address_master
		sqlStatementabm1 := `insert into accounts_billing_address_master(
			accountid,
			billingid,
			street,
			city,
			stateprovince,
			postalcode,
			country)
			select
			ld.accountid,
			lba.billingid,
			lba.street,
			lba.city,
			lba.stateprovince,
			lba.postalcode,
			lba.country
			from
			cms_leads_billing_address_master lba
			inner join
			cms_leads_master ld on ld.leadid=lba.leadid
			where ld.leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementabm1, lead.LeadId)
		fmt.Println("account Contact data is inserted into accounts_billing_address_master")
		//Insert into accounts_shipping_address_master
		sqlStatementasm1 := `insert into accounts_shipping_address_master(
			accountid,
			shippingid,
			street,
			city,
			stateprovince,
			postalcode,
			country)
			select
			ld.accountid,
			lsa.shippingid,
			lsa.street,
			lsa.city,
			lsa.stateprovince,
			lsa.postalcode,
			lsa.country
			from
			cms_leads_shipping_address_master lsa
			inner join
			cms_leads_master ld on ld.leadid=lsa.leadid
			where ld.leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementasm1, lead.LeadId)
		fmt.Println("account Contact data is inserted into accounts_shipping_address_master")

		if err != nil {
			logger.Info(err.Error())
			return nil, err
			//return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		fmt.Println(res)
		return &model.AccountDetailsResponse{
			StatusCode: 200,
			Payload:    res,
		}, nil
		//return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil
	}

	res1, _ := json.Marshal("Success")
	fmt.Println(res1)
	//return events.APIGatewayProxyResponse{200, headers, nil, string(res1), false}, nil

	return &model.AccountDetailsResponse{
		StatusCode: 200,
		Payload:    res1,
	}, nil
}

func (r accountRepo) InsertIntoAccountMaster(ctx context.Context, AccountMaster entity.AccountMaster) error {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to insert into account master table")
	if err := r.db.Create(&AccountMaster).Error; err != nil {
		logger.Error("error while inserting into Accont master ", err.Error())
		return err
	}
	return nil
}

func (r accountRepo) GetAllAccountDetails(ctx context.Context) ([]*model.GetAllAccountsResponseBody, error) {
	result := make([]*model.GetAllAccountsResponseBody, 0)
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to insert into account master table FROM accounts_master")
	if err := r.db.Raw(`SELECT accountname,aliases,accounttypeid,account_owner,masterstatus FROM dbo.accounts_master`).Scan(&result).Error; err != nil {
		logger.Error("error while get all account details from Accont master ", err.Error())
		return nil, err
	}
	//if err := r.db.Raw(`SELECT "accountname","aliases","accounttypeid","account_owner","masterstatus" FROM "accounts_master"`).Scan(&result).Error; err != nil {
	//	logger.Error("error while get all account details from Accont master ", err.Error())
	//	return nil, err
	//}
	return result, nil
}

func (r accountRepo) GetAllLeadAccounts(ctx context.Context) ([]*model.GetAlleadsResonseBody, error) {
	var result []*model.GetAlleadsResonseBody
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	rows, err := r.db.Raw(`SELECT "accountname","aliases","contactfirstname","contactlastname","phone","email","approvalstatus" FROM "cms_leads_master"`).Rows()

	if err != nil {
		logger.Error(err.Error())
	}
	for rows.Next() {
		var account model.GetAlleadsResonseBody
		err = rows.Scan(&account.AccountName, &account.Aliases, &account.Contactfirstname, &account.Contactlastname, &account.Phone, &account.Email, &account.ApprovalStatus)
		result = append(result, &account)
	}
	res, _ := json.Marshal(result)
	fmt.Println(string(res))
	return result, nil
}

func (r accountRepo) GetAllQuoteLineItems(ctx context.Context, req *model.GetAllQuoteLineRequest) ([]model.GetAllQouteLineItemsResponseBody, error) {
	var result []model.GetAllQouteLineItemsResponseBody
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	sqlStatement := getAllQuotesLineQuery()
	if err := r.db.Raw(sqlStatement, req.QuoteId).Scan(&result).Error; err != nil {
		logger.Error("error while get all account details from Accont master ", err.Error())
		return nil, err
	}
	return result, nil
}

func (r accountRepo) GetAllQoutes(ctx context.Context, req *model.GetAllQoutesRequestBody) ([]*model.QuoteDetails, error) {
	var result []*model.QuoteDetails
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all leadAccounts")
	sqlStatement := getAllQoutesQuery(req.Type)
	if req.Type == MYQOUTES {
		if err := r.db.Raw(sqlStatement, req.Createdby).Scan(&result).Error; err != nil {
			logger.Error("error while get all account details from Accont master ", err.Error())
			return nil, err
		}
		return result, nil
	}
	if err := r.db.Raw(sqlStatement).Scan(&result).Error; err != nil {
		logger.Error("error while get all account details from Accont master ", err.Error())
		return nil, err
	}
	return result, nil
}

func getAllQuotesLineQuery() string {
	const sqlStatement = `select q.quoteitemid,q.quoteid,q.sampleid,q.expectedorderqty, s.categoryname,q.packcategorytypeid,q.packweightid, q.packupcid
	from cms_quote_item_master q
	INNER JOIN cms_prod_pack_category s on q.packcategoryid = s.categoryid
	where q.quoteid=$1 order by q.createddate desc`
	return sqlStatement
}

func getAllQoutesQuery(queryType string) string {
	var sqlStatement string
	if queryType == PENDINGQUOTES {
		sqlStatement = `SELECT q.quoteid,a.accountname,s.status,u.firstname as createdby,q.createddate 
		FROM crm_quote_master q
		INNER JOIN accounts_master a ON q.accountid = a.accountid 
		INNER JOIN userdetails_master u ON q.createdby = u.userid
		INNER JOIN cms_allstatus_master s ON q.statusid = s.id
		where q.statusid=2 order by createddate desc`
		return sqlStatement
	}
	if queryType == MYQOUTES {
		sqlStatement = `SELECT q.quoteid,a.accountname,s.status,u.firstname as createdby,q.createddate 
		FROM crm_quote_master q
		INNER JOIN accounts_master a ON q.accountid = a.accountid 
		INNER JOIN userdetails_master u ON q.createdby = u.userid
		INNER JOIN cms_allstatus_master s ON q.statusid = s.id
		where q.createdby=$1 order by createddate desc`
		return sqlStatement
	}
	sqlStatement = `SELECT q.quoteid,a.accountname,s.status,u.firstname as createdby,q.createddate 
								FROM crm_quote_master q
								INNER JOIN accounts_master a ON q.accountid = a.accountid 
								INNER JOIN userdetails_master u ON q.createdby = u.userid
								INNER JOIN cms_allstatus_master s ON q.statusid = s.id order by createddate desc`
	return sqlStatement
}

//func (r accountRepo) GetQuotationCreationInfo(ctx context.Context, req *model.GetQuoatotionReqBody) (interface{}, error) {
//	if req.Type == "incoterms" {
//		var allInCoterms []model.InCotermsInfo
//		log.Println("get incoterms", req.Type)
//		sqlStatement := `SELECT * FROM "cms_incoterms_master"`
//		err := r.db.Raw(sqlStatement).Scan(&allInCoterms).Error
//		if err != nil {
//			return nil, err
//
//		}
//
//	}
//	return nil, nil
//}
