package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = ""
	port     = 12
	user     = ""
	password = "1"
	dbname   = "1"
)

type AccountDetails struct {
	LeadId               int    `json:"leadid"`
	Role                 string `json:"role"`
	ConvertLeadToAccount bool   `json:"convertleadtoaccount"`
	Approve              bool   `json:"approve"`
	Reject               bool   `json:"reject"`
	Comments             string `json:"comments"`
}

func insertLeadIntoAccount(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var lead AccountDetails
	err := json.Unmarshal([]byte(request.Body), &lead)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
	}
	defer db.Close()

	// check db
	err = db.Ping()

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
	}

	fmt.Println("Connected!")
	var rows *sql.Rows

	if (lead.Role == "Marketing Executive") && (lead.ConvertLeadToAccount) {

		sqlStatementl1 := `UPDATE CMS_LEADS_MASTER
						  	  SET 
						  	  masterstatus='Pending Approval'
						      WHERE 
					    	  leadid=$1`
		rows, err = db.Query(sqlStatementl1, lead.LeadId)
		log.Println("Updated Status to Pending Approval")

		if err != nil {
			log.Println(err.Error())
			return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil

		// LEAD REJECTION MODULE
	} else if (lead.Role == "Managing Director") && (lead.Reject) {

		sqlStatementmdr1 := `UPDATE CMS_LEADS_MASTER
							SET 
							masterstatus='Rejected',
							comments=$1
	 						WHERE 
	  						leadid=$2`
		rows, err = db.Query(sqlStatementmdr1, lead.Comments, lead.LeadId)
		log.Println("Updated Status to Rejected")

		if err != nil {
			log.Println(err.Error())
			return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil
		// APPROVAL BY MD
	} else if (lead.Role == "Managing Director") && (lead.ConvertLeadToAccount || lead.Approve) {

		sqlStatementacn1 := `UPDATE CMS_LEADS_MASTER
							  SET 
							  masterstatus='account Created',
					  		  accountid=(select floor(100 + random() * 899)::numeric)
					  		  WHERE 
					  		  leadid=$1`
		rows, err = db.Query(sqlStatementacn1, lead.LeadId)
		log.Println("account ID assigned")
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
		rows, err = db.Query(sqlStatementina1, lead.LeadId)
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

		rows, err = db.Query(sqlStatementstat1, lead.Comments, lead.LeadId)
		sqlStatementapp1 := `UPDATE CMS_LEADS_MASTER
					  	  SET 
					  	  masterstatus='Appoved'
						  comments=$1	
					      WHERE 
					      leadid=$2`

		rows, err = db.Query(sqlStatementapp1, lead.Comments, lead.LeadId)
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
		rows, err = db.Query(sqlStatementcon1, lead.LeadId)
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
		rows, err = db.Query(sqlStatementabm1, lead.LeadId)
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
		rows, err = db.Query(sqlStatementasm1, lead.LeadId)
		fmt.Println("account Contact data is inserted into accounts_shipping_address_master")

		if err != nil {
			log.Println(err.Error())
			return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil
	}

	res1, _ := json.Marshal("Success")
	return events.APIGatewayProxyResponse{200, headers, nil, string(res1), false}, nil
}

func main() {
	lambda.Start(insertLeadIntoAccount)
}
