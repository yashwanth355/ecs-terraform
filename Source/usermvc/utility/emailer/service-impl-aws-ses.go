package emailer

import (
	"fmt"
	apputils "usermvc/utility/apputils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type AwsSesEmailSender struct{}

/*
*
 */
func buildConsumptionConfigForAwsSes(sendEmailRequestVO SendEmailRequestVO,
	allParamsMap map[string]string,
	credentialsInfoMap map[string]string) (*InputConfigForServiceConsumption, error) {

	emailTypeObj := &sendEmailRequestVO.Type
	if emailTypeObj != nil {
		cloudRegion := emailTypeObj.CallerITInfraRegion
		if cloudRegion == "" {
			cloudRegion = "ap-south-1" // get from config, eval based on CallerFunctionalityRef
		}
		allParamsMap["REGION"] = cloudRegion
	}
	// to do: fetch/get Credetials based on emailTypeObj
	// and using App CRED / SECRET / CONFIG management pratice
	// standards of the system
	// consider other ways of providing creds or APP / System identity to Provider

	credentialsInfoMap["Credential1"] = "AKIAW4SF47I56UGUFEG5"
	credentialsInfoMap["Credential2"] = "pOUhTTXsawxywBClrLJk+7a4flIqOupQMZIkgqxN"
	credentialsInfoMap["Credential3"] = ""

	return &InputConfigForServiceConsumption{
		AllParams:          allParamsMap,
		CredentialsInfoMap: credentialsInfoMap,
	}, nil
}

/*
*
 */
func sendWithAwsSes(
	sendEmailRequestVO SendEmailRequestVO,
	consumptionConfig *InputConfigForServiceConsumption) (bool, error) {

	var processingErr error
	var awsSession *session.Session
	awsSession, processingErr = createSession(consumptionConfig)
	if processingErr == nil {
		_, processingErr = invokeSendBasedOnMessageType(awsSession, sendEmailRequestVO)
		if processingErr == nil {
			return true, nil
		}
	}
	return false, processingErr
}

/*
*
 */
func invokeSendBasedOnMessageType(awsSession *session.Session,
	sendEmailRequestVO SendEmailRequestVO) (bool, error) {

	fmt.Println("[[Email Sender]] -> AwsSes Service Impl -> Number of Attachments -> ", len(sendEmailRequestVO.Attachments))

	var processingErr error
	var sendToTargets *ses.Destination
	sendToTargets, processingErr = buildSendTargets(sendEmailRequestVO)

	if len(sendEmailRequestVO.Attachments) > 0 {

		fmt.Println("[[Email Sender]] -> AwsSes Service Impl -> calling sendRawMessage...")
		_, processingErr = sendRawMessage(awsSession, sendToTargets, sendEmailRequestVO)

	} else {

		if processingErr == nil {
			template := sendEmailRequestVO.Template
			fmt.Println("[[Email Sender]] -> AwsSes Service Impl -> template -> ", template.TemplateRef)

			if template.TemplateRef == "" {
				fmt.Println("[[Email Sender]] -> AwsSes Service Impl -> calling sendSimpleMessage...")
				_, processingErr = sendSimpleMessage(awsSession, sendToTargets, sendEmailRequestVO)
			} else if template.TemplateRef != "" {
				fmt.Println("[[Email Sender]] -> AwsSes Service Impl -> calling sendTemplatedMessage...")
				_, processingErr = sendTemplatedMessage(awsSession, sendToTargets, sendEmailRequestVO)
			}
		}
	}
	if processingErr == nil {
		return true, nil
	}
	return false, processingErr
}

/*
*
 */
func createSession(consumptionConfig *InputConfigForServiceConsumption) (*session.Session, error) {
	consumptionCredsMap := consumptionConfig.CredentialsInfoMap

	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(consumptionConfig.AllParams["REGION"]),
		Credentials: credentials.NewStaticCredentials(consumptionCredsMap["Credential1"], consumptionCredsMap["Credential2"], consumptionCredsMap["Credential3"]),
	})
	return awsSession, err
}

/*
*	To do: log using ServiceLogger
 */
func buildSendTargets(sendEmailRequestVO SendEmailRequestVO) (*ses.Destination, error) {

	targetRecipients := sendEmailRequestVO.TargetRecipients
	var ccIdPointersList []*string
	var bccIdPointersList []*string
	ccList := targetRecipients.CcList
	if ccList == nil || len(ccList) == 0 {
		ccIdPointersList = []*string{}
	} else {
		ccIdPointersList = apputils.StringArrToStringPointersArr(ccList)
	}
	bccList := targetRecipients.BccList
	if bccList == nil {
		bccIdPointersList = []*string{}
	} else {
		bccIdPointersList = apputils.StringArrToStringPointersArr(bccList)
	}
	destination := &ses.Destination{

		ToAddresses:  apputils.StringArrToStringPointersArr(targetRecipients.ToList),
		CcAddresses:  ccIdPointersList,
		BccAddresses: bccIdPointersList,
	}
	return destination, nil
}

/*
*
 */
func buildReplyTo(sendEmailRequestVO SendEmailRequestVO) []*string {

	return []*string{&sendEmailRequestVO.SenderDetails.ReplyToIdentity}
}

/*
*
 */
func checkSendError(err error) error {

	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case ses.ErrCodeMessageRejected:
			fmt.Println("ErrCodeMessageRejected: ", ses.ErrCodeMessageRejected, aerr.Error())

		case ses.ErrCodeMailFromDomainNotVerifiedException:
			fmt.Println("ErrCodeMailFromDomainNotVerifiedException: ", ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())

		case ses.ErrCodeConfigurationSetDoesNotExistException:
			fmt.Println("ErrCodeConfigurationSetDoesNotExistException: ", ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())

		default:
			if "Throttling" == aerr.Code() {
				fmt.Println("From AwsSesEmailSender checkSendError -> API Call Limits Reached")
				return SendingRateReachedErr
			} else {
				fmt.Println("From AwsSesEmailSender checkSendError: ", aerr.Code(), " <--Code  ErrMessage-->", aerr.Error())
			}
		}
	} else {
		fmt.Println("From AwsSesEmailSender checkSendError: ", err.Error())
	}
	return err
}

/*
*
 */
func sendSimpleMessage(awsSession *session.Session, sendToTargets *ses.Destination, sendEmailRequestVO SendEmailRequestVO) (bool, error) {

	var processingErr error
	if processingErr == nil {

		var message *ses.Message
		message, processingErr = buildSimpleMessage(sendEmailRequestVO)
		if processingErr == nil {

			var input *ses.SendEmailInput
			input, processingErr = buildSimpleEmailInput(sendEmailRequestVO, message, sendToTargets)
			if processingErr == nil {
				fmt.Println("AWS SES Sender -> sendSimpleMessage -> calling awsSession.SendEmail.. ")
				_, processingErr = ses.New(awsSession).SendEmail(input)
			}
		}
	}
	if processingErr == nil {
		return true, nil
	}
	processingErr = checkSendError(processingErr)
	return false, processingErr
}

/*
*
 */
func buildSimpleEmailInput(sendEmailRequestVO SendEmailRequestVO,
	message *ses.Message,
	sendToTargets *ses.Destination) (*ses.SendEmailInput, error) {

	input := &ses.SendEmailInput{
		Destination:      sendToTargets,
		Message:          message,
		Source:           &sendEmailRequestVO.SenderDetails.SendFromIdentity,
		ReplyToAddresses: buildReplyTo(sendEmailRequestVO),
	}
	return input, nil
}

/*
*
 */
func buildSimpleMessage(sendEmailRequestVO SendEmailRequestVO) (*ses.Message, error) {

	emailBody := sendEmailRequestVO.Message

	messageToSend := &ses.Message{
		Body: &ses.Body{
			Html: simpleMessageHtmlBody(emailBody.BodyHtml),
			Text: simpleMessageTextBody(emailBody.BodyText),
		},
		Subject: &ses.Content{
			Charset: aws.String(sendEmailRequestVO.Subject.CharSet),
			Data:    aws.String(sendEmailRequestVO.Subject.Data),
		},
	}
	return messageToSend, nil
}

/*
*
 */
func simpleMessageTextBody(textBody Content) *ses.Content {
	var textBodyContent ses.Content
	if (textBody == Content{}) {
		return nil
	}
	textBodyContent = ses.Content{
		Charset: aws.String(textBody.CharSet),
		Data:    aws.String(textBody.Data),
	}
	return &textBodyContent
}

/*
*
 */
func simpleMessageHtmlBody(htmlBody Content) *ses.Content {
	var textBodyContent ses.Content
	if (htmlBody == Content{}) {
		return nil
	}
	textBodyContent = ses.Content{
		Charset: aws.String(htmlBody.CharSet),
		Data:    aws.String(htmlBody.Data),
	}
	return &textBodyContent
}

/*
*
 */
func sendTemplatedMessage(awsSession *session.Session, sendToTargets *ses.Destination, sendEmailRequestVO SendEmailRequestVO) (bool, error) {

	var processingErr error
	var templatedEmailInput *ses.SendTemplatedEmailInput
	templatedEmailInput, processingErr = buildTemplatedEmailInput(sendToTargets, sendEmailRequestVO)

	//fmt.Println("input to sendTemplatedMessage :", templatedEmailInput.GoString())
	_, processingErr = ses.New(awsSession).SendTemplatedEmail(templatedEmailInput)
	if processingErr == nil {
		return true, nil
	}
	processingErr = checkSendError(processingErr)
	return false, processingErr
}

/*
*
 */
func buildTemplatedEmailInput(sendToTargets *ses.Destination,
	sendEmailRequestVO SendEmailRequestVO) (*ses.SendTemplatedEmailInput, error) {

	templateDataFeedJsonString := apputils.StringsMapToJsonString(sendEmailRequestVO.Template.TemplateDataFeed)
	input := &ses.SendTemplatedEmailInput{

		Source:           &sendEmailRequestVO.SenderDetails.SendFromIdentity,
		Template:         &sendEmailRequestVO.Template.TemplateRef,
		Destination:      sendToTargets,
		TemplateData:     &templateDataFeedJsonString,
		ReplyToAddresses: buildReplyTo(sendEmailRequestVO),
	}
	return input, nil
}

/*
*
 */
func sendRawMessage(awsSession *session.Session, sendToTargets *ses.Destination,
	sendEmailRequestVO SendEmailRequestVO) (bool, error) {

	var processingErr error
	var rawMsgString string
	rawMsgString, processingErr = buildHandCraftedRawMessage(sendEmailRequestVO)
	if processingErr == nil {
		//fmt.Println("Raw essage: \n\n", rawMsgString)
		var rawMsgInput *ses.SendRawEmailInput
		rawMsgInput, processingErr = buildRawEmailInput(rawMsgString, sendToTargets, sendEmailRequestVO)

		if processingErr == nil {
			fmt.Println("Sending with Input to sendRawMessage ", rawMsgInput.GoString())
			_, processingErr = ses.New(awsSession).SendRawEmail(rawMsgInput)
		}
	}
	if processingErr == nil {
		return true, nil
	}
	processingErr = checkSendError(processingErr)
	return false, processingErr
}

/*
*
 */
func buildRawEmailInput(rawMsg string,
	sendToTargets *ses.Destination,
	sendEmailRequestVO SendEmailRequestVO) (*ses.SendRawEmailInput, error) {

	var processingErr error

	rawEmailInput := &ses.SendRawEmailInput{
		//Source: &sendEmailRequestVO.SenderDetails.SendFromIdentity,
		//Destinations: apputils.StringArrToStringPointersArr(sendEmailRequestVO.TargetRecipients.ToList),
		RawMessage: &ses.RawMessage{
			Data: []byte(rawMsg),
		},
	}
	return rawEmailInput, processingErr
}
