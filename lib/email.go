package lib

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to_email string, from_email string, from_name string, html string, subject string) {
	LogInfo.Println("Starting email sending process")

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
	)

	LogSuccess.Println("Authentication set up")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte("From: " + from_name + " <" + from_email + ">\r\n" +
		"To: " + to_email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
		"\r\n" +
		html + "\r\n")

	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         os.Getenv("SMTP_HOST"),
	}

	LogSuccess.Println("TLS configuration set up")

	c, err := smtp.Dial(os.Getenv("SMTP_HOST") + ":587")
	if err != nil {
		LogError.Println("Error dialing SMTP server:", err)
		return
	}

	LogSuccess.Println("Connected to SMTP server")

	if err = c.StartTLS(tlsconfig); err != nil {
		LogError.Println("Error starting TLS:", err)
		return
	}

	LogInfo.Println("TLS started")

	if err = c.Auth(auth); err != nil {
		LogError.Println("Error authenticating with SMTP server:", err)
		return
	}

	LogSuccess.Println("Authenticated with SMTP server")

	if err = c.Mail(from_email); err != nil {
		LogError.Println("Error setting sender address:", err)
		return
	}

	LogSuccess.Println("Sender address set")

	if err = c.Rcpt(to_email); err != nil {
		LogError.Println("Error setting recipient address:", err)
		return
	}

	LogSuccess.Println("Recipient address set")

	w, err := c.Data()
	if err != nil {
		LogError.Println("Error getting write closer:", err)
		return
	}

	LogSuccess.Println("Got write closer")

	_, err = w.Write(msg)
	if err != nil {
		LogError.Println("Error writing message:", err)
		return
	}

	LogSuccess.Println("Message written")

	err = w.Close()
	if err != nil {
		LogError.Println("Error closing write closer:", err)
		return
	}

	LogSuccess.Println("Write closer closed")

	c.Quit()

	LogSuccess.Println("Email sent successfully")
}
