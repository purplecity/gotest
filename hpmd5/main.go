package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	pw :="123456"
	ph := "13100000000"
	data := []byte(pw)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)

	fmt.Printf("%v\n",fmt.Sprintf("%x",md5.Sum([]byte(md5str+"HP"+ph))))
}

