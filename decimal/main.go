package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

//在我们顶多2位小数的情况下 就算是8位也够了
//而且我们的运算都是相同小数位的运算

func HPAdd(x,y float64)  float64 {
	ret,_ := decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y)).Float64()
	return ret
}

func HPSub(x,y float64)  float64 {
	ret, _ := decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y)).Float64()
	return ret
}

func HPround(x float64,y int32) float64 {
	ret,_ := decimal.NewFromFloat(x).Round(y).Float64()
	return ret
}

func main() {
	//fmt.Printf("%T",decimal.NewFromFloat(1.34).Floor())
	//fmt.Println(decimal.NewFromFloat(HPSub(90,1200)).GreaterThanOrEqual(decimal.NewFromFloat(HPSub(3890,5000))))
	fmt.Println(HPround(5.66,5))
	a := fmt.Sprintf("%.5f",5.66)
	fmt.Println(a)
}