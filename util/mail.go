package util

import (
	"fmt"
	"net/smtp"
	"os"
	"sync"
)

func SendMail(to []string, message []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	from := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent successfully!")
}
