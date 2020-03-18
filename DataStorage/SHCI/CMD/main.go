package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"gotest/DataStorage/Operation"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	SHCIBaseURL 		= "https://api.wmcloud.com/xmlfileread/v1/api/market2/getTickRTSnapshotL2.json?securityID=000001.XSHG&field="
	SHCIToken 			= "Bearer b6275e695f9442f2b08f2f7604e42607e1f9d24157e52d21e0abc9f9947b0bdc"
	starttime = int64(1111)
	price = float64(0)
	Mu sync.Mutex
	signalChannel = make(chan struct{})
	done = make(chan error)
	ReconnctMaxTime = 3
	LastPrice float64 = 0
	datamap = make(map[int64]float64)
	Firstclosehour = 11
	Firstclosemin = 30
	Firstopenhour = 9
	Firstopenmin = 30
	Secondopenmin   = 0
	Secondopenhour  = 13
	Secondclosemin  = 0
	Secondclosehour  = 15
)

func ExcuteTask() {

	tt := time.Now()

	//取消延后一miao是因为select的偶然性怕 中断消息优先被执行而遗漏了最后1s的行情
	//执行task提前2s是想开盘的第一秒数据是有的 因为pull task会停1s再执行
	taskExcuteFirst := toolbox.NewTask("ExcuteFirst", getString(Firstopenmin,Firstopenhour), startHandle)
	taskFirstCancel := toolbox.NewTask("FirstCancel", fmt.Sprintf("1 %d %d * * 1,2,3,4,5",Firstclosemin,Firstclosehour),cancelHandle)
	taskExcuteSecond := toolbox.NewTask("ExcuteSecond", getString(Secondopenmin,Secondopenhour), startHandle)
	taskSecondCancel := toolbox.NewTask("SecondCancel", fmt.Sprintf("1 %d %d * * 1,2,3,4,5",Secondclosemin,Secondclosehour), cancelHandle)
	toolbox.AddTask("ExcuteFirst", taskExcuteFirst)
	toolbox.AddTask("FirstCancel", taskFirstCancel)
	toolbox.AddTask("ExcuteSecond", taskExcuteSecond)
	toolbox.AddTask("SecondCancel", taskSecondCancel)
	toolbox.StartTask()


	Et1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstclosehour, Firstclosemin, 1, 0, tt.Location()).Unix()
	Et2 := time.Date(tt.Year(), tt.Month(), tt.Day(), Secondclosehour, Secondclosemin, 1, 0, tt.Location()).Unix()
	Bt1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstopenhour, Firstopenmin, 0, 0, tt.Location()).Unix()
	Bt2 := time.Date(tt.Year(), tt.Month(), tt.Day(), Secondopenhour, Secondopenmin, 0, 0, tt.Location()).Unix()


	if (tt.Unix() > (Bt1 -30)  && tt.Unix() < Et1 && tt.Weekday() <=5 && tt.Weekday() >= 1 ) || ( tt.Unix() > (Bt2 -30 ) && tt.Unix() < Et2 && tt.Weekday() <=5 && tt.Weekday() >= 1) {
		go pullData()
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
	request, err := http.NewRequest("GET",SHCIBaseURL,nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n",err)
	}
	request.Header.Set("Authorization",SHCIToken)
	request.Header.Set("accept-encoding","gzip")
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
			for count < ReconnctMaxTime {
				resp, err := client.Do(request) //超时设置
				if err != nil  {
					count++
					log.Printf("ERROR----request xshg failed----err:%v\n", err)
					continue
				} else {
					respmap := map[string]interface{}{}
					readBytes, _ := ioutil.ReadAll(resp.Body)
					json.Unmarshal([]byte(readBytes), &respmap)
					//fmt.Printf("%v\n",respmap)
					if  _,ok := respmap["retMsg"];!ok || respmap["retMsg"] != "Success" {
						count++
						log.Printf("ERROR----xshg reponse err----resp:%v\n", respmap)
						continue
					} else {
						Mu.Lock()
						LastPrice = respmap["xmlfileread"].([]interface{})[0].(map[string]interface{})["lastPrice"].(float64)
						Mu.Unlock()
						resp.Body.Close()
						break
					}
				}
			}

			if count == ReconnctMaxTime  {

				log.Printf("ERROR----request shci 3th failed, close shci server----err:%v\n", err)

			}
			count = 0
		case <-signalChannel:
			return
		}
	}
}

func getResult(st int64) int64 {
	tt := time.Unix(st,0)
	m := int64(0)
	issue := int64(0)
	Et1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstclosehour, Firstclosemin, 0, 0, tt.Location()).Unix()
	Et2 := time.Date(tt.Year(), tt.Month(), tt.Day(), Secondclosehour, Secondclosemin, 0, 0, tt.Location()).Unix()
	Bt1 := time.Date(tt.Year(), tt.Month(), tt.Day(), Firstopenhour, Firstopenmin, 0, 0, tt.Location()).Unix()
	Bt2 := time.Date(tt.Year(), tt.Month(), tt.Day(), Secondopenhour, Secondopenmin, 0, 0, tt.Location()).Unix()

	if (st > Bt1  && st <= Et1  && (st - Bt1)%90 == 0)   {
		m = Bt1 + ((st - Bt1)/90 - 1)*90

	} else if ( st > Bt2 && st <= Et2  && (st -Bt2) %90 == 0) {
		m = Bt2 + ((st - Bt2)/90 - 1)*90

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
	a := fmt.Sprintf("%.2f",p)


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
	var pricepath= "/root/DataStorage/shciprice.csv"
	pricefile, err := os.OpenFile(pricepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer pricefile.Close()
	pricefile.WriteString("\xEF\xBB\xBF")
	pricew := csv.NewWriter(pricefile)
	pricew.Write([]string{"时间","价格"})
	pricew.Flush()

	var tradepath = "/root/DataStorage/shcitrade.csv"
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
			log.Printf("issue:%+v,ov:%+v,price:%+v,ud:%+v,sdp:%+v\n",issue,ov,price,ud,sdp)
			tradew.Write([]string{fmt.Sprint("%+v",issue),fmt.Sprintf("%+v",ov),fmt.Sprintf("%+v",price),fmt.Sprintf("%+v",ud),fmt.Sprintf("%+v",sdp)})
			tradew.Flush()
		}
	}

}