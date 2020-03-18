package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//SoundNotifyURL := "http://127.0.0.1:80/api/setNotices"
	SoundNotifyURL := "http://app-hpoption-admin.azfaster.com/api/setNotices"
	x := map[string]int{}
	x["type"] = 2
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	rn, _ := http.NewRequest("POST", SoundNotifyURL, bytes.NewBuffer(jsonStr))
	rn.Header.Set("Content-Type", "application/json")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(rn)
	if err != nil {
		log.Printf("ERROR----deposit sounds notify failed----err:%+v\n", err)
	}
	defer rn.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	respmap := map[string]interface{}{}
	json.Unmarshal([]byte(data), &respmap)
	log.Printf("trade::%+v,%+v\n",resp.StatusCode,respmap)

	/*
		SoundNotifyURL := "http://127.0.0.1:9000/api/setNotices"
		xmlfileread := url.Values{}
		xmlfileread.Set("type", string(2))
		dn, _ := http.PostForm(SoundNotifyURL, xmlfileread)
		responseMap := map[string]interface{}{}
		dataBytes, _ := ioutil.ReadAll(dn.Body)
		json.Unmarshal([]byte(dataBytes), &responseMap)
		log.Printf("trade::%+v,%+v\n",dn.StatusCode,responseMap)
	*/
}
