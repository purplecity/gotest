package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	phlist := []string{"13788888880", "13788888881","13788888882","13788888883","13788888885","13788888886","13788888887"}
	pw :="7722777"
	for _, ph := range phlist {
		data := []byte(pw)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)

		fmt.Printf("%+v,%+v\n",ph,fmt.Sprintf("%x",md5.Sum([]byte(md5str+"HP"+ph))))
	}
}
