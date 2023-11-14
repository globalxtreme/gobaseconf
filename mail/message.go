package mail

import (
	"gopkg.in/mail.v2"
)

type MailMessage interface {
	Message() *mail.Message
}
