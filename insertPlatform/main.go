package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"net/url"
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
	xmlfileread := url.Values{}
	xmlfileread.Set("version","1.0")
	xmlfileread.Set("id","ABC")

	en := "2+XaOq4XB+hqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL/+ZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM="
	xmlfileread.Set("xmlfileread",url.QueryEscape(en))
	log.Println(xmlfileread)

	 */


	/*
	src := []byte("hehetest11")
	key := []byte("44D0RJXFTJXRR0464248H4B624T80RZH")
	dst , _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	fmt.Println(string(dst))
	tmp := base64.StdEncoding.EncodeToString(dst)
	fmt.Println(tmp)  // yXVUkR45PFz0UfpbDB8/ew==
	tmp2,_ := base64.StdEncoding.DecodeString(tmp)
	fmt.Println(string(tmp2))
	dst , _ = openssl.AesECBDecrypt(tmp2, key, openssl.PKCS7_PADDING)
	fmt.Println(string(dst)) // 123456

	 */


	/*
	TPlatKey := "44D0RJXFTJXRR0464248H4B624T80RZH"
	name := "hehetest11"
	jsMap := map[string]string{"playerName":name}
	srcBytes,_ := json.Marshal(jsMap)
	keyBytes := []byte(TPlatKey)
	dst, _ := openssl.AesECBEncrypt(srcBytes,keyBytes,openssl.PKCS7_PADDING)
	fdst := base64.StdEncoding.EncodeToString(dst)
	fmt.Println(fdst)
	ddst,_ := base64.StdEncoding.DecodeString(fdst)
	ffdst , _ := openssl.AesECBDecrypt(ddst, keyBytes, openssl.PKCS7_PADDING)
	fmt.Println(string(ffdst)) // 123456

	 */

	//4MPT0Tri9f+i0GX4B9im37w85GEdjUhl9R5+2LLFfAQ=

	/*
	name := "hehetest11"
	jsMap := map[string]string{"playerName":name}
	srcBytes,_ := json.Marshal(jsMap)
	key := []byte("44D0RJXFTJXRR0464248H4B624T80RZH")

	str := EcbEncrypt(srcBytes,key)
	res := base64.StdEncoding.EncodeToString(str)
	fmt.Println(res)

	dres,_ := base64.StdEncoding.DecodeString("I0BY5qdEViFGE8xnAYa/ug==")
	str = EcbDecrypt(dres,key)
	fmt.Println(string(str))

	 */


	fdst := url.QueryEscape("2+XaOq4XB+hqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL/+ZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM=")
	fmt.Println(fdst=="2%2BXaOq4XB%2BhqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL%2F%2BZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM%3D")
	dst,_ := url.QueryUnescape("2%2BXaOq4XB%2BhqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL%2F%2BZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM%3D")
	fmt.Println(dst=="2+XaOq4XB+hqDMCHBAr4Z1pCXnaLHcyZapdQiDM168dzL/+ZcbMNteN1sMhHYKiOynobPY4X4rTYo3X29EMuDVDNspeh2XKHVUVXR8qPNdM=")
}

