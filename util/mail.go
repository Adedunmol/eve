package util

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"sync"

	"github.com/jordan-wright/email"
	gomail "gopkg.in/mail.v2"
)

func SendMail(to []string, message []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	from := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")

	fmt.Println(from, password)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	fmt.Println("working: before sending the mail")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("working: after sending the mail")

	fmt.Println("Email sent successfully!")

	return
}

func SendMailV2(to string, wg *sync.WaitGroup) {
	defer wg.Done()
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", os.Getenv("EMAIL_USERNAME"))

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	fmt.Println("working")

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}

func SendMailV3(to []string) {
	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", "Adedunmola", os.Getenv("EMAIL_USERNAME"))
	e.Subject = "testing"
	e.HTML = []byte("testing")
	e.To = to

	smtpAuth := smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")

	fmt.Println(os.Getenv("EMAIL_USERNAME"))
	fmt.Println(os.Getenv("EMAIL_PASSWORD"))
	fmt.Println(smtpAuth)
	fmt.Println(to)

	err := e.Send("smtp.gmail.com:587", smtpAuth)

	fmt.Println("working")

	if err != nil {
		fmt.Println("err hand", err)
		panic(err)
	}

	return
}
