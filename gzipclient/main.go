package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	url := "http://127.0.0.1:8888/helloworld"
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
	/*
	data := baseResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data)
	log.Printf("get response %+v\n",data)

	 */

	reader,err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("http resp unzip is failed,err: ", err)
	}
	var buf2 []byte
	readBytes, _ := reader.Read(buf2)
	your_to_byte, _ := base64.StdEncoding.DecodeString(string(readBytes))
	reqStruct := map[string]interface{}{}
	json.Unmarshal([]byte(your_to_byte),&reqStruct)
	log.Printf("Info----get response----%+v\n",reqStruct)
	defer reader.Close()
	defer resp.Body.Close()
	req.Body.Close()


}
