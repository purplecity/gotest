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

/*
func main() {
	//a := []float64{1.012,1.011,1.024,1.024,1.025,1.026,1.012,1.011,1.024,1.024,1.025,1.026}

	a := []float64{1345,1301,1368,1322,1310,1370,1318,1350,1303,1299}
	rs := float64(0)
	for _,x := range a {
		rs = Operation.HPAdd(rs,x)
	}
	rs = Operation.HPDivInt(rs,int64(len(a)))
	fmt.Println("avg ",rs)
	//rs = Operation.HPround(rs,16)
	fmt.Println("avg ",rs)

	rs2 := float64(0)
	for _,x := range a {
		rs2 = Operation.HPAdd(rs2,Operation.HPMul(Operation.HPSub(x,rs),Operation.HPSub(x,rs)))
	}
	fmt.Println("double ",rs2)
	//rs2 = Operation.HPround(rs2,16)
	fmt.Println("double ",rs2)

	rs3 := float64(0)
	rs3 = Operation.HPDivInt(rs2,int64(len(a)-1))
	fmt.Println("div ",rs3)
	//rs2 = Operation.HPround(rs3,16)
	fmt.Println("div ",rs3)

	fmt.Println(math.Sqrt(rs3))
}


 */

func main() {
	//fmt.Println(Operation.HPAdd(1.6,Operation.HPTrunc(2.757,2)))
	//scoreMap[v.Contributorid] = Operation.HPAdd(scoreMap[v.Contributorid], Operation.HPMul(Operation.HPAbs(Operation.HPSub(v.Outputamount,v.Amount)),CommonConf.ScoreMap[v.Type]))
	a := 2
	b := a
	a = 3
	println(a,b)
}