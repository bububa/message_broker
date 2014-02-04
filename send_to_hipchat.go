package main

import (
	"fmt"
	"github.com/bububa/hipchat"
)

func sendToHipchat(msg string) error {
	logger.Debug(fmt.Sprintf("Sending message to hipchat: %s", msg))
	c := hipchat.Client{AuthToken: hipchatToken}
	req := hipchat.MessageRequest{
		RoomId:        *roomFlag,
		From:          *fromFlag,
		Message:       msg,
		Color:         *colorFlag,
		MessageFormat: *formatFlag,
		Notify:        true,
	}
	if err := c.PostMessage(req); err != nil {
		logger.Warn(err)
		return err
	}
	logger.Debug(fmt.Sprintf("Sent message to hipchat: %s", msg))
	return nil
}
