package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	r "math/rand"
)




type hpPingPong struct {
	Operation  string `json:"op"`  // sub  ping pong
}

type SubRequest struct {
	Operation  string `json:"op"`  // sub  unsub
	Token  	   string `json:"to"`
}

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}
type loginResponse struct {
	baseResponse
	Username string
	Token  string
	InvitationCode string
	Bal     map[string]float64
	Symbol  string
	cnName  string
}

func genValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	x := len(numeric)
	r.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ r.Intn(x) ])
	}
	return sb.String()
}

func main() {

	ph := "02" + genValidateCode(10)
	x := map[string]string{}
	x["pn"] = ph
	x["pw"] = ph
	x["ic"] = "AGnOFq"
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	url := "http://47.244.212.51:8888/register"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_,err := client.Do(req)
	if err != nil {
		log.Printf("ERROR----regist failed----err:%+v\n", err)
	}

	y := map[string]string{}
	y["pn"] = ph
	y["pw"] = ph
	y["v"] = "0.2.0"
	n, _ := json.Marshal(y)
	var jsonStr2= []byte(n)
	url2 := "http://47.244.212.51:8888/loginByPassword"
	req2, _ := http.NewRequest("POST", url2, bytes.NewBuffer(jsonStr2))
	req2.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req2)
	if err != nil {
		log.Printf("ERROR----login failed----err:%+v\n", err)
	}
	data := loginResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data)
	token := data.Token

	endpoint := "ws://47.244.212.51:55555/ws/BTCUSDT"
	hpdial := &websocket.Dialer{}
	wsConn, _, err := hpdial.Dial(endpoint, nil)
	if err != nil {
		log.Printf("ERROR----dial  ws failed----err:%+v\n", err)
	}

	subdata := SubRequest{Operation: "sub", Token: token}
	dataByte, _ := json.Marshal(subdata)
	wsConn.WriteMessage(websocket.TextMessage, dataByte)
	_, _, _ = wsConn.ReadMessage()

	doneC := make(chan struct{})

	var hpresp map[string]interface{}
	go func() {
		defer wsConn.Close()
		defer close(doneC)

		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Printf("ERROR----read  message failed----err:%+v\n", err.Error())
				return
			}
			json.Unmarshal(message, &hpresp)
			if v, ok := hpresp["op"]; ok && v.(string) == "ping" {
				v, _ := json.Marshal(hpPingPong{Operation: "pong"})
				wsConn.WriteMessage(websocket.TextMessage, v)
				delete(hpresp, "op")
			}
		}
	}()



	now := time.Now()
	st := time.Unix(1565866500,0)
	time.Sleep(st.Sub(now))
	count := 1
	for count <= 60 {
		z := map[string]interface{}{}
		z["am"] = 10
		z["si"] = 1
		z["in"] = 30
		z["sy"] = "BTC"
		z["ts"] = time.Now().Unix()
		z["at"] = 1
		o, _ := json.Marshal(z)
		var jsonStr3= []byte(o)
		url3 := "http://47.244.212.51:8888/trade"
		req3, _ := http.NewRequest("POST", url3, bytes.NewBuffer(jsonStr3))
		req3.Header.Set("Content-Type", "application/json")
		req3.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
		req3.Header.Set("accept-encoding","gzip")
		_,err = client.Do(req3)
		if err != nil {
			log.Printf("ERROR----trade failed----err:%+v\n", err)
		}
		time.Sleep(time.Second*1)
		count++
	}


	<-doneC
}

/*
func main() {
	doneC := make(chan struct{})
	for i:=0;i<10;i++ {
		go do()
	}
	<- doneC
}*/