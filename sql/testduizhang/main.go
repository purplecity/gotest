package main

import (
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"time"
)

var reconcile = func() error {
	assetInfo := []mysql.Asset{}
	mysql.GetAllRecord("Asset", map[string]interface{}{},&assetInfo)
	now := time.Now()
	et := time.Date(now.Year(), now.Month(), now.Day(), 2, 30, 0, 0, now.Location()).Unix()
	last := now.Add(time.Hour * -24)
	bt := time.Date(last.Year(), last.Month(), last.Day(), 2, 30, 0, 0, last.Location()).Unix()
	for _, v := range assetInfo {
		userInfo := mysql.AdminUsers{}
		mysql.GetOneRecord("AdminUsers", map[string]interface{}{"Uid":v.Uid},&userInfo)
		wa := float64(0)
		da := float64(0)
		ta := float64(0)
		rwa := float64(0)
		rla := float64(0)
		if userInfo.Valid == 1 {
			withdrawFinInfo := []mysql.Withdrawrecord{}
			mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid":v.Uid,"Createtime__gte":bt*1000,"Finishtime__gte":bt*1000,"Finishtime__lt":et*1000,"Status":1},&withdrawFinInfo)
			withdrawUnFinInfo := []mysql.Withdrawrecord{}
			mysql.GetAllRecord("Withdrawrecord", map[string]interface{}{"Uid":v.Uid,"Createtime__gte":bt*1000,"Createtime__lt":et*1000,"Status":0},&withdrawUnFinInfo)
			depositInfo := []mysql.Depositrecord{}
			mysql.GetAllRecord("Depositrecord", map[string]interface{}{"Uid":v.Uid,"Finishtime__gte":bt*1000,"Finishtime__lt":et*1000,"Status":1},&depositInfo)
			takeScoreInfo := []mysql.Takescorerecord{}
			mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid":v.Uid,"Handletime__gte":bt*1000,"Handletime__lt":et*1000},&takeScoreInfo)
			realTradeWinInfo := []mysql.Realtrade{}
			mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":v.Uid,"Settletime__gte":bt*1000,"Settletime__lt":et*1000,"Orderresult":1},&realTradeWinInfo)
			realTradeLoseInfo := []mysql.Realtrade{}
			mysql.GetAllRecord("RealTrade", map[string]interface{}{"Uid":v.Uid,"Settletime__gte":bt*1000,"Settletime__lt":et*1000,"Orderresult":0},&realTradeLoseInfo)
			for _,x := range withdrawFinInfo {
				wa = Operation.HPAdd(wa,x.Amount)
			}

			for _,x := range withdrawUnFinInfo {
				wa = Operation.HPAdd(wa,x.Amount)
			}

			for _, x := range depositInfo {
				da = Operation.HPAdd(da,x.Amount)
			}
			for _,x := range takeScoreInfo {
				ta = Operation.HPAdd(ta,x.Amount)
			}
			for _,x := range realTradeWinInfo {
				rwa = Operation.HPAdd(rwa,Operation.HPSub(x.Outputamount,x.Inputamount))
			}

			for _,x := range realTradeWinInfo {
				rla = Operation.HPAdd(rla,x.Inputamount)
			}
		}
		reconcileInfo := mysql.Reconciliation{}
		mysql.GetOneRecord("Reconciliation", map[string]interface{}{"Uid":v.Uid},&reconcileInfo)
		mysql.UpdateByCond("Reconciliation", map[string]interface{}{"Uid":v.Uid}, map[string]interface{}{
			"Balance":v.Balance,
			"Win":Operation.HPAdd(reconcileInfo.Win,rwa),
			"Lose":Operation.HPAdd(reconcileInfo.Lose,rla),
			"Deposit":Operation.HPAdd(reconcileInfo.Deposit,da),
			"Withdraw":Operation.HPAdd(reconcileInfo.Withdraw,wa),
			"Score":Operation.HPAdd(reconcileInfo.Score,ta),
			"Handletime":now.Unix(),
		})
		if !Operation.HPlte(Operation.HPSub(v.Balance,reconcileInfo.Balance),Operation.HPSub(Operation.HPSum(rwa,da,ta),Operation.HPSum(rla,wa))) {
			mysql.UpdateByCond("AdminUsers",map[string]interface{}{"Uid":v.Uid}, map[string]interface{}{"Valid":0})
		}
	}
	return nil
}
