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

	eurusdsy = "EURUSD"
	eurusdtype = "Forex"
	eurusdisopen = 1


	eurjpysy = "EURJPY"
	eurjpytype = "Forex"
	eurjpyisopen = 1

	usdjpysy = "USDJPY"
	usdjpytype = "Forex"
	usdjpyisopen = 1



	version = "0.6.8"
	isforce = 1

)


func main() {

	mysql.AddOneRecord(&mysql.Subject{Symbol:btcsy,Type:btctype,Isopen:btcisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurusdsy,Type:eurusdtype,Isopen:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurjpysy,Type:eurjpytype,Isopen:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:usdjpysy,Type:usdjpytype,Isopen:usdjpyisopen})

	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:eurusdsy,Type:eurusdtype,Udisopen:eurusdisopen,Sdpisopen:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:eurjpysy,Type:eurjpytype,Udisopen:eurjpyisopen,Sdpisopen:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:usdjpysy,Type:usdjpytype,Udisopen:usdjpyisopen,Sdpisopen:usdjpyisopen})

	mysql.AddOneRecord(&mysql.Clientversion{Version:version,Isforce:isforce,Createtime:time.Now().Unix()})


	mysql.AddOneRecord(&mysql.Depositway{Way:1,Isopen:1})
	mysql.AddOneRecord(&mysql.Depositway{Way:2,Isopen:0})
	mysql.AddOneRecord(&mysql.Depositway{Way:3,Isopen:1})
}