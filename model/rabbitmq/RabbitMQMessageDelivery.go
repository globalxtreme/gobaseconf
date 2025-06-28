package model

import "github.com/globalxtreme/gobaseconf/model"

type RabbitMQMessageDelivery struct {
	model.RabbitMQBaseModel
	MessageId        uint                           `gorm:"column:messageId;type:bigint;null"`
	ConsumerService  string                         `gorm:"column:consumerService;type:varchar(100);null"`
	StatusId         int                            `gorm:"column:statusId;type:tinyint;null"`
	NeedNotification bool                           `gorm:"column:needNotification;type:tinyint;null"`
	Responses        *model.ArrayMapInterfaceColumn `gorm:"column:responses;type:json;default:null"`
}

func (RabbitMQMessageDelivery) TableName() string {
	return "message_deliveries"
}
