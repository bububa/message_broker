package main

import (
	//"github.com/bububa/gomemcache/memcache"
	//"github.com/bububa/wechat"
	"github.com/bububa/weixin"
	"errors"
    "encoding/json"
    "net/http"
    "net/url"
	"strings"
)

const (
    _API_HOST = "http://api.weixin.xibao100.com/xibao"
)

func sendToWechat(msg string) error {
	if *toFlag == "" {
		return errors.New("Please tell me the receivers")
	}
	toStr := *toFlag
	receivers := strings.Split(toStr, ",")
    article := []*weixin.Article{&weixin.Article{
        Title: "Xibao Muinin Notification",
        Description: msg,
        PicUrl: "",
        Url: "",
    }}
    js, err := json.Marshal(article)
    if err != nil {
        return err
    }
    for _, receiver := range receivers {
        _, err := http.PostForm(_API_HOST + "/message/custom/send", url.Values{"touser":{receiver},"msgtype":{"news"},"articles":{string(js)}})
        if err != nil {
            logger.Warn(err)
        }
	    logger.Infof("Sent message to wechat: %s", receiver)
    }
	return nil

    /*
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
    */
}
