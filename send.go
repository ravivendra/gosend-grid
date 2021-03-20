package gosendgrid

import (
	"errors"
	"fmt"
	"net/http"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	// Method : http method
	Method = "POST"

	// SendgridAPIURL : base URL of Sendgrid
	SendgridAPIURL = "https://api.sendgrid.com"

	// SendgridAPIEndpoint : endpoint of Sendgrid
	SendgridAPIEndpoint = "/v3/mail/send"
)

// DoSend : do send email using Sendgrid
func DoSend(key, format, from, fromName, to, toName, subject, body string) (string, string, error) {
	var messageID, resultDesc string

	if format != "text/plain" && format != "text/html" {
		return messageID, resultDesc, errors.New("Invalid mail format you provide")
	}

	sgMailV3 := mail.NewV3Mail()

	sender := mail.NewEmail(fromName, from)

	content := mail.NewContent(format, body)

	sgMailV3.SetFrom(sender)
	sgMailV3.AddContent(content)

	personalization := mail.NewPersonalization()

	receiver := mail.NewEmail(toName, to)

	personalization.AddTos(receiver)

	personalization.Subject = subject

	sgMailV3.AddPersonalizations(personalization)

	request := sendgrid.GetRequest(key, SendgridAPIEndpoint, SendgridAPIURL)

	request.Method = Method
	request.Body = mail.GetRequestBody(sgMailV3)

	response, err := sendgrid.API(request)

	if (response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted) || err != nil {
		return messageID, response.Body, err
	}

	messageID = fmt.Sprintf("%s", response.Headers["X-Message-Id"])

	return messageID, response.Body, nil
}
