package  mysql

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MysqlUserName = "root"
	//MysqlPassWord = "k"
	MysqlPassWord = "HP@123"
	MysqlIP = "47.244.212.51"
	//MysqlIP = "127.0.0.1"
	MysqlPort = 3306
	MysqlDefaultDatabase = "test3"
)

func init() {

	_ = orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&allowNativePasswords=true",
			MysqlUserName, MysqlPassWord, MysqlIP, MysqlPort, MysqlDefaultDatabase))
	//注册模型
	orm.RegisterModel(new(Subject),new(Clientversion),
		new(Odds),new(OddsInfo))
	//自动创建表 参数二为是否drop然后创建表   参数三是否打印创建表过程
	orm.RunSyncdb("default",false,true)
}

var hpOrm orm.Ormer

func getOrm() orm.Ormer {
	if hpOrm == nil {
		hpOrm = orm.NewOrm()
	}
	return hpOrm
}

func AddOneRecord(record interface{}) {
	o := getOrm()
	o.Insert(record)
}