package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}

func main() {
	x := map[string]string{}
	x["pn"] = "test"
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	b64Str := base64.StdEncoding.EncodeToString([]byte(jsonStr))
	buf := bytes.NewBufferString(b64Str)
	/*
	buf := new(bytes.Buffer)
	wr := gzip.NewWriter(buf)
	wr.Write([]byte(b64Str))

	 */
	url := "https://app-hpoption-web.azfaster.com:8081/helloworld"
	req, _ := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Accept-Encoding", "gzip")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	// client.Do(req)

	resp,err := client.Do(req)

	if err != nil {
		log.Printf("ERROR----helloworld failed----err:%+v\n", err)
	}

	gr,err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("http resp unzip is failed,err: ", err)
	}
	readBytes, _ := ioutil.ReadAll(gr)
	your_to_byte, _ := base64.StdEncoding.DecodeString(string(readBytes))
	reqStruct := map[string]interface{}{}
	json.Unmarshal([]byte(your_to_byte),&reqStruct)
	log.Printf("Info----get response----%+v\n",reqStruct)
	defer resp.Body.Close()
	req.Body.Close()


}
