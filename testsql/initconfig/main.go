package main

import (
	"gotest/testsql/mysql"
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



	version = "0.7.0"
	isforce = 1

)


func main() {
	mysql.AddOneRecord(&mysql.Subject{Symbol:btcsy,Type:btctype,Isopen:btcisopen,Pisopen:btcisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurusdsy,Type:eurusdtype,Isopen:eurusdisopen,Pisopen:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurjpysy,Type:eurjpytype,Isopen:eurjpyisopen,Pisopen:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:usdjpysy,Type:usdjpytype,Isopen:usdjpyisopen,Pisopen:usdjpyisopen})

	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:btcsy,Type:btctype,Udisopen:btcisopen,Sdpisopen:btcisopen,Symutex:btcisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:eurusdsy,Type:eurusdtype,Udisopen:eurusdisopen,Sdpisopen:eurusdisopen,Symutex:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:eurjpysy,Type:eurjpytype,Udisopen:eurjpyisopen,Sdpisopen:eurjpyisopen,Symutex:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:usdjpysy,Type:usdjpytype,Udisopen:usdjpyisopen,Sdpisopen:usdjpyisopen,Symutex:usdjpyisopen})

	mysql.AddOneRecord(&mysql.Clientversion{Version:version,Isforce:isforce,Createtime:time.Now().Unix()})


	mysql.AddOneRecord(&mysql.Depositway{Way:1,Isopen:1})
	mysql.AddOneRecord(&mysql.Depositway{Way:2,Isopen:0})
	mysql.AddOneRecord(&mysql.Depositway{Way:3,Isopen:1})
}