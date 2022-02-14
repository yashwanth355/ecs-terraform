package controller

import (
	"fmt"
	"strconv"
	"strings"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/accountrepo"
	"usermvc/repositories/leadrepo"
	"usermvc/repositories/userrepo"
	emailer "usermvc/utility/emailer"
	logger2 "usermvc/utility/logger"

	"github.com/gin-gonic/gin"
)

type LeadController interface {
	GetAllLeadsDetails(ctx *gin.Context)
	GetAllQuoteLineItems(ctx *gin.Context)
	GetAllQuotes(ctx *gin.Context)
	GetLeadCreationInfo(ctx *gin.Context)
	GetLeadDetails(ctx *gin.Context)
	GetLeadsInfo(ctx *gin.Context)
	GetQuotationCreationInfo(ctx *gin.Context)
	GetQuoteInformation(ctx *gin.Context)
	InsertLeadDetails(ctx *gin.Context)
}

type leadController struct {
	accountrepo accountrepo.AccountRepo
	leadRepo    leadrepo.LeadRepo
}

func newLeadController() LeadController {
	return &leadController{
		accountrepo: accountrepo.NewAccountRepo(),
		leadRepo:    leadrepo.NewLeadRepo(),
	}
}

func (lc leadController) GetAllLeadsDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := lc.leadRepo.GetAllLeadDetails(ctx)
	if err != nil {
		logger.Error("error while getting all leadaccounts", err.Error())
		ctx.JSON(503, err.Error())
	}
	logger.Info("getting response from get accountrepo.GetAllLeadAccounts ", res)
	ctx.JSON(200, res)
}

func (lc leadController) GetAllQuoteLineItems(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var getAllQoutedRequestBody *model.GetAllQuoteLineRequest
	logger.Info("start hadling leadController.leadController")
	if err := ctx.ShouldBindJSON(&getAllQoutedRequestBody); err != nil {
		logger.Error("error while parsing the qoutationItems", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	res, err := lc.accountrepo.GetAllQuoteLineItems(ctx, getAllQoutedRequestBody)
	if err != nil {
		logger.Error("error while getting all QuoteLineItens", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	for k, QuoteLineItem := range res {
		if QuoteLineItem.Packcategorytypeid != nil {
			categoryName, err := lc.leadRepo.GetProdPackcategoryName(ctx, *QuoteLineItem.Packcategorytypeid)
			if err != nil {
				logger.Error("error while getting category name where category_id is ", QuoteLineItem.Packcategorytypeid)
				ctx.JSON(503, err.Error())
				return
			}
			res[k].CategoryType = *categoryName
		}
		if QuoteLineItem.Packweightid != nil {
			weight, err := lc.leadRepo.GetProdPackCategoryWeight(ctx, *QuoteLineItem.Packweightid)
			if err != nil {
				logger.Error("error while getting  weight ", QuoteLineItem.Packweightid)
				ctx.JSON(503, err.Error())
				return
			}
			res[k].Weight = *weight
		}
	}
	ctx.JSON(200, res)
}

func (lc leadController) GetAllQuotes(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	validate, err := lc.validateGetAllQuotes(ctx)
	if err != nil {
		logger.Error("error while validating request", err.Error())
		ctx.JSON(503, "error while validating request")
		return
	}
	if !validate {
		logger.Error("request couls not be validate")
		ctx.JSON(404, "invalid request")
		return
	}
	var getAllQoutedRequestBody *model.GetAllQoutesRequestBody
	if err := ctx.ShouldBindJSON(&getAllQoutedRequestBody); err != nil {
		logger.Error("error while parsing the GetAllQuoteRequest", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	res, err := lc.accountrepo.GetAllQoutes(ctx, getAllQoutedRequestBody)
	if err != nil {
		logger.Error("error while getting allQuoteItems", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	ctx.JSON(200, res)
}

// need to work on it

func (lc leadController) GetLeadCreationInfo(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	validate, err := lc.validateGetLeadCreationInfo(ctx)
	if err != nil {
		logger.Error("error while validating request", err.Error())
		ctx.JSON(503, "error while validating request")
		return
	}
	if !validate {
		logger.Error("request could not be validate")
		ctx.JSON(404, "invalid request")
		return
	}
	var getLeadCreationInfoRequest *model.GetLeadCreationInfoRequest
	if err := ctx.ShouldBindJSON(&getLeadCreationInfoRequest); err != nil {
		logger.Error("error while parsing the GetLeadCreationInfoRequest", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	res, err := lc.leadRepo.GetLeadCreationInfo(ctx, getLeadCreationInfoRequest)
	if err != nil {
		logger.Error("error while getting GetLeadCreationInfo", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	ctx.JSON(200, res)
}

func (lc leadController) GetLeadDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	validate, err := lc.validateGetLeadDetails(ctx)
	if err != nil {
		logger.Error("error while validating request", err.Error())
		ctx.JSON(403, "error while validating request")
		return
	}
	if !validate {
		logger.Error("request could not be validate")
		ctx.JSON(403, "invalid request")
		return
	}
	var getLeadDetailsRequestBody model.GetLeadDetailsRequestBody
	if err := ctx.ShouldBindJSON(&getLeadDetailsRequestBody); err != nil {
		logger.Error("error while parsing the getLeadDetailsRequest", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	res, err := lc.leadRepo.GetCmsLeads(ctx, getLeadDetailsRequestBody)
	if err != nil {
		logger.Error("error while getting GetLeadCreationInfo", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	leadDetailResp := cmsLeadMasterToLeadDetail(res)
	Salutationid, err := lc.leadRepo.GetSalutation(ctx, leadDetailResp.ContactSalutationid)
	if err != nil {
		logger.Error("error while getting GetSalutation", err.Error())
		ctx.JSON(503, err.Error())
		return
	}

	leadDetailResp.Salutations = model.Salutations{
		Salutationid: Salutationid.Salutationid,
		Salutation:   Salutationid.Salutation,
	}

	cmsLeadsBillingAddressMaster, err := lc.leadRepo.GetCmsLeadsBillingAddress(ctx, getLeadDetailsRequestBody.Id)
	if err != nil {
		logger.Error("error while getting GetCmsLeadsBillingAddressMaster", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	leadDetailResp.BillingStreetAddress = cmsLeadsBillingAddressMaster.Street
	leadDetailResp.BillingCity = cmsLeadsBillingAddressMaster.City
	leadDetailResp.BillingState = cmsLeadsBillingAddressMaster.Stateprovince
	leadDetailResp.BillingPostalCode = cmsLeadsBillingAddressMaster.Postalcode
	leadDetailResp.BillingCountry = cmsLeadsBillingAddressMaster.Country

	shippingAddress, err := lc.leadRepo.GetCmsLeadsShippingAddress(ctx, getLeadDetailsRequestBody.Id)
	if err != nil {
		logger.Error("error while getting GetleadsShippingAddressMaster  ")
		ctx.JSON(503, err.Error())
		return

	}
	leadDetailResp.ContactStreetAddress = shippingAddress.Street
	leadDetailResp.ContactCity = shippingAddress.City
	leadDetailResp.ContactState = shippingAddress.Stateprovince
	leadDetailResp.ContactPostalCode = shippingAddress.Postalcode
	leadDetailResp.ContactCountry = shippingAddress.Country
	// if leadDetailResp.Productsegmentid != "" {
	// 	segmentIds := strings.Split(leadDetailResp.Productsegmentid, ",")
	// 	for _, segmentId := range segmentIds {
	// 		segmentIdToInt, err := strconv.Atoi(segmentId)
	// 		if err != nil {
	// 			logger.Info("invalid segment id ", segmentId)
	// 			continue
	// 		}
	// 		logger.Info("getting segment from segmentID repo", segmentId)
	// 		Productsegment, err := lc.leadRepo.GetCmsAccountProductSegment(ctx, segmentIdToInt)
	// 		if err != nil {
	// 			logger.Error("error while getting GetCmsAccountProductSegmentMaster", err.Error())
	// 			ctx.JSON(503, err.Error())
	// 			return
	// 		}
	// 		leadDetailResp.Productsegment = append(leadDetailResp.Productsegment, model.ProductSegments{
	// 			Productsegmentid: Productsegment.Productsegmentid,
	// 			Productsegment:   Productsegment.Productsegment,
	// 		})
	// 	}
	// }

	if leadDetailResp.Coffeetypeid != "" {
		coffeIds := strings.Split(leadDetailResp.Coffeetypeid, ",")
		for _, coffeId := range coffeIds {
			cofeeIdInt, err := strconv.ParseInt(coffeId, 10, 64)
			if err != nil {
				logger.Info("invalid cofeeId id ", coffeId)
				continue
			}
			logger.Info("getting segment from segmentID repo", coffeId)
			coffeType, err := lc.leadRepo.GetcmsCoffeetype(ctx, cofeeIdInt)
			if err != nil {
				logger.Error("error while getting GetCmsAccountProductSegmentMaster", err.Error())
				ctx.JSON(503, err.Error())
				return
			}
			leadDetailResp.CoffeeTypes = append(leadDetailResp.CoffeeTypes, model.CoffeeTypes{
				CoffeeType:   coffeType.Coffeetype,
				CoffeeTypeId: strconv.FormatInt(coffeType.ID, 10),
			})
		}

	}
	if leadDetailResp.Accounttypeid != "" {
		Accounttypeids := strings.Split(leadDetailResp.Accounttypeid, ",")
		for _, Accounttypeid := range Accounttypeids {
			AccounttypeidInt, err := strconv.ParseInt(Accounttypeid, 10, 64)
			if err != nil {
				logger.Info("invalid Accounttypeid id ", Accounttypeid)
				continue
			}
			logger.Info("getting GetCmsAccountTypeMaster  repo", Accounttypeid)
			AccountTypesResp, err := lc.leadRepo.GetCmsAccountType(ctx, AccounttypeidInt)
			if err != nil {
				logger.Error("error while getting GetCmsAccountTypeMaster", err.Error())
				ctx.JSON(503, err.Error())
				return
			}
			leadDetailResp.AccountTypes = append(leadDetailResp.AccountTypes, model.AccountsInformation{
				Accounttypeid: strconv.FormatInt(AccountTypesResp.Accounttypeid, 10),
				Accounttype:   AccountTypesResp.Accounttype,
			})
		}
	}

	if leadDetailResp.Contact_extid != 0 {
		contactId := leadDetailResp.Contact_extid
		if err != nil {
			logger.Info("invalid contact id")
		} else {
			contactResp, err := lc.leadRepo.GetCmsPhonecodes(ctx, contactId)
			if err != nil {
				logger.Error("error while getting contact detals", err.Error())
			} else {
				leadDetailResp.Contact_ext = model.PhoneCodes{}
				leadDetailResp.Contact_ext.Id = int(contactResp.Id)
				leadDetailResp.Contact_ext.Countryname = contactResp.Countryname
				leadDetailResp.Contact_ext.Dialcode = contactResp.Dial
			}
		}

	}
	ctx.JSON(200, leadDetailResp)
}

func cmsLeadMasterToLeadDetail(cmsLeads entity.CmsLeadsMaster) model.LeadDetails {
	return model.LeadDetails{
		Accountname:                cmsLeads.Accountname,
		Accounttypeid:              cmsLeads.Accounttypeid,
		ContactMobile:              cmsLeads.ContactMobile,
		ContactEmail:               cmsLeads.Email,
		Approximativeannualrevenue: cmsLeads.Approxannualrev,
		Website:                    cmsLeads.Website,
		Productsegmentid:           cmsLeads.Productsegmentid,
		Contactfirstname:           cmsLeads.Contactfirstname,
		Contactlastname:            cmsLeads.Contactlastname,
		Manfacunit:                 cmsLeads.Manfacunit,
		Instcoffee:                 cmsLeads.Instcoffee,
		Price:                      cmsLeads.Price,
		ContactSalutationid:        cmsLeads.ContactSalutationid,
		ContactPosition:            cmsLeads.ContactPosition,
		ContactPhone:               cmsLeads.ContactMobile,
		ShippingContinent:          cmsLeads.ShippingContinent,
		ShippingCountry:            cmsLeads.ShippingCountry,
		Coffeetypeid:               cmsLeads.Coffeetypeid,
		Aliases:                    cmsLeads.Aliases,
		OtherInformation:           cmsLeads.Otherinformation,
		Status:                     cmsLeads.Masterstatus,
		Leadscore:                  cmsLeads.Leadscore,
		//Contact_extid:              cmsLeads.ContactExt,
	}
}
func (lc leadController) GetQuotationCreationInfo(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var getQuoatotionCreateInfoReq model.GetQuoatotionCreateInfoReq
	if err := ctx.ShouldBindJSON(&getQuoatotionCreateInfoReq); err != nil {
		logger.Error("error while parsing the getQuoatotionCreateInfoReq", err.Error())
		ctx.JSON(403, err.Error())
		return
	}

	LeadsInfo, err := lc.leadRepo.GetQuoatotionCreateInfoReq(ctx, getQuoatotionCreateInfoReq)
	if err != nil {
		logger.Error("error while getting GetCmsLeadsBillingAddressMaster", err.Error())
		ctx.JSON(503, err.Error())
		return
	}

	ctx.JSON(200, LeadsInfo)
}

func (lc leadController) validateGetLeadsInfo(ctx *gin.Context) (bool, error) {
	return true, nil
}

func (lc leadController) validateGetLeadCreationInfo(ctx *gin.Context) (bool, error) {
	return true, nil
}

func (lc leadController) validateGetAllQuotes(ctx *gin.Context) (bool, error) {
	return true, nil
}

func (lc leadController) validateGetLeadDetails(ctx *gin.Context) (bool, error) {
	return true, nil
}

func (lc leadController) GetQuoteInformation(ctx *gin.Context) {

}

//func (lc leadController) 	GetQuotationCreationInfo(ctx *gin.Context) {
////	1) looger method
//	logger := logger2.GetLoggerWithContext(ctx)
//	validate, err := lc.validateGetLeadCreationInfo(ctx)
//	if err != nil {
//		logger.Error("error while validating request", err.Error())
//		ctx.JSON(503, "error while validating request")
//		return
//	}
//	if !validate {
//		logger.Error("request could not be validate")
//		ctx.JSON(404, "invalid request")
//		return
//	}
//	var getAllQoutedRequestBody *model.GetQuoatotionReqBody
//	if err := ctx.ShouldBindJSON(&getAllQoutedRequestBody); err != nil {
//		logger.Error("error while parsing the GetAllQuoteRequest", err.Error())
//		ctx.JSON(403, err.Error())
//		return
//	}
//	res, err := lc.accountrepo.GetQuotationCreationInfo(ctx, getAllQoutedRequestBody)
//	if err != nil {
//		logger.Error("error while getting GetLeadCreationInfo", err.Error())
//		ctx.JSON(503, err.Error())
//		return
//	}
//	ctx.JSON(200, &model.GetQuoationResp{
//		Status:  200,
//		Payload: res,
//	})
////}

/*
*
 */
func (lc leadController) InsertLeadDetails(ctx *gin.Context) {

	logger := logger2.GetLoggerWithContext(ctx)
	var insertLeadDetailsRequestBody model.InsertLeadDetailsRequest
	if err := ctx.ShouldBindJSON(&insertLeadDetailsRequestBody); err != nil {
		logger.Error("error while parsing the InsertLeadDetailsRequest", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	var persistenceOpErr error
	if insertLeadDetailsRequestBody.Update {

		persistenceOpErr = lc.leadRepo.UpdateLead(ctx, insertLeadDetailsRequestBody)

	} else {
		var leadExists bool
		leadExists, persistenceOpErr = lc.leadRepo.LeadExists(ctx,
			insertLeadDetailsRequestBody)
		if !leadExists && persistenceOpErr == nil {
			persistenceOpErr = lc.leadRepo.CreateNewLead(ctx, insertLeadDetailsRequestBody)
			go emailConfirmUserAboutTheCreatedLead(ctx, insertLeadDetailsRequestBody)
		} else if leadExists && persistenceOpErr == nil {
			ctx.JSON(230, "Lead Name already exists")
		}
	}
	if persistenceOpErr != nil {
		ctx.JSON(500, persistenceOpErr.Error())
		return
	}
	ctx.JSON(200, "SUCCESS")
}

/*
*
 */
func emailConfirmUserAboutTheCreatedLead(ctx *gin.Context, requestPayload model.InsertLeadDetailsRequest) {

	userRepo := userrepo.NewUserRepo()
	userInfo, err := userRepo.GetNameAndUsernameByUserId(ctx, requestPayload.CreatedUserid)
	//fmt.Println("Got user info: ", userInfo.Username, userInfo.Firstname)
	if err == nil {

		dataFeedToTemplate := make(map[string]string)
		dataFeedToTemplate["MessageToUser"] = "New lead has been created."
		dataFeedToTemplate["AccountName"] = requestPayload.Accountname
		dataFeedToTemplate["AccountCountry"] = requestPayload.ShippingCountry
		dataFeedToTemplate["AccountOwner"] = userInfo.Firstname + " " + userInfo.Lastname + " (" + userInfo.Username + ") "

		sendEmailRequestInput := emailer.SendEmailRequestVO{
			SenderDetails: emailer.Sender{
				SendFromIdentity: "itsupport@continental.coffee", // nice to have: pick dynamically
			},
			TargetRecipients: emailer.Recipients{
				ToList: []string{requestPayload.CreatorsEmail},
			},
			Template: emailer.EmailTemplate{
				TemplateRef:      "EmailOnLeadCreation",
				TemplateDataFeed: dataFeedToTemplate,
			},
		}
		var emailRequestAccepted bool
		emailRequestAccepted, err = emailer.Send(ctx, sendEmailRequestInput)
		fmt.Println("Email request OK? ", emailRequestAccepted)
	}
	if err != nil {
		//fmt.Println("Error while trying to send email confirmation to 'New Lead Creator' ", err.Error())
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error while trying to send email confirmation to 'New Lead Creator' ", err.Error())
	}
}

/*
*
 */
func (controllerRef leadController) GetLeadsInfo(ctx *gin.Context) {

	logger := logger2.GetLoggerWithContext(ctx)

	isAuthorised, _ := func(ctx *gin.Context) (bool, error) {
		authorised := true
		return authorised, nil
	}(ctx)

	if isAuthorised {
		var getLeadsReqPayload model.ProvideLeadsInfoReqContext
		if err := ctx.ShouldBindJSON(&getLeadsReqPayload); err != nil {

			logger.Error("Error while mapping / binding Request Payload to  model.ProvideLeadsInfoReqContext in LeadController ", err.Error())
			ctx.JSON(500, err.Error())
			return
		}
		leadsData, err := controllerRef.leadRepo.ProvideLeadsData(ctx, getLeadsReqPayload)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, leadsData)
	}

	//emailer.TestSendingSimpleEmail(ctx)
	//emailer.TestSendingTemplatedEmailWithoutAttachments(ctx)

	//emailer.TestSendingRawMessege_EmptyBody_Without_Attachments(ctx) // 2 tests done

	//emailer.TestSendingRawMessege_TextBodyOnly_Without_Attachments(ctx) // 1 test done (to field to be populated in raw msg)

	//emailer.TestSendingRawMessege_HtmlBodyOnly_Without_Attachments(ctx) //1 test with multiple to addresses, dynamically

	//emailer.TestSendingRawMessege_BothTextAndHtmlBody_Without_Attachments(ctx)
	//emailer.TestSendingRawMessege_BothTextAndHtmlBody_With_Attachments(ctx)
}
