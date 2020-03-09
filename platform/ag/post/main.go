package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	//      "fmt"
)


//var baseUrl = "http://wlrevamp.a1go.org/login.aspx?username=XY_666666&langcode=zh-cn&sign=666666&v=XY&cr=RMB"




var param = "Hc6zETdNWOqhDA3uNrF3kB+mucsUUwolnPdvhMSLbhuIsyXhtEeFVzTeJD2DRV33rbBm+pSeKzAue+7sL8UClJPCYTKdgnlq7MGdjqDmoJl9RBruGTIySw8CQJMrHBKR1QFKVtc0Kjo2P2tYmCJslY2lSDH0G02SLoJC5of0zTztR7ZxCnPHE4aHUBPsoFhVw/5DLgnx+JBY+Tikjsq7JY+ZI/QAz0X+yaKxIyrMg1LNdkyKCl+7KllHN1dTfDxGcqCGi2nzA2o="
var key = "51f650c608873a4c75ccc696a0bbeb7a"
//var baseUrl = "https://gci.spark222.com/forwardGame.do?params="+param+"&key="+key
var baseUrl = "https://gci.spark222.com/forwardGame.do"

func main() {
	pp := url.Values{}
	pp.Set("params",param)
	pp.Set("key",key)

	rn, _ := http.NewRequest("POST", baseUrl, strings.NewReader(pp.Encode()))
	rn.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(rn)
	if err != nil {
		log.Printf("err:%+v\n", err)
	}
	readBytes, _ := ioutil.ReadAll(resp.Body)

	log.Printf("return: %+v\n",string(readBytes))

	defer resp.Body.Close()
}
