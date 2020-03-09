package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var  (
	YSBPrefix = "1429"
	YSBsecretkey = "N37Wq9cfPSz24fyb"
	YSBDATAURL = "http://webapi.a1go.org/WhiteLabelApi/api/Vendor"
)

func main() {
	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}


	type ysbparam struct {
		MsgId string
		DateFrom string
		DateTo string
		VendorId string
		HashPassword string
	}
	bgst := int64(1583719200) //030910
	bget := int64(1583737200)  //030915
	data := map[string]interface{}{}
	param := ysbparam{}
	utcLoc,_ := time.LoadLocation("")
	fromtime := time.Unix(bgst,0).In(utcLoc).Format("06-01-02T15:04")
	totime := time.Unix(bget,0).In(utcLoc).Format("06-01-02T15:04")

	param.MsgId = "BD"
	param.DateFrom = fromtime
	param.DateTo = totime
	param.VendorId = YSBPrefix
	srcbytes := []byte("BD"+fromtime+totime+YSBPrefix+YSBsecretkey)
	log.Println(string(srcbytes))
	md5str := fmt.Sprintf("%x", md5.Sum(srcbytes))
	param.HashPassword = strings.ToUpper(md5str)
	data["Params"] = param
	log.Printf("%+v\n",param)
	jsonstr, _ := json.Marshal(data)

	r, _ := http.NewRequest("POST", YSBDATAURL,bytes.NewBuffer(jsonstr)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("get ysb  order data failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)
	y := map[string]interface{}{}
	json.Unmarshal(readBytes,&y)
	log.Printf("get  ysb  order data return: %+v\n",string(readBytes))

}
