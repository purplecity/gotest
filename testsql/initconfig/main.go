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

)


func main() {

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

	mysql.AddOneRecord(&mysql.Clientversion{Version:version,Isforce:isforce,Createtime:time.Now().Unix()})


	mysql.AddOneRecord(&mysql.Depositway{Way:1,Isopen:1})
	mysql.AddOneRecord(&mysql.Depositway{Way:2,Isopen:0})
	mysql.AddOneRecord(&mysql.Depositway{Way:3,Isopen:1})

	mysql.AddOneRecord(&mysql.Payamount{Payway:1,One:156,Two:298,Three:498,Four:998,Five:2098,Six:4908})

	//mysql.UpdateByCond("AdminUsers", map[string]interface{}{}, map[string]interface{}{"Valid":1})
}