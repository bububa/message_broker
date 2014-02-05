package main

import (
	"errors"
	"fmt"
	"github.com/bububa/email"
	"net/smtp"
	"strings"
)

type EmailAuth struct {
	User     string
	Passwd   string
	SMTPHost string
	SMTPPort string
}

func sendToEmail(title string, msg string) error {
	logger.Infof("Sending message to email: %s", msg)
	if len(title) == 0 {
		title = "Message From Xibao Message Broker"
	}
	mailTo := *mailToFlag
	list := strings.Split(mailTo, ";")
	if len(list) == 0 {
		err := errors.New("Need receivers")
		logger.Error(err)
		return err
	}
	m := email.NewMessage(title, msg, *formatFlag)
	m.From = emailAuth.User
	m.To = list
	err := email.Send(fmt.Sprintf("%s:%s", emailAuth.SMTPHost, emailAuth.SMTPPort), smtp.PlainAuth("", emailAuth.User, emailAuth.Passwd, emailAuth.SMTPHost), m)
	if err != nil {
		logger.Warn(err)
		return err
	}
	logger.Infof("Sent message to email: %s", msg)
	return nil
}
