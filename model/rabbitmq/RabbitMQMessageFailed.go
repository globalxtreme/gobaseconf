package model

import "github.com/globalxtreme/gobaseconf/model"

type RabbitMQMessageFailed struct {
	model.RabbitMQBaseModel
	MessageId uint                     `gorm:"column:messageId;type:bigint;not null"`
	Sender    string                   `gorm:"column:sender;type:varchar(250);not null"`
	Consumer  string                   `gorm:"column:consumer;type:varchar(250);not null"`
	Key       string                   `gorm:"column:key;type:varchar(250);default:null"`
	Payload   []byte                   `gorm:"column:payload;type:json;default:null"`
	Exception model.MapInterfaceColumn `gorm:"column:exception;type:json;default:null"`
}

func (RabbitMQMessageFailed) TableName() string {
	return "message_faileds"
}
