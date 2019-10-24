package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"sync"
)

var BTCBaseURL 		= "wss://stream.binance.com:9443/ws"
var Mu sync.Mutex
var Event WsTradeEvent

type WsTradeEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	TradeID       int64  `json:"t"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	BuyerOrderID  int64  `json:"b"`
	SellerOrderID int64  `json:"a"`
	TradeTime     int64  `json:"T"`
	IsBuyerMaker  bool   `json:"m"`
	Placeholder   bool   `json:"M"` // add this field to avoid case insensitive unmarshaling
}


func main() {
	endpoint := fmt.Sprintf("%s/%s@trade", BTCBaseURL, strings.ToLower("BTCUSDT"))
	hpdial := &websocket.Dialer{}
	count := 1

	for {
		wsConn,_,err := hpdial.Dial(endpoint,nil)
		if err != nil {
			log.Printf("ERROR----dial binance ws failed----err:%+v\n",err)
		}
		doneC := make(chan struct{})
		go func() {
			defer wsConn.Close()
			defer close(doneC)

			for {
				_, message, err := wsConn.ReadMessage() //如果币安维护 将祖斯在这
				if err != nil {
					log.Printf("ERROR----read binance message failed----err:%+v\n",err.Error())
					return
				}
				json.Unmarshal(message,&Event)
				fmt.Printf("%+v,%+v\n",Event.TradeTime/1e3,Event.Price)

			}
		}()
		<-doneC
		log.Printf("WARNING----the %v time to reconnect binance ws\n",count)
		count++
	}
}