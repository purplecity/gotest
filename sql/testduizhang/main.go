package main

import (
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"time"
)

/*
func main() {
 handletime := 1571515200
 rs := []mysql.Reconciliation{}
 mysql.GetAllRecord("Reconciliation", map[string]interface{}{"Handletime":handletime},&rs)
 for _, x := range rs {
	xs := mysql.GetCountByCond("Reconciliation",map[string]interface{}{"Uid":x.Uid})
	if xs != 3 {
		fmt.Println(x.Uid,xs)
	}

 }

}

 */

/*
func main() {
	uid := "1185835680218832896"
	now := time.Now()
	//et := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, now.Location()).Unix()
	//last := now.Add(time.Hour * -24)
	//bt := time.Date(last.Year(), last.Month(), last.Day(), 3, 0, 0, 0, last.Location()).Unix()
	//realTradeInfo := []mysql.Realtrade{}
	//mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": uid,  "Settletime__lt": now.Unix()}, &realTradeInfo)

	wa := float64(0)
	da := float64(0)
	ta := float64(0)
	rwa := float64(0)
	rla := float64(0)

	withdrawInfo := []mysql.Withdrawrecord{}
	mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid": uid, "Createtime__lt": now.Unix() * 1000, "Status__in": []int{0, 1}}, &withdrawInfo)
	depositInfo := []mysql.Depositrecord{}
	mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid": uid, "Finishtime__lt": now.Unix()  * 1000, "Status": 1}, &depositInfo)
	takeScoreInfo := []mysql.Takescorerecord{}
	mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid": uid, "Handletime__lt": now.Unix() }, &takeScoreInfo)
	realTradeWinInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": uid, "Settletime__lt": now.Unix() , "Orderresult": 3}, &realTradeWinInfo)
	realTradeLoseInfo := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": uid, "Settletime__lt": now.Unix() , "Orderresult": 2}, &realTradeLoseInfo)

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

	fmt.Printf("%+v,%+v\n",len(realTradeWinInfo),len(realTradeLoseInfo))
	fmt.Printf("%+v,%+v,%+v,%+v,%+v\n",wa,da,ta,rwa,rla)
}
*/


/*
mysql.GetOneRecord("Reconciliation", map[string]interface{}{"Uid":uid},&reconcileInfo)
mysql.UpdateByCond("Reconciliation", map[string]interface{}{"Uid":uid}, map[string]interface{}{
	"Balance":assetInfo.Balance,
	"Win":Operation.HPAdd(reconcileInfo.Win,rwa),
	"Lose":Operation.HPAdd(reconcileInfo.Lose,rla),
	"Deposit":Operation.HPAdd(reconcileInfo.Deposit,da),
	"Withdraw":Operation.HPAdd(reconcileInfo.Withdraw,wa),
	"Score":Operation.HPAdd(reconcileInfo.Score,ta),
	"Handletime":now.Unix(),
})
*/





func main() {

	now := time.Now()
	et := time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, now.Location()).Unix()
	last := now.Add(time.Hour * -24)
	bt := time.Date(last.Year(), last.Month(), last.Day(), 4, 0, 0, 0, last.Location()).Unix()

	userInfo := []mysql.AdminUsers{}
	mysql.GetAllRecord("AdminUsers", map[string]interface{}{"type__in": []string{"Director", "Player", "Parter"}}, &userInfo)
	for _, x := range userInfo {
		assetInfo := mysql.Asset{}
		mysql.GetOneRecord("Asset", map[string]interface{}{"Uid": x.Uid}, &assetInfo)


		withdrawInfo2 := []mysql.Withdrawrecord{}
		mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid": x.Uid,"Createtime__lt": et * 1000,  "Status__in": []int{0, 1}}, &withdrawInfo2)
		depositInfo2 := []mysql.Depositrecord{}
		mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid": x.Uid, "Finishtime__lt": et * 1000, "Status": 1}, &depositInfo2)
		takeScoreInfo2 := []mysql.Takescorerecord{}
		mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid": x.Uid, "Handletime__lt": et}, &takeScoreInfo2)
		realTradeWinInfo2 := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": x.Uid, "Settletime__lt": et, "Orderresult": 3}, &realTradeWinInfo2)
		realTradeLoseInfo2 := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": x.Uid,  "Settletime__lt": et, "Orderresult": 2}, &realTradeLoseInfo2)

		wa2 := float64(0)
		da2 := float64(0)
		ta2 := float64(0)
		rwa2 := float64(0)
		rla2 := float64(0)
		for _, x := range withdrawInfo2 {
			wa2 = Operation.HPAdd(wa2, x.Amount)
		}

		for _, x := range depositInfo2 {
			da2 = Operation.HPAdd(da2, x.RealAmount)
		}
		for _, x := range takeScoreInfo2 {
			ta2 = Operation.HPAdd(ta2, x.Amount)
		}
		for _, x := range realTradeWinInfo2 {
			rwa2 = Operation.HPAdd(rwa2, Operation.HPSub(x.Outputamount, x.Inputamount))
		}

		for _, x := range realTradeLoseInfo2 {
			rla2 = Operation.HPAdd(rla2, x.Inputamount)
		}

		a2 :=  Operation.HPSub(Operation.HPSum(rwa2, da2, ta2), Operation.HPAdd(rla2, wa2))


		wa := float64(0)
		da := float64(0)
		ta := float64(0)
		rwa := float64(0)
		rla := float64(0)


		withdrawInfo := []mysql.Withdrawrecord{}
		mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid": x.Uid, "Createtime__gte": bt * 1000, "Createtime__lt": et * 1000, "Status__in": []int{0, 1}}, &withdrawInfo)
		depositInfo := []mysql.Depositrecord{}
		mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid": x.Uid, "Finishtime__gte": bt * 1000, "Finishtime__lt": et * 1000, "Status": 1},  &depositInfo)
		takeScoreInfo := []mysql.Takescorerecord{}
		mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid": x.Uid, "Handletime__gte": bt, "Handletime__lt": et},&takeScoreInfo)
		realTradeWinInfo := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": x.Uid, "Settletime__gt": bt, "Settletime__lt": et, "Orderresult": 3},  &realTradeWinInfo)
		realTradeLoseInfo := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": x.Uid, "Settletime__gt": bt, "Settletime__lt": et, "Orderresult": 2},  &realTradeLoseInfo)

		//fmt.Printf("%+v\n",withdrawInfo)
		//fmt.Printf("%+v\n",depositInfo)
		//fmt.Printf("%+v\n",takeScoreInfo)
		//fmt.Printf("%+v\n",realTradeWinInfo)
		//fmt.Printf("%+v\n",realTradeLoseInfo)
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

		rt := []mysql.Hpreconciliation{}
		mysql.GetTopByCondStruct("Hpreconciliation","Handletime",1,map[string]interface{}{"Uid":x.Uid},&rt)
		b := float64(0)
		if len(rt) != 0 {
			b = rt[0].Balance
		}
		mysql.AddOneRecord(&mysql.Hpreconciliation{
			Uid:         x.Uid,
			Balance:     a2,
			Lastbalance: b,
			Win:         rwa,
			Lose:        rla,
			Deposit:     da,
			Withdraw:    wa,
			Score:       ta,
			Handletime:  et,
		})

	}
}

		/*
		withdrawInfo2 := []mysql.Withdrawrecord{}
		mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid": x.Uid,"Createtime__gte": bt1 * 1000,  "Status__in": []int{0, 1}}, &withdrawInfo2)
		depositInfo2 := []mysql.Depositrecord{}
		mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid": x.Uid, "Finishtime__gte": bt1 * 1000, "Status": 1}, &depositInfo2)
		takeScoreInfo2 := []mysql.Takescorerecord{}
		mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid": x.Uid, "Handletime__gte": bt1}, &takeScoreInfo2)
		realTradeWinInfo2 := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": x.Uid, "Settletime__gte": bt1, "Orderresult": 3}, &realTradeWinInfo2)
		realTradeLoseInfo2 := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": x.Uid,  "Settletime__gte": bt1, "Orderresult": 2}, &realTradeLoseInfo2)

		wa2 := float64(0)
		da2 := float64(0)
		ta2 := float64(0)
		rwa2 := float64(0)
		rla2 := float64(0)
		for _, x := range withdrawInfo2 {
			wa2 = Operation.HPAdd(wa2, x.Amount)
		}

		for _, x := range depositInfo2 {
			da2 = Operation.HPAdd(da2, x.RealAmount)
		}
		for _, x := range takeScoreInfo2 {
			ta2 = Operation.HPAdd(ta2, x.Amount)
		}
		for _, x := range realTradeWinInfo2 {
			rwa2 = Operation.HPAdd(rwa2, Operation.HPSub(x.Outputamount, x.Inputamount))
		}

		for _, x := range realTradeLoseInfo2 {
			rla2 = Operation.HPAdd(rla2, x.Inputamount)
		}

		a2 :=  Operation.HPSub(Operation.HPSum(rwa2, da2, ta2), Operation.HPAdd(rla2, wa2))
		if !Operation.HPEqual(Operation.HPAdd(a1,a2),assetInfo.Balance) {
			fmt.Println("%+v,%+v,%+v\n",assetInfo.Uid,a1,a2)
			break
		}

		 */



	/*
	uid := "1177608163068293120"
	Amount := float64(222.5)
	now := time.Now()
	bt1 := time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, now.Location()).Unix()
	//last := now.Add(time.Hour * -24)
	//bt := time.Date(last.Year(), last.Month(), last.Day(), 4, 0, 0, 0, last.Location()).Unix()
	llast := now.Add(time.Hour * -24 )
	et := time.Date(llast.Year(), llast.Month(), llast.Day(), 4, 0, 0, 0, llast.Location()).Unix()


	//win := float64(0)
	//lose := float64(0)


	for _, x := range realTradeInfo {

		Amount = Operation.HPAdd(Amount, Operation.HPSub(x.Outputamount, x.Inputamount))
		fmt.Printf("%+v,%+v\n", Amount, x.Settletime)
		if x.Handletime > 1571305459 && x.Handletime < et {
			fmt.Println("testtest")
		}
	}
	*/
	//fmt.Printf("%+v,%+v\n",win,lose)