package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"gotest/rabbitmqtest/rabbitmq"
	"math"
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
	x := map[string]interface{}{}
	x["am"] = int64(math.Floor(1.6))
	x["si"] = 1
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	publishMessage(rabbitmq.WebCenoddsChannel,rabbitmq.BTCExchange,rabbitmq.BTCCenoddsRK,jsonStr)
	forever := make(chan bool)
	<- forever
}