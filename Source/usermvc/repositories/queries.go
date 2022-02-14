package repositories

import "fmt"

const (
	UpdateQuery = `UPDATE CMS_LEADS_MASTER
						  	  SET 
						  	  masterstatus='Pending Approval'
						      WHERE 
					    	  leadid=$1`
)

func getAllDetailsQuery(leadId int) string {
	return fmt.Sprintf("%s,%d", UpdateQuery, leadId)
}
