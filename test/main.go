package main

import "fmt"

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
	fmt.Println(1.7e308 > float64(24*3600*1000000000))
}