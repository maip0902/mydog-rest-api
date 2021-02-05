package mail

import (
	"net/smtp"
	"fmt"
)

func ConnectSMTP() (*smtp.Client, error) {
	c, err := smtp.Dial("mail:1025")
	if err != nil {
		fmt.Printf("handle: connect smtp error: %s\n", err.Error())
	}
	return c, err
}

func SendTemporaryRegisterMail(c *smtp.Client, m string) (*smtp.Client, error) {
	if err := c.Mail("sender@example.org"); err != nil {
		fmt.Println(err.Error())
	}
	if err := c.Rcpt(m); err != nil {
		fmt.Println(err.Error())
	}
	wc, err := c.Data()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = wc.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	return c, err
}