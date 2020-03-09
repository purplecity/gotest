package main

import (

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (

	baseurl = "https://staging.tgpaccess.com/"
	access_token="Z8SYrTxeg7z4u4KxYg3vi78D7FVTHaxHFEoeiU0rwuYLLqeHjgqKFs3EEIRnMxzkK" //9.45


)


func main() {


	starttime :=url.QueryEscape( time.Unix(1583719200,0).Format("2006-01-02T15:04:05"))
	endtime :=url.QueryEscape( time.Unix(1583755200,0).Format("2006-01-02T15:04:05"))


	heheurl := baseurl+"/api/history/bets?startdate="+starttime+"&enddate="+endtime+"&includetestplayers=true&issettled=true"
	dn, _ := http.NewRequest("GET", heheurl, nil)
	dn.Header.Set("X-Tgp-Accept", "json")
	dn.Header.Set("Authorization","Bearer "+access_token)
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR----get shenbo data failed----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("get data  return: %+v\n",string(readBytes))



}
