package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (

	baseurl = "https://staging.tgpaccess.com/"
	access_token="Lu14gKJRXiTcqZhdzRAunNo9QRtwRR8nSAuEXdEw99qad6M43cSIFagPxUxBCPwBZ" //9.45


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
		log.Printf("ERROR----get shenbo xmlfileread failed----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	y := []interface{}{}
	json.Unmarshal(readBytes,&y)
	for k,v := range y[0].(map[string]interface{}) {
		log.Printf("%+v,%T",k,v)
	}

	fmt.Printf("get xmlfileread  return: %+v\n",string(readBytes))

	/*
	2020/03/09 23:35:54 beton,string
	2020/03/09 23:35:54 turnover,float64
	2020/03/09 23:35:54 validbet,float64
	2020/03/09 23:35:54 winamt,float64
	2020/03/09 23:35:54 bettype,string
	2020/03/09 23:35:54 txid,string
	2020/03/09 23:35:54 timestamp,string
	2020/03/09 23:35:54 betupdatedon,string
	2020/03/09 23:35:54 gameprovidercode,string
	2020/03/09 23:35:54 gamename,string
	2020/03/09 23:35:54 ugsbetid,string
	2020/03/09 23:35:54 betid,string
	2020/03/09 23:35:54 username,string
	2020/03/09 23:35:54 gameprovider,string
	2020/03/09 23:35:54 riskamt,float64
	2020/03/09 23:35:54 cur,string
	2020/03/09 23:35:54 platformtype,string
	2020/03/09 23:35:54 ipaddress,string
	2020/03/09 23:35:54 playtype,string
	2020/03/09 23:35:54 playertype,float64
	2020/03/09 23:35:54 betclosedon,string
	2020/03/09 23:35:54 roundstatus,string
	2020/03/09 23:35:54 postbal,float64
	2020/03/09 23:35:54 gameid,string
	2020/03/09 23:35:54 roundid,string
	2020/03/09 23:35:54 userid,string
	2020/03/09 23:35:54 winloss,float64
	2020/03/09 23:35:54 beforebal,float64
	*/




}
