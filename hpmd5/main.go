package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	phlist := []string{"13788888888","13975368407","15173364957","15886327401","13866668888","19186803207","19186803302"}
	pw :="qaqa2277"
	for _, ph := range phlist {
		data := []byte(pw)
		has := md5.Sum(data)
		md5str := fmt.Sprintf("%x", has)

		fmt.Printf("%+v,%+v\n",ph,fmt.Sprintf("%x",md5.Sum([]byte(md5str+"HP"+ph))))
	}
}
