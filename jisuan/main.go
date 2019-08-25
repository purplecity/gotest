package main

import (
	"fmt"
	"gotest/jisuan/Operation"
)

var(
	ab = float64(0.81946)
	ba = Operation.HPSub(float64(1),ab)

	ac = float64(0.81547)
	ca = Operation.HPSub(float64(1),ac)

	bc = float64(0.81935)
	cb = Operation.HPSub(float64(1),bc)


)

/*
AB BC AC
AB AC BC

BC AB AC
BC AC AB

AC BC AB
AC AB BC

AB BC
AB CB AC
AB AC


BC AB
CB AB AC
BC AC AB
CB AC

AC BC AB
AC CB
AC AB
*/

func main() {

	a := Operation.HPSum(Operation.HPSumMul(ab,ac,float64(4)),Operation.HPSumMul(ab,bc,float64(2)),Operation.HPSumMul(ac,cb,float64(2)))
	b := Operation.HPSum(Operation.HPSumMul(ba,bc,float64(4)),Operation.HPSumMul(ba,ac,float64(2)),Operation.HPSumMul(bc,ca,float64(2)))
	c := Operation.HPSum(Operation.HPSumMul(ca,cb,float64(4)),Operation.HPSumMul(cb,ba,float64(2)),Operation.HPSumMul(ca,ab,float64(2)))
	fmt.Println(a,b,c)
	fmt.Println(Operation.HPSum(a,b,c))
	fmt.Println(Operation.HPdiv(Operation.HPMul(a,float64(100)),float64(6)))


	/*
		a := Operation.HPSum(Operation.HPMul(ab,ac),Operation.HPMul(ab,bc),Operation.HPMul(ac,cb))
		b := Operation.HPSum(Operation.HPMul(ba,bc),Operation.HPMul(ba,ac),Operation.HPMul(bc,ca))
		c := Operation.HPSum(Operation.HPMul(ca,cb),Operation.HPMul(cb,ba),Operation.HPMul(ca,ab))
	*/
}