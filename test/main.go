package main

import (
	"fmt"
	"time"
)


func getIssue(tt int64) {
	t := time.Unix(tt,0)
	tt = time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,t.Location()).Unix()

	w := 0
	d := 0
	a := (tt - 1577030400)/(24*3600)
	//有几个周日
	sunNum :=  int(a / 7)
	fmt.Printf("---- %+v,%+v\n",a,sunNum)

	w = 20001 + sunNum - 1
	d = 10001 + sunNum*5-1

	if int(t.Weekday()) == 6 {
		d = d + 5
	} else if int(t.Weekday()) >= 1 && int(t.Weekday()) <=5 {
		d = d + int(t.Weekday()) - 1
	} else if int(t.Weekday()) == 0 {
		w = w + 1
		d = d + 5
	}

	fmt.Printf("%+v,%+v\n",w,d)

}
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
	//tm := time.Now()
	//fmt.Println(tm.Year(),int(tm.Month()),tm.Day(),tm.Hour(),tm.Minute(),tm.Second())
	//fmt.Println(1.7e308 > float64(24*3600*1000000000))
	getIssue(1578164400)
}