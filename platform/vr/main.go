package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"
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


)


// 经测试在centos6.10和mac都ok


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


func testBasic() {
	name := "hehetest11"
	jsMap := map[string]string{"playerName":name}
	srcBytes,_ := json.Marshal(jsMap)
	key := []byte("44D0RJXFTJXRR0464248H4B624T80RZH")

	str := EcbEncrypt(srcBytes,key)
	res := base64.StdEncoding.EncodeToString(str)
	fmt.Println(res)

	dres,_ := base64.StdEncoding.DecodeString(res)
	str = EcbDecrypt(dres,key)
	fmt.Println(string(str))
}


func regPlat(name string) error {
	platurl := "https://fe.vrbetdemo.com/Account/CreateUser"
	data := url.Values{}
	data.Set("version",TPlatVersion)
	data.Set("id",TPlatAPPID)

	jsMap := map[string]string{"playerName":name}
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
		readBytes, _ := ioutil.ReadAll(resp.Body)
		if readBytes != nil {
			fmt.Printf("reg plag failed %+v,%+v\n",err,string(readBytes))
			return err
		}
		fmt.Printf("reg plag failed %+v\n",err)
		return err
	}

	readBytes, _ := ioutil.ReadAll(resp.Body)
	dres,_ :=  base64.StdEncoding.DecodeString(string(readBytes))
	str := EcbDecrypt(dres,keyBytes)
	resmap := map[string]interface{}{}
	json.Unmarshal([]byte(str),&resmap)
	fmt.Println(resmap)
	v,ok := resmap["errorCode"]
	if !ok || v.(float64) != 0 {
		return nil
	}
	return nil
}

func loginPlat(name string) {
	platurl := "https://fe.vrbetdemo.com/Account/LoginValidate"
	data := url.Values{}
	data.Set("version",TPlatVersion)
	data.Set("id",TPlatAPPID)

	utcLoc,_ := time.LoadLocation("")
	srcString := "playerName="+name+"&loginTime="+time.Now().In(utcLoc).Format("2006-01-02T15:04:05Z")
	srcBytes := []byte(srcString)
	keyBytes := []byte(TPlatKey)
	dst := EcbEncrypt(srcBytes,keyBytes)
	fdst := base64.StdEncoding.EncodeToString(dst)
	fmt.Println(fdst)
	fdst = url.QueryEscape(fdst)
	//fdst := url.QueryEscape(base64.StdEncoding.EncodeToString(dst))
	//xmlfileread.Set("xmlfileread",fdst)
	data.Set("xmlfileread",fdst)

	fmt.Printf("%+v,%+v,%+v\n",srcString,data,fdst)
	fmt.Println("%+v\n",platurl+"?version="+TPlatVersion+"&id="+TPlatAPPID+"&xmlfileread="+fdst)
	/*
	r, _ := http.NewRequest("POST", platurl, strings.NewReader(xmlfileread.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("login plag failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)
	//mmp := map[string]interface{}{}
	//json.Unmarshal(readBytes,&mmp)
	//fmt.Printf("login plag return: %+v\n",mmp)
	fmt.Printf("login plag return: %+v\n",string(readBytes))
	defer resp.Body.Close()

	 */
}


func depwit(name,uid string,ty int,am float64) error {
	platurl := "https://fe.vrbetdemo.com/UserWallet/Transaction"
	data := url.Values{}
	data.Set("version",TPlatVersion)
	data.Set("id",TPlatAPPID)

	utcLoc,_ := time.LoadLocation("")
	timeString := time.Now().In(utcLoc).Format("2006-01-02T15:04:05Z")

	jsMap := map[string]interface{}{"playerName":name,"serialNumber":uid,"type":ty,"amount":am,"createTime":timeString}
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
		readBytes, _ := ioutil.ReadAll(resp.Body)
		if readBytes != nil {
			fmt.Printf("deposit failed %+v,%+v\n",err,string(readBytes))
			return err
		}
		fmt.Printf("deposit failed %+v\n",err)
		return err
	}

	readBytes, _ := ioutil.ReadAll(resp.Body)
	dres,_ :=  base64.StdEncoding.DecodeString(string(readBytes))
	str := EcbDecrypt(dres,keyBytes)
	resmap := map[string]interface{}{}
	json.Unmarshal([]byte(str),&resmap)
	fmt.Println(resmap)
	return nil
}

func getBal(name string) {
	platurl := "https://fe.vrbetdemo.com/UserWallet/Balance"
	data := url.Values{}
	data.Set("version",TPlatVersion)
	data.Set("id",TPlatAPPID)

	jsMap := map[string]string{"playerName":name}
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
		log.Printf("reg plag failed %+v\n",err)
	}

	readBytes, _ := ioutil.ReadAll(resp.Body)
	dres,_ :=  base64.StdEncoding.DecodeString(string(readBytes))
	str := EcbDecrypt(dres,keyBytes)
	resmap := map[string]interface{}{}
	json.Unmarshal([]byte(str),&resmap)
	log.Println(resmap)
	// map[balance:18414.477 games:[] playerName:hehetest14]
	//map[balance:-1 games:[] playerName:hehetest27]
}



func main () {

	/*
		err := regPlat("hehetest14")
		if err != nil {
			return
		} else {
			time.Sleep(time.Second*5)
			loginPlat("hehetest14")
		}


	 */

	//regPlat("hehetest15")
	//loginPlat("hehetest14")
	//testAesECBEnc("hehetest11")
	//testBasic()
	//depwit("hehetest14","1115",0,50000.111)
	//depwit("hehetest12","1112",1,109.111)
	/*
	x1 := url.QueryEscape("2+XaOq4XB+hqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL/+ZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM=")
	x2 := "2%2BXaOq4XB%2BhqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL%2F%2BZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM%3D"
	fmt.Println(x1==x2)

	 */
	//getBal("hehetest14")
}
