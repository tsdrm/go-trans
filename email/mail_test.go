package email

import (
	"github.com/Centny/gwf/log"
	"testing"
)

func TestSendEmail(t *testing.T) {
	DefaultMailSender = NewAuth("smtp.126.com:25", "xxx@126.com", "xxx", "xxx@qq.com")

	// send text to a person.
	var err error
	err = SendEmail("email from golang", "xxx@126.com", []string{"xxx@qq.com"}, MAIL_TEXT, "你猜我能不能把邮件发送出去")
	if err != nil {
		t.Error(err)
		return
	}

	// send text to 3 person.
	err = SendEmail("email from golang", "xxx@126.com", []string{"xxx@qq.com", "xxx@qq.com", "xxx@qq.com"}, MAIL_TEXT, "你猜我能不能把邮件发送出去")
	if err != nil {
		t.Error(err)
		return
	}

	// send html to a person
	err = SendEmail("email from golang", "xxx@126.com", []string{"xxx@qq.com"}, MAIL_HTML, "<h1>你好啊朋友</h1>")
	if err != nil {
		t.Error(err)
		return
	}
	log.D("TestSendEmail test success")
}
