package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/astaxie/beego/toolbox"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"gotest/DataStorage/Operation"
)

var (
	EURJPYBaseURL 		= "https://webrates.truefx.com/rates/connect.html?p=anystring&f=csv&s=n&c=EUR/JPY"
	starttime = int64(1111)
	price = float64(0)
	Mu sync.Mutex
	signalChannel = make(chan struct{})
	done = make(chan error)
	ForexReconnctMaxTime = 3
	LastPrice float64 = 0
	datamap = make(map[int64]float64)
	Firstclosehour = 2
	Firstclosemin = 30
	Firstopenhour = 9
	Firstopenmin = 30
)

func ExcuteTask() {

	tt := time.Now()

	//取消延后一miao是因为select的偶然性怕 中断消息优先被执行而遗漏了最后1s的行情
	//执行task提前2s是想开盘的第一秒数据是有的 因为pull task会停1s再执行
	taskExcuteFirst := toolbox.NewTask("ExcuteFirst", getString(Firstopenmin,Firstopenhour), startHandle)
	taskFirstCancel := toolbox.NewTask("FirstCancel", fmt.Sprintf("1 %d %d * * 2,3,4,5,6",Firstclosemin,Firstclosehour),cancelHandle)
	toolbox.AddTask("ExcuteFirst", taskExcuteFirst)
	toolbox.AddTask("FirstCancel", taskFirstCancel)
	toolbox.StartTask()

	Et1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstclosehour, Firstclosemin, 0, 0, tt.Location()).Unix()
	Bt1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstopenhour, Firstopenmin, 0, 0, tt.Location()).Unix()

	if ((tt.Unix() < Et1 || tt.Unix() > Bt1 - 30) && tt.Weekday() <=5 && tt.Weekday() >= 2) || (tt.Weekday() == 6 && tt.Unix() < Et1 - 3600) || (tt.Unix() > Bt1 - 30  && tt.Weekday() == 1) {
		go pullData()
		//go handleOdds()
	}

	err := <- done
	log.Printf("ERROR----err:%+v\n",err.Error())
}

func getString(min,hour int) (ef string) {
	if min == 0 {
		ef = fmt.Sprintf("30 %d %d * * 1,2,3,4,5",59,hour-1)
	} else {
		ef = fmt.Sprintf("30 %d %d * * 1,2,3,4,5",min-1,hour)
	}
	return
}

var startHandle = func() error {
	go pullData()
	//go handleOdds()
	return nil
}

var cancelHandle = func() error {
	//oddsChannel <- struct{}{}
	signalChannel <- struct{}{}
	return nil
}


var pullData = func() {
	request, err := http.NewRequest("GET",EURJPYBaseURL,nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n",err)
	}
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	count := 0


	tick := time.Tick(1000 * time.Millisecond)

	for {
		select {
		case <-tick:
			for count < ForexReconnctMaxTime {
				resp, err := client.Do(request) //超时设置
				if err != nil  {
					count++
					log.Printf("ERROR----request eurjpy failed----err:%v\n", err)
					continue
				} else {
					readBytes, _ := ioutil.ReadAll(resp.Body)
					dataSlice := strings.Split(string(readBytes),",")
					if  len(dataSlice) != 9 {
						count++
						log.Printf("ERROR----eurjpy reponse err----resp:%+v\n", dataSlice)
						continue
					} else {
						big,_ := strconv.ParseFloat(dataSlice[2]+dataSlice[3],64)
						pips,_ := strconv.ParseFloat(dataSlice[4]+dataSlice[5],64)
						Mu.Lock()
						LastPrice = Operation.HPMul(Operation.HPAdd(big,pips),float64(0.5))
						fmt.Printf("%+v,%+v\n",time.Now(),LastPrice)
						Mu.Unlock()
						resp.Body.Close()
						break
					}
				}
			}

			if count == ForexReconnctMaxTime {
				log.Printf("ERROR----request eurjpy 3th failed, close eurjpy server----err:%v\n", err)

			}
			count = 0
		case <-signalChannel:
			return
		}
	}
}

func getResult(st int64) int64 {
	tt := time.Unix(st,0)
	Et1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstclosehour, Firstclosemin, 0, 0, tt.Location()).Unix()
	Bt1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstopenhour, Firstopenmin, 0, 0, tt.Location()).Unix()
	last := tt.Add(time.Hour * -24)
	bt1 := time.Date(last.Year(), last.Month(), last.Day(), Firstopenhour, Firstopenmin, 0, 0, last.Location()).Unix()

	m := int64(0)
	issue := int64(0)
	if tt.Weekday() >= 1 && tt.Weekday() <=5 && st > Bt1 && (st - Bt1)%90 == 0{
		m = Bt1 + ((st - Bt1)/90 - 1)*90
		//ordertime = ts + 30
	} else if tt.Weekday() >= 2 && tt.Weekday() <= 5 && st <= Et1 && (st - bt1)%90 == 0{
		m = bt1 + ((st - bt1)/90 -1)*90
		//ordertime = ts + 30
	} else if tt.Weekday() == 6  && st <= Et1 - 3600 && (st - bt1)%90 == 0 {
		m = bt1 + ((st - bt1)/90 -1)*90
	}

	if m != 0 {
		month,day,hour,min,sec := "","","","",""
		tm := time.Unix(m,0)
		if int(tm.Month()) < 10 {
			month = fmt.Sprintf("%+v%+v",0,int(tm.Month()))
		} else {
			month = fmt.Sprintf("%+v",int(tm.Month()))
		}

		if tm.Day() < 10 {
			day = fmt.Sprintf("%+v%+v",0,tm.Day())
		} else {
			day = fmt.Sprintf("%+v",tm.Day())
		}

		if tm.Hour() < 10 {
			hour = fmt.Sprintf("%+v%+v",0,tm.Hour())
		} else {
			hour = fmt.Sprintf("%+v",tm.Hour())
		}

		if tm.Minute() < 10 {
			min = fmt.Sprintf("%+v%+v",0,tm.Minute())
		} else {
			min = fmt.Sprintf("%+v",tm.Minute())
		}

		if tm.Second() < 10 {
			sec = fmt.Sprintf("%+v%+v",0,tm.Second())
		} else {
			sec = fmt.Sprintf("%+v",tm.Second())
		}

		tss := month+day+hour+min+sec
		issue, _ = strconv.ParseInt(tss, 10, 64)
	}
	return issue
}

func getSDP(p float64) int {
	a := fmt.Sprintf("%.3f",p)


	b,_:= strconv.Atoi(a[len(a)-2:len(a)])
	x := b/10
	y := b - x*10
	if x ==y {
		if y == 0 {
			return 5//00
		} else if y % 2 == 0 {
			return 1 //对双
		} else {
			return 2 //对单
		}
	} else {
		if y % 2 == 0 {
			return 3 //双
		} else {
			return 4 //单
		}
	}
}

func main() {
	var pricepath= "/root/DataStorage/eurjpyprice.csv"
	pricefile, err := os.OpenFile(pricepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer pricefile.Close()
	pricefile.WriteString("\xEF\xBB\xBF")
	pricew := csv.NewWriter(pricefile)
	pricew.Write([]string{"时间","价格"})
	pricew.Flush()

	var tradepath = "/root/DataStorage/eurjpytrade.csv"
	tradefile, err := os.OpenFile(tradepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer tradefile.Close()
	tradefile.WriteString("\xEF\xBB\xBF")
	tradew := csv.NewWriter(tradefile)
	tradew.Write([]string{"期号","下单价格","结算价格","涨跌结果","单双对结果"})
	tradew.Flush()

	go ExcuteTask()

	now := time.Now()
	st := time.Unix(starttime,0)
	st = st.Add(time.Millisecond*-900)
	time.Sleep(st.Sub(now))
	tick := time.Tick(1000 * time.Millisecond)

	for range tick {
		tt := time.Now()
		ts := tt.Unix()
		Mu.Lock()
		price = Operation.HPround(LastPrice,3)
		Mu.Unlock()
		pricew.Write([]string{fmt.Sprint("%+v",ts),fmt.Sprintf("%+v",price)})
		pricew.Flush()
		datamap[ts] = price
		if len(datamap) < 61 {
			continue
		} else if len(datamap) > 61 {
			delete(datamap,ts-61)
		}

		issue := getResult(ts)
		if issue > 0 {
			sdp := getSDP(price)
			ud := 0
			ov := datamap[ts-60]
			if Operation.HPgt(price,ov) {
				ud = 1
			} else if Operation.HPEqual(price,ov) {
				ud = 2
			}
			tradew.Write([]string{fmt.Sprint("%+v",issue),fmt.Sprintf("%+v",ov),fmt.Sprintf("%+v",price),fmt.Sprintf("%+v",ud),fmt.Sprintf("%+v",sdp)})
			tradew.Flush()
		}
	}

}