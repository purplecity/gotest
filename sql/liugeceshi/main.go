package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)



type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}


func main() {
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}

	ph := int(16880001)
	for ph <= 168801000 {
		x := map[string]string{}
		s := strconv.Itoa(ph)
		x["pn"] =  s
		x["pw"] =  s
		x["ic"] = "1KxlvP"
		m, _ := json.Marshal(x)
		var jsonStr= []byte(m)
		url := "https://app-hpoption-web.azfaster.com:8081/register"
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		registrs,err := client.Do(req)
		if err != nil {
			log.Printf("ERROR----regist failed----err:%+v\n", err)
		}
		registdata := baseResponse{}
		registbody, _ := ioutil.ReadAll(registrs.Body)
		json.Unmarshal([]byte(registbody), &registdata)
		log.Printf("regist %+v,%+v\n",ph,registdata)
		req.Body.Close()
		time.Sleep(time.Second*1)
		ph++
	}
}