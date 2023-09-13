package configs

import (
	"crypto/tls"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

var Email *gomail.Dialer

func SetupEmail() {
	port, err := strconv.Atoi(Config.SMTPPort)
	if err != nil {
		log.Fatal("Invalid SMTP port")
	}
	tlsMode, err := strconv.ParseBool(Config.SMTPTLS)
	if err != nil {
		log.Fatal("SMTP_TLS should be boolean type")
	}

	Email = gomail.NewDialer(Config.SMTPHOST, port, Config.SMTPUser, Config.SMTPPass)

	if tlsMode {
		Email.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
}
