package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"strings"
	"time"
)

var Event WsTradeEvent
var BTCBaseURL 		= "wss://stream.binance.com:9443/ws"
var timeLayout = "2006-01-02 15:04:05"

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

	var filepath = "/root/dowload/ba.csv"
	file, err := os.OpenFile(filepath,os.O_CREATE|os.O_RDWR,0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n",err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"时间","价格"})
	w.Flush()


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
				w.Write([]string{time.Unix(Event.TradeTime/1e3,0).Format(timeLayout),Event.Price})
				w.Flush()
			}
		}()
		<-doneC
		log.Printf("WARNING----the %v time to reconnect binance ws\n",count)
		count++
	}
}
