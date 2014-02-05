package main

import (
	"errors"
	"github.com/bububa/gomemcache/memcache"
	"github.com/bububa/wechat"
	"strings"
)

func sendToWechat(msg string) error {
	if *toFlag == "" {
		return errors.New("Please tell me the receivers")
	}
	cache := memcache.NewClient(CacheHosts)
	wechatClient := wechat.NewClient(wechatAppID, wechatAppKey, cache)
	toStr := *toFlag
	receivers := strings.Split(toStr, ",")
	logger.Infof("Sending message to wechat: %s", msg)
	wechatMsg := &wechat.Message{
		Type:    "text",
		Content: msg,
	}
	wechatClient.CreateMessage(wechatMsg)
	_, err := wechatClient.SendMessage(wechatMsg.MessageId, receivers, true, msg)
	if err != nil {
		logger.Warn(err)
		return err
	}
	logger.Infof("Sent message to wechat: %s", msg)
	return nil
}
