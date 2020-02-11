package main

import (
	"fmt"
	"net/http"
)

func main() {
	var baseUrl = "http://wlrevamp.a1go.org/login.aspx?username=555555&langcode=zh-cn&sign=55555&v=XY&cr=RMB"
	resp, err := http.Get(baseUrl)
	if err != nil {
		fmt.Printf("%+v\n",err)
		return
	}
	defer resp.Body.Close()
}