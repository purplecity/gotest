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

/*
AG 13900000000 123456

YSB 13800000000 123456

申博 13700000000 - 13700000009 123456
*/

func main() {
	phlist := []string{"13900000000","13800000000","13700000000","13700000001","13700000002","13700000003","13700000004","13700000005","13700000006","13700000007","13700000008","13700000009"}
	pw :="123456"
	invitationcode := "GH7116"
	baseurl := "http://app-hpoption-web-test.azfaster.com:8081/AddPlayer"
	for _, ph := range phlist {
		data := []byte(pw)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)

		pwd := fmt.Sprintf("%x",md5.Sum([]byte(md5str+"HP"+ph)))

		x := map[string]string{}
		x["pn"] = ph
		x["pw"] = pwd
		x["ic"] = invitationcode
		m, _ := json.Marshal(x)
		var jsonStr= []byte(m)
		rn, _ := http.NewRequest("POST", baseurl, bytes.NewBuffer(jsonStr))
		rn.Header.Set("hpoption","1688")
		if rn != nil &&rn.Body != nil {
			defer rn.Body.Close()
		}
		rn.Header.Set("Content-Type", "application/json")
		trans := http.Transport{
			DisableKeepAlives: true,
		}
		client := &http.Client{
			Transport: &trans,
		}
		rnsp, err := client.Do(rn)
		readbytes,_ := ioutil.ReadAll(rnsp.Body)
		y := map[string]interface{}{}
		json.Unmarshal([]byte(readbytes),y)

		if rnsp != nil &&rnsp.Body != nil {
			defer rnsp.Body.Close()
		}
		if err != nil {
			log.Printf("reg failed----err:%+v\n", err)
		}
		log.Printf("%+v\n",y)
		time.Sleep(time.Second)

	}
}
