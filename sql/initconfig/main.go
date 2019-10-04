package main

import (
	"gotest/sql/mysql"
	"time"
)

//启动shci和szci之前要先插入数据  所以这里要modle
//邀请码要自己操作
//subject clientversion odds oddsInfo

var (
	btcsy = "BTC"
	btctype = "CryptoCurrency"
	btcisopen = 1


	eurusd = "EURUSD"
	eurusdtype = "Forex"
	eurusdisopen = 1

	eurjpy = "EURJPY"
	eurjpytype = "Forex"
	eurjpyisopen = 1

	usdjpy = "USDJPY"
	usdjpytype = "Forex"
	usdjpyisopen = 1


	version = "0.6.3"
	isforce = 1





)


func main() {



	mysql.AddOneRecord(&mysql.Subject{Symbol:btcsy,Type:btctype,Isopen:btcisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurusd,Type:eurusdtype,Isopen:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurjpy,Type:eurjpytype,Isopen:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:usdjpy,Type:usdjpytype,Isopen:usdjpyisopen})
	//mysql.AddOneRecord(&mysql.Subject{Symbol:shcisy,Type:shcitype,Isopen:shciisopen})
	//mysql.AddOneRecord(&mysql.Subject{Symbol:szcisy,Type:szcitype,Isopen:szciisopen})

	mysql.AddOneRecord(&mysql.Clientversion{Version:version,Isforce:isforce,Createtime:time.Now().Unix()})

	/*
	mysql.AddOneRecord(&mysql.Odds{Symbol:btcoddssy,Upodds:btcoddsUpodds,Downodds:btcoddsDownodds})
	mysql.AddOneRecord(&mysql.Odds{Symbol:shcioddssy,Upodds:shcioddsUpodds,Downodds:shcioddsDownodds})
	mysql.AddOneRecord(&mysql.Odds{Symbol:szcioddssy,Upodds:szcioddsUpodds,Downodds:szcioddsDownodds})

	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:btcioisy,Level:btcoiLevelOne,Mindv:btcoiOneMindv,Maxdv:btcoiOneMaxdv,
		Greaterodds:btcoiOneGreaterodds,Lessodds:btcoiOneLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:btcioisy,Level:btcoiLevelTwo,Mindv:btcoiTwoMindv,Maxdv:btcoiTwoMaxdv,
		Greaterodds:btcoiTwoGreaterodds,Lessodds:btcoiTwoLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:btcioisy,Level:btcoiLevelThree,Mindv:btcoiThreeMindv,Maxdv:btcoiThreeMaxdv,
		Greaterodds:btcoiThreeGreaterodds,Lessodds:btcoiThreeLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:btcioisy,Level:btcoiLevelFour,Mindv:btcoiFourMindv,
		Greaterodds:btcoiFourGreaterodds,Lessodds:btcoiFourLessodds})


	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:shcioisy,Level:shcioiLevelOne,Mindv:shcioiOneMindv,Maxdv:shcioiOneMaxdv,
		Greaterodds:shcioiOneGreaterodds,Lessodds:shcioiOneLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:shcioisy,Level:shcioiLevelTwo,Mindv:shcioiTwoMindv,Maxdv:shcioiTwoMaxdv,
		Greaterodds:shcioiTwoGreaterodds,Lessodds:shcioiTwoLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:shcioisy,Level:shcioiLevelThree,Mindv:shcioiThreeMindv,Maxdv:shcioiThreeMaxdv,
		Greaterodds:shcioiThreeGreaterodds,Lessodds:shcioiThreeLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:shcioisy,Level:shcioiLevelFour,Mindv:shcioiFourMindv,
		Greaterodds:shcioiFourGreaterodds,Lessodds:shcioiFourLessodds})

	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:szcioisy,Level:szcioiLevelOne,Mindv:szcioiOneMindv,Maxdv:szcioiOneMaxdv,
		Greaterodds:szcioiOneGreaterodds,Lessodds:szcioiOneLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:szcioisy,Level:szcioiLevelTwo,Mindv:szcioiTwoMindv,Maxdv:szcioiTwoMaxdv,
		Greaterodds:szcioiTwoGreaterodds,Lessodds:szcioiTwoLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:szcioisy,Level:szcioiLevelThree,Mindv:szcioiThreeMindv,Maxdv:szcioiThreeMaxdv,
		Greaterodds:szcioiThreeGreaterodds,Lessodds:szcioiThreeLessodds})
	mysql.AddOneRecord(&mysql.OddsInfo{Symbol:szcioisy,Level:szcioiLevelFour,Mindv:szcioiFourMindv,
		Greaterodds:szcioiFourGreaterodds,Lessodds:szcioiFourLessodds})
	*/
	mysql.AddOneRecord(&mysql.Depositway{Way:1,Isopen:1})
	mysql.AddOneRecord(&mysql.Depositway{Way:2,Isopen:0})
	mysql.AddOneRecord(&mysql.Depositway{Way:3,Isopen:1})

	//mysql.AddOneRecord(&mysql.Payamount{Payway:1,One:156,Two:298,Three:498,Four:998,Five:2098,Six:4908})

	//mysql.UpdateByCond("AdminUsers", map[string]interface{}{}, map[string]interface{}{"Valid":1})
}