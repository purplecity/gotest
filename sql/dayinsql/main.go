package main

import (
	"fmt"
	"gotest/sql/mysql"
)

func main() {
	/*
	now := time.Now()
	et := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, now.Location()).Unix()
	last := now.Add(time.Hour * -24)
	bt := time.Date(last.Year(), last.Month(), last.Day(), 3, 0, 0, 0, last.Location()).Unix()
	withdrawInfo := []mysql.Withdrawrecord{}
	Uid := "uid3333333333"
	mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid":Uid,"Createtime__gte":bt*1000,"Createtime__lt":et*1000,"Status__in":[]int{0,1}},&withdrawInfo)
	depositInfo := []mysql.Depositrecord{}
	mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid":Uid,"Finishtime__gte":bt*1000,"Finishtime__lt":et*1000,"Status":1},&depositInfo)
	takeScoreInfo := []mysql.Takescorerecord{}
	mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid":Uid,"Handletime__gte":bt,"Handletime__lt":et},&takeScoreInfo)
	realTradeWinInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et,"Orderresult":3},&realTradeWinInfo)
	realTradeLoseInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et,"Orderresult":2},&realTradeLoseInfo)


	Uid := "1164438374518382592"
	bt := 1566500400
	et := 1566533225
	withdrawInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et},&withdrawInfo)
	out := float64(0)
	in := float64(0)
	for _, x := range withdrawInfo {
		out = Operation.HPAdd(out,x.Outputamount)
		in = Operation.HPAdd(in,x.Inputamount)
	}
	fmt.Println(out,in)

	fmt.Println(11150-15910 + 3000 - 28)
	*/
	zhangyunUidList := []string{"1164766833442557952","1164767299832389632","1164767465893277696",
		"1164767636932804608","1164767724895752192","1164767931939184640","1164768087195545600",
		"1164768165499011072","1164768240447033344","1164768329466945536",
	}
	assetInfo := []mysql.Asset{}
	mysql.GetAllRecord("Asset", map[string]interface{}{"Uid__in":zhangyunUidList},&assetInfo)
	for _, x := range assetInfo {
		fmt.Printf("%+v\n", x)
	}
}
