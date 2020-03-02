package main

import (
	"log"
	"net/http"
	"strings"
)

var baseUrl = "http://mwlrevamp.a1go.org/login.aspx?username?username=XY_55555&langcode=zh-cn&sign=55555&v=XY&cr=RMB"

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