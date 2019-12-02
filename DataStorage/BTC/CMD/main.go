package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gotest/DataStorage/Operation"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	BTCBaseURL 		= "wss://stream.binance.com:9443/ws"
	starttime = int64(1111)
	price = float64(0)
	Mu sync.Mutex
	Event WsTradeEvent
	AssignData map[string]interface{} = map[string]interface{}{"ts":int64(0),"price":float64(0)}
	LastEvent map[string]interface{} = map[string]interface{}{"ts":int64(0),"price":float64(0)}
	datamap = make(map[int64]float64)
	Firstclosehour = 2
	Firstclosemin = 30
	Firstopenhour = 9
	Firstopenmin = 30
)

type WsTradeEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	TradeID       int64  `json:"t"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	BuyerOrderID  int64  `json:"b"`
	SellerOrderID int64  `json:"a"`
	TradeTime     int64  `json:"T"`
	IsBuyerMaker  bool   `json:"m"`
	Placeholder   bool   `json:"M"` // add this field to avoid case insensitive unmarshaling
}


func PullBinanceData(symbol string) {
	endpoint := fmt.Sprintf("%s/%s@trade",BTCBaseURL, strings.ToLower(symbol))
	hpdial := &websocket.Dialer{}
	count := 1

	for {
		wsConn,_,err := hpdial.Dial(endpoint,nil)
		if err != nil {
			log.Printf("ERROR----dial binance ws failed----err:%+v\n",err)
			time.Sleep(time.Second*1)
			continue
		}
		doneC := make(chan struct{})
		go func() {
			defer wsConn.Close()
			defer close(doneC)

			for {
				_, message, err := wsConn.ReadMessage() //如果币安维护 将祖斯在这
				if err != nil {
					log.Printf("ERROR----read binance message failed----err:%+v\n",err.Error())
					return
				}
				if message != nil {
					Mu.Lock()
					json.Unmarshal(message,&Event)
					if Event.TradeTime/1e3 >= LastEvent["ts"].(int64) + 1 {
						AssignData["ts"] = LastEvent["ts"]
						AssignData["price"] = LastEvent["price"]
					}
					LastEvent["ts"] = Event.TradeTime/1e3
					LastEvent["price"],_ = strconv.ParseFloat(Event.Price,64)
					Mu.Unlock()
				}
			}
		}()
		<-doneC
		log.Printf("WARNING----the %v time to reconnect binance ws\n",count)
		count++
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
	if st > Bt1 && (st - Bt1)%90 == 0{
		m = Bt1 + ((st - Bt1)/90 - 1)*90
		//ordertime = ts + 30
	} else if st <= Et1 && (st - bt1)%90 == 0{
		m = bt1 + ((st - bt1)/90 -1)*90
		//ordertime = ts + 30
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
	var pricepath= "/root/DataStorage/btcprice.csv"
	pricefile, err := os.OpenFile(pricepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer pricefile.Close()
	pricefile.WriteString("\xEF\xBB\xBF")
	pricew := csv.NewWriter(pricefile)
	pricew.Write([]string{"时间","价格"})
	pricew.Flush()

	var tradepath = "/root/DataStorage/btctrade.csv"
	tradefile, err := os.OpenFile(tradepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer tradefile.Close()
	tradefile.WriteString("\xEF\xBB\xBF")
	tradew := csv.NewWriter(tradefile)
	tradew.Write([]string{"期号","下单价格","结算价格","涨跌结果","单双对结果"})
	tradew.Flush()

	go PullBinanceData("BTCUSDT")
	now := time.Now()
	st := time.Unix(starttime,0)
	st = st.Add(time.Millisecond*-900)
	time.Sleep(st.Sub(now))
	tick := time.Tick(1000 * time.Millisecond)

	for range tick {
		tt := time.Now()
		ts := tt.Unix()
		st := ts - 1
		Mu.Lock()
		if ts == LastEvent["ts"] {
			price = Operation.HPround(AssignData["price"].(float64),2)
		} else {
			price = Operation.HPround(LastEvent["price"].(float64),2)
		}
		Mu.Unlock()
		pricew.Write([]string{fmt.Sprint("%+v",st),fmt.Sprintf("%+v",price)})
		pricew.Flush()
		datamap[st] = price
		if len(datamap) < 61 {
			continue
		} else if len(datamap) > 61 {
			delete(datamap,st-61)
		}

		issue := getResult(st)
		if issue > 0 {
			sdp := getSDP(price)
			ud := 0
			ov := datamap[st-60]
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