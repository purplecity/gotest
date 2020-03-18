package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	TPlatVersion="1.0"
	TPlatAPPID = "SPARK"
	TPlatKey = "44D0RJXFTJXRR0464248H4B624T80RZH"
	TPBaseurl = "https://fe.vrbetdemo.com/"


)


func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}



func main() {
	/*
	platurl := "https://fe.vrbetdemo.com/MerchantQuery/GameBet"
	xmlfileread := url.Values{}
	xmlfileread.Set("version",TPlatVersion)
	xmlfileread.Set("id",TPlatAPPID)

	bgst := int64(1583719200) //030910
	bget := int64(1583737200)  //030915
	utcLoc,_ := time.LoadLocation("")
	fromtime := time.Unix(bgst,0).In(utcLoc).Format("2006-01-02T15:04:05Z")
	totime := time.Unix(bget,0).In(utcLoc).Format("2006-01-02T15:04:05Z")

	jsMap := map[string]interface{}{"startTime":fromtime,"endTime":totime,"type":0,"recordPage":0,"recordCountPerPage":500}
	srcBytes,_ := json.Marshal(jsMap)
	keyBytes := []byte(TPlatKey)
	dst := EcbEncrypt(srcBytes,keyBytes)
	fdst := base64.StdEncoding.EncodeToString(dst)
	xmlfileread.Set("xmlfileread",fdst)

	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}


	r, _ := http.NewRequest("POST", platurl, strings.NewReader(xmlfileread.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("getdata failed %+v\n",err)
	}

	readBytes, _ := ioutil.ReadAll(resp.Body)
	dres,_ :=  base64.StdEncoding.DecodeString(string(readBytes))
	str := EcbDecrypt(dres,keyBytes)
	resmap := map[string]interface{}{}
	json.Unmarshal([]byte(str),&resmap)
	for k,v := range resmap {
		log.Printf("%+v,%T",k,v)
	}
	for k,v := range resmap["betRecords"].([]interface{})[0].(map[string]interface{}) {
		log.Printf("%+v,%T",k,v)
	}

	log.Println(string(str))
	log.Println(resmap)

	 */

	/*
		020/03/09 22:14:38 betRecords,[]interface {}
		2020/03/09 22:14:38 recordCountPerPage,float64
		2020/03/09 22:14:38 recordPage,float64
		2020/03/09 22:14:38 totalRecords,float64
		2020/03/09 22:14:38 merchantCode,string
		2020/03/09 22:14:38 prize,float64
		2020/03/09 22:14:38 updateTime,string
		2020/03/09 22:14:38 createTime,string
		2020/03/09 22:14:38 channelName,string
		2020/03/09 22:14:38 serialNumber,string
		2020/03/09 22:14:38 playerName,string
		2020/03/09 22:14:38 cost,float64
		2020/03/09 22:14:38 channelId,float64
	*/


	//8

	platurl := "https://fe.vrbetdemo.com/MerchantQuery/Bet"
	data := url.Values{}
	data.Set("version",TPlatVersion)
	data.Set("id",TPlatAPPID)

	bgst := int64(1583719200) //030910
	bget := int64(1583737200)  //030915
	utcLoc,_ := time.LoadLocation("")
	fromtime := time.Unix(bgst,0).In(utcLoc).Format("2006-01-02T15:04:05Z")
	totime := time.Unix(bget,0).In(utcLoc).Format("2006-01-02T15:04:05Z")

	jsMap := map[string]interface{}{"startTime":fromtime,"endTime":totime,"channelId":-1,"state":-1,"recordPage":0,"recordCountPerPage":500}
	srcBytes,_ := json.Marshal(jsMap)
	keyBytes := []byte(TPlatKey)
	dst := EcbEncrypt(srcBytes,keyBytes)
	fdst := base64.StdEncoding.EncodeToString(dst)
	data.Set("xmlfileread",fdst)

	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}


	r, _ := http.NewRequest("POST", platurl, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("getdata failed %+v\n",err)
	}

	readBytes, _ := ioutil.ReadAll(resp.Body)
	dres,_ :=  base64.StdEncoding.DecodeString(string(readBytes))
	str := EcbDecrypt(dres,keyBytes)
	resmap := map[string]interface{}{}
	json.Unmarshal([]byte(str),&resmap)
	for k,v := range resmap {
		log.Printf("%+v,%T",k,v)
	}
	for k,v := range resmap["betRecords"].([]interface{})[0].(map[string]interface{}) {
		log.Printf("%+v,%T",k,v)
	}
	log.Println(string(str))
	log.Println(resmap)

	/*
	2020/03/09 22:16:51 recordCountPerPage,float64
	2020/03/09 22:16:51 recordPage,float64
	2020/03/09 22:16:51 totalRecords,float64
	2020/03/09 22:16:51 betRecords,[]interface {}
	2020/03/09 22:16:51 subState,float64
	2020/03/09 22:16:51 prizeDetail,[]interface {}
	2020/03/09 22:16:51 updateTime,string
	2020/03/09 22:16:51 serialNumber,string
	2020/03/09 22:16:51 cost,float64
	2020/03/09 22:16:51 state,float64
	2020/03/09 22:16:51 count,float64
	2020/03/09 22:16:51 createTime,string
	2020/03/09 22:16:51 note,<nil>
	2020/03/09 22:16:51 betTypeName,string
	2020/03/09 22:16:51 channelName,string
	2020/03/09 22:16:51 odds,string
	2020/03/09 22:16:51 playerPrize,float64
	2020/03/09 22:16:51 merchantPrize,float64
	2020/03/09 22:16:51 multiple,float64
	2020/03/09 22:16:51 position,string
	2020/03/09 22:16:51 unit,float64
	2020/03/09 22:16:51 channelId,float64
	2020/03/09 22:16:51 channelCode,string
	2020/03/09 22:16:51 issueNumber,string
	2020/03/09 22:16:51 number,string
	2020/03/09 22:16:51 playerName,string
	2020/03/09 22:16:51 merchantCode,string
	2020/03/09 22:16:51 lossPrize,float64
	2020/03/09 22:16:51 playerOdds,float64
	2020/03/09 22:16:51 winningNumber,string
	*/

}
