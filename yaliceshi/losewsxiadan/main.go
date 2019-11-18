package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	r "math/rand"
	"net/http"
	"strings"
	"time"
)


type tradeResponse struct {
	baseResponse
	Tid     string  // 0 order id
	Bal     map[string]float64
	Pe  	float64
	Ht 		int64
	Si 		int32
}

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

	ph := "0101" + genValidateCode(10)
	x := map[string]string{}
	x["pn"] = ph
	x["pw"] = ph
	x["ic"] = "c3gX8n"
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	url := "http://app-hpoption-web-test.azfaster.com:8081/register"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	registrs,err := client.Do(req)
	if err != nil {
		log.Printf("ERROR----regist failed----err:%+v\n", err)
	}
	registdata := baseResponse{}
	registbody, _ := ioutil.ReadAll(registrs.Body)
	json.Unmarshal([]byte(registbody), &registdata)
	log.Printf("regist %+v,%+v\n",ph,registdata)
	req.Body.Close()


	y := map[string]string{}
	y["pn"] = ph
	y["pw"] = ph
	y["v"] = "0.7.0"
	n, _ := json.Marshal(y)
	var jsonStr2= []byte(n)
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/loginByPassword"
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
	req2.Body.Close()
	doneC := make(chan struct{})
	go func() {
		endpoint := "ws:// app-hpoption-web-test.azfaster.com:55555/ws/BTCUSDT"
		hpdial := &websocket.Dialer{}
		wsConn, _, err := hpdial.Dial(endpoint, nil)
		if err != nil {
			log.Printf("ERROR----dial  ws failed----err:%+v\n", err)
		}

		subdata := SubRequest{Operation: "sub", Token: token}
		dataByte, _ := json.Marshal(subdata)
		wsConn.WriteMessage(websocket.TextMessage, dataByte)
		_, _, _ = wsConn.ReadMessage()




		defer wsConn.Close()
		defer close(doneC)

		for {
			var hpresp map[string]interface{}
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Printf("ERROR----read  message failed----err:%+v\n", err.Error())
				return
			}
			json.Unmarshal(message, &hpresp)
			if v, ok := hpresp["op"]; ok && v.(string) == "ping" {
				log.Printf("ERROR----dial  ws failed----err:%+v\n", err)
				v, _ := json.Marshal(hpPingPong{Operation: "pong"})
				wsConn.WriteMessage(websocket.TextMessage, v)
			} else if v,ok := hpresp["op"];ok && v.(string) == "BTCCenOddsNotify" {
				//本来想着隔90s下单的但是实际上没必要 要的就是这一期内疯狂下单赔率变化 直到不能下单为止
				time.Sleep(time.Second*1) //限流
				z := map[string]interface{}{}
				z["am"] = 500
				z["si"] = 1
				z["in"] = 60
				z["sy"] = "BTC"
				z["ts"] = time.Now().Unix()
				z["at"] = 1
				z["ve"] = "0.7.0"
				z["odds"] = hpresp["cuo"].(float64)
				z["m"] = "centralism"
				o, _ := json.Marshal(z)
				var jsonStr3= []byte(o)
				url3 := "http://app-hpoption-web-test.azfaster.com:8081/trade"
				req3, _ := http.NewRequest("POST", url3, bytes.NewBuffer(jsonStr3))
				req3.Header.Set("Content-Type", "application/json")
				req3.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
				req3.Header.Set("accept-encoding","gzip")
				traderesp,err := client.Do(req3)
				if err != nil {
					log.Printf("ERROR----trade failed----err:%+v\n", err)
				}
				tradedata := tradeResponse{}
				tradebody, _ := ioutil.ReadAll(traderesp.Body)
				json.Unmarshal([]byte(tradebody), &tradedata)
				log.Printf("trade::%+v,%+v\n",ph,tradedata)
				req3.Body.Close()
			}

		}
	}()





	now := time.Now()
	st := time.Unix(1573576560,0)
	time.Sleep(st.Sub(now))

	z := map[string]interface{}{}
	z["am"] = 500
	z["si"] = 1
	z["in"] = 60
	z["sy"] = "BTC"
	z["ts"] = st.Unix()
	z["at"] = 1
	z["ve"] = "0.7.0"
	z["odds"] = 0.85
	z["m"] = "centralism"
	o, _ := json.Marshal(z)
	var jsonStr3= []byte(o)
	url3 := "http://app-hpoption-web-test.azfaster.com:8081/trade"
	req3, _ := http.NewRequest("POST", url3, bytes.NewBuffer(jsonStr3))
	req3.Header.Set("Content-Type", "application/json")
	req3.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	req3.Header.Set("accept-encoding","gzip")
	traderesp,err := client.Do(req3)
	if err != nil {
		log.Printf("ERROR----trade failed----err:%+v\n", err)
	}
	tradedata := tradeResponse{}
	tradebody, _ := ioutil.ReadAll(traderesp.Body)
	json.Unmarshal([]byte(tradebody), &tradedata)
	log.Printf("trade::%+v,%+v\n",ph,tradedata)
	req3.Body.Close()
	//time.Sleep(time.Second*1)



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