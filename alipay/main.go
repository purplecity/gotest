package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	uid := "testtest"
	tid := "testtest"
	Amount := 50
	AliPayID := "34b70ed6481645b496ef46a0a97dc446"
	AliPayAPPID := "8c87527e6c494e04b890fad1791a746c"
	AliPayKey := "e1b6aeb75a71418ab10c04191d37a754"
	AliPayURL := "http://order.order202.com:8800/api/v1/trades/v2page/"
	x := map[string]interface{}{}
	x["app_user"] = AliPayID
	x["app_id"] = AliPayAPPID
	x["out_trade_sn"] = tid
	x["coin_type"] = "107"
	x["quantity"] = strconv.Itoa(Amount)
	x["account_type"] = int(3)
	waitString := fmt.Sprintf("account_type=%+v&app_id=%+v&app_user=%+v&coin_type=%+v&out_trade_sn=%+v&quantity=%+v",
		3,AliPayAPPID,uid,107,tid,Amount)
	h := hmac.New(sha256.New, []byte(AliPayKey))
	h.Write([]byte(waitString))
	x["sign"] = hex.EncodeToString(h.Sum(nil))
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", AliPayURL, bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	depositResp,err := client.Do(dn)
	if err != nil {
		log.Printf("ERROR----request deposit failed----err:%+v\n", err)
	}

	responseMap := map[string]interface{}{}
	dataBytes, _ := ioutil.ReadAll(depositResp.Body)
	json.Unmarshal([]byte(dataBytes), &responseMap)
	log.Printf("DEPOSIT----tid:%+v,%+v\n",tid,responseMap)
}
