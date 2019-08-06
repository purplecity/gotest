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

func main() {
	fmt.Println(0.5-0.3)
}