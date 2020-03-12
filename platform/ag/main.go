package main

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
)


type agreqproperties struct {
	Info string `xml:"info,attr"` // 读取flag属性
	Msg string `xml:"msg,attr"`
}


func EntryptDesECB(data, key []byte) string {
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		log.Printf("EntryptDesECB newCipher error[%v]", err)
		return ""
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		log.Printf("EntryptDesECB Need a multiple of the blocksize")
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	return base64.StdEncoding.EncodeToString(out)
}


func DecryptDESECB(d string, key []byte) string {
	data, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		log.Printf("DecryptDES Decode base64 error[%v]", err)
		return ""
	}
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		log.Printf("DecryptDES NewCipher error[%v]", err)
		return ""
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		log.Printf("DecryptDES crypto/cipher: input not full blocks")
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)

	return string(out)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


var (
	DM =  "http://azfaster.com"
	loginurl = "https://gci.spark222.com/forwardGame.do?"
	baseurl = "https://gi.spark222.com/doBusiness.do"
	cagent = "GI5_AGIN"
	md5key = "DHE3L5rf46YW"
	deskey = "SchB2nnA"
	//deskey = "12341234"
	fenge = "/\\\\\\\\/"
)

func  getLoginURL() {
	loginid := "555556666677714"
	Platusername := "HPpehdF"
	PLTPassword := "123456"

	data := "cagent="+cagent+fenge + "loginname=" + Platusername+fenge+ "actype=1" + fenge+ "password=" + PLTPassword + fenge+"dm="+DM+fenge  + "sid="+cagent+loginid+ fenge+ "lang=1"+fenge+"gameType=18"+fenge+"oddtype=A" + fenge+ "cur=CNY"

	srcBytes := []byte(data)
	keyBytes := []byte(deskey)
	param := EntryptDesECB(srcBytes,keyBytes)




	md5data := param+md5key
	key := fmt.Sprintf("%+x",md5.Sum([]byte(md5data)))

	fmt.Println(loginurl+"params="+param+"&key="+key)
}

//Key MD5(params +” MD5_Encrypt_key”);

/*

	data := []byte(pw)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
*/

func main() {

	//检查创建账户

	/*
	//data := "cagent="+cagent+fenge + "loginname=" + "agtest5"+fenge + "method=lg" + fenge + "actype=1" + fenge + "password=" + "123456" + fenge + "oddtype=A" + fenge+ "cur=CNY"
	data :=  "cagent="+cagent+fenge + "loginname=" + "agtest3"+fenge + "method=gb" + fenge + "actype=1" + fenge + "password=" + "123456" + fenge + "cur=CNY"
	//data :=  "cagent="+cagent+fenge  +"method=tc" + fenge + "loginname=" + "agtest3"+ fenge +  "billno="+cagent+"12345612345612" +fenge + "type=IN" + fenge + "credit=500.00" + fenge + "actype=1" + fenge + "password=" + "123456" + fenge + "cur=CNY"
	//data :=  "cagent="+cagent+fenge  + "method=tc" + fenge + "loginname=" + "agtest3"+fenge + "billno="+cagent+"12345612345613" +fenge + "type=OUT" + fenge + "credit=50.00" + fenge + "actype=1" + fenge + "password=" + "123456" + fenge + "cur=CNY"
	//data :=  "cagent="+cagent+fenge   + "loginname=" + "agtest3"+fenge + "method=tcc" + fenge+ "billno="+cagent+"12345612345613" +fenge + "type=IN" + fenge + "credit=500.00" + fenge + "actype=1" + fenge + "flag=1" +fenge  + "password=" + "123456" + fenge + "cur=CNY"

	log.Println(data)
	srcBytes := []byte(data)
	keyBytes := []byte(deskey)
	param := EntryptDesECB(srcBytes,keyBytes)
	log.Printf("param is %+v\n",param)



	md5data := param+md5key
	key := fmt.Sprintf("%+x",md5.Sum([]byte(md5data)))
	log.Printf("key is %+v\n",key)

	pp := url.Values{}
	pp.Set("params",param)
	pp.Set("key",key)



	r, _ := http.NewRequest("POST", baseurl, strings.NewReader(pp.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("test ag failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)
	teststruct := agreqproperties{}
	xml.Unmarshal(readBytes,&teststruct)
	fmt.Printf("test ag return: %+v,%+v\n",string(readBytes),teststruct)
	defer resp.Body.Close()

//<?xml version="1.0" encoding="utf-8"?><result info="0" msg=""/>
//test ag return: <?xml version="1.0" encoding="utf-8"?><result info="error" msg="error:60001,Account not exist with this currency value or account hierarchical error!"/>,{Info:error Msg:error:60001,Account not exist with this currency value or account hierarchical error!}


	 */



	getLoginURL()
}