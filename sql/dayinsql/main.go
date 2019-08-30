package main

import (
	"fmt"
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"time"
)

func main() {

	now := time.Now()
	//et := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, now.Location()).Unix()
	//last := now.Add(time.Hour * -24)
	//bt := time.Date(last.Year(), last.Month(), last.Day(), 3, 0, 0, 0, last.Location()).Unix()

	bt := 1566460274
	et := now.Unix()
	withdrawInfo := []mysql.Withdrawrecord{}
	Uid := "1164444907327197184"
	mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid":Uid,"Createtime__gte":bt*1000,"Createtime__lt":et*1000,"Status__in":[]int{0,1}},&withdrawInfo)
	depositInfo := []mysql.Depositrecord{}
	mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid":Uid,"Finishtime__gte":bt*1000,"Finishtime__lt":et*1000,"Status":1},&depositInfo)
	takeScoreInfo := []mysql.Takescorerecord{}
	mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid":Uid,"Handletime__gte":bt,"Handletime__lt":et},&takeScoreInfo)
	realTradeWinInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et,"Orderresult":3},&realTradeWinInfo)
	realTradeLoseInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et,"Orderresult":2},&realTradeLoseInfo)

	wa := float64(0)
	da := float64(0)
	ta := float64(0)
	rwa := float64(0)
	rla := float64(0)
	for _, x := range withdrawInfo {
		wa = Operation.HPAdd(wa, x.Amount)
	}

	for _, x := range depositInfo {
		da = Operation.HPAdd(da, x.RealAmount)
	}
	for _, x := range takeScoreInfo {
		ta = Operation.HPAdd(ta, x.Amount)
	}
	for _, x := range realTradeWinInfo {
		rwa = Operation.HPAdd(rwa, Operation.HPSub(x.Outputamount, x.Inputamount))
	}

	for _, x := range realTradeLoseInfo {
		rla = Operation.HPAdd(rla, x.Inputamount)
	}

	fmt.Printf("in:%+v,out:%+v\n",Operation.HPSum(rwa, da, ta),Operation.HPAdd(rla, wa))

}
