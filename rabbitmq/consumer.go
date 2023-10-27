package rabbitmq

var (
	RabbitMQConsumer map[string]ConsumerInterface
)

type ConsumerInterface interface {
	Consume(message any) error
}

type Consumer struct {
}

func (Consumer) Set(consumers map[string]ConsumerInterface) {
	RabbitMQConsumer = consumers
}

func (Consumer) Get(key string) ConsumerInterface {
	consumer := RabbitMQConsumer[key]
	return consumer
}
