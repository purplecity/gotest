package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	endpoint := "ws://app-hpoption-ws-test.azfaster.com:55556/ws/SHCI"
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
	rs := map[string]interface{}{"op":"sub","to":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MDg3MzMsInVpZCI6IjExOTAxMDI1MTY1MDk1MjgwNjQifQ.pNaMzcgnPj-PGFcZ37_Cy1VgCuqcmwAoa1-ASAlHymY"}
	dataByte,_ := json.Marshal(rs)
	conn.WriteMessage(websocket.TextMessage,dataByte)
	time.Sleep(5*time.Second)
	if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second)); err != nil {
		log.Println("ping:", err)
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