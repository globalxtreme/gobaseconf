package model

import "github.com/globalxtreme/gobaseconf/model"

type RabbitMQConfiguration struct {
	model.RabbitMQBaseModel
	Name  string `gorm:"column:name;type:varchar(200);null"`
	Value string `gorm:"column:value;type:varchar(250);null"`
}

func (RabbitMQConfiguration) TableName() string {
	return "configurations"
}
