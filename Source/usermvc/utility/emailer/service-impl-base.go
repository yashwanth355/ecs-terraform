package emailer

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	apputils "usermvc/utility/apputils"
	logger2 "usermvc/utility/logger"

	"go.uber.org/zap"
)

type (
	SendingInfraServiceStub struct {
		SendServiceRef                string                            //AWS_SES
		ServiceConsumptionInputParams *InputConfigForServiceConsumption // input params
		ServiceLogger                 ServiceLogger
	}
	InputConfigForServiceConsumption struct {
		AllParams          map[string]string
		CredentialsInfoMap map[string]string
	}
	ServiceLogger struct {
		Logger *zap.SugaredLogger
	}
)

/*
*	Internal Known Errors ( implementation related )
*	that may occur
*
 */
var ServiceStubBuildingErr = errors.New("Fatal - Could not build a Stub for the Email Sending Infrastructre Service.")

/*
*
 */
func (sendingServiceRef SendingInfraServiceStub) SendEmail(ctx context.Context,
	sendEmailRequestVO SendEmailRequestVO) (bool, error) {

	var isSuccess bool = false
	var processingErr error
	var serviceStub SendingInfraServiceStub
	serviceStub, processingErr = getServiceStubInstance(ctx, sendEmailRequestVO)
	// log info for observing, troubleshooting
	logDebugInfo(sendEmailRequestVO, serviceStub)
	if processingErr == nil {
		processingErr = validateSanitize(&sendEmailRequestVO, serviceStub)
		if processingErr == nil {
			// SEND, picking the right EMAIL SENDING INFRA SERVICE PROVIDER e.g. AWS-SES
			isSuccess, processingErr = sendUsingRelevantProvider(sendEmailRequestVO, serviceStub)
		}
	}
	return isSuccess, processingErr
}

/*
*
 */
func getServiceStubInstance(ctx context.Context, sendEmailRequestVO SendEmailRequestVO) (SendingInfraServiceStub, error) {

	var stub SendingInfraServiceStub
	sendThruServiceRef := sendEmailRequestVO.SendUsingService
	emailTypeObj := &sendEmailRequestVO.Type

	if sendThruServiceRef == "" {
		if emailTypeObj != nil {
			sendThruServiceRef = emailTypeObj.SendThoughServiceRef
		}
		if sendThruServiceRef == "" {
			sendThruServiceRef = EMAIL_SENDER_SERVICE_AWS_SES
		}
	}
	consumptionConfig, configErr := buildConfigForServiceConsumption(sendEmailRequestVO, sendThruServiceRef)

	if configErr != nil {
		return stub, ServiceStubBuildingErr
	}

	stub = SendingInfraServiceStub{
		SendServiceRef:                sendThruServiceRef,
		ServiceConsumptionInputParams: consumptionConfig,
		ServiceLogger: ServiceLogger{
			Logger: logger2.GetLoggerWithContext(ctx),
		},
	}
	return stub, nil
}

/*
*
 */
func sendUsingRelevantProvider(sendEmailRequestVO SendEmailRequestVO, serviceStub SendingInfraServiceStub) (bool, error) {

	var sendSucceeded bool = false
	var processingErr error

	switch serviceStub.SendServiceRef {

	case EMAIL_SENDER_SERVICE_AWS_SES:

		sendSucceeded, processingErr = sendWithAwsSes(sendEmailRequestVO, serviceStub.ServiceConsumptionInputParams)
	default:
		processingErr = UnknownEmailSendingInfraServiceErr
	}
	return sendSucceeded, processingErr
}

/*
*
 */
func buildConfigForServiceConsumption(sendEmailRequestVO SendEmailRequestVO, sendThruServiceRef string) (*InputConfigForServiceConsumption, error) {

	var inputParamsToConsumeService *InputConfigForServiceConsumption

	var credentialsInfoMap = make(map[string]string)
	var allParamsMap = make(map[string]string)

	switch sendThruServiceRef {

	case EMAIL_SENDER_SERVICE_AWS_SES:

		inputParamsToConsumeService, _ = buildConsumptionConfigForAwsSes(sendEmailRequestVO, allParamsMap, credentialsInfoMap)

	default:
		return nil, UnknownEmailSendingInfraServiceErr
	}
	return inputParamsToConsumeService, nil
}

/*
*
 */
func logDebugInfo(sendEmailRequestVO SendEmailRequestVO, serviceStub SendingInfraServiceStub) {

	var servicesLogger ServiceLogger = serviceStub.ServiceLogger

	servicesLogger.Logger.Info("[[Email Sender]] will process email send request using Provider -> ", serviceStub.SendServiceRef)

	fmt.Println("[[Email Sender]] will process email send request using Provider -> ", serviceStub.SendServiceRef)
}

/*
*
*	to do: attachments total size validation (nice to have)
*
*
 */
func validateSanitize(sendEmailRequestVO *SendEmailRequestVO, serviceStub SendingInfraServiceStub) error {

	var sanityCheckErr error
	if sendEmailRequestVO.MimeVersion == "" {
		sendEmailRequestVO.MimeVersion = "1.0"
	}
	template := sendEmailRequestVO.Template
	if template.TemplateRef == "" {
		checkBody(&sendEmailRequestVO.Message)

	} else if !template.DoesNotNeedsDataPopulating && len(template.TemplateDataFeed) == 0 {
		sanityCheckErr = NoTemplatePopulatingDataErr
	}
	if len(sendEmailRequestVO.TargetRecipients.ToList) == 0 {
		sanityCheckErr = NoTargetRecipientsErr
	}
	senderDetails := sendEmailRequestVO.SenderDetails
	sanityCheckErr = checkSenderDetails(&senderDetails)
	if sanityCheckErr != nil {
		return sanityCheckErr
	}
	fmt.Println("[[Email Sender]] service base impl -> validateSanitize -> BodyHtml: ", sendEmailRequestVO.Message.BodyHtml)
	fmt.Println("[[Email Sender]] service base impl -> validateSanitize -> BodyText: ", sendEmailRequestVO.Message.BodyText)
	fmt.Println("[[Email Sender]] service base impl -> validateSanitize -> Subject: ", sendEmailRequestVO.Subject)
	fmt.Println("[[Email Sender]] service base impl -> validateSanitize -> SendFromIdentity: ", sendEmailRequestVO.SenderDetails.SendFromIdentity)
	fmt.Println("[[Email Sender]] service base impl -> validateSanitize -> ReplyToIdentity: ", sendEmailRequestVO.SenderDetails.ReplyToIdentity)
	return nil
}

/*
*
 */
func checkSenderDetails(senderDetails *Sender) error {

	var sanityCheckErr error
	if senderDetails.SendFromIdentity == "" {
		sanityCheckErr = NoSendFromIdentityErr
	}
	if senderDetails.ReplyToIdentity == "" {
		senderDetails.ReplyToIdentity = senderDetails.SendFromIdentity
	}
	if senderDetails.DisplayFromName == "" {
		senderDetails.DisplayFromName = senderDetails.SendFromIdentity
	}
	if sanityCheckErr != nil {
		return sanityCheckErr
	}
	return nil
}

/*
*
 */
func checkBody(body *EmailBody) error {

	var sanityCheckErr error = nil

	if body.BodyText.Data != "" && body.BodyText.CharSet == "" {
		body.BodyText.CharSet = CHARSET_UTF_8
		body.BodyText.Type = CONTENT_TYPE_TEXT_PLAIN
	}
	if body.BodyHtml.Data != "" && body.BodyHtml.CharSet == "" {
		body.BodyText.CharSet = CHARSET_UTF_8
		body.BodyHtml.Type = CONTENT_TYPE_TEXT_HTML
	}
	var containsBothTextNHtml bool = false
	if (body.BodyText != Content{} && body.BodyHtml != Content{}) || (body.BodyText.Data != "" && body.BodyHtml.Data != "") {
		containsBothTextNHtml = true
	}
	body.HasBothTextAndHtml = containsBothTextNHtml
	/*if (*body == EmailBody{}) {
		fmt.Println("[[Email Sender]] service base impl -> NO EMAIL BODY or Message")
		body = &EmailBody{
			BodyText: Content{
				CharSet: CHARSET_UTF_8,
				Type:    CONTENT_TYPE_TEXT_PLAIN,
				Data:    "",
			},
		}
	} */
	return sanityCheckErr
}

/*
*
 */
func makeRawMsgHeadersBeforeCTypeHeader(sendEmailRequestVO SendEmailRequestVO) string {

	var builder strings.Builder
	builder.WriteString("From: '" + sendEmailRequestVO.SenderDetails.DisplayFromName + "' <" + sendEmailRequestVO.SenderDetails.SendFromIdentity + ">\n")
	builder.WriteString("Subject: " + sendEmailRequestVO.Subject.Data + "\n")
	builder.WriteString("To: " + strings.Join(sendEmailRequestVO.TargetRecipients.ToList[:], ",") + "\n")
	builder.WriteString("MIME-Version: " + sendEmailRequestVO.MimeVersion + "\n")
	return builder.String()
}

/*
*
 */
func makeCTypeHeaderForRawMsgWithAttachments() (string, string) {

	var builder strings.Builder
	rootBoundaryId := generateBoundaryId("MESSAGE-WITH-ATTACHMENTs")
	builder.WriteString("Content-Type: multipart/mixed; boundary=\"" + rootBoundaryId + "\"\n\n")
	builder.WriteString("--" + rootBoundaryId + "\n")
	return rootBoundaryId, builder.String()
}

/*
*
 */
func buildHandCraftedRawMessage(sendEmailRequestVO SendEmailRequestVO) (string, error) {

	var buildErr error = nil
	var rawMsgBuilder strings.Builder
	rawMsgBuilder.WriteString(makeRawMsgHeadersBeforeCTypeHeader(sendEmailRequestVO))
	var rasMsgPart string
	var rootBoundaryId string = ""
	if len(sendEmailRequestVO.Attachments) > 0 {
		rootBoundaryId, rasMsgPart = makeCTypeHeaderForRawMsgWithAttachments()
		rawMsgBuilder.WriteString(rasMsgPart)
	}
	body := sendEmailRequestVO.Message
	if body.HasBothTextAndHtml {
		if rootBoundaryId == "" {
			bodyBoundary := generateBoundaryId("multipart/alternative")
			rawMsgBuilder.WriteString("Content-Type: multipart/alternative; boundary=\"" + bodyBoundary + "\"\n\n")
			rawMsgBuilder.WriteString(makeRawMsgBodyPart(body, "--"+bodyBoundary))
		} else {
			rawMsgBuilder.WriteString("Content-Type: multipart/alternative; boundary=\"sub_" + rootBoundaryId + "\"\n\n")
			rawMsgBuilder.WriteString(makeRawMsgBodyPart(body, "--sub_"+rootBoundaryId))
		}
	} else {
		if rootBoundaryId != "" {
			rawMsgBuilder.WriteString(makeRawMsgBodyPart(body, "--sub_"+rootBoundaryId))
		} else {
			rawMsgBuilder.WriteString(makeRawMsgBodyPart(body, rootBoundaryId))
		}
	}
	if len(sendEmailRequestVO.Attachments) > 0 {

		rasMsgPart, buildErr = makeAttachmentsPartOfRawMessage(sendEmailRequestVO, "--"+rootBoundaryId)
		if buildErr != nil {
			return "", buildErr
		}
		rawMsgBuilder.WriteString(rasMsgPart)
	}
	if rootBoundaryId != "" {
		rawMsgBuilder.WriteString("\n\n--" + rootBoundaryId + "--")
	}
	return rawMsgBuilder.String(), buildErr
}

/*
*
 */
func makeAttachmentsPartOfRawMessage(sendEmailRequestVO SendEmailRequestVO, boundary string) (string, error) {

	var builder strings.Builder
	var processingErr error = nil
	var oneAttachmentPartOfRawMsg string
	attachments := sendEmailRequestVO.Attachments
	for _, attachment := range attachments {
		oneAttachmentPartOfRawMsg, processingErr = makeRawMsgPartForOneAttachment(attachment, boundary)
		if processingErr != nil {
			break
		}
		builder.WriteString(oneAttachmentPartOfRawMsg)
	}
	return builder.String(), processingErr
}

/*
*
 */
func makeRawMsgPartForOneAttachment(attachment Attachment, boundary string) (string, error) {

	var ioErr error
	var attachmentContentString string

	if attachment.FQPath != "" {
		attachmentContentString, ioErr = fileContentAsBase64EncodedString(attachment.FQPath)
		if ioErr != nil {
			return "", ioErr
		}
	} else if len(attachment.ContentBytes) > 0 {
		attachmentContentString = base64.StdEncoding.EncodeToString(attachment.ContentBytes)
	} else if attachment.ContentString != "" {
		attachmentContentString = base64.StdEncoding.EncodeToString([]byte(attachment.ContentString))
	} else {
		return "", AttachmentContentAccessErr
	}
	var builder strings.Builder
	builder.WriteString("\n\n" + boundary + "\n")
	builder.WriteString("Content-Type: " + attachment.ContentType + "; name=\"" + attachment.Filename + "\"\n")
	builder.WriteString("Content-Transfer-Encoding: base64\n")
	builder.WriteString("Content-Disposition: attachment;filename=\"" + attachment.Filename + "\"\n\n")
	builder.WriteString(attachmentContentString)
	return builder.String(), nil
}

/*
*
*
 */
func fileContentAsBase64EncodedString(filePath string) (string, error) {

	var processingErr error = nil
	contentBytes, processingErr := ioutil.ReadFile(filePath)
	if processingErr != nil {
		fmt.Println("Attachment reading error: ", processingErr)
	}
	contentStringAsBase64 := base64.StdEncoding.EncodeToString(contentBytes)
	//replace(/([^\0]{76})/g, "$1\n") + "\n\n";
	//regexp.MustCompile(`([^\0]{76})`).ReplaceAllString(contentStringAsBase64, `$1\n`)
	return contentStringAsBase64, processingErr
}

/*
*
 */
func makeRawMsgBodyPart(body EmailBody, boundary string) string {

	var builder strings.Builder
	if (body == EmailBody{}) {
		if boundary != "" {
			builder.WriteString(boundary + "\n")
		}
		builder.WriteString("Content-Type: text/plain; charset=" + CHARSET_UTF_8 + "\n")
		builder.WriteString("Content-Length: 0\n")
		if boundary != "" {
			builder.WriteString(boundary + "--")
		}
		return builder.String()
	}
	textBody := body.BodyText
	if (textBody != Content{}) {
		if boundary != "" {
			builder.WriteString(boundary + "\n")
		}
		builder.WriteString("Content-Type: " + CONTENT_TYPE_TEXT_PLAIN + "; charset=" + textBody.CharSet + "\n")
		builder.WriteString("Content-Transfer-Encoding: quoted-printable\n\n")
		builder.WriteString(textBody.Data + "\n\n")
	}
	htmlBody := body.BodyHtml
	if (htmlBody != Content{}) {
		if boundary != "" {
			builder.WriteString(boundary + "\n")
		}
		builder.WriteString("Content-Type: " + CONTENT_TYPE_TEXT_HTML + "; charset=" + htmlBody.CharSet + "\n")
		builder.WriteString("Content-Transfer-Encoding: quoted-printable\n\n")
		builder.WriteString(htmlBody.Data + "\n\n")
	}
	if boundary != "" {
		builder.WriteString(boundary + "--")
	}
	return builder.String()
}

/*
*
 */
func generateBoundaryId(inputHint string) string {

	return fmt.Sprint(apputils.Crc32OfString(inputHint))
}
