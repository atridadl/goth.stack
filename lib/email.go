package lib

import (
	"fmt"
	"os"

	"github.com/resendlabs/resend-go"
)

var client *resend.Client

// init function
func init() {
	client = resend.NewClient(os.Getenv("RESEND_API_KEY"))
}

func SendEmail(to_email string, from_email string, from_name string, html string, subject string) {
	params := &resend.SendEmailRequest{
		From:    from_name + "<" + from_email + ">",
		To:      []string{to_email},
		Html:    html,
		Subject: subject,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sent.Id)
}
