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
	log.Println(length)
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
	platurl := "https://fe.vrbetdemo.com/MerchantQuery/GameBet"
	data := url.Values{}
	data.Set("version",TPlatVersion)
	data.Set("id",TPlatAPPID)

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
	data.Set("data",fdst)

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
	log.Println(len(readBytes))
	dres,_ :=  base64.StdEncoding.DecodeString(string(readBytes))
	str := EcbDecrypt(dres,keyBytes)
	resmap := map[string]interface{}{}
	json.Unmarshal([]byte(str),&resmap)
	log.Println(string(str))
	log.Println(resmap)
	//8
	/*
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
	data.Set("data",fdst)

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
	log.Println(string(str))
	log.Println(resmap)

	 */
}
