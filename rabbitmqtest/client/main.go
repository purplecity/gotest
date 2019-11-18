package main

import (
	"fmt"

	//"encoding/json"
	//"fmt"
	"gotest/rabbitmqtest/rabbitmq"
	"log"
	"time"

	//"time"
)

type cenOddsReq struct {
	Amount int64 	`json:"am"`
	Side   int32	`json:"si"`
}

func main() {
	rabbitmq.BTCinit()
	messages, err := rabbitmq.BTCCenoddsChannel.Consume(rabbitmq.BTCCenoddsQueue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("Consume queue: %s faild\n", rabbitmq.BTCCenoddsQueue)
	}

	forever := make(chan bool)


	go func() {
		for message := range messages {
			log.Printf("Receive message: %d - %s\n", time.Now().UnixNano(), message.Body)
			//var hehe = map[string]interface{}{}
			//err = json.Unmarshal(message, &hehe)
			/*
			var req = cenOddsReq{}
			err = json.Unmarshal(message.Body, &req)
			if err != nil {
				fmt.Printf("Marshal mq message faild")
			}

			 */
			fmt.Printf("%+v\n",message)
		}
	}()


	<-forever
}
