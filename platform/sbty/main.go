package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var  (
	baseurl = "http://tsa.a0214.sport-test.azfaster.com/api/"
	vendor_id = "uhixgxbdhn"
	currencyID =20
	OperatorId = "Spark"
	OddsType = "2"
	Currency = "20"
	MinTransfer = "1"
	MaxTransfer = "5000000"
)


func main() {

	/* 注册
	data := url.Values{}a
	data.Set("vendor_id",vendor_id)
	data.Set("Vendor_Member_ID","555555")
	data.Set("OperatorId",OperatorId)
	data.Set("UserName","555555")
	data.Set("OddsType",OddsType)
	data.Set("Currency",Currency)
	data.Set("MaxTransfer",MaxTransfer)
	data.Set("MinTransfer",MinTransfer)


	r, _ := http.NewRequest("POST", baseurl+"CreateMember", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("reg sbty failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("reg sbty return: %+v\n",string(readBytes))
	defer resp.Body.Close()

	 */




	//登录

	/*
	data := url.Values{}
	data.Set("vendor_id",vendor_id)
	data.Set("Vendor_Member_ID","HPYFdZ6")



	r, _ := http.NewRequest("POST", baseurl+"Login", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("login sbty failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	url222 := "http://sbtest.a0214.sport-test.azfaster.com/deposit_processlogin.aspx?lang=zhcn&webskintype=2&token="

	d := map[string]interface{}{}
	json.Unmarshal(readBytes,&d)
	fmt.Printf("login sbty return:%+v, %+v,%+v\n",time.Now(),string(readBytes),url222+d["Data"].(string))

	defer resp.Body.Close()


	 */









	/*充值
	data := url.Values{}
	data.Set("vendor_id",vendor_id)
	data.Set("Vendor_Member_ID","555555")
	data.Set("vendor_trans_id","5555551")
	data.Set("amount","500")
	data.Set("currency",Currency)
	data.Set("direction","1")
	data.Set("wallet_id","1")





	r, _ := http.NewRequest("POST", baseurl+"FundTransfer", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("deposit sbty failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("deposit sbty return: %+v\n",string(readBytes))
	defer resp.Body.Close()

	 */




	/*
	data := url.Values{}
	data.Set("vendor_id",vendor_id)
	data.Set("Vendor_Member_ID","555555")
	data.Set("vendor_trans_id","5555552")
	data.Set("amount","50")
	data.Set("currency",Currency)
	data.Set("direction","0")
	data.Set("wallet_id","1")





	r, _ := http.NewRequest("POST", baseurl+"FundTransfer", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("deposit sbty failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("deposit sbty return: %+v\n",string(readBytes))
	defer resp.Body.Close()


	 */


	//获取下单数目




	/*
	data := url.Values{}
	data.Set("vendor_id",vendor_id)
	data.Set("version_key","1111")




	r, _ := http.NewRequest("POST", baseurl+"GetBetDetail", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("get bet sbty failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	d := map[string]interface{}{}
	json.Unmarshal(readBytes,&d)
	fmt.Printf("%+T",d["Data"].(map[string]interface{})["BetDetails"])
	_,ok := d["Data"].(map[string]interface{})["BetDetails"]
	if ok {
		fmt.Printf("get bet sbty return: %+v\n",len(d["Data"].(map[string]interface{})["BetDetails"].([]interface{})))
		fmt.Printf("get bet sbty return: %+v\n",string(readBytes))
	} else {
		fmt.Printf("get bet sbty return: %+v\n",string(readBytes))
	}
	defer resp.Body.Close()

	 */


	/*
	x := map[string]interface{}{}
	x["vendor_id"] = vendor_id
	x["vendor_member_id"] = "HPYFdZ6"
	y := []map[string]interface{}{}
	y = append(y, map[string]interface{}{"sport_type":"1","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"2","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"3","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"5","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"8","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"11","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"43","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"99","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"161","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"180","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})
	y = append(y, map[string]interface{}{"sport_type":"190","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":10000000})

	x["bet_setting"] = y



	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"SetMemberBetSetting", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR---- SetMemberBetSettingfailed----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("SetMemberBetSetting  return: %+v\n",string(readBytes))

	defer dn.Body.Close()

	 */





	data := url.Values{}
	data.Set("vendor_id",vendor_id)
	data.Set("vendor_member_id","HPYFdZ6")
	y := []map[string]interface{}{}
	y = append(y, map[string]interface{}{"sport_type":"1","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"2","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"3","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"5","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"8","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"11","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"43","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"99","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"161","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"180","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"190","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"10","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})
	y = append(y, map[string]interface{}{"sport_type":"99MP","min_bet":0,"max_bet":10000000,"max_bet_per_match":0,"max_bet_per_ball":5000000})

	da,_ := json.Marshal(y)
	data.Set("bet_setting",string(da))


	r, _ := http.NewRequest("POST", baseurl+"SetMemberBetSetting?vendor_id="+vendor_id+"&vendor_member_id=HPYFdZ6", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "text/json")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf(" bet sbty failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	d := map[string]interface{}{}
	json.Unmarshal(readBytes,&d)
	fmt.Printf(" sbty return: %+v\n",string(readBytes))
	defer resp.Body.Close()







	/*

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	version_key := "146077"

	tick := time.Tick(time.Second*120)
	for range tick {
		data := url.Values{}
		data.Set("vendor_id",vendor_id)
		fmt.Println(time.Now(),version_key)
		data.Set("version_key",version_key)




		r, _ := http.NewRequest("POST", baseurl+"GetBetDetail", strings.NewReader(data.Encode())) // URL-encoded payload
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")



		resp, err := client.Do(r)
		if err != nil {
			fmt.Printf("get bet sbty failed %+v\n",err)
			return
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		d := map[string]interface{}{}
		json.Unmarshal(readBytes,&d)
		fmt.Printf("%+T",d["Data"].(map[string]interface{})["BetDetails"])
		_,ok := d["Data"].(map[string]interface{})["BetDetails"]
		version_key = fmt.Sprintf("%+v\n",d["Data"].(map[string]interface{})["last_version_key"])
		if ok {
			fmt.Printf("get bet sbty return: %+v\n",len(d["Data"].(map[string]interface{})["BetDetails"].([]interface{})))
			fmt.Printf("get bet sbty return: %+v\n",string(readBytes))
		} else {
			fmt.Printf("get bet sbty return: %+v\n",string(readBytes))
		}
		fmt.Println("+++++++++++++++++++++++++===================================")
		defer resp.Body.Close()

	}

	 */

}
