package lib

import (
	"bytes"
	"html/template"
	"server/configs"

	"gopkg.in/gomail.v2"
)

func renderTemplate(name string) (string, error) {
	data := struct{ Name string }{Name: name}

	template, err := template.ParseFiles("data/template.html")
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err = template.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func SendEmail(email string, name string) error {
	body, err := renderTemplate(name)
	if err != nil {
		return err
	}

	message := gomail.NewMessage()
	message.SetAddressHeader("From", configs.Config.SMTPUser, "Mars工作室")
	message.SetHeader("To", email)
	message.SetHeader("Subject", name+"同学你好")
	message.SetBody("text/html", body)

	if err := configs.Email.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
