package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	agencyloginID = "sparkagency"
	agencypassword = "sparkagency"
	sn = "of06"
	secretkey = "252C03BBD1E34DEFBA851E77375D065A"
	baseurl = "http://n1api.linirn.com/open-cloud/api/"
	agencyID = "166527433"
	secretCode = "m4FnACywgs3Ed1W7MTBYgJLIQoI="

	/*
	 h := sha1.New()
	 h.Write([]byte(agencypassword))
	 fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	*/


)

func main() {
	//创建代理账号


		x := map[string]interface{}{}
		x["random"] = "111111"

		data := []byte("111111"+sn+agencyloginID+secretkey)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["sign"] = md5str
		x["sn"] = sn
		x["loginId"] = agencyloginID
		x["password"] = agencypassword
		x["fromIp"] = "47.244.217.66"
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.agent.create"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.agent.create", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- create agency----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("create agency  return: %+v\n",string(readBytes))

		defer dn.Body.Close()





	//创建账号

	/*

		x := map[string]interface{}{}
		x["random"] = "111112"

		data := []byte("111112"+sn+secretCode)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["digest"] = md5str
		x["sn"] = sn
		x["loginId"] = "testspark2"
		x["nickname"] = "testspark2"
		x["agentLoginId"] = agencyloginID
		x["fromIp"] = "47.244.217.66"
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.user.create"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.user.create", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- create user----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("create user  return: %+v\n",string(readBytes))

		defer dn.Body.Close()

	 */

	/*
	正确信息
	create user
	return :{
		"id": "5555",
		"result": {
			"loginId": "testspark2",
			"snType": 2,
			"success": true,
			"nickname": "testspark2",
			"userId": 167452885,
			"regType": "n"
		},
		"error": null,
		"jsonrpc": "2.0"
	}

	错误信息
	create user
	return :{
		"result": "0",
		"error": {
			"code": "2206",
			"sn": "1583134737140.532797050375454738",
			"message": "登录名已存在,注册用户失败",
			"reason": "loginId:testspark",
			"action": "null"
		},
		"id": "5555",
		"jsonrpc": "2.0",
		"request": {
			"cxt": "/open-cloud",
			"method": "open.user.create",
			"params": {
				"random": "111112",
				"agentLoginId": "sparkagency",
				"loginId": "testspark",
				"reqIp": "47.244.217.66",
				"loginIp": "47.244.217.66",
				"digest": "6f0819bd0a7bc4090fbce8e5682c9425",
				"nickname": "testspark",
				"fromIp": "47.244.217.66",
				"sn": "am00",
				"opIp": "47.244.217.66"
			},
			"uri": null,
			"redirectUri": null,
			"elapsed": 6,
			"from": "47.244.217.66",
			"server": "10.113.1.116",
			"result": null
		}
	}
	*/

		//"loginId":"testspark","snType":2,"success":true,"nickname":"testspark","userId":166533588,"regType":"n"



	//启用会员账号

	/*
		x := map[string]interface{}{}
		x["random"] = "111113"

		data := []byte("111113"+sn+"testspark"+secretCode)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["digest"] = md5str
		x["sn"] = sn
		x["loginId"] = "testspark"
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.user.enable"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.user.enable", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- enable user----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("enable user  return: %+v\n",string(readBytes))

		defer dn.Body.Close()


	 */


	//查询用户状态

	/*

		x := map[string]interface{}{}
		x["random"] = "111114"

		data := []byte("111114"+sn+"testspark"+secretCode)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["digest"] = md5str
		x["sn"] = sn
		x["loginId"] = "testspark"
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.user.get"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.user.get", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- get user----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(readBytes,&y)
		fmt.Printf("login fy return: %+v\n",y)
		fmt.Printf("get user  return: %+v\n",string(readBytes))

		defer dn.Body.Close()

	 */





	/*
	错误的时候
	get user
	return :{
		"result": "0",
		"error": {
			"code": "2213",
			"sn": "1583130463267.8883978014011799935",
			"message": "当前用户ID不存在auth数据,登录失败",
			"reason": "open-api loginId:testspark2, from sn:am00",
			"action": "null"
		},
		"id": "5555",
		"jsonrpc": "2.0",
		"request": {
			"cxt": "/open-cloud",
			"method": "open.user.get",
			"params": {
				"random": "111114",
				"loginId": "testspark2",
				"reqIp": "47.244.217.66",
				"loginIp": "47.244.217.66",
				"digest": "ca7d402f59bc7e7bd5f1b363ba9521bf",
				"fromIp": "47.244.217.66",
				"sn": "am00",
				"opIp": "47.244.217.66"
			},
			"uri": null,
			"redirectUri": null,
			"elapsed": 4,
			"from": "47.244.217.66",
			"server": "10.113.1.115",
			"result": null
		}
	}

	正确的时候
	get user return :{
		"id": "5555",
		"result": {
			"birthday": null,
			"certType": null,
			"userStatus": 1,
			"loginId": "testspark",
			"gender": null,
			"loginLastUpdateTime": "2020-03-01 08:06:34",
			"parentPath": "/166527433",
			"awardPoint": 0,
			"memo": "734544",
			"remark": null,
			"idNumber": null,
			"parentPathIncSelf": "/166527433/166533588/",
			"userImage": null,
			"balance": null,
			"loginIp": "101.244.46.146",
			"nickname": "testspark",
			"tel": null,
			"currency": "1",
			"sn": "am00",
			"wechatName": null,
			"realSn": "am00",
			"regType": "n",
			"email": null,
			"passportNumber": null,
			"qq": null,
			"alipay": null,
			"address": null,
			"mobile": null,
			"wechat": null,
			"alipayName": null,
			"userId": 166533588,
			"parentId": 166527433,
			"regIp": "47.244.217.66",
			"loginCount": 4,
			"recommendUserId": 166527433,
			"certNumber": null,
			"unreadNotice": 0,
			"regTime": "2020-02-27 10:19:48",
			"name": "9TVRx9T9VgHOp93t7375Mvg==",
			"payPassword": null,
			"loginMobile": null,
			"status": 1
		},
		"error": null,
		"jsonrpc": "2.0"
	}
	*/

	//查询余额

	/*
		x := map[string]interface{}{}
		x["random"] = "111115"

		data := []byte("111115"+sn+"testspark"+secretCode)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["digest"] = md5str
		x["sn"] = sn
		x["loginId"] = "testspark05"
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.balance.get"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.balance.get", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- get user balance----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("get user balance  return: %+v\n",string(readBytes))

		defer dn.Body.Close()

		//get user balance  return: {"id":"5555","result":6.4,"error":null,"jsonrpc":"2.0"}

	 */


	//获取视讯地址


	/*

	x := map[string]interface{}{}
	x["random"] = "111118"

	data := []byte("111118"+sn+"testspark"+secretCode)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)


	x["digest"] = md5str
	x["sn"] = sn
	x["loginId"] = "testspark"
	x["isMobileUrl"]=1
	x["isHttpsUrl"] = 0
	y := map[string]interface{}{}
	y["id"] = "5555"
	y["method"] = "open.video.game.url"
	y["jsonrpc"] = "2.0"
	y["params"] = x

	m, _ := json.Marshal(y)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"open.video.game.url", bytes.NewBuffer(jsonStr))
	dn.Header.Set("Content-Type", "application/json")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp,err := client.Do(dn)

	if err != nil {
		log.Printf("ERROR---- get url----err:%+v\n", err)
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("get url  return: %+v\n",string(readBytes))

	defer dn.Body.Close()

	 */









	//额度转换

	/*
		x := map[string]interface{}{}
		x["random"] = "111120"

		data := []byte("111121"+sn+"testspark"+fmt.Sprintf("%+v",200)+secretCode)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["digest"] = md5str
		x["sn"] = sn
		x["loginId"] = "testspark"
		x["amount"] = 200
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.balance.transfer"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.balance.transfer", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- zhuanzhang----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("zhuanzhang  return: %+v\n",string(readBytes))

		defer dn.Body.Close()


	 */




	//转入转出
	/*
		x := map[string]interface{}{}
		x["random"] = "111119"

		data := []byte("111119"+sn+"testspark"+fmt.Sprintf("%+v",200)+secretkey)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)


		x["sign"] = md5str
		x["sn"] = sn
		x["userId"] = 166533588
		x["loginId"] = "testspark"
		x["amount"] = 200
		y := map[string]interface{}{}
		y["id"] = "5555"
		y["method"] = "open.operator.user.transfer"
		y["jsonrpc"] = "2.0"
		y["params"] = x

		m, _ := json.Marshal(y)
		var jsonStr= []byte(m)
		dn, _ := http.NewRequest("POST", baseurl+"open.operator.user.transfer", bytes.NewBuffer(jsonStr))
		dn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives:true,
		}
		client := &http.Client{
			Transport:&trans,
		}
		resp,err := client.Do(dn)

		if err != nil {
			log.Printf("ERROR---- zhuanzhang----err:%+v\n", err)
		}


		readBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("zhuanzhang  return: %+v\n",string(readBytes))

		defer dn.Body.Close()

	*/

}
