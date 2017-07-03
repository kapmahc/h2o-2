package job

import (
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// Push push task to queue
func Push(ch *amqp.Channel, name string, prv uint8, typ string, body []byte) error {

	qu, err := ch.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return ch.Publish("", qu.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		MessageId:    uuid.New().String(),
		Priority:     prv,
		Body:         body,
		Timestamp:    time.Now(),
		Type:         typ,
	})
}
