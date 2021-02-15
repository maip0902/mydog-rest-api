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

func SendTemporaryRegisterMail(c *smtp.Client, m string, t string) (*smtp.Client, error) {
	// 送信元
	if err := c.Mail("sender@example.org"); err != nil {
		fmt.Println(err.Error())
	}
	// 送信先
	if err := c.Rcpt(m); err != nil {
		fmt.Println(err.Error())
	}
	wc, err := c.Data()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 本文１
	_, err = fmt.Fprintf(wc, "仮登録完了のお知らせ\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 本文２
	_, err = fmt.Fprintf(wc, "以下のリンクをタップすることで本登録完了となります\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 本文3
	_, err = fmt.Fprintf(wc, "http://localhost/email?verify_token=" + t)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = wc.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	return c, err
}