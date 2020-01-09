package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotest/testsql/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	r "math/rand"
)

var (
	trans = http.Transport{
		DisableKeepAlives:true,
	}
	client = &http.Client{
		Transport:&trans,
	}

	BTCCenList = []int64{20,50,100,200,500,1000,2000}
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

type tradeResponse struct {
	baseResponse
	Tid     string  // 0 order id
	Bal     map[string]float64
	Pe  	float64
	Ht 		int64
	Ot   	int64
	Si 		int32
	Issue   int64
	Ts  	[]interface{}
}


func genNumber() int {
	r.Seed(time.Now().UnixNano())
	return r.Intn(7)
}

func login(ph string) (token string) {
	y := map[string]string{}
	y["pn"] = ph
	y["pw"] = ph
	y["v"] = "0.7.5"
	n, _ := json.Marshal(y)
	var jsonStr2= []byte(n)
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/loginByPassword"
	req2, _ := http.NewRequest("POST", url2, bytes.NewBuffer(jsonStr2))
	req2.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req2)
	if err != nil {
		log.Printf("ERROR----login failed----err:%+v\n", err)
	}
	data := loginResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data)
	token = data.Token
	req2.Body.Close()
	log.Printf("================LoginFinished===============\n")
	return
}

func reg(token,apiname string,)  {
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/"+apiname
	req, _ := http.NewRequest("POST", url2, bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR----dayGame failed----err:%+v\n", err)
	}
	data := map[string]interface{}{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data)
	log.Printf("%+v---- %+v\n",apiname,data)
	req.Body.Close()
	log.Printf("================RegGameFinished===============\n")
}


func trade( token,symbol string) {
	side := genNumber() //0 1 2 3 4对应金额 20 30 40 50 60 意味着2000块钱至少可以下30次 1800秒 半小时
	amount := BTCCenList[side]

	z := map[string]interface{}{}
	z["am"] = amount
	z["si"] = side
	z["in"] = 60
	z["sy"] = symbol
	z["ts"] = time.Now().Unix()
	z["ve"] = "0.7.5"
	z["at"] = 0
	z["m"] = "centralism"

	o, _ := json.Marshal(z)
	var jsonStr3= []byte(o)

	url3 := "http://app-hpoption-web-test.azfaster.com:8081/tradeSDP"

	trareq, _ := http.NewRequest("POST", url3, bytes.NewBuffer(jsonStr3))
	trareq.Header.Set("Content-Type", "application/json")
	trareq.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	traderesp,err := client.Do(trareq)
	if err != nil {
		log.Printf("ERROR----trade failed----err:%+v\n", err)
	}
	tradedata := tradeResponse{}
	tradebody, _ := ioutil.ReadAll(traderesp.Body)
	json.Unmarshal([]byte(tradebody), &tradedata)
	log.Printf("trade::%+v::%+v\n",tradedata,side)
	trareq.Body.Close()
}

func main() {
	//睡眠到某一期
	adminInfo := []mysql.AdminUsers{}
	mysql.GetAllRecord("AdminUsers", map[string]interface{}{"phonenumber__contains":"hp"},&adminInfo) //不到1000条可以容的下
	fmt.Printf("len:%+v\n",len(adminInfo))
	for _, x := range adminInfo {
		//登录
		token := login(x.Phonenumber)
		//报名
		reg(token,"dayGame")
		time.Sleep(time.Millisecond*500)
		//玩一把  能玩得区间 随机选择金额 同时下单刚好测试grpc的性能能力 但是有赔率的问题



	}
}