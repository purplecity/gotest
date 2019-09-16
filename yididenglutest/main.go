package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	tt := time.Now()
	et := time.Date(tt.Year(), tt.Month(), tt.Day(), 12, 0, 0, 0, tt.Location()).Unix()
	t := time.Now().Unix()
	for t < et {
		endpoint := "wss://app-hpoption-ws.azfaster.com:55555/ws/BTCUSDT"
		hpdial := &websocket.Dialer{}
		wsConn,_,err := hpdial.Dial(endpoint,nil)
		if err != nil {
			log.Printf("ERROR----dial  ws failed----err:%+v\n",err)
		}

		rs := map[string]interface{}{"op":"sub","to":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njg2MjA1OTksInVpZCI6IjExNzE2OTkyODIxMjY4MzU3MTIifQ.UnGPNkPXPwaMEr4z4sUeH7eXj2K1LBkaFu7ekOz7lyo"}
		dataByte,_ := json.Marshal(rs)
		wsConn.WriteMessage(websocket.TextMessage,dataByte)

		_, data, _ := wsConn.ReadMessage()
		respmap := map[string]interface{}{}
		json.Unmarshal([]byte(data), &respmap)
		fmt.Printf("%+v\n",respmap)

		wsConn.Close()

		time.Sleep(time.Second*1)
	}
}

