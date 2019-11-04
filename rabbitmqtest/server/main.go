package main

import (
	"github.com/streadway/amqp"
	"gotest/rabbitmqtest/rabbitmq"
)



func publishMessage(rabbitmqchan *amqp.Channel, exchange, rk string, message []byte) error {
	err := rabbitmqchan.Publish(exchange, rk, false, false, amqp.Publishing{ContentType: "text/plain", Body: message})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	rabbitmq.Webinit()
	publishMessage(rabbitmq.WebCenoddsChannel,)

}
