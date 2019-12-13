package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	r "math/rand"
	"net/http"
	"strings"
	"time"
)


type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}
type loginResponse struct {
	baseResponse
	Username string
	Token  string
	InvitationCode string
	Bal     map[string]float64
	Symbol  string
	cnName  string
}

func genValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	x := len(numeric)
	r.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ r.Intn(x) ])
	}
	return sb.String()
}



var (

	trans = http.Transport{
		DisableKeepAlives:true,
	}
	client = &http.Client{
		Transport:&trans,
	}
)


func printLog(token string) {
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/gameSingleUnFinRank"
	prireq, _ := http.NewRequest("POST", url2, bytes.NewBuffer([]byte{}))
	prireq.Header.Set("Content-Type", "application/json")
	prireq.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	resp, err := client.Do(prireq)
	if err != nil {
		log.Printf("ERROR----gameSingleUnFinRank failed----err:%+v\n", err)
	}
	data := map[string]interface{}{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data) //打印出邀请码 另用于注册账号 打印虚拟金额是否变化
	log.Printf("gameSingleUnFinRank::%+v\n",data)
	prireq.Body.Close()
}

func main() {

	//注册
	ph := "0104" + genValidateCode(10)
	x := map[string]string{}
	x["pn"] = ph
	x["pw"] = ph
	x["ic"] = "wzW2hg"
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	url := "http://app-hpoption-web-test.azfaster.com:8081/register"
	regreq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	regreq.Header.Set("Content-Type", "application/json")


	registrs,err := client.Do(regreq)
	if err != nil {
		log.Printf("ERROR----regist failed----err:%+v\n", err)
	}
	registdata := baseResponse{}
	registbody, _ := ioutil.ReadAll(registrs.Body)
	json.Unmarshal([]byte(registbody), &registdata)
	log.Printf("regist %+v,%+v\n",ph,registdata)
	regreq.Body.Close()
	log.Printf("================RegistFinished===============\n")



	//登录
	y := map[string]string{}
	y["pn"] = ph
	y["pw"] = ph
	y["v"] = "0.7.4"
	n, _ := json.Marshal(y)
	var jsonStr2= []byte(n)
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/loginByPassword"
	logreq, _ := http.NewRequest("POST", url2, bytes.NewBuffer(jsonStr2))
	logreq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(logreq)
	if err != nil {
		log.Printf("ERROR----login failed----err:%+v\n", err)
	}
	data := loginResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data) //打印出邀请码 另用于注册账号 打印虚拟金额是否变化
	token := data.Token
	logreq.Body.Close()
	doneC := make(chan struct{})
	log.Printf("================LoginFinished===============\n")



	//到固定的时间 不同标的物下不同的单 循环
	now := time.Now()
	st := time.Unix(1576218790,0)
	time.Sleep(st.Sub(now))

	for {
		time.Sleep(time.Second*95)
		printLog(token)
	}
	<-doneC
}
