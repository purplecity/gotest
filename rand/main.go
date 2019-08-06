package main

import (
	"fmt"
	"time"
)

/*
func genNumber() int {
	r.Seed(time.Now().UnixNano())
	return r.Intn(9)
}*/

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(fmt.Sprintf("%+v",time.Now().UnixNano()))
}
