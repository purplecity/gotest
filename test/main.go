package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	a := []float64{108.66,119.836,1.10188,119.836,108.66,1.10239,119.836,108.661,1.10228,108.661 ,119.836 ,1.10257}

	b := []float64{119.836 ,108.66 ,1.10239 ,119.836 ,108.661 ,1.10228 ,108.661 ,119.836 ,1.10257 ,108.661 ,119.836 ,1.10169}

	c := []float64{119.836 ,108.661 ,1.10228 ,108.661 ,119.836 ,1.10257 ,108.661 ,119.836 ,1.10169 ,108.661 ,119.836 ,1.10277}
	sort.Float64s(a[:])
	sort.Float64s(b[:])
	sort.Float64s(c[:])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	 */
	//a,_ := strconv.ParseFloat("3.89e-06",64)
	//fmt.Printf("%+v\n",a)
	tm := time.Now()
	fmt.Println(int(tm.Month()),tm.Day(),tm.Hour(),tm.Minute(),tm.Second())

	tick := time.Tick(time.Second*1)
	for range tick {
		fmt.Println("hehe")
		time.Sleep(10000*time.Second)
	}
}