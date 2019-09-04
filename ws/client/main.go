package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	endpoint := "ws://127.0.0.1:55555/ws/BTCUSDT"
	hpdial := &websocket.Dialer{}
	wsConn,_,err := hpdial.Dial(endpoint,nil)
	if err != nil {
		log.Printf("ERROR----dial  ws failed----err:%+v\n",err)
	}
	go testread(wsConn)
	go testwrite(wsConn)
	doneC := make(chan struct{})
	<-doneC
}


func testread(conn *websocket.Conn) {
	rs := map[string]interface{}{"client":"client"}
	dataByte,_ := json.Marshal(rs)
	for i:=0;i<100;i++ {
		conn.WriteMessage(websocket.TextMessage,dataByte)
		time.Sleep(time.Second*5)
	}

}

func testwrite(conn *websocket.Conn) {
	for {
		_, data, _ := conn.ReadMessage()
		respmap := map[string]interface{}{}
		json.Unmarshal([]byte(data), &respmap)
		fmt.Printf("%+v\n",respmap)
	}

}