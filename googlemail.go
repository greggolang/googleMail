package googlemail

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(to, from, subject, body, serverAddress, serverUser, serverPassword string, serverPort int) error {

	m := gomail.NewMessage()

	m.SetHeader("From", from)

	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(serverAddress, serverPort, serverUser, serverPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
