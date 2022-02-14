package emailer

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
)

/*
*
From: '' <deb.work.related@gmail.com>
Subject: TestSendingRawMessege_TextBodyOnly_Without_Attachments - SUBJECT
To: banani.karma@gmail.com,debdas.sinha@gmail.com
MIME-Version: 1.0
Content-Type: text/plain; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

Text body of Raw Messege
*/
func TestSendingRawMessege_TextBodyOnly_Without_Attachments(ctx context.Context) {

	/*attachment1 := Attachment{
		Name:      "attach1",
		SourceUrl: "some-readable-file-path",
	}
	attachments := []Attachment{attachment1} */

	sendEmailRequestInput := SendEmailRequestVO{
		Type: EmailType{
			CallerITInfraRegion: "us-east-1",
			IsSystemEmail:       true,
		},
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{"banani.karma@gmail.com", "debdas.sinha@gmail.com"},
		},
		Message: EmailBody{
			BodyText: Content{
				CharSet: "iso-8859-1",
				Data:    "Text body of Raw Messege",
			},
		},
		Subject: Content{
			CharSet: "iso-8859-1",
			Data:    "TestSendingRawMessege_TextBodyOnly_Without_Attachments - SUBJECT",
		},
		//Attachments: attachments,
	}

	fmt.Println("[[Email Sender]] Tester TestSendingRawMessege_TextBodyOnly_Without_Attachments -> Num Attach -> ", len(sendEmailRequestInput.Attachments))

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success ->  ", emailRequestProcessingSuccess)
	}

}

/*
*
From: '' <deb.work.related@gmail.com>
Subject: TestSendingRawMessege_HtmlBodyOnly_Without_Attachments - SUBJECT
To: banani.karma@gmail.com,debdas.sinha@gmail.com
MIME-Version: 1.0
Content-Type: text/html; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

Html body of Raw Messege

*/
func TestSendingRawMessege_HtmlBodyOnly_Without_Attachments(ctx context.Context) {

	sendEmailRequestInput := SendEmailRequestVO{
		Type: EmailType{
			CallerITInfraRegion: "us-east-1",
			IsSystemEmail:       true,
		},
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{"banani.karma@gmail.com", "debdas.sinha@gmail.com"},
		},
		Message: EmailBody{
			BodyHtml: Content{
				CharSet: "iso-8859-1",
				Data:    "Html body of Raw Messege",
			},
		},
		Subject: Content{
			CharSet: "iso-8859-1",
			Data:    "TestSendingRawMessege_HtmlBodyOnly_Without_Attachments - SUBJECT",
		},
	}

	fmt.Println("[[Email Sender]] Tester TestSendingRawMessege_HtmlBodyOnly_Without_Attachments -> Num Attach -> ", len(sendEmailRequestInput.Attachments))

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success ->  ", emailRequestProcessingSuccess)
	}

}

/*
*
From: '' <deb.work.related@gmail.com>
Subject: TestSendingRawMessege_BothTextAndHtmlBody_Without_Attachments - SUBJECT
To: banani.karma@gmail.com,debdas.sinha@gmail.com
MIME-Version: 1.0
Content-Type: multipart/alternative; boundary="1867535975"

--1867535975
Content-Type: text/plain; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

Text body of Raw Messege

--1867535975
Content-Type: text/html; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

<h1>Html body of <br><ul><li>11</li><li>22</li></ul>Raw Messege</h1>

--1867535975--

*/
func TestSendingRawMessege_BothTextAndHtmlBody_Without_Attachments(ctx context.Context) {

	sendEmailRequestInput := SendEmailRequestVO{
		Type: EmailType{
			CallerITInfraRegion: "us-east-1",
			IsSystemEmail:       true,
		},
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{"banani.karma@gmail.com", "debdas.sinha@gmail.com"},
		},
		Message: EmailBody{
			BodyText: Content{
				CharSet: "iso-8859-1",
				Data:    "Text body of Raw Messege",
			},
			BodyHtml: Content{
				CharSet: "iso-8859-1",
				Data:    "<h1>Html body of <br><ul><li>11</li><li>22</li></ul>Raw Messege</h1>",
			},
		},
		Subject: Content{
			CharSet: "iso-8859-1",
			Data:    "TestSendingRawMessege_BothTextAndHtmlBody_Without_Attachments - SUBJECT",
		},
		//Attachments: attachments,
	}

	fmt.Println("[[Email Sender]] Tester TestSendingRawMessege_BothTextAndHtmlBody_Without_Attachments -> Num Attach -> ", len(sendEmailRequestInput.Attachments))

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success ->  ", emailRequestProcessingSuccess)
	}

}

/*
*

From: '' <deb.work.related@gmail.com>
Subject: TestSendingRawMessege_BothTextAndHtmlBody_With_Attachments - SUBJECT
To: banani.karma@gmail.com,debdas.sinha@gmail.com
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="3330259215"

--3330259215
Content-Type: multipart/alternative; boundary="sub_3330259215"

--sub_3330259215
Content-Type: text/plain; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

Text body of Raw Messege With Attachments

--sub_3330259215
Content-Type: text/html; charset=iso-8859-1
Content-Transfer-Encoding: quoted-printable

Html body of <br><ul><li>11</li><li>22</li></ul>Raw Messege with Attachments

--sub_3330259215--

--3330259215
Content-Type: application/pdf; name="Test.pdf"
Content-Transfer-Encoding: base64
Content-Disposition: attachment;filename="Test.pdf"

JVBERi0xLjUNCiW1tbW1DQoxIDAgb2JqDQo8PC9UeXBlL0NhdGFsb2cvUGFnZXMgMiAwIFIvTGFuZyhlbi1JTikgPj4NCmVuZG9iag0KMiAwIG9iag0KPDwvVHlwZS9QYWdlcy9Db3VudCA3L0tpZHNbIDMgMCBSIDI2IDAgUiAzMCAwIFIgMzQgMCBSIDM2IDAgUiAzOCAwIFIgNDIgMCBSXSA+Pg0KZW5kb2JqDQozIDAgb2JqDQo8PC9UeXBlL1BhZ2UvUGFyZW50IDIgMCBSL1Jlc291cmNlczw8L0ZvbnQ8PC9GMSA1IDAgUi9GMiA5IDAgUi9GMyAxMyAwIFIvRjQgMTUgMCBSL0Y1IDIwIDAgUi9GNiAyMiAwIFIvRjcgMjQgMCBSPj4vRXh0R1N0YXRlPDwvR1M3IDcgMCBSL0dTOCA4IDAgUj4+L1hPYmplY3Q8PC9JbWFnZTExIDExIDAgUj4+L1Byb2NTZXRbL1BERi9UZXh0L0ltYWdlQi9JbWFnZUMvSW1hZ2VJXSA+Pi9NZWRpYUJveFsgMCAwIDYxMiA3OTJdIC9Db250ZW50cyA

--3330259215--
*
*
*/
func TestSendingRawMessege_BothTextAndHtmlBody_With_Attachments(ctx context.Context) {

	fmt.Println("----------------------\n")
	dirOfAttachedFile, err := os.Getwd()
	if err != nil {
		fmt.Println("FATAL: Error Getting Directory of current executable: ", err)
	}
	fmt.Println("Directory where attachment file is put for testing: ", dirOfAttachedFile)
	fmt.Println("----------------------\n")

	contentBytes, err := ioutil.ReadFile(dirOfAttachedFile + "/TestAttachment.pdf")
	if err != nil {
		panic(err)
		return
	}
	attachment1 := Attachment{
		Filename:     "TestAttachment.pdf",
		ContentType:  ATTACHMENT_CONTENT_TYPE_PDF,
		ContentBytes: contentBytes,
	}
	attachments := []Attachment{attachment1}

	sendEmailRequestInput := SendEmailRequestVO{

		Attachments: attachments,

		Type: EmailType{
			CallerITInfraRegion: "us-east-1",
			IsSystemEmail:       true,
		},
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{"banani.karma@gmail.com", "debdas.sinha@gmail.com"},
		},
		Message: EmailBody{
			BodyText: Content{
				CharSet: "iso-8859-1",
				Data:    "Text body of Raw Messege With Attachments",
			},
			BodyHtml: Content{
				CharSet: "iso-8859-1",
				Data:    "Html body of <br><ul><li>11</li><li>22</li></ul>Raw Messege with Attachments",
			},
		},
		Subject: Content{
			CharSet: "iso-8859-1",
			Data:    "TestSendingRawMessege_BothTextAndHtmlBody_With_Attachments - SUBJECT",
		},
	}

	fmt.Println("[[Email Sender]] Tester TestSendingRawMessege_BothTextAndHtmlBody_With_Attachments -> Num Attach -> ", len(sendEmailRequestInput.Attachments))

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success ->  ", emailRequestProcessingSuccess)
	}

}

