package util

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"sync"
)

func SendMail(to string, message []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	// https://stackoverflow.com/questions/57063411/go-smtp-unable-to-send-email-through-gmail-getting-eof
	from := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")

	fmt.Println(from, password)

	smtpHost := "smtp.gmail.com"
	smtpPort := "465"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := smtp.NewClient(conn, smtpHost)

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = c.Auth(auth); err != nil {
		fmt.Println(err)
	}

	if err = c.Mail(from); err != nil {
		fmt.Println(err)
	}

	if err = c.Rcpt(to); err != nil {
		fmt.Println(err)
	}

	w, err := c.Data()
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(message)
	if err != nil {
		fmt.Println(err)
	}

	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}

	c.Quit()

	fmt.Println("Email sent successfully!")

	return
}
