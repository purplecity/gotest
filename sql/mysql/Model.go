package mysql

//由于orm的约定 结构体字段第一个字母必须为大写 其他字母必须全为小写
//表名如果有多个大小写也会以_分开在数据库中显示
//默认都是not null
//以后要改为区分大小写

//实际账户下单记录表


//内部操作 增加配置或者更新配置
type Subject struct {
	Id 				int			`orm:"pk;auto"`
	Symbol 			string      `orm:"index" description:"标的物名称"`
	Type   			string 		`orm:"index" description:"标的物所属种类"`
	Isopen			int 		`orm:"index" description:"是否开启该标的物竞猜服务"` //0 关闭 1开启
	Firstopenhour 	int   	`description:"每天第一次开启小时"`
	Firstopenmin 	int   	`description:"每天第一次开启分钟"`
	Firstclosehour	int   	`description:"每天第一次关闭小时"`
	Firstclosemin 	int   	`description:"每天第一次开启分钟"`
	Secondopenhour 	int   	`description:"每天第二次开启小时"`
	Secondopenmin 	int   	`description:"每天第二次开启分钟"`
	Secondclosehour int   	`description:"每天第二次关闭小时"`
	Secondclosemin  int   	`description:"每天第二次关闭分钟"`
}

type Clientversion struct {
	Id 			int			`orm:"pk;auto"`
	Version  	string 		`orm:"index" description:"版本号"`
	Isforce  	int 		`orm:"index" description:"是否强制更新0否1是"`
	Createtime  int64      	`orm:"index" description:"创建版本时间"`
}


//默认等级为1 赔率也要写
type Odds struct {
	Id          int         	`orm:"pk;auto"`
	Symbol 		string      `orm:"index" description:"标的物名称"`
	Upodds 			float64		`orm:"digits(12);decimals(2)" description:"看涨赔率"`
	Downodds		float64		`orm:"digits(12);decimals(2)" description:"看跌赔率"`
}

type OddsInfo struct {
	Id          int         	`orm:"pk;auto"`
	Symbol 		string      `orm:"index" description:"标的物名称"`
	Level 		int      	`orm:"index" description:"风控等级"`
	Mindv  		float64			`orm:"digits(12);decimals(2)" description:"最小差值"`
	Maxdv  		float64			`orm:"digits(12);decimals(2)" description:"最大差值"`
	Greaterodds 		float64			`orm:"digits(12);decimals(2)" description:"大方赔率"`
	Lessodds	float64			`orm:"digits(12);decimals(2)" description:"小方赔率"`
}

