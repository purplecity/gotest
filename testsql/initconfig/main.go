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



	version = "0.7.4"
	isforce = 1

)


func main() {


	mysql.AddOneRecord(&mysql.Subject{Symbol:btcsy,Type:btctype,Isopen:btcisopen,Pisopen:btcisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:"SHCI",Type:"Stock",Isopen:1,Pisopen:1})
	mysql.AddOneRecord(&mysql.Subject{Symbol:"SZCI",Type:"Stock",Isopen:1,Pisopen:1})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurusdsy,Type:eurusdtype,Isopen:eurusdisopen,Pisopen:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:eurjpysy,Type:eurjpytype,Isopen:eurjpyisopen,Pisopen:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subject{Symbol:usdjpysy,Type:usdjpytype,Isopen:usdjpyisopen,Pisopen:usdjpyisopen})

	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:btcsy,Type:btctype,Udisopen:btcisopen,Sdpisopen:btcisopen,Symutex:btcisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:"SHCI",Type:"Stock",Udisopen:1,Sdpisopen:1,Symutex:1})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:"SZCI",Type:"Stock",Udisopen:1,Sdpisopen:1,Symutex:1})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:eurusdsy,Type:eurusdtype,Udisopen:eurusdisopen,Sdpisopen:eurusdisopen,Symutex:eurusdisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:eurjpysy,Type:eurjpytype,Udisopen:eurjpyisopen,Sdpisopen:eurjpyisopen,Symutex:eurjpyisopen})
	mysql.AddOneRecord(&mysql.Subjecttrade{Symbol:usdjpysy,Type:usdjpytype,Udisopen:usdjpyisopen,Sdpisopen:usdjpyisopen,Symutex:usdjpyisopen})

	mysql.AddOneRecord(&mysql.Clientversion{Version:version,Isforce:isforce,Createtime:time.Now().Unix()})

	mysql.AddOneRecord(&mysql.Depositway{Way:1,Isopen:1})
	mysql.AddOneRecord(&mysql.Depositway{Way:2,Isopen:0})
	mysql.AddOneRecord(&mysql.Depositway{Way:3,Isopen:1})

	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"deposit",Opisopen:1})
	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"withdraw",Opisopen:1})
	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"takescore",Opisopen:1})
	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"other",Opisopen:1})
	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"daygame",Opisopen:0})
	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"weekgame",Opisopen:0})
	mysql.AddOneRecord(&mysql.Gracefuldown{Op:"yeargame",Opisopen:0})




}