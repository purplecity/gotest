package main

import (
	"fmt"
	"log"

	//"github.com/shopspring/decimal"
	"time"
)


/*
func HPTrunc(x float64,y int32) float64 {
	ret,_ := decimal.NewFromFloat(x).Truncate(y).Float64()
	return ret
}

 */

const(
	InnerError	uint = 1001 + iota
	PhonenumberExisted
	PhonenumberNotExist
	PasswordNotMatchPhonenumber
	NoToken
	TokenInvalid
	BalanceNotEnough
	VerifyCodeNotRight
	CanNotFindOrder
	InvitationCodeInvalid
	InvalidRemoteHost
	CanNotFindDepositOrder
	DontRepeatEnsure
	SendValidCodeFailed
	SymbolInvalid
	InvalidTradeTime
	SymbolServerNotOpen
	TradeRequestTimeOut
	SMSExceedHourLimit
	SMSExceedDayLimit
	SMSExceedMinuteLimit
	ClientVersionNeedUpdate
	NotAllowParter
	NotTakeScoreTime
	RealTimeReconcileNotRight
	HaveNotDepositBankCard
	AbnormalAccount
	CurOddsNotAllowTrade
	ExceedTradeLimit
	NotWithdrawTime
	DepositWayNotOpen
	NotEnoughTradeTimes
	WithdrawExceedHourLimit
	WithdrawExceedDayLimit
	NOTALLOWEDDPR
	CantCanelWithdraw
	UserNameExist
	NotAllowApproveCancel
	GetCashierURLFailed
	InvalidCallback
	DepositZero
	AliPayDepositTimeLimit
	InvalidHandleTime
	CurOddsOnlyTradeOnce
	RequestTooFast
	NotAllowApproveDeduct
	BalanceCantDeduct
	PlatformLossTooMuch
	UpDownNotAllow
	SDPNotAllow
	NotAllowTrade
	NotExceedWithdrawMin
	InvalidOdds
	InvalidCentralismTime
	InvalidTradeMode
	AtPercentFive
	ForexNotAllowCenTrade
	TemporarilyCloseService
	SingleSymbolExceed
	InvalidTradeAmount
	GameNotOpen
	NotRegistrationTime
	NotUpToRequiredRank
	RepeatedReg
	HaveNotReg
	OutOfGameTime
	GameSessionClose
	ImageTypeError
	GeneratePosterFailed
	NeedTradOnce
	EmtPhoneNum
	TmpNotAllowAutoTransfer
	YDlogin
)

func getIssue(tt int64) {
	t := time.Unix(tt,0)
	tt = time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,t.Location()).Unix()

	w := 0
	d := 0
	a := (tt - 1577030400)/(24*3600)
	//有几个周日
	sunNum :=  int(a / 7)
	fmt.Printf("---- %+v,%+v\n",a,sunNum)

	w = 20001 + sunNum - 1
	d = 10001 + sunNum*5-1

	if int(t.Weekday()) == 6 {
		d = d + 5
	} else if int(t.Weekday()) >= 1 && int(t.Weekday()) <=5 {
		d = d + int(t.Weekday()) - 1
	} else if int(t.Weekday()) == 0 {
		w = w + 1
		d = d + 5
	}

	fmt.Printf("%+v,%+v\n",w,d)

}

type testhehe struct {
	A string
	B string
	C string
}
func main() {
	/*
	a := []float64{108.66,119.836,1.10188,119.836,108.66,1.10239,119.836,108.661,1.10228,108.661 ,119.836 ,1.10257}

	b := []float64{119.836 ,108.66 ,1.10239 ,119.836 ,108.661 ,1.10228 ,108.661 ,119.836 ,1.10257 ,108.661 ,119.836 ,1.10169}

	c := []float64{119.836 ,108.661 ,1.10228 ,108.661 ,119.836 ,1.10257 ,108.661 ,119.836 ,1.10169 ,108.661 ,119.836 ,1.10277}
	sort.Float64s(a[:])
	sort.Float64s(b[:])
	sort.Float64s(c[:])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	 */
	//a,_ := strconv.ParseFloat("3.89e-06",64)
	//fmt.Printf("%+v\n",a)
	//tm := time.Now()
	//fmt.Println(tm.Year(),int(tm.Month()),tm.Day(),tm.Hour(),tm.Minute(),tm.Second())
	//fmt.Println(1.7e308 > float64(24*3600*1000000000))
	//getIssue(1578164400)
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(fmt.Sprintf("%.0f",float64(1.69158209e+08)))
	/*
	utcLoc,_ := time.LoadLocation("America/New_York")
	fmt.Println(time.Now().In(utcLoc).Format("2006-01-02T15:04"))

	 */

	/*
	srctime := "202003120805"


	layout :="200601021504"

	//重要：获取时区 

	loc,_ := time.LoadLocation("Local")

	//使用模板在对应时区转化为time.time类型

	dsttime, err :=time.ParseInLocation(layout,srctime,loc)
	if err != nil {
		fmt.Printf("%+v\n",err)
	}
	//转化为时间戳 类型是int64  
	fmt.Printf("%+v\n",dsttime)

	 */

	/*
	str := "2019-12-10T06:58:21.193"
	str2 := "2019-12-10T06:58:21"
	fmt.Printf("%+v\n,%+v\n",strings.Split(str,"."),strings.Split(str2,"."))

	 */

	/*
	layout := "2006-01-02T15:04:05+08:00"
	srctime := "2020-03-18T15:32:01+08:00"
	dsttime, _ :=time.ParseInLocation(layout,srctime,time.Local)
	fmt.Println(dsttime.Unix())
	fmt.Println(YDlogin)

	 */

	t := time.Now()
	bt2 := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,t.Location())

	if t>bt2 {
		log.Print("egege")
	}

}