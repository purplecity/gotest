package main

import (
	"fmt"
	"time"
)

func main() {
	var timeLayout = "2006-01-02 15:04:05"
	fmt.Printf("%+T\n",time.Unix(time.Now().Unix(),0).Format(timeLayout))
}