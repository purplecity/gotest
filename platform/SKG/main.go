package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	clientID = "SKG"
	client_secret = "9c2dF9QsBAe2Q4DSQNQLjujMF1F909tWyemkOMSqTGi2"
	baseurl = "https://staging.tgpaccess.com/"
	zhuanzhangurl = "https://staging.tgpasia.com/"
	access_token="2HEBKy1FgkD7I0xzeSAfPuCM8KIYxgdB7vzpqRC8RMM49yPvv1asyaky8J22P8gTo" //9.45

	ipaddress = "47.244.217.66"
	authtoken = "Z8SYrTxeg7z4u4KxYg3vi78D7FVTHaxHFEoeiU0rwuYLLqeHjgqKFs3EEIRnMxzkK"

)

func  main() {

	//提现

	/*
	x := map[string]interface{}{}
	x["userid"] = "111111"
	x["amt"] = 50
	x["cur"] = "RMB"
	x["txid"] = "333333"
	//utcLoc,_ := time.LoadLocation("")
	//timeString := time.Now().In(utcLoc).Format("2006-01-02T15:04:05Z")+"+00:00"
	timeString := time.Now().Format("2006-01-02T15:04:05Z")
	x["timestamp"] = timeString


	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"api/wallet/debit", bytes.NewBuffer(jsonStr))
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

	 */

	//加钱


	/*
	x := map[string]interface{}{}
	x["userid"] = "111112"
	x["amt"] = 1000
	x["cur"] = "RMB"
	x["txid"] = "222223"
	//utcLoc,_ := time.LoadLocation("")
	//timeString := time.Now().In(utcLoc).Format("2006-01-02T15:04:05Z")+"+00:00"
	timeString := time.Now().Format("2006-01-02T15:04:05Z")
	x["timestamp"] = timeString


	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"api/wallet/credit", bytes.NewBuffer(jsonStr))
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

	 */







	//获取余额



	/*
	nurl := baseurl+"/api/player/balance?userid=111120&cur=RMB"
	fmt.Println(nurl)
	dn, _ := http.NewRequest("GET", nurl, nil)
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


	 */







	/*
	dn, err := http.Get( zhuanzhangurl+"/api/player/balance?userid=111111&cur=RMB")
	if err != nil {
		log.Printf("ERROR----get balance failed----err:%+v\n", err)
	}
	http.Heade

	readBytes, _ := ioutil.ReadAll(dn.Body)

	fmt.Printf("get balance  return: %+v\n",string(readBytes))

	defer dn.Body.Close()


	 */






	//获取authtoken


	/*
	x := map[string]interface{}{}
	x["ipaddress"] = ipaddress
	x["username"] = "Sparktestskg2"
	x["userid"] = "111112"
	x["lang"] = "zh-CN"
	x["cur"] = "RMB"
	x["betlimitid"] = 1
	x["istestplayer"] = true
	x["platformtype"] = 1

	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"api/player/authorize", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	dn.Header.Set("Authorization","Bearer "+access_token+"1")
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


	//{"authtoken":"0KMR453pdOLnepF8BbTFcTjL6yOKegE9JzgCiVNSsjqSwUjzSlIsZsSkvN3wmXW6FO8lYQhNQa5Yewj6ZkIIkmUAIVVIyZoJGg49IvCS2ZILtjHjnUiBqLN3DzCdNVhWG","isnew":false}
	//{"err":12,"errdesc":"Cannot decrypt brand token; INNER 1: Length of the data to decrypt is invalid."}




	//获取access token




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







	// {"access_token":"RU4VSWG8hL3EnHfBZiWeUUB6v5JTnYU8SEQqQoU4VU17CgM7Z0we0xnn36KVKFhUi","token_type":"Bearer","expires_in":3600,"scope":"playerapi"}




}
