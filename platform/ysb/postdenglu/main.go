package main

import (
	"log"
	"net/http"
	"strings"
)

var baseUrl = "http://mwlrevamp.a1go.org/login.aspx?username=XY_HPpehdF&langcode=zh-cn&sign=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQxNTg4NzMsImh0eXBlIjoiRGlyZWN0b3IiLCJ1aWQiOiIxMjM1NDQ5ODg1OTc3NzUxNTUyIn0.57qFstAb5yqjcQKYC-qTs6-jo60sR6hyztMgVAqoW0c&v=XY&cr=RMB"

/*
func main() {
	resp, err := http.Get(baseUrl)
	if err != nil {
		fmt.Printf("%+v\n",err)
		return
	}
	defer resp.Body.Close()
}

 */


func main() {

	rn, _ := http.NewRequest("POST", baseUrl, strings.NewReader(""))
	rn.Header.Set("Content-Type", "text/xml")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	_,err := client.Do(rn)
	if err != nil {
		log.Printf("ERROR----chongzhi 1 req----err:%+v\n", err)
	}

	defer rn.Body.Close()
}