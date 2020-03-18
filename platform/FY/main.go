package main

import (
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
	xmlfileread := url.Values{}
	xmlfileread.Set("UserName","spark01")
	xmlfileread.Set("password","testspark")



	

	r, _ := http.NewRequest("POST", baseurl+"api/user/register", strings.NewReader(xmlfileread.Encode())) // URL-encoded payload
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


	/*
	xmlfileread := url.Values{}
	xmlfileread.Set("UserName","spark01")




	r, _ := http.NewRequest("POST", baseurl+"api/user/login", strings.NewReader(xmlfileread.Encode())) // URL-encoded payload
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


	 */




	//查询余额

	/*

	xmlfileread := url.Values{}
	xmlfileread.Set("UserName","spark01")




	r, _ := http.NewRequest("POST", baseurl+"api/user/balance", strings.NewReader(xmlfileread.Encode())) // URL-encoded payload
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
	xmlfileread := url.Values{}
	xmlfileread.Set("UserName","spark01")
	xmlfileread.Set("Money","50")
	xmlfileread.Set("Type","OUT")
	xmlfileread.Set("ID","111113")




	r, _ := http.NewRequest("POST", baseurl+"api/user/transfer", strings.NewReader(xmlfileread.Encode())) // URL-encoded payload
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

	//查询转账结果
	data := url.Values{}
	data.Set("ID","111113")




	r, _ := http.NewRequest("POST", baseurl+"api/user/transferinfo", strings.NewReader(data.Encode())) // URL-encoded payload
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
		fmt.Printf("transferinfo fy failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("transferinfo fy return: %+v\n",string(readBytes))
	defer resp.Body.Close()


}