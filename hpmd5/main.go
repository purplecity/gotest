package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	phlist := []string{"13285550886"}
	pw :="111111"
	for _, ph := range phlist {
		data := []byte(pw)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)

		fmt.Printf("%+v,%+v\n",ph,fmt.Sprintf("%x",md5.Sum([]byte(md5str+"HP"+ph))))
	}
}
