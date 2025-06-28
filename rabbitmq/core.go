package rabbitmq

import (
	rabbitmqmodel "github.com/globalxtreme/gobaseconf/model/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"time"
)

const RABBITMQ_CONNECTION_GLOBAL = "global"
const RABBITMQ_CONNECTION_LOCAL = "local"

const RABBITMQ_MESSAGE_DELIVERY_STATUS_PENDING_ID = 1
const RABBITMQ_MESSAGE_DELIVERY_STATUS_PENDING = "Pending"
const RABBITMQ_MESSAGE_DELIVERY_STATUS_FINISH_ID = 2
const RABBITMQ_MESSAGE_DELIVERY_STATUS_FINISH = "Finish"
const RABBITMQ_MESSAGE_DELIVERY_STATUS_ERROR_ID = 3
const RABBITMQ_MESSAGE_DELIVERY_STATUS_ERROR = "Error"

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

type RabbitMQDeliveryResponse struct {
	Status rabbitMQDeliveryResponseStatus `json:"status"`
	Error  rabbitMQDeliveryResponseError  `json:"error"`
	Result interface{}                    `json:"result"`
}

type rabbitMQDeliveryResponseStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type rabbitMQDeliveryResponseError struct {
	Message string `json:"message"`
	Trace   string `json:"trace"`
}

type AsyncTransactionForm struct {
	MessageId  uint   `json:"messageId"`
	SenderId   string `json:"senderId"`
	SenderType string `json:"senderType"`
}

type RabbitMQMessageDeliveryStatus struct{}

func (cons RabbitMQMessageDeliveryStatus) OptionIDNames() map[int]string {
	return map[int]string{
		RABBITMQ_MESSAGE_DELIVERY_STATUS_PENDING_ID: RABBITMQ_MESSAGE_DELIVERY_STATUS_PENDING,
		RABBITMQ_MESSAGE_DELIVERY_STATUS_FINISH_ID:  RABBITMQ_MESSAGE_DELIVERY_STATUS_FINISH,
		RABBITMQ_MESSAGE_DELIVERY_STATUS_ERROR_ID:   RABBITMQ_MESSAGE_DELIVERY_STATUS_ERROR,
	}
}

func (cons RabbitMQMessageDeliveryStatus) IDAndName(id int) map[string]interface{} {
	return map[string]interface{}{
		"id":   id,
		"name": cons.Display(id),
	}
}

func (cons RabbitMQMessageDeliveryStatus) Display(id int) string {
	idNames := cons.OptionIDNames()
	if name, ok := idNames[id]; ok {
		return name
	}
	return ""
}
