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