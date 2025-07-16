// utils/email/email.go
package email

import (
	"log"

	"github.com/resend/resend-go/v2"
)

var apiKey = "re_EMhVfENW_MmpG85YwSX6vUiYi9kfp73tX"
var client = resend.NewClient(apiKey)

func SendEmail(to, subject, html string) {
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
		return
	}
	log.Println("Email sent to", to, "with ID:", sent.Id)
}
