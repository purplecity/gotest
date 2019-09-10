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
	btcFirstopenhour = 9
	btcFirstopenmin = 30
	btcFirstclosehour = 2
	btcFirstclosemin = 30

	shcisy = "SHCI"
	shcitype = "Stock"
	shciisopen = 1
	shciFirstopenhour = 9
	shciFirstopenmin = 30
	shciFirstclosehour = 11
	shciFirstclosemin = 30
	shciSecondopenhour = 13
	shciSecondopenmin = 0
	shciSecondclosehour = 15
	shciSecondclosemin = 0


	szcisy = "SZCI"
	szcitype = "Stock"
	szciisopen = 1
	szciFirstopenhour = 9
	szciFirstopenmin = 30
	szciFirstclosehour = 11
	szciFirstclosemin = 30
	szciSecondopenhour = 13
	szciSecondopenmin = 0
	szciSecondclosehour = 15
	szciSecondclosemin = 0

	version = "0.4.5"
	isforce = 1

	btcoddssy = "BTC"
	btcoddsUpodds = 0.9
	btcoddsDownodds = 0.9

	shcioddssy = "SHCI"
	shcioddsUpodds = 0.9
	shcioddsDownodds = 0.9

	szcioddssy = "SZCI"
	szcioddsUpodds = 0.9
	szcioddsDownodds = 0.9


	btcioisy = "BTC"
	btcoiLevelOne = 1
	btcoiOneMindv = float64(0)
	btcoiOneMaxdv = float64(50000)
	btcoiOneGreaterodds = 0.9
	btcoiOneLessodds = 0.9

	btcoiLevelTwo = 2
	btcoiTwoMindv = float64(50000)
	btcoiTwoMaxdv = float64(100000)
	btcoiTwoGreaterodds = 1.2
	btcoiTwoLessodds = 0.6

	btcoiLevelThree = 3
	btcoiThreeMindv = float64(100000)
	btcoiThreeMaxdv = float64(200000)
	btcoiThreeGreaterodds = 1.7
	btcoiThreeLessodds = 0.1

	btcoiLevelFour = 4
	btcoiFourMindv = float64(200000)
	btcoiFourGreaterodds = 1.8
	btcoiFourLessodds = float64(0)

	shcioisy = "SHCI"
	shcioiLevelOne = 1
	shcioiOneMindv = float64(0)
	shcioiOneMaxdv = float64(50000)
	shcioiOneGreaterodds = 0.9
	shcioiOneLessodds = 0.9

	shcioiLevelTwo = 2
	shcioiTwoMindv = float64(50000)
	shcioiTwoMaxdv = float64(100000)
	shcioiTwoGreaterodds = 1.2
	shcioiTwoLessodds = 0.6

	shcioiLevelThree = 3
	shcioiThreeMindv = float64(100000)
	shcioiThreeMaxdv = float64(200000)
	shcioiThreeGreaterodds = 1.7
	shcioiThreeLessodds = 0.1

	shcioiLevelFour = 4
	shcioiFourMindv = float64(200000)
	shcioiFourGreaterodds = 1.8
	shcioiFourLessodds = float64(0)

	szcioisy = "SZCI"
	szcioiLevelOne = 1
	szcioiOneMindv = float64(0)
	szcioiOneMaxdv = float64(50000)
	szcioiOneGreaterodds = 0.9
	szcioiOneLessodds = 0.9

	szcioiLevelTwo = 2
	szcioiTwoMindv = float64(50000)
	szcioiTwoMaxdv = float64(100000)
	szcioiTwoGreaterodds = 1.2
	szcioiTwoLessodds = 0.6

	szcioiLevelThree = 3
	szcioiThreeMindv = float64(100000)
	szcioiThreeMaxdv = float64(200000)
	szcioiThreeGreaterodds = 1.7
	szcioiThreeLessodds = 0.1

	szcioiLevelFour = 4
	szcioiFourMindv = float64(200000)
	szcioiFourGreaterodds = 1.8
	szcioiFourLessodds = float64(0)
)


func main() {


	/*
	mysql.AddOneRecord(&mysql.Subject{Symbol:btcsy,Type:btctype,Isopen:btcisopen,Firstopenhour:btcFirstopenhour,
		Firstopenmin:btcFirstopenmin,Firstclosehour:btcFirstclosehour,Firstclosemin:btcFirstclosemin})
	mysql.AddOneRecord(&mysql.Subject{Symbol:shcisy,Type:shcitype,Isopen:shciisopen,Firstopenhour:shciFirstopenhour,
		Firstopenmin:shciFirstopenmin,Firstclosehour:shciFirstclosehour,Firstclosemin:shciFirstclosemin,
		Secondopenhour:shciSecondopenhour,Secondopenmin:shciSecondopenmin,Secondclosehour:shciSecondclosehour,
	Secondclosemin:shciSecondclosemin})
	mysql.AddOneRecord(&mysql.Subject{Symbol:szcisy,Type:szcitype,Isopen:szciisopen,Firstopenhour:szciFirstopenhour,
		Firstopenmin:szciFirstopenmin,Firstclosehour:szciFirstclosehour,Firstclosemin:szciFirstclosemin,
		Secondopenhour:szciSecondopenhour,Secondopenmin:szciSecondopenmin,Secondclosehour:szciSecondclosehour,
		Secondclosemin:szciSecondclosemin})

	 */

	mysql.AddOneRecord(&mysql.Clientversion{Version:version,Isforce:isforce,Createtime:time.Now().Unix()})

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

	mysql.AddOneRecord(&mysql.Depositway{Way:1,Isopen:1})
	mysql.AddOneRecord(&mysql.Depositway{Way:2,Isopen:0})
	mysql.AddOneRecord(&mysql.Depositway{Way:3,Isopen:1})

	mysql.AddOneRecord(&mysql.Payamount{Payway:1,One:156,Two:298,Three:498,Four:998,Five:2098,Six:4908})

	//mysql.UpdateByCond("AdminUsers", map[string]interface{}{}, map[string]interface{}{"Valid":1})
}