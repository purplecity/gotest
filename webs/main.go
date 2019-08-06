package  main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
)


//var baseURL = "wss://stream.binance.com:9443/ws"
var baseURL = "ws://47.244.217.66:55555/ws"
var symbol = "BTCUSDT"


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

func handleTradeMsg(msg []byte) {
	event := new([]WsTradeEvent)
	json.Unmarshal(msg,event)
	fmt.Printf("%+v\n",*event)
}



func main() {
	endpoint := fmt.Sprintf("%s/%s", baseURL, strings.ToUpper(symbol))
	hpdial := &websocket.Dialer{}

	wsConn,_,err := hpdial.Dial(endpoint,nil)
	if err != nil {
		fmt.Printf("dial failed %+v",err)
		return
	}
	doneC := make(chan struct{})
	stopC := make(chan struct{})
	go func() {
		defer func() {
			err := wsConn.Close()
			if err != nil {
				fmt.Printf("close wsconn %+v",err)
			}
		}()
		//defer close(doneC)

		for {
			select {
			case <-stopC:
				return
			default:
				_, message, err := wsConn.ReadMessage()
				if err != nil {
					fmt.Printf("read message %+v",err.Error())
					return
				}
				handleTradeMsg(message)
				//fmt.Printf("%+v\n",message)
			}
		}
	}()
	<-doneC
	return
}