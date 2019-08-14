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

func AddMultiRecord(num int, record interface{}){
	o := getOrm()
	o.InsertMulti(num,record)
}

func Exist(table,filed string, value interface{}) bool {
	o := getOrm()
	return o.QueryTable(table).Filter(filed,value).Exist()
}

//经测试Exist会把int相关类型跟string相等 就算设置了exact也不行
func MultiExist(table string, cond map[string]interface{}) bool {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	return qs.Exist()
}

func UpdateByCond(table string,cond,updateMap map[string]interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.Update(orm.Params(updateMap))
}

func GetOneRecord(table string,cond map[string]interface{},resultStruct interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.One(resultStruct)
}

func GetAllRecord(table string,cond map[string]interface{},resultStruct interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.All(resultStruct)
}

func GetSortAllRecord(table, orderFiled string,cond map[string]interface{},resultStruct interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.OrderBy("-"+orderFiled).All(resultStruct)
}


func GetTopByCondStruct(table,orderFiled string, num int,cond map[string]interface{},resultStruct interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value).OrderBy("-"+orderFiled).Limit(num)
	}
	qs.All(resultStruct)
}


func GetRecordMap(table string,cond map[string]interface{},resultMap *[]orm.Params) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.Values(resultMap)
}

func GetSpecialSlice(table string,cond map[string]interface{},resultList *[]orm.ParamsList) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.ValuesList(resultList)
}

func GetTopMap(table,orderFiled string, num int,resultMap *[]orm.Params) {
	o := getOrm()
	qs := o.QueryTable(table)
	qs = qs.OrderBy("-"+orderFiled).Limit(num)
	qs.Values(resultMap)
}

func GetTopByCondMap(table,orderFiled string, num int,cond map[string]interface{},resultMap *[]orm.Params) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value).OrderBy("-"+orderFiled).Limit(num)
	}
	qs.Values(resultMap)
}

func GetCountByCond(table string,cond map[string]interface{}) (cnt int64) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	cnt, _ = qs.Count()
	return
}

func GetOffsetByCondStruct(table,orderFiled string, m,skip int,cond map[string]interface{},resultStruct interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)

	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	qs.OrderBy("-"+orderFiled).Limit(m,skip).All(resultStruct)
}

func GetGroupOneList(table,groupFiled,orderFiled string,resultStruct interface{}) {
	o := getOrm()
	qs := o.QueryTable(table)
	qs.GroupBy(groupFiled).OrderBy("-"+orderFiled).All(resultStruct)
