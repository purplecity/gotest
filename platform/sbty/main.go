package main

import (
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
	data := url.Values{}
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


	/*
	//登录
	data := url.Values{}
	data.Set("vendor_id",vendor_id)
	data.Set("Vendor_Member_ID","555555")



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

	fmt.Printf("login sbty return: %+v\n",string(readBytes))
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


}
