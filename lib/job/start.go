package job

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

// Start start background worker
func Start(ch *amqp.Channel, name string) error {
	log.Infof("waiting for messages, to exit press CTRL+C")
	if err := ch.Qos(1, 0, false); err != nil {
		return err
	}
	qu, err := ch.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(qu.Name, name, false, false, false, false, nil)
	if err != nil {
		return err
	}
	for m := range msgs {
		m.Ack(false)
		log.Infof("receive message %s@%s", m.MessageId, m.Type)
		now := time.Now()
		if hnd, ok := handlers[m.Type]; ok {
			if er := hnd(m.Body); er == nil {
				log.Infof("done %s time %s", m.MessageId, time.Now().Sub(now))
			} else {
				log.Errorf("failed %s time %s with error [%s]", m.MessageId, time.Now().Sub(now), er)
			}

			return err
		}
		log.Errorf("unknown message type %s", m.Type)
	}
	return nil
}
