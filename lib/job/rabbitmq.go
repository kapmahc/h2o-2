package job

import (
	"fmt"

	"github.com/streadway/amqp"
)

// RabbitMQ rabbitmq
type RabbitMQ struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Virtual  string `toml:"virtual"`
}

func (p RabbitMQ) String() string {
	return fmt.Sprintf(
		"postgres://%s@%s:%d/%s",
		p.User,
		p.Host,
		p.Port,
		p.Virtual,
	)
}

// Do do
func (p *RabbitMQ) Do(f func(*amqp.Channel) error) error {
	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Virtual,
	))
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	return f(ch)
}
