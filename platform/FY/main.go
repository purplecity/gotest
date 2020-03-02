package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	apikey = "d35b9daaa34042469093826186e511da"
	baseurl = "http://api.bw-gaming.com/"

)

func main() {
	// 注册

	/*
	data := url.Values{}
	data.Set("UserName","spark01")
	data.Set("password","testspark")




	r, _ := http.NewRequest("POST", baseurl+"api/user/register", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization",apikey)

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("reg fy failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)
	y := map[string]interface{}{}
	json.Unmarshal(readBytes,&y)
	fmt.Printf("login fy return: %T,%T,%+vn",y["success"],y["info"],y["info"])
	fmt.Printf("reg fy return: %+v,%+v\n",string(readBytes),y)
	defer resp.Body.Close()


	 */



	//登录


	data := url.Values{}
	data.Set("UserName","spark01")




	r, _ := http.NewRequest("POST", baseurl+"api/user/login", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization",apikey)

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("login fy failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	y := map[string]interface{}{}
	json.Unmarshal(readBytes,&y)
	fmt.Printf("login fy return: %T,%T,%+vn",y["success"],y["info"],y["info"])
	fmt.Printf("login fy return: %+v,%+v\n",string(readBytes),y)
	defer resp.Body.Close()





	//查询余额

	/*

	data := url.Values{}
	data.Set("UserName","spark01")




	r, _ := http.NewRequest("POST", baseurl+"api/user/balance", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization",apikey)

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("balance fy failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("balance fy return: %+v\n",string(readBytes))
	defer resp.Body.Close()

	 */



	//转账
	/*
	data := url.Values{}
	data.Set("UserName","spark01")
	data.Set("Money","50")
	data.Set("Type","OUT")
	data.Set("ID","111112")




	r, _ := http.NewRequest("POST", baseurl+"api/user/transfer", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization",apikey)

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("transfer fy failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("transfer fy return: %+v\n",string(readBytes))
	defer resp.Body.Close()

	 */

}