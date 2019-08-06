package  main

import (
	"HPOptionServer/Quotation/Stock/AShareIndex/Conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)
var Mu sync.Mutex
var LastPrice float64
var BaseURL = "https://api.wmcloud.com/data/v1/api/market/getTickRTSnapshotL2.json?securityID=000001.XSHG&field="

func main() {
	request, err := http.NewRequest("GET",Conf.BaseURL,nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n",err)
	}
	request.Header.Set("Authorization",Conf.Token)
	request.Header.Set("accept-encoding","gzip")
	client := http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	count := 0

	now := time.Now()
	st := time.Unix(Conf.Starttime,0)
	time.Sleep(st.Sub(now))
	tick := time.Tick(Conf.TaskInterval * time.Microsecond)
	for range tick {
		fmt.Printf("time:%v---",time.Now())
		for count < Conf.ReconnctMaxTime {
			resp, err = http.DefaultClient.Do(request) //超时设置
			defer resp.Body.Close()
			if err != nil  {
				count++
				log.Printf("ERROR----request xshg failed----err:%v\n", err)
				continue

			} else {
				respmap := map[string]interface{}{}
				readBytes, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal([]byte(readBytes), &respmap)
				Mu.Lock()
				LastPrice = respmap["data"].([]interface{})[0].(map[string]interface{})["lastPrice"].(float64)
				Mu.Unlock()
				fmt.Printf("price:%v\n",LastPrice)
				break
			}
		}

		if count == Conf.ReconnctMaxTime {
			log.Panicf("ERROR----request  xshg 3th failed----err:%v\n", err)
			//停止web下单a股
		}
		count = 0
	}
}
