package model

import "github.com/globalxtreme/gobaseconf/model"

type RabbitMQConnection struct {
	model.RabbitMQBaseModel
	Connection string `gorm:"column:connection;type:varchar(50);null"`
	Service    string `gorm:"column:service;type:varchar(150);null"`
}

func (RabbitMQConnection) TableName() string {
	return "connections"
}
