package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	r "math/rand"
)

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
	enName  string
	Bgm     int
	Butttonsound int
	Tradehint    int
	Index   interface{}
	IndexList interface{}
	Mode   	string
}



func main() {

	//注册
	ivcode := "111111" //我们自己的邀请码
	ph := "0105" + genValidateCode(10)
	regmap := map[string]string{}
	regmap["pn"] = ph
	regmap["pw"] = ph
	regmap["ic"] = ivcode
	regbina, _ := json.Marshal(regmap)
	var jsonStr= []byte(regbina)
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
	//已经注册查询
	//登录
	//报名

}
