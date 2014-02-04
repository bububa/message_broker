package main

import (
	"flag"
	"github.com/bububa/goconfig/config"
	"github.com/bububa/hipchat"
	"github.com/codeforsystemsbiology/verboselogger"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	_CONFIG_FILE = "/var/code/go/config.cfg"
)

var (
	logFlag     = flag.String("log", "", "set log path")
	methodFlag  = flag.String("method", "hipchat", "use which method")
	roomFlag    = flag.String("room", "munin", "send to which room")
	fromFlag    = flag.String("from", "xibao.com", "from who")
	toFlag      = flag.String("to", "", "to who")
	titleFlag   = flag.String("title", "", "title")
	mailToFlag  = flag.String("mailto", "", "mailto")
	formatFlag  = flag.String("format", hipchat.FormatText, "text or html")
	colorFlag   = flag.String("color", hipchat.ColorRed, "color")
	messageFlag = flag.String("message", "", "message to send")
	logger      *log4go.VerboseLogger

	hipchatToken string
	CacheHosts   []string
	wechatAppID  string
	wechatAppKey string
	emailAuth    *EmailAuth
)

func init() {
	cfg, _ := config.ReadDefault(_CONFIG_FILE)
	hctoken, err := cfg.String("hipchat", "notification")
	if err != nil {
		log.Fatalln(err)
	}
	hipchatToken = hctoken
	cacheHosts, err := cfg.String("memcached", "host")
	if err != nil {
		log.Fatalln(err)
	}
	CacheHosts = strings.Split(cacheHosts, ",")
	wechatAppID, _ = cfg.String("wechat", "appid")
	wechatAppKey, _ = cfg.String("wechat", "appkey")
	emailUser, err := cfg.String("message_email", "user")
	if err != nil {
		log.Fatalln(err)
	}
	emailPasswd, err := cfg.String("message_email", "passwd")
	if err != nil {
		log.Fatalln(err)
	}
	emailSmtpHost, err := cfg.String("message_email", "smtphost")
	if err != nil {
		log.Fatalln(err)
	}
	emailSmtpPort, err := cfg.String("message_email", "smtpport")
	if err != nil {
		log.Fatalln(err)
	}
	emailAuth = &EmailAuth{User: emailUser, Passwd: emailPasswd, SMTPHost: emailSmtpHost, SMTPPort: emailSmtpPort}
}

func SetGlobalLogger(logPath string) *log4go.VerboseLogger {
	logger := log4go.NewVerboseLogger(true, nil, "")
	if len(logPath) > 0 {
		logf, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
		if err != nil {
			return logger
		}
		logger = log4go.NewVerboseLogger(true, logf, "")
	}
	return logger
}

func main() {

	flag.Parse()

	logger = SetGlobalLogger(*logFlag)

	var msg string
	if *messageFlag == "" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logger.Warn(err)
			return
		}
		msg = string(bytes)
	} else {
		msg = *messageFlag
	}
	msg = strings.TrimSpace(msg)

	methodsFlag := *methodFlag
	methods := strings.Split(methodsFlag, ",")
	for _, mt := range methods {
		switch mt {
		case "hipchat":
			sendToHipchat(msg)
		case "wechat":
			sendToWechat(msg)
		case "email":
			sendToEmail(*titleFlag, msg)
		}
	}
}
