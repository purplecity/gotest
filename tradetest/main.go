

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	y := map[string]string{}
	y["pn"] = "13760377012"
	y["pw"] = "a276c21e4cfda3efc0d91dad6a5b63d9"
	y["v"] = "0.4.9"
	n, _ := json.Marshal(y)
	var jsonStr2= []byte(n)
	url2 := "https://app-hpoption-webapi.azfaster.com:8081/loginByPassword"
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

	tt := time.Now()
	et := time.Date(tt.Year(), tt.Month(), tt.Day(), 11, 30, 0, 0, tt.Location()).Unix()
	t := tt.Unix()
	for t < et {
		z := map[string]interface{}{}
		z["am"] = 10
		z["si"] = 1
		z["in"] = 30
		z["sy"] = "SHCI"
		z["ts"] = time.Now().Unix() - 2
		z["at"] = 1
		z["ve"] = "0.4.9"
		o, _ := json.Marshal(z)
		var jsonStr3 = []byte(o)
		url3 := "https://app-hpoption-webapi.azfaster.com:8081/trade"
		req3, _ := http.NewRequest("POST", url3, bytes.NewBuffer(jsonStr3))
		req3.Header.Set("Content-Type", "application/json")
		req3.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		traderesp, err := client.Do(req3)
		if err != nil {
			log.Printf("ERROR----trade failed----err:%+v\n", err)
		}
		tradedata := tradeResponse{}
		tradebody, _ := ioutil.ReadAll(traderesp.Body)
		json.Unmarshal([]byte(tradebody), &tradedata)
		log.Printf("trade::%+v\n", tradedata)
		req3.Body.Close()
		time.Sleep(time.Second * 10)
		t += 10
	}
}

