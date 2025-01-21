package rabbitmq

import (
	rabbitmqmodel "github.com/globalxtreme/gobaseconf/model/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"time"
)

const RABBITMQ_CONNECTION_GLOBAL = "global"
const RABBITMQ_CONNECTION_LOCAL = "local"

var (
	RabbitMQSQL  *gorm.DB
	RabbitMQConf rabbitmqconf

	RabbitMQConnectionDial  map[string]*amqp091.Connection
	RabbitMQConnectionCache map[string]rabbitmqmodel.RabbitMQConnection
)

type rabbitmqconf struct {
	Queue      string
	Connection map[string]RabbitMQConnectionConf
	Exchange   RabbitMQExchangeConf
	Timeout    time.Duration
}

type RabbitMQConnectionConf struct {
	Host     string
	Port     string
	Username string
	Password string
}

type RabbitMQExchangeConf struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp091.Table
}
