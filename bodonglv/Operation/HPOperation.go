package Operation

import (
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

func HPMul(x,y float64) float64 {
	ret,_ := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y)).Float64()
	return ret
}


//     -1 if x <  y
//      0 if x == y
//     +1 if x >  y
func HPCmp(x,y float64) int {
	ret := decimal.NewFromFloat(x).Cmp(decimal.NewFromFloat(y))
	return ret
}

func HPEqual(x,y float64) bool {
	ret := decimal.NewFromFloat(x).Equal(decimal.NewFromFloat(y))
	return ret
}

func HPgt(x,y float64) bool {
	ret := decimal.NewFromFloat(x).GreaterThan(decimal.NewFromFloat(y))
	return ret
}

func HPgte(x,y float64) bool {
	ret := decimal.NewFromFloat(x).GreaterThanOrEqual(decimal.NewFromFloat(y))
	return ret
}

func HPlt(x,y float64) bool {
	ret := decimal.NewFromFloat(x).LessThan(decimal.NewFromFloat(y))
	return ret
}

func HPlte(x,y float64) bool {
	ret := decimal.NewFromFloat(x).LessThanOrEqual(decimal.NewFromFloat(y))
	return ret
}

func HPpos(x float64) bool {
	ret := decimal.NewFromFloat(x).IsPositive()
	return ret
}

func HPneg(x float64) bool {
	ret := decimal.NewFromFloat(x).IsNegative()
	return ret
}

func HPzero(x float64) bool {
	ret := decimal.NewFromFloat(x).IsZero()
	return ret
}

func HPintpart(x float64) int64 {
	ret := decimal.NewFromFloat(x).IntPart()
	return ret
}

func HPstring(x float64) string {
	ret := decimal.NewFromFloat(x).String()
	return ret
}

func HPSum(first float64, then ...float64) float64 {
	total := first
	for _,v := range then {
		total = HPAdd(total,v)
	}
	return total
}

func HPAbs(x float64) float64 {
	ret,_ := decimal.NewFromFloat(x).Abs().Float64()
	return ret
}

func HPround(x float64,y int32) float64 {
	ret,_ := decimal.NewFromFloat(x).Round(y).Float64()
	return ret
}

func HPSumMul(first float64,then ...float64) float64 {
	total := first
	for _,v := range then {
		total = HPMul(total,v)
	}
	return total
}

func HPDiv(x float64,y float64) float64 {
	ret,_ := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y)).Float64()
	return ret
}


func HPDivInt(x float64, y int64) float64 {
	ret,_ := decimal.NewFromFloat(x).Div(decimal.New(y,0)).Float64()
	return ret
}