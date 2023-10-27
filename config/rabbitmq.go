package config

import (
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"time"
)

var (
	RabbitMQSQL  *gorm.DB
	RabbitMQConf rabbitmqconf
)

type rabbitmqconf struct {
	Queue      string
	Connection RabbitMQConnection
	Exchange   RabbitMQExchange
	Timeout    time.Duration
}

type RabbitMQConnection struct {
	Host     string
	Port     string
	Username string
	Password string
}

type RabbitMQExchange struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp091.Table
}
