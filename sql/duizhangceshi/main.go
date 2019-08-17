package main

import (
	"HPOptionServer/Common/CommonConf"
	"fmt"
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"time"
)

var (
	Uid = "1162186916003659776"
)
func main() {
	now := time.Now()
	et := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, now.Location()).Unix()
	last := now.Add(time.Hour * -24)
	bt := time.Date(last.Year(), last.Month(), last.Day(), 3, 0, 0, 0, last.Location()).Unix()
	userInfo := mysql.AdminUsers{}
	mysql.GetOneRecord(CommonConf.TableUser, map[string]interface{}{"Uid":Uid},&userInfo)
	wa := float64(0)
	da := float64(0)
	ta := float64(0)
	rwa := float64(0)
	rla := float64(0)

	withdrawInfo := []mysql.Withdrawrecord{}
	mysql.GetAllRecord(CommonConf.TableWithdrawrecord, map[string]interface{}{"Uid":Uid,"Createtime__gte":bt*1000,"Createtime__lt":et*1000},&withdrawInfo)
	depositInfo := []mysql.Depositrecord{}
	mysql.GetAllRecord(CommonConf.TableDepositRecord, map[string]interface{}{"Uid":Uid,"Finishtime__gte":bt*1000,"Finishtime__lt":et*1000,"Status":CommonConf.DepositSuccess},&depositInfo)
	takeScoreInfo := []mysql.Takescorerecord{}
	mysql.GetAllRecord(CommonConf.TableTakescorerecord, map[string]interface{}{"Uid":Uid,"Handletime__gte":bt,"Handletime__lt":et},&takeScoreInfo)
	realTradeWinInfo := []mysql.Realtrade{}
	mysql.GetAllRecord(CommonConf.TableRealTrade, map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et,"Orderresult":CommonConf.HP_win},&realTradeWinInfo)
	realTradeLoseInfo := []mysql.Realtrade{}
	mysql.GetAllRecord(CommonConf.TableRealTrade, map[string]interface{}{"Uid":Uid,"Settletime__gt":bt,"Settletime__lt":et,"Orderresult":CommonConf.HP_lose},&realTradeLoseInfo)
	for _,x := range withdrawInfo {
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

	for _,x := range realTradeLoseInfo {
		rla = Operation.HPAdd(rla,x.Inputamount)
	}

	assetInfo := mysql.Asset{}
	mysql.GetAllRecord(CommonConf.TableAsset, map[string]interface{}{"Uid":Uid},&assetInfo)
	fmt.Printf("%+v,%+v,%+v,%+v,%+v\n",rwa,da,ta,rla,wa)
	fmt.Println(Operation.HPSub(Operation.HPSum(rwa,da,ta),Operation.HPSum(rla,wa)))
	fmt.Println(assetInfo.Balance,assetInfo.Freezebalance)

}