/*
*
Worked 1:
From: '' <deb.work.related@gmail.com>
Subject: TestSendingRawMessege_EmptyBody_Without_Attachments - SUBJECT
To: banani.karma@gmail.com
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Length: 0
*/
func TestSendingRawMessege_EmptyBody_Without_Attachments(ctx context.Context) {

	sendEmailRequestInput := SendEmailRequestVO{
		Type: EmailType{
			CallerITInfraRegion: "us-east-1",
			IsSystemEmail:       true,
		},
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{"banani.karma@gmail.com", "debdas.sinha@gmail.com"},
		},
		Subject: Content{
			CharSet: "iso-8859-1",
			Data:    "TestSendingRawMessege_EmptyBody_Without_Attachments 2 - SUBJECT",
		},
	}

	fmt.Println("[[Email Sender]] Tester TestSendingRawMessege_EmptyBody_Without_Attachments -> Num Attach -> ", len(sendEmailRequestInput.Attachments))

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success ->  ", emailRequestProcessingSuccess)
	}

}

/*
*
 */
func TestSendingSimpleEmail(ctx context.Context) {

	toEmail := "banani.karma@gmail.com"

	sendEmailRequestInput := SendEmailRequestVO{
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{toEmail},
		},
		Message: EmailBody{
			BodyText: Content{
				CharSet: "iso-8859-1",
				Data:    "Simple Email Test - Test Body Text",
			},
		},
		Subject: Content{
			CharSet: "iso-8859-1",
			Data:    "Simple Email Test - SUBJECT",
		},
	}

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success ->  ", emailRequestProcessingSuccess)
	}

}

/*
*
 */
func TestSendingTemplatedEmailWithoutAttachments(ctx context.Context) {

	toEmail := "banani.karma@gmail.com"

	dataFeedToTemplate := make(map[string]string)
	dataFeedToTemplate["MessageToUser"] = "New lead has been created."
	dataFeedToTemplate["AccountName"] = "Prospect Name"
	dataFeedToTemplate["AccountCountry"] = "India"
	dataFeedToTemplate["AccountOwner"] = "Banani"

	sendEmailRequestInput := SendEmailRequestVO{
		SenderDetails: Sender{
			SendFromIdentity: "deb.work.related@gmail.com",
		},
		TargetRecipients: Recipients{
			ToList: []string{toEmail},
		},
		Template: EmailTemplate{
			TemplateRef:      "EmailOnLeadCreation",
			TemplateDataFeed: dataFeedToTemplate,
		},
	}

	emailRequestProcessingSuccess, err := Send(ctx, sendEmailRequestInput)

	if err != nil {
		fmt.Println("[[Email Sender]] Processing Error -> ", err.Error())
	} else {
		fmt.Println("[[Email Sender]] Processing Success -> ", emailRequestProcessingSuccess)
	}

}
