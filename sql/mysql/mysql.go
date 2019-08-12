package  mysql

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MysqlUserName = "root"
	MysqlPassWord = "k"
	//MysqlPassWord = "HP@123"
	//MysqlIP = "47.244.212.51"
	MysqlIP = "127.0.0.1"
	MysqlPort = 3306
	MysqlDefaultDatabase = "test2"
	//MysqlDefaultDatabase = "test3"
)

func init() {

	_ = orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&allowNativePasswords=true",
			MysqlUserName, MysqlPassWord, MysqlIP, MysqlPort, MysqlDefaultDatabase))
	//注册模型
	orm.RegisterModel(new(AdminUsers),new(Realtrade),new(Vitualtrade),
		new(Asset),new(Parter),new(Director),new(Player),new(Score),
		new(Scorerecord),new(Depositrecord),new(Withdrawrecord),
		new(BankInfo),new(Subject),new(DepositEnsureRecord),new(Clientversion),
		new(AdminRoleUsers),new(Lastconnect),new(Depositbank),new(Odds),new(OddsInfo),
		new(Takescorerecord),new(Reconciliation))
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

func UpdateByCond(table string,cond,updateMap map[string]interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.Update(orm.Params(updateMap))
}