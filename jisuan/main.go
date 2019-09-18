package main

import (
	"fmt"
	"gotest/jisuan/Operation"
	"strconv"
)

func getSDP(p float64) int {
	a := fmt.Sprintf("%.2f",Operation.HPround(p,2))
	b,_:= strconv.Atoi(a[len(a)-2:len(a)])
	fmt.Printf("%v\n",b)
	x := b/10
	y := b - x*10
	fmt.Printf("%v,%v\n",x,y)
	if x ==y {
		if y % 2 == 0 {
			return 1 //对双
		} else {
			return 2 //对单
		}
	} else {
		if y % 2 == 0 {
			return 3 //双
		} else {
			return 4 //单
		}
	}
}
func main() {
	fmt.Println(getSDP(55.11))
}