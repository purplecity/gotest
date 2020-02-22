package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	clientID = "SKG"
	client_secret = "9c2dF9QsBAe2Q4DSQNQLjujMF1F909tWyemkOMSqTGi2"
	baseurl = "https://staging.tgpaccess.com/"
	zhuanzhangurl = "https://staging.tgpasia.com/"
	access_token="AblKu06zBu7Pj6IlE8yN55oB4UFhX4etuSXbcEmM3KFoOvDygjYKwKFxu6nrb1n0K" //9.45

	ipaddress = "47.244.217.66"
	authtoken = "PnXuypcLETFXGDxuVDhH26EXPnvMjtfpw4cLIveDL5mRrMmjVEupfoRCscOWvOGcrwlcEOsBru6RRY5Alov80prb7H41eU76JEa2jifxb56XAmdHNZFpCnFUwheLT9EIr1"

)

func  main() {

	//加钱

	x := map[string]interface{}{}
	x["userid"] = "111111"
	x["amt"] = 500
	x["cur"] = "RMB"
	x["txid"] = "222222"
	utcLoc,_ := time.LoadLocation("")
	timeString := time.Now().In(utcLoc).Format("2006-01-02T15:04:05Z")+"+00:00"
	x["timestamp"] = timeString


	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", zhuanzhangurl+"/api/wallet/credit", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	dn.Header.Set("Authorization","Bearer "+access_token)
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR----deposit shenbo failed----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("deposit shenbo  return: %+v\n",string(readBytes))

	defer dn.Body.Close()



	//获取余额
	/*
	x := map[string]interface{}{}
	x["userid"] = "111111"
	x["cur"] = "RMB"

	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("GET", zhuanzhangurl+"/api/player/balance", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	dn.Header.Set("Authorization","Bearer "+access_token)
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR----get balance failed----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("get balance  return: %+v\n",string(readBytes))

	defer dn.Body.Close()

	 */


	/*
	dn, err := http.Get( zhuanzhangurl+"/api/player/balance?userid=111111&cur=RMB")
	if err != nil {
		log.Printf("ERROR----get balance failed----err:%+v\n", err)
	}

	readBytes, _ := ioutil.ReadAll(dn.Body)

	fmt.Printf("get balance  return: %+v\n",string(readBytes))

	defer dn.Body.Close()

	 */



	//获取authtoken

	/*
	x := map[string]interface{}{}
	x["ipaddress"] = ipaddress
	x["username"] = "Sparktestskg1"
	x["userid"] = "111111"
	x["lang"] = "zh-CN"
	x["cur"] = "RMB"
	x["betlimitid"] = 1
	x["istestplayer"] = true
	x["platformtype"] = 0

	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"api/player/authorize", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	dn.Header.Set("Authorization","Bearer "+access_token)
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR----get auth token failed----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("get access token  return: %+v\n",string(readBytes))

	defer dn.Body.Close()

	 */




	//获取access token

	/*
	data := url.Values{}
	data.Set("client_id",clientID)
	data.Set("client_secret",client_secret)
	data.Set("grant_type","client_credentials")
	data.Set("scope","playerapi")





	r, _ := http.NewRequest("POST", baseurl+"api/oauth/token", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("get access token failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("get access token  return: %+v\n",string(readBytes))
	defer resp.Body.Close()

	 */


}
