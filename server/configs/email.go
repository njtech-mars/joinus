package configs

import "gopkg.in/gomail.v2"

var Email *gomail.Dialer

func SetupEmail() {
	Email = gomail.NewDialer("smtp.office365.com", 587, Config.OutlookUser, Config.OutlookPassword)
}
