package email

import (
	"net/smtp"
	"strings"
)

const MAIL_HTML = "html"
const MAIL_TEXT = "text"

type Auth struct {
	SMTP      string // just like ip:port, e.g smtp.example.com:25
	Username  string
	Password  string
	Receivers []string // email receivers.
	auth      smtp.Auth
}

var DefaultMailSender = NewAuth("", "", "")

func NewAuth(addr string, username string, password string, receiver ...string) *Auth {
	var domain string
	addrStrings := strings.Split(addr, ":")
	if len(addrStrings) > 0 {
		domain = addrStrings[0]
	}
	auth := smtp.PlainAuth("", username, password, domain)

	var receivers []string
	if len(receiver) > 0 {
		receivers = append(receivers, receiver...)
	}

	return &Auth{
		SMTP:      addr,
		Username:  username,
		Password:  password,
		Receivers: receivers,
		auth:      auth,
	}
}

func SendEmail(subject string, from string, to []string, mailType string, message string) error {
	return DefaultMailSender.SendEmail(subject, from, to, mailType, message)
}

func (a *Auth) SendEmail(subject string, from string, to []string, mailType string, message string) error {
	var contentType = "text/plain; charset=UTF-8"
	if mailType == MAIL_HTML {
		contentType = "text/html; charset=UTF-8"
	}
	var msg = "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + contentType + "\r\n\r\n" +
		message + "\r\n"
	return smtp.SendMail(a.SMTP, a.auth, from, to, []byte(msg))
}

// Return email sender username
func From() string {
	return DefaultMailSender.From()
}
func (a *Auth) From() string {
	return a.Username
}

// Return email receivers
func To() []string {
	return DefaultMailSender.To()
}
func (a *Auth) To() []string {
	return a.Receivers
}
