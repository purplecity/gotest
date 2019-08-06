package main

import (
	"fmt"
)

func main() {

	/*
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	fmt.Println(t.Unix() + 1)
	t3 := t2.AddDate(0, 0, 1)
	fmt.Println(t3)

	now := time.Now()
	fmt.Println(now)
	// 计算下一个零点
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	time.NewTimer()
	fmt.Println(next)

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	fmt.Printf("----%+v\n",currentLocation)
	fmt.Printf("----%+v\n",currentYear)
	fmt.Printf("----%+v\n",currentMonth)

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)

	fmt.Println(firstOfMonth.Unix())
	fmt.Println(lastOfMonth.Unix())

	timeStr := time.Now().Format("2006-01-02")
	t1, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	t2 := t1.AddDate(0, 0, -1)

	fmt.Println(t1.Unix())
	fmt.Println(t2.Unix())

	var Starttime int64 = 1562033860000
	var Endtime int64 = 1558663200
	st := time.Unix(Starttime,0)
	et := time.Unix(Endtime,0)

	fmt.Printf("%v\n",int(st.Sub(et).Seconds()))



	tt := time.Now()
	bt1 := time.Date(tt.Year(), tt.Month(), tt.Day(), 9, 30, 0, 0, tt.Location())
	et1 := time.Date(tt.Year(), tt.Month(), tt.Day(), 11, 30, 0, 0, tt.Location())
	bt2 := time.Date(tt.Year(), tt.Month(), tt.Day(), 13, 0, 0, 0, tt.Location())
	et2 := time.Date(tt.Year(), tt.Month(), tt.Day(), 15, 0, 0, 0, tt.Location())
	fmt.Printf("%v,%v\n",bt1.Unix(),et1.Unix(),bt2.Unix(),et2.Unix())
	in := 3800
	h := in/3600
	m := (in - h*3600)/60

	fmt.Println(m)
	var Starttime int64 = 1562033860000
	var Endtime int64 = 1558663200
	st := time.Unix(Starttime,0)
	et := time.Unix(Endtime,0)

	fmt.Printf("%v\n",int(st.Sub(et).Seconds()))
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)
	begin := firstOfMonth.Unix() * 1000
	end := lastOfMonth.Unix() * 1000
	fmt.Println(begin,end)
	et := time.Unix(1558663200,0)
	st := et.Add(time.Second * -30).Unix()
	fmt.Println(et,st)*

	for range time.Tick(time.Millisecond * 1e3) {
		fmt.Println(time.Now().UnixNano())
	}*/

	nt := 1564416000
	//nt1 := time.Unix(int64(nt),0)
	et := 1564302876
	///et1 := time.Unix(int64(et),0)

	fmt.Printf("%v\n",(nt-et)/86400)
	//fmt.Printf("%v,\n",int64(nt1.Sub(et1).Hours()))
}