package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"time"
)


//var Mu sync.Mutex
//var Event WsTradeEvent
var Event = map[string]interface{}{}

var BTCBaseURL = "wss://stream.binance.com:9443/ws/btcusdt@depth10"


func main() {
	var filepath= "/root/go/src/gotest/btcdepth/btcdepth.csv"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"时间","买1价","买2价","买3价","买4价","买5价","买6价","买7价","买8价","买9价","买10价", "卖1价","卖2价","卖3价","卖4价","卖5价","卖6价","卖7价","卖8价","卖9价","卖10价",
		"买1量","买2量","买3量","买4量","买5量","买6量","买7量","买8量","买9量","买10量", "卖1量","卖2量","卖3量","卖4量","卖5量","卖6量","卖7量","卖8量","卖9量","卖10量"})
	w.Flush()

	hpdial := &websocket.Dialer{}
	count := 1

	for {
		wsConn,_,err := hpdial.Dial(BTCBaseURL,nil)
		if err != nil {
			log.Printf("ERROR----dial binance ws failed----err:%+v\n",err)
			time.Sleep(time.Second*1)
			continue
		}
		doneC := make(chan struct{})
		go func() {
			defer wsConn.Close()
			defer close(doneC)

			for {
				_, message, err := wsConn.ReadMessage() //如果币安维护 将祖斯在这
				if err != nil {
					log.Printf("ERROR----read binance message failed----err:%+v\n", err.Error())
					return
				}
				if message != nil {
					json.Unmarshal(message, &Event)
					w.Write([]string{fmt.Sprintf("%+v", time.Now().Unix()),
						Event["bids"].([]interface{})[0].([]interface{})[0].(string),
						Event["bids"].([]interface{})[1].([]interface{})[0].(string),
						Event["bids"].([]interface{})[2].([]interface{})[0].(string),
						Event["bids"].([]interface{})[3].([]interface{})[0].(string),
						Event["bids"].([]interface{})[4].([]interface{})[0].(string),
						Event["bids"].([]interface{})[5].([]interface{})[0].(string),
						Event["bids"].([]interface{})[6].([]interface{})[0].(string),
						Event["bids"].([]interface{})[7].([]interface{})[0].(string),
						Event["bids"].([]interface{})[8].([]interface{})[0].(string),
						Event["bids"].([]interface{})[9].([]interface{})[0].(string),
						Event["asks"].([]interface{})[0].([]interface{})[0].(string),
						Event["asks"].([]interface{})[1].([]interface{})[0].(string),
						Event["asks"].([]interface{})[2].([]interface{})[0].(string),
						Event["asks"].([]interface{})[3].([]interface{})[0].(string),
						Event["asks"].([]interface{})[4].([]interface{})[0].(string),
						Event["asks"].([]interface{})[5].([]interface{})[0].(string),
						Event["asks"].([]interface{})[6].([]interface{})[0].(string),
						Event["asks"].([]interface{})[7].([]interface{})[0].(string),
						Event["asks"].([]interface{})[8].([]interface{})[0].(string),
						Event["asks"].([]interface{})[9].([]interface{})[0].(string),
						Event["bids"].([]interface{})[0].([]interface{})[1].(string),
						Event["bids"].([]interface{})[1].([]interface{})[1].(string),
						Event["bids"].([]interface{})[2].([]interface{})[1].(string),
						Event["bids"].([]interface{})[3].([]interface{})[1].(string),
						Event["bids"].([]interface{})[4].([]interface{})[1].(string),
						Event["bids"].([]interface{})[5].([]interface{})[1].(string),
						Event["bids"].([]interface{})[6].([]interface{})[1].(string),
						Event["bids"].([]interface{})[7].([]interface{})[1].(string),
						Event["bids"].([]interface{})[8].([]interface{})[1].(string),
						Event["bids"].([]interface{})[9].([]interface{})[1].(string),
						Event["asks"].([]interface{})[0].([]interface{})[1].(string),
						Event["asks"].([]interface{})[1].([]interface{})[1].(string),
						Event["asks"].([]interface{})[2].([]interface{})[1].(string),
						Event["asks"].([]interface{})[3].([]interface{})[1].(string),
						Event["asks"].([]interface{})[4].([]interface{})[1].(string),
						Event["asks"].([]interface{})[5].([]interface{})[1].(string),
						Event["asks"].([]interface{})[6].([]interface{})[1].(string),
						Event["asks"].([]interface{})[7].([]interface{})[1].(string),
						Event["asks"].([]interface{})[8].([]interface{})[1].(string),
						Event["asks"].([]interface{})[9].([]interface{})[1].(string),
					})
					w.Flush()
				} else {
					time.Sleep(time.Second*1)
				}
			}
		}()
		<-doneC
		log.Printf("WARNING----the %v time to reconnect binance ws\n",count)
		count++
	}
}

/*
func main() {
        var filepath= "/Users/ludongdong/zaqizaba/btcdepth.csv"
        file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
        if err != nil {
                fmt.Println("open file failed err:%+v\n", err)
        }
        defer file.Close()
        file.WriteString("\xEF\xBB\xBF")
        w := csv.NewWriter(file)
        w.Write([]string{"时间","买1价","买2价","买3价","买4价","买5价","买6价","买7价","买8价","买9价","买10价","卖1价","卖2价","卖3价","卖4价","卖5价","卖6价","卖7价","卖8价","卖9价","卖10价",
                "买1量","买2量","买3量","买4量","买5量","买6量","买7量","买8量","买9量","买10量","卖1量","卖2量","卖3量","卖4量","卖5量","卖6量","卖7量","卖8量","卖9量","卖10量"})
        w.Flush()





}

*/