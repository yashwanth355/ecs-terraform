package emailer

import (
	"context"
	"errors"
	"time"
)

/*
*	Machine Clients needing to Send an Email from their code / functionality
*	can invoke this PUBLIC / EXPORTed api / method
*	by appropriately forming the "sendEmailRequestInput" input VO (value object)
*
*	Do refer the above Schema / Struct / Structure definitions
*
*
*	This program file is the only one needed for CALLERs to understand the contract
*	and consume the email sending capability.
*
 */
func Send(ctx context.Context,
	sendEmailRequestVO SendEmailRequestVO) (bool, error) {

	var sendingServiceStub SendingInfraServiceStub
	return sendingServiceStub.SendEmail(ctx, sendEmailRequestVO)
}

// Callers or Consumers of Email Sender Capability,
// need to consult the below INPUT contract / schema / structure
type (
	/*
		*	The root coarse grained INPUT parameter / payload / structure to
		*	use/consume

			"emailer.Send"

		*
		*	capability
	*/
	SendEmailRequestVO struct {
		TargetRecipients Recipients

		Subject Content // optional if Type / EmailType is provided or if Template is provided

		SenderDetails Sender // optional if Type / EmailType is provided or if Template is provided

		Message     EmailBody    // optional for EMPTY MESSAGE BODY scenarios
		Attachments []Attachment // optional

		Template EmailTemplate // optional

		Type EmailType // optional

		SendUsingService string // e.g. AWS-SES (optional), also available in EmailType (need to be available in one of the places)

		IsBulk bool // optional

		Priority     SendPriority    // optional
		SendSchedule SendingSchedule // optional

		MimeVersion string
	}

	/*
	*
	 */
	Content struct {
		CharSet string // optional,  defaulted to CHARSET_UTF_8 / "UTF-8"
		Data    string

		// optional, service will evaluate based
		// on EmailBody (BodyText / BodyHtml),
		// if not defaulted to MIME_TYPE_TEXT_PLAIN, "text/plain"
		Type string

		MimeVersion      string
		TransferEncoding string
	}
	/*
	*
	*
	 */
	EmailType struct {
		Kind                 string // e.g. "Account-Activation"
		IsSystemEmail        bool   // application event triggered emails
		IsUserGeneratedEmail bool   // application user sending email

		// identifier of calling application/functionality
		// e.g. "CRM", "Leads Management"
		CallerFunctionalityRef string

		SendThoughServiceRef string //  e.g. AWS-SES
		CallerITInfraRegion  string // "us-east4-b" GCP, "ap-south-1" AWS

	}

	/*
	*
	 */
	EmailTemplate struct {
		TemplateRef                string
		TemplateDataFeed           map[string]string
		DoesNotNeedsDataPopulating bool
	}

	/*
	*
	 */
	Recipients struct {
		ToList  []string
		CcList  []string
		BccList []string
	}

	/*
	*
	*
	 */
	Sender struct {
		DisplayFromName     string
		DisplayFromIdentity string // e.g. from email to display
		SendFromIdentity    string // e.g. actual sender email to use
		ReplyToName         string
		ReplyToIdentity     string
	}

	/*
	*
	*
	 */
	EmailBody struct {
		BodyText Content
		BodyHtml Content

		HasBothTextAndHtml bool // optional

	}
	/*
	*
	*
	 */
	Attachment struct {
		FQPath      string // full file path e.g. /home/application/docs/Invoice.pdf
		Filename    string // e.g. Invoice.pdf
		ContentType string // e.g. application/pdf, application/msword

		ContentBytes  []byte // can provide a byte array of content from an input source (file, network stream etc)
		ContentString string // attachment's content as one string
	}

	/*
	*
	*
	 */
	SendPriority struct {
		Priority int8 // greater number means higher priority e.g. 10, 6, 4, 2, 1
	}

	/*
	*
	*
	 */
	SendingSchedule struct {
		SendOnInvocation bool
		SendOnDateTime   time.Time
	}
)

/*
*	Caller may refer these static values
*	as part of the
*	Client-Service contract
*
*	should it need to pass any
 */
const (
	EMAIL_SENDER_SERVICE_AWS_SES = "AWS-SES"
	//EMAIL_SENDER_SERVICE_PROVIDER_REF = "OTHER-PROVIDER_REF"

	CHARSET_UTF_8 = "UTF-8"

	CONTENT_TYPE_TEXT_PLAIN = "text/plain"
	CONTENT_TYPE_TEXT_HTML  = "text/html"

	ATTACHMENT_CONTENT_TYPE_PLAIN  = "text/plain"
	ATTACHMENT_CONTENT_TYPE_JPG    = "image/jpeg"
	ATTACHMENT_CONTENT_TYPE_MSWORD = "application/msword"
	ATTACHMENT_CONTENT_TYPE_PDF    = "application/pdf"
)

/*
*	Known Errors, that may occur
*	and those CALLER may like to be aware of
*	to adjust input
 */
var NoSendFromIdentityErr = errors.New("Fatal - No 'sendFrom' identity is provided.")
var NoTargetRecipientsErr = errors.New("Fatal - No 'emailto' identities are provided.")

var NoTemplatePopulatingDataErr = errors.New("Fatal - Request tells Template Needs Populating with Dynamic Data, but TemplateDataFeed or required data is not available in the request.")
var AttachmentContentAccessErr = errors.New("Fatal - Could not extract Attachment Content. Provide as a []byte or string or a reachable Full File Path.")
var UnknownEmailSendingInfraServiceErr = errors.New("Fatal - Could not build input config due to UnknownEmailSendingInfraService.")
var SendingRateReachedErr = errors.New("Email Send Rate Quota Reached.")
var SendingLimitReachedErr = errors.New("Email Sending Limit Quota Reached.")
