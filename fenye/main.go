package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var MysqlUserName = "root"
var MysqlPassWord = "HP@123"
var MysqlIP = "47.244.212.51"
var MysqlPort = 3306
var MysqlDefaultDatabase = "test7"

func init() {

	_ = orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&allowNativePasswords=true",
			MysqlUserName, MysqlPassWord, MysqlIP, MysqlPort,MysqlDefaultDatabase))
	//注册模型
	orm.RegisterModel(new(User))
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


func GetCountByCond(table string,cond map[string]interface{}) (cnt int64) {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	cnt, _ = qs.Count()
	return
}

/*
type Realtrade struct {
	Id 				int			`orm:"pk;auto"`
	Userid 			string		`orm:"index" description:"用户id"`
	Tid				string		`orm:"index" description:"订单id"`
	Handletime 		int64		`description:"下单处理时间"`
	Setteltime 		int64		`orm:"default(0)" description:"下单结算时间"`
	Inputamount 	float64		`orm:"digits(12);decimals(2)" description:"下单金额"`
	Outputamount	float64		`orm:"digits(12);decimals(2)" description:"结算金额"`
	Ordervalue		float64		`orm:"digits(12);decimals(2)" description:"下单指数"`
	Settlevalue		float64		`orm:"digits(12);decimals(2)" description:"结算指数"`
	Side   			int			`description:"看涨看跌"`
	Interval		int64		`description:"下单周期时间"`
	Symbol			string		`orm:"index" description:"标的物"`
	Orderresult		int			`orm:"index;default(1)" description:"下单结果"`
}

//虚拟账户下单记录表
type Vitualtrade struct {
	Id 				int			`orm:"pk;auto"`
	Userid 			string		`orm:"index" description:"用户id"`
	Tid				string		`orm:"index" description:"订单id"`
	Handletime 		int64		`description:"下单处理时间"`
	Setteltime 		int64		`orm:"default(0)" description:"下单结算时间"`
	Inputamount 	float64		`orm:"digits(12);decimals(2)" description:"下单金额"`
	Outputamount	float64		`orm:"digits(12);decimals(2)" description:"结算金额"`
	Ordervalue		float64		`orm:"digits(12);decimals(2)" description:"下单指数"`
	Settlevalue		float64		`orm:"digits(12);decimals(2)" description:"结算指数"`
	Side   			int			`description:"看涨看跌"`
	Interval		int64		`description:"下单周期时间"`
	Symbol			string		`orm:"index" description:"标的物"`
	Orderresult		int			`orm:"index;default(1)" description:"下单结果"`
}*/

//用户表
type User struct {
	Id 					int			`orm:"pk;auto"`
	Userid				string		`orm:"index" description:"用户id"`
	Username			string		`orm:"index" description:"用户名"`
	Phonenumber			string		`orm:"index" description:"手机号码"`
	Password			string		`description:"密码"`
	Invitationcode	string			`orm:"index" description:"邀请码"` //如果是h普通玩家填的是二级代理的邀请码 二级代理填的是大代理的邀请码 大代理的邀请码是自己
	Type 				int8		`orm:"index" description:"玩家类型"` // 0 普通玩家 1 二级代理 2 大代理
	Typename  			string		`orm:"index" description:"玩家类型所属昵称"` //普通玩家 探花  童生 进士等等
	Registtime 			int64		`description:"注册时间"`
	Lastlogintime		int64		`description:"最后登录时间"`
}

/*
//资金表
type Asset struct {
	Id 					int			`orm:"pk;auto"`
	Userid				string		`orm:"index" description:"用户id"`
	Balance				float64		`orm:"digits(12);decimals(2)" description:"实际账户余额"`
	Freezebalance		float64		`orm:"digits(12);decimals(2)" description:"实际账户冻结余额"`
	Vitualbalance 		float64     `orm:"digits(12);decimals(2)" description:"虚拟账户余额"`
	Vitualfreezebalance	float64		`orm:"digits(12);decimals(2)" description:"虚拟账户冻结余额"`
}

//代理关系表  对于二级代理 LevelTwo为空字符串 对于大代理levelOne LevelTwo都是空字符串
type Agency struct {
	Id 				int			`orm:"pk;auto"`
	Userid			string		`orm:"index" description:"用户id"`
	LevelOne		string		`orm:"index" description:"所属大代理id"`
	LevelTwo		string		`orm:"index" description:"所属二级代理id"`
}

//代理用户佣金收入表
type Commissionrecord struct {
	Id 			int			`orm:"pk;auto"`
	Userid		string		`orm:"index" description:"用户id"`
	Settletime 	int64  		`description:"下单结算时间"`
	Amount		float64  	`orm:"digits(12);decimals(2)" description:"下单金额"`
	Agencyid	string   	`description:"贡献玩家的id"`
	Tid   		string   	`description:"订单号"`
}

type Depositrecord struct {
	Id              int             `orm:"pk;auto"`
	Userid          string          `orm:"index" description:"用户id"`
	Payway          int8            `orm:"index" description:"支付方式"` // 1 支付宝 2微信
	Amount          float64			`orm:"index;digits(12);decimals(2)" description:"充值数量"`
	Createtime      int64			`description:"创建订单时间"`
	Finishtime  	int64			`description:"确认充值订单时间"`
	Tid             string      	`orm:"index" description:"充值订单id" `
	Status     	int  			`orm:"index" description:"充值状态"`// 1 pending  2 成功
}


type Withdrawrecord struct {
	Id          int             `orm:"pk;auto"`
	Userid      string          `orm:"index" description:"用户id"`
	Amount		float64			`orm:"index;digits(12);decimals(2)" description:"提现数量"`
	Bank		string			`description:"银行名称"`
	Banknumber 	string 			`description:"银行卡号"`
	Bankname	string 			`description:"姓名"`
	City		string			`description:"开户城市"`
	Province   	string			`description:"开户省份"`
	Bankbranch 	string			`description:"开户支行"`
	Createtime  int64			`description:"创建订单时间"`
	Finishtime  int64			`description:"确认提现订单时间"`
	Tid         string      	`orm:"index" description:"提现订单id" `
	Status  	int				`orm:"index" description:"提现状态"`
}

type DepositEnsureRecord struct {
	Id              int             `orm:"pk;auto"`
	Tid             string      	`orm:"index" description:"充值订单id" `
	Userid          string          `orm:"index" description:"用户id"`
	Gid				string			`orm:"index" description:"商家下单时传递的商品id"`
	Amount          float64			`orm:"index;digits(12);decimals(2)" description:"充值数量"`
	Ptid			string			`orm:"index" description:"平台订单流水号"`
	Sn 				string  		`orm:"index" descripton:"签名验证值"`
	Createtime      int64			`description:"接收回调时间"`

}

//内部操作 增加配置或者更新配置
type Subject struct {
	Id 			int			`orm:"pk;auto"`
	Symbol 		string      `orm:"index" description:"标的物名称"`
	Isopen		int 		`orm:"index" description:"是否开启该标的物竞猜服务"` //0 关闭 1开启
}*/


func Exist(table,filed string, value interface{}) bool {
	o := getOrm()
	return o.QueryTable(table).Filter(filed,value).Exist()
}

func MultiExist(table string, cond map[string]interface{}) bool {
	o := getOrm()
	qs := o.QueryTable(table)
	for key,value := range cond {
		qs = qs.Filter(key,value)
	}
	return qs.Exist()
}

func main() {

	if MultiExist("User",map[string]interface{}{"Invitationcode": "5nELL1","Type":1}) {
		fmt.Printf("true")
	}
}
