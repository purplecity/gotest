package mysql

//由于orm的约定 结构体字段第一个字母必须为大写 其他字母必须全为小写
//表名如果有多个大小写也会以_分开在数据库中显示
//默认都是not null
//以后要改为区分大小写

//实际账户下单记录表
type Realtrade struct {
	Id 				int			`orm:"pk;auto"`
	Uid 			string		`orm:"index" description:"用户id"`
	Tid				string		`orm:"unique" description:"订单id"`
	Handletime 		int64		`description:"下单处理时间"`
	Settletime 		int64		`orm:"default(0)" description:"下单结算时间"`
	Inputamount 	float64		`orm:"digits(12);decimals(2)" description:"下单金额"`
	Outputamount	float64		`orm:"digits(12);decimals(2)" description:"结算金额"`
	Ordervalue		float64		`orm:"digits(12);decimals(2)" description:"下单指数"`
	Settlevalue		float64		`orm:"digits(12);decimals(2)" description:"结算指数"`
	Side   			int32		`description:"看涨看跌"`
	Interval		int32		`description:"下单周期时间"`
	Symbol			string		`orm:"index" description:"标的物"`
	Orderresult		int32		`orm:"index;default(1)" description:"下单结果"`
	Settlereason  	string  	`description:"结算原因"`
	Odds 			float64     `orm:"digits(12);decimals(2)" description:"赔率"`
}

//虚拟账户下单记录表
type Vitualtrade struct {
	Id 				int			`orm:"pk;auto"`
	Uid 			string		`orm:"index" description:"用户id"`
	Tid				string		`orm:"unique" description:"订单id"`
	Handletime 		int64		`description:"下单处理时间"`
	Settletime 		int64		`orm:"default(0)" description:"下单结算时间"`
	Inputamount 	float64		`orm:"digits(12);decimals(2)" description:"下单金额"`
	Outputamount	float64		`orm:"digits(12);decimals(2)" description:"结算金额"`
	Ordervalue		float64		`orm:"digits(12);decimals(2)" description:"下单指数"`
	Settlevalue		float64		`orm:"digits(12);decimals(2)" description:"结算指数"`
	Side   			int32		`description:"看涨看跌"`
	Interval		int32		`description:"下单周期时间"`
	Symbol			string		`orm:"index" description:"标的物"`
	Orderresult		int32		`orm:"index;default(1)" description:"下单结果"`
	Settlereason  	string  	`description:"结算原因"`
	Odds 			float64     `orm:"digits(12);decimals(2)" description:"赔率"`
}

//用户表 包括合伙人 总监 玩家
type AdminUsers struct {
	Id 					int			`orm:"pk;auto"`
	Uid				string		`orm:"unique" description:"用户id"`
	Username			string		`orm:"unique" description:"用户名"`
	Phonenumber			string		`orm:"unique" description:"手机号码"`
	Password			string		`description:"密码"`
	Invitationcode	string			`orm:"index" description:"邀请码"`
	Type 				string		`orm:"index" description:"玩家类型"`
	Registtime 			int64		`description:"注册时间"`
	Lastlogintime		int64		`description:"最后登录时间"`
	Valid    			int			`description:"是否有效"` //0 无效 1有效
	RememberToken				string		`description:"代理token"`
}

//资金表 包括总监 玩家
type Asset struct {
	Id 					int			`orm:"pk;auto"`
	Uid					string		`orm:"unique" description:"用户id"`
	Balance				float64		`orm:"digits(12);decimals(2)" description:"实际账户余额"`
	Freezebalance		float64		`orm:"digits(12);decimals(2)" description:"实际账户冻结余额"`
	Vitualbalance 		float64     `orm:"digits(12);decimals(2)" description:"虚拟账户余额"`
	Vitualfreezebalance	float64		`orm:"digits(12);decimals(2)" description:"虚拟账户冻结余额"`
}

//合伙人信息
type Parter struct {
	Id 			int			`orm:"pk;auto"`
	Uid			string		`orm:"unique" description:"合伙人id"`
	Invitationcode	string			`orm:"index" description:"邀请码"`
	Isvip		int		`orm:"index" description:"是否vip"` //1 是 0 不是
}

// 总监代理关系
type Director struct {
	Id 			int			`orm:"pk;auto"`
	Directorid	string		`orm:"index" description:"总监id"`
	Partnerid	string		`orm:"index" description:"所属合伙人id"`
	Isvip		int		`orm:"index" description:"所属合伙人是否vip"`
}

//玩家代理关系
type Player struct {
	Id 			int			`orm:"pk;auto"`
	Uid			string		`orm:"unique" description:"用户id"`
	Playerid	string		`orm:"index" description:"所属玩家id"`
	Directorid	string		`orm:"index" description:"所属总监id"`
}

//积分记录
type Scorerecord struct {
	Id 				int			`orm:"pk;auto"`
	Directorid		string		`orm:"index" description:"总监id"`
	Playerid		string		`orm:"index" description:"玩家id"`
	Settletime 		int64  		`description:"下单结算时间"`
	Amount			float64  	`orm:"digits(12);decimals(2)" description:"下单金额"`
	Contributorid	string   	`orm:"index" description:"贡献玩家的id"`
	Tid   			string   	`description:"订单号"`
}

//每天24点05分计算 所以24点-1点不可提取积分 否则会出现 total != hadwithdraw + remain
type Score struct {
	Id 				int			`orm:"pk;auto"`
	Uid			string			`orm:"unique" description:"用户id"`
	Total 			float64		`orm:"digits(12);decimals(2)" description:"总积分"`
	Hadwithdraw		float64		`orm:"digits(12);decimals(2)" description:"已经提取积分"`
	Remain 			float64		`orm:"digits(12);decimals(2)" description:"剩余积分"`
}


type Takescorerecord struct {
	Id 				int			`orm:"pk;auto"`
	Uid			string			`orm:"index" description:"用户id"`
	Handletime 		int64		`description:"提取积分时间"`
	Amount 	float64		`orm:"digits(12);decimals(2)" description:"提取数量"`

}

type Depositway struct {
	Id 				int			`orm:"pk;auto"`
	Way 			int 		`orm:"index" description:"充值方式"` // 1 支付宝 2微信 3 银行卡
	Isopen 			int  		`orm:"index" description:"充值开关"` // 0 关闭 1开启
}


type Depositrecord struct {
	Id              int             `orm:"pk;auto"`
	Uid          	string          `orm:"index" description:"用户id"`
	Payway          int8            `orm:"index" description:"支付方式"` // 1 支付宝 2微信 3 银行卡
	Amount          float64			`orm:"index;digits(12);decimals(2)" description:"辨识充值数量"`
	RealAmount		float64			`orm:"index;digits(12);decimals(2)" description:"真实充值数量"`
	Createtime      int64			`description:"创建订单时间"`
	Finishtime  	int64			`description:"确认充值订单时间"`
	Tid             string      	`orm:"unique" description:"充值订单id" `
	Status     		int  			`orm:"index" description:"充值状态"`// 0 pending  1 成功
	Isclick			int				`orm:"index" description:"是否点击"`// 0 未点击  1 点击
	Bank		string			`description:"银行名称"`
	Banknumber 	string 			`orm:"index" description:"银行卡号"`
	Bankname	string 			`orm:"index" description:"开户人"`
	City		string			`description:"开户城市"`
	Province   	string			`description:"开户省份"`
	Bankbranch 	string			`description:"开户支行"`
	Postscript  string			`orm:"index" description:"附言" `
}


type Withdrawrecord struct {
	Id          int             `orm:"pk;auto"`
	Uid      	string          `orm:"index" description:"用户id"`
	Amount		float64			`orm:"index;digits(12);decimals(2)" description:"提现数量"`
	Bank		string			`description:"银行名称"`
	Banknumber 	string 			`orm:"index" description:"银行卡号"`
	Bankname	string 			`orm:"index" description:"开户人"`
	City		string			`description:"开户城市"`
	Province   	string			`description:"开户省份"`
	Bankbranch 	string			`description:"开户支行"`
	Createtime  int64			`description:"创建订单时间"`
	Finishtime  int64			`description:"确认提现订单时间"`
	Tid         string      	`orm:"unique" description:"提现订单id" `
	Status  	int				`orm:"index" description:"提现状态"`
}


type BankInfo struct {
	Id          int         	`orm:"pk;auto"`
	Uid         string      	`orm:"index" description:"用户id"`
	Bank		string			`description:"银行名称"`
	Banknumber 	string 			`orm:"index" description:"银行卡号"`
	Bankname	string 			`orm:"index" description:"姓名"`
	City		string			`description:"开户城市"`
	Province   	string			`description:"开户省份"`
	Bankbranch 	string			`orm:"index" description:"开户支行"`
}

//内部操作 增加配置或者更新配置
type Subject struct {
	Id 				int			`orm:"pk;auto"`
	Symbol 			string      `orm:"index" description:"标的物名称"`
	Type   			string 		`orm:"index" description:"标的物所属种类"`
	Isopen			int 		`orm:"index" description:"是否开启该标的物竞猜服务"` //0 关闭 1开启
}

type Clientversion struct {
	Id 			int			`orm:"pk;auto"`
	Version  	string 		`orm:"index" description:"版本号"`
	Isforce  	int 		`orm:"index" description:"是否强制更新0否1是"`
	Createtime  int64      	`orm:"index" description:"创建版本时间"`
}

type AdminRoleUsers struct {
	Id 			int		`orm:"pk;auto"`
	RoleId		int  	`orm:"index"`
	UserId		int 	`orm:"index"`
}

type Lastconnect struct {
	Id          int         	`orm:"pk;auto"`
	Uid         string      	`orm:"unique" description:"用户id"`
	Symbol 		string      `orm:"index" description:"标的物名称"`
}

type Depositbank struct {
	Id          int         	`orm:"pk;auto"`
	Bank		string			`description:"银行名称"`
	Banknumber 	string 			`orm:"unique" description:"银行卡号"`
	Bankname	string 			`orm:"index" description:"开户人"`
	City		string			`description:"开户城市"`
	Province   	string			`description:"开户省份"`
	Bankbranch 	string			`orm:"index" description:"开户支行"`
}

//每天1点结算
type Reconciliation struct {
	Id          int         	`orm:"pk;auto"`
	Uid         string      	`orm:"unique" description:"用户id"`
	Balance 	float64			`orm:"digits(12);decimals(2)" description:"3点余额"`
	Lastbalance float64		`orm:"digits(12);decimals(2)" description:"昨天3点余额"`
	Win 		float64			`orm:"digits(12);decimals(2)" description:"当天1点盈利收入"`
	Lose 		float64			`orm:"digits(12);decimals(2)" description:"当天1点累计亏损"`
	Deposit 	float64 		`orm:"digits(12);decimals(2)" description:"当天1点累计充值收入"`
	Withdraw  	float64			`orm:"digits(12);decimals(2)" description:"当天1点累计提现支出"`
	Score  		float64  		`orm:"digits(12);decimals(2)" description:"当天1点累计提取积分收入"`

	Handletime 		int64		`description:"对账时间"`
}

type Remarks struct {
	Id 					int			`orm:"pk;auto"`
	Uid				string		`orm:"unique" description:"用户id"` //可能是合伙人总监玩家
	Directorremark 	string		`description:"总监备注"`
	Parterremark 	string		`description:"合伙人备注"`
	Csremark 		string		`description:"客服备注"`
	Csid 			string   	`orm:"index" description:"客服id"`
}

type Sounds struct {
	Id 					int			`orm:"pk;auto"`
	Uid				string		`orm:"unique" description:"用户id"`
	Bgm  			int  		`description:"bgm开关"`
	Buttonsound     int 		`description:"按钮声音"`
}

type Payamount struct {
	Id 					int			`orm:"pk;auto"`
	Payway				int		`orm:"unique" description:"充值方式"` //1支付宝 2微信 3银行卡
}

type Alipayensure struct {
	Id              int             `orm:"pk;auto"`
	Tid             string      	`orm:"unique" description:"充值订单id" `
	Uid          	string          `orm:"index" description:"用户id"`
	Amount          float64			`orm:"index;digits(12);decimals(2)" description:"充值数量"`
	Sn 				string  		`orm:"index" descripton:"回调返回的签名验证值"`
	Ts   			string  		`orm:"index" descripton:"请求平台时候返回的交易号"`
	Createtime      int64			`description:"接收回调时间"`
	Status   		int  			`description:"状态"`
}