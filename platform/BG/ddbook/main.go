package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	agencyloginID = "sparkagency"
	agencypassword = "sparkagency"
	sn = "am00"
	secretkey = "8153503006031672EF300005E5EF6AEF"
	baseurl = "http://am.bgvip55.com/open-cloud/api/"
	agencyID = "166527433"
	secretCode = "m4FnACywgs3Ed1W7MTBYgJLIQoI="

	/*
	 h := sha1.New()
	 h.Write([]byte(agencypassword))
	 fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	*/


)

func main() {
	//注单查询
	//测试拿到的是所有账户的 注单并不是我们平台的注单
	/*
	x := map[string]interface{}{}
	x["random"] = "111112"

	data := []byte("111112"+sn+secretkey)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)


	x["sign"] = md5str
	x["sn"] = sn
	ts := time.Now()
	t := time.Date(2020,3,1,0,0,0,0,ts.Location())
	x["startTime"] = t.Format("2006-01-02 15:04:05")
	y := map[string]interface{}{}
	y["id"] = "5555"
	y["method"] = "open.order.query"
	y["jsonrpc"] = "2.0"
	y["params"] = x

	m, _ := json.Marshal(y)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"open.order.query", bytes.NewBuffer(jsonStr))
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
	z := map[string]interface{}{}
	json.Unmarshal(readBytes,&z)
	fmt.Printf("return: %+v\n",string(readBytes))
	fmt.Printf(" return: %+v\n",z)

	defer dn.Body.Close()

	 */


	//代理注单查询
	//只能查到一天的范围 而且文档没有说明state transid为啥为空

	x := map[string]interface{}{}
	x["random"] = "111112"

	data := []byte("111112"+sn+secretCode)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)


	x["digest"] = md5str
	x["sn"] = sn
	x["agentLoginId"] = agencyloginID
	ts := time.Now()
	t := time.Date(2020,3,1,0,0,0,0,ts.Location())
	t2 := time.Date(2020,3,2,0,0,0,0,ts.Location())
	x["startTime"] = t.Format("2006-01-02 15:04:05")
	x["endTime"] = t2.Format("2006-01-02 15:04:05")
	y := map[string]interface{}{}
	y["id"] = "5555"
	y["method"] = "open.order.agent.query"
	y["jsonrpc"] = "2.0"
	y["params"] = x

	m, _ := json.Marshal(y)
	var jsonStr= []byte(m)
	dn, _ := http.NewRequest("POST", baseurl+"open.order.agent.query", bytes.NewBuffer(jsonStr))
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
	z := map[string]interface{}{}
	json.Unmarshal(readBytes,&z)
	fmt.Printf("return: %+v\n",string(readBytes))
	fmt.Printf(" return: %+v\n",z)

	defer dn.Body.Close()


}
