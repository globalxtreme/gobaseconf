package mail

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

type Mail struct {
	queue  string
	dialer *mail.Dialer
}

func (m *Mail) Dial(conf config.MailConf) *Mail {
	m.dialer = mail.NewDialer(conf.Host, mailPort(conf.Port), conf.Username, conf.Password)

	return m
}

func (m *Mail) Send(msg MailMessage) error {
	content := msg.Message()
	content.SetHeader("From", content.FormatAddress(os.Getenv("MAIL_FROM_ADDRESS"), os.Getenv("MAIL_FROM_NAME")))

	if err := m.dialer.DialAndSend(content); err != nil {
		xtremelog.Error(fmt.Sprintf("Error sending email: %v", err), true)
		return err
	}

	return nil
}

func mailPort(port string) int {
	p, _ := strconv.Atoi(port)
	return p
}
