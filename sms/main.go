package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var url = "http://sms.nuage.asia/validcodesmsapi.php"
//var url = "http://sms.nuage.asia/validcodesmsapi.php?checksum=8279d4e7d7e2ab29dd3e93f76c1a8017&name=leeotp&pwd=8b3e5f796fc0&receiver=8613760377012&validcode=666666"
func main() {
	req,_ := http.NewRequest("GET",url,nil)


	q := req.URL.Query()
	vc := "【HPOption】您的验证码为:666666"
	ph := "8613760377012"
	na := "leeotp"
	pwd := "8b3e5f796fc0"
	k := "c9534532d44f"
	q.Add("validcode",vc)
	q.Add("receiver",ph)
	q.Add("name",na)
	q.Add("pwd",pwd)
	data := []byte(vc+ph+k)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	fmt.Println(md5str)
	q.Add("checksum",strings.ToLower(md5str))
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	respmap := map[string]interface{}{}
	readBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(readBytes), &respmap)
	fmt.Printf("-----%v----",respmap)
}
