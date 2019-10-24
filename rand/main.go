package main

import (
	"fmt"
	"math/rand"
	"time"
)


func genNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20)
}

func main() {
	/*
	s := fmt.Sprintf("1.10%+v",genNumber())
	price,_ := strconv.ParseFloat(s,64)

	fmt.Println(price)

	tick := time.Tick(CommonConf.ForexQuotationInterval * time.Millisecond)
	stoptime := int64(1571224740)
	starttime := int64(1571224920)
	for range tick {
		t := time.Now().Unix()
		if t < stoptime || t > starttime {
			rand.Seed(time.Now().UnixNano())
			s := fmt.Sprintf("1.1%+v",rand.Intn(3000))
			price,_ := strconv.ParseFloat(s,64)
			Mu.Lock()
			LastPrice = price
			fmt.Printf("%+v,%+v\n",time.Now(),LastPrice)
			Mu.Unlock()
		} else {
			Mu.Lock()
			LastPrice = float64(1.10111)
			fmt.Printf("%+v,%+v\n",time.Now(),LastPrice)
			Mu.Unlock()
		}
	}

	 */
	//fmt.Println(genNumber())
	fmt.Println(65/100)
}
