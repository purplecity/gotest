package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	HandshakeTimeout: 3 * time.Second,
}

func Hpws(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Printf("ERROR----upgrade conn %v to websocket failed----err:%v\n",r.RemoteAddr,err)
		return
	}
	log.Printf("PCTEST----receive addr %v upgrade ws request\n",wsConn.RemoteAddr())

	go testread(wsConn)
	go testwrite(wsConn)
}

func testread(conn *websocket.Conn) {
	rs := map[string]interface{}{"server":"server"}
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

