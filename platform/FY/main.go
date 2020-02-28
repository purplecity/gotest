package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	apikey = "d35b9daaa34042469093826186e511da"
	baseurl = "http://api.bw-gaming.com/"

)

func main() {
	//
	x := map[string]interface{}{}
	x["UserName"] = "111111"
	x["password"] = "testspark"


	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"api/user/register", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	dn.Header.Set("Authorization",apikey)
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR---- reg fy----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("reg fy  return: %+v\n",string(readBytes))

	defer dn.Body.Close()
}