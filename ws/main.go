package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)


var Event WsTradeEvent


type HistoryData struct {
	P float64	`json:"hp"`
	C int64		`json:"hc"`
	T int64		`json:"ht"`
}

type WsTradeEvent struct {
	Price 		float64	`json:"p"`
	Count		int64	`json:"c"`
	Ts  		int64	`json:"t"`
	OnlineMen	int64	`json:"om"`
	UpPercent	float64	`json:"up"`
	DownPercent	float64	`json:"dp"`
	HistoryData	[]HistoryData	`json:"hd"`
}

type SubResult struct {
	Operation  string `json:"op"`
	Msg        string `json:"msg"`
}

type SubRequest struct {
	Operation  string `json:"op"`  // sub  unsub
}

func main() {
	endpoint := "ws://47.244.216.181:55555/ws/BTCUSDT"
	hpdial := &websocket.Dialer{}
	wsConn,_,err := hpdial.Dial(endpoint,nil)
	if err != nil {
		log.Printf("ERROR----dial  ws failed----err:%+v\n",err)
	}

	req := SubRequest{Operation:"sub"}
	data, _ := json.Marshal(req)
	wsConn.WriteMessage(websocket.TextMessage,data)
	_,msg,err := wsConn.ReadMessage()
	if err != nil {
		fmt.Printf("---err:%v",err)
	}
	resp := SubResult{}
	json.Unmarshal(msg,&resp)
	fmt.Printf("%v\n",resp)
	doneC := make(chan struct{})

	var hpresp  SubRequest
	go func() {
		defer wsConn.Close()
		defer close(doneC)

		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Printf("ERROR----read  message failed----err:%+v\n",err.Error())
				return
			}
			json.Unmarshal(message,&hpresp)
			if hpresp.Operation == "ping" {
				fmt.Printf("receiving ping %v,%v\n",hpresp,time.Now())
				v,_ := json.Marshal(SubRequest{Operation:"pong"})
				wsConn.WriteMessage(websocket.TextMessage,v)
				hpresp.Operation = ""
			} else {
				json.Unmarshal(message,&Event)
				fmt.Printf("%v\n",Event)
			}

		}
	}()
	<-doneC
}