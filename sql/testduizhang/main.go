package main

import (
	"fmt"
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"time"
)

func main() {
	uid := "1161667228077133824"
	assetInfo := mysql.Asset{}
	mysql.GetOneRecord("Asset", map[string]interface{}{"Uid": uid}, &assetInfo)
	now := time.Now()
	et := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, now.Location()).Unix()
	last := now.Add(time.Hour * -24)
	bt := time.Date(last.Year(), last.Month(), last.Day(), 3, 0, 0, 0, last.Location()).Unix()
	userInfo := mysql.AdminUsers{}
	mysql.GetOneRecord("AdminUsers", map[string]interface{}{"Uid": uid}, &userInfo)
	wa := float64(0)
	da := float64(0)
	ta := float64(0)
	rwa := float64(0)
	rla := float64(0)
	if userInfo.Valid == 1 {
		withdrawInfo := []mysql.Withdrawrecord{}
		mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid": uid, "Createtime__gte": bt * 1000, "Createtime__lt": et * 1000}, &withdrawInfo)
		depositInfo := []mysql.Depositrecord{}
		mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid": uid, "Finishtime__gte": bt * 1000, "Finishtime__lt": et * 1000, "Status": 1}, &depositInfo)
		takeScoreInfo := []mysql.Takescorerecord{}
		mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid": uid, "Handletime__gte": bt, "Handletime__lt": et}, &takeScoreInfo)
		realTradeWinInfo := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": uid, "Settletime__gte": bt, "Settletime__lt": et, "Orderresult": 3}, &realTradeWinInfo)
		realTradeLoseInfo := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid": uid, "Settletime__gte": bt, "Settletime__lt": et, "Orderresult": 2}, &realTradeLoseInfo)
		for _, x := range withdrawInfo {
			wa = Operation.HPAdd(wa, x.Amount)
		}

		for _, x := range depositInfo {
			da = Operation.HPAdd(da, x.Amount)
		}
		for _, x := range takeScoreInfo {
			ta = Operation.HPAdd(ta, x.Amount)
		}
		for _, x := range realTradeWinInfo {
			rwa = Operation.HPAdd(rwa, Operation.HPSub(x.Outputamount, x.Inputamount))
		}

		for _, x := range realTradeWinInfo {
			rla = Operation.HPAdd(rla, x.Inputamount)
		}
	}
	reconcileInfo := mysql.Reconciliation{}
	mysql.GetOneRecord("Reconciliation", map[string]interface{}{"Uid":uid},&reconcileInfo)
	a := Operation.HPSub(assetInfo.Balance, reconcileInfo.Balance)
	b := Operation.HPSub(Operation.HPSum(rwa, da, ta), Operation.HPSum(rla, wa))
	fmt.Println(a,b)
	if !Operation.HPlte(a,b) {
		fmt.Println("err")
	}
}
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

