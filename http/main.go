package  main

import (
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}

func main() {
	url := "https://app-hpoption-webapi.azfaster.com:8081/loginByPassword"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
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
	body, err := gzip.NewReader(registrs.Body)
	registbody, _ := ioutil.ReadAll(body)
	your_to_byte, _ := base64.StdEncoding.DecodeString(string(registbody))
	your_string := string(your_to_byte)
	log.Printf("string %+v\n",your_string)
	req.Body.Close()
}
