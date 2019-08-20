package main

import (
	"HPOptionServer/Quotation/Stock/SHCI/Conf"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)



var Mu sync.Mutex
var LastPrice float64

var BaseURL = "https://api.wmcloud.com/data/v1/api/market/getTickRTSnapshotL2.json?securityID=399001.XSHE&field="
func main() {
	request, err := http.NewRequest("GET",BaseURL,nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n",err)
	}
	request.Header.Set("Authorization",Conf.Token)
	request.Header.Set("accept-encoding","gzip")
	client := http.Client{}
	count := 0


	tick := time.Tick(1000 * time.Microsecond)
	for range tick {

		for count < Conf.ReconnctMaxTime {
			resp, err := client.Do(request) //超时设置
			if err != nil  {
				count++
				log.Printf("ERROR----request xshg failed----err:%v\n", err)
				continue

			} else {
				respmap := map[string]interface{}{}
				readBytes, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal([]byte(readBytes), &respmap)
				if respmap["retMsg"].(string) != "Success" {
					count++
					log.Printf("ERROR----xshg reponse err----resp:%v\n", respmap)
					continue
				} else {
					Mu.Lock()
					log.Printf("%+v,%+v\n", time.Now().Unix(),respmap)
					//log.Printf("%T\n", respmap["data"])
					//LastPrice = respmap["data"].([]interface{})[0].(map[string]interface{})["lastPrice"].(float64)
					//log.Printf("%+v\n", LastPrice)
					Mu.Unlock()
					resp.Body.Close()
					break
				}
			}
		}

		if count == Conf.ReconnctMaxTime {
			log.Panicf("ERROR----request xshg 3th failed----err:%v\n", err)
			//停止web下单a股
		}
		count = 0
	}
}
