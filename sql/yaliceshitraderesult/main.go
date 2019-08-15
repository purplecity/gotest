package main

import (
	"fmt"
	"gotest/sql/Operation"
	"gotest/sql/mysql"
)

func main() {

	userlist := []mysql.AdminUsers{}
	mysql.GetAllRecord("AdminUsers", map[string]interface{}{"Phonenumber__startswith":"2"},&userlist)
	for _,x := range userlist {

		tradeInfo := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":x.Uid},&tradeInfo)
		in := float64(0)
		out := float64(0)
		for _,x := range tradeInfo {
			in = Operation.HPAdd(in,x.Inputamount)
			out = Operation.HPAdd(out,x.Outputamount)
		}
		assetInfo := mysql.Asset{}
		mysql.GetOneRecord("Asset", map[string]interface{}{"Uid":x.Uid},&assetInfo)
		if !Operation.HPEqual(assetInfo.Freezebalance,float64(0)) || !Operation.HPEqual(Operation.HPSub(assetInfo.Balance,float64(5000)),Operation.HPSub(out,in)) {
			fmt.Printf("asset %+v\n",x.Uid)
			break
		}

		tradeInfo2 := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":x.Uid,"Orderresult":1},&tradeInfo)
		if len(tradeInfo2) != 0 {
			for _, x := range tradeInfo2 {
				fmt.Printf("trade %+v\n",x)
			}
		}

		scInfo := []mysql.Scorerecord{}
		mysql.GetAllRecord("Scorerecord", map[string]interface{}{"Playerid":x.Uid},&scInfo)
		if len(scInfo) != len(tradeInfo) {
			fmt.Printf("score %+v\n",x.Uid)
			break
		}

	}
}
