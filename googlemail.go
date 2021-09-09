package googlemail

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

type MailConfig struct {
	Address  string
	Port     int
	User     string
	Password string
}

func (c MailConfig) SendEmail(to, from, subject, body string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", from)

	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(c.Address, c.Port, c.User, c.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
