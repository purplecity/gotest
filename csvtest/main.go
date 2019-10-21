package main

import (
	"encoding/csv"
	"fmt"
	"gotest/csvtest/Cache"
	"gotest/csvtest/Operation"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var mEURUSDChan = make(chan float64)
var nEURUSDChan = make(chan float64)
var EURUSDPriceCheck = "EURUSDPriceCheck"

func checkIsNeedCloseEURUSD() {
	for price := range mEURUSDChan {
		lenList,_ := Cache.RedisRPUSH(EURUSDPriceCheck,fmt.Sprintf("%v",price))
		if lenList > 12 {
			Cache.RedisLPop(EURUSDPriceCheck)
			s,_ := Cache.RedisLRange(EURUSDPriceCheck,0,-1)
			pl := []float64{}
			for _, x := range s {
				t,_ := strconv.ParseFloat(x,64)
				pl = append(pl,t)
			}


			avg := float64(0)
			for _,x := range pl {
				avg = Operation.HPAdd(avg,x)
			}
			avg = Operation.HPDivInt(avg,int64(len(pl)))

			rs := float64(0)
			for _,x := range pl {
				rs = Operation.HPAdd(rs,Operation.HPMul(Operation.HPSub(x,avg),Operation.HPSub(x,avg)))
			}

			rs2 := float64(0)
			rs2 = Operation.HPDivInt(rs,int64(len(pl)-1))

			rs3 := math.Sqrt(rs2)
			nEURUSDChan <- Operation.HPround(rs3,8)
			//log.Printf("test check price ---- pl:%+v,avg:%+v,rs2:%+v,rs3:%+v\n",pl,avg,rs2,rs3)
		}
	}
}

func writeCsvEURUSD() {
	var filepath= "/Users/ludongdong/zaqizaba/eurusd.csv"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"波动率"})
	w.Flush()
	for price := range nEURUSDChan {
		w.Write([]string{fmt.Sprintf("%+v\n",price)})
		w.Flush()
	}
}



func EURUSD() {
	var EURUSDBaseURL= "https://webrates.truefx.com/rates/connect.html?p=anystring&f=csv&s=n&c=EUR/USD"
	request, err := http.NewRequest("GET", EURUSDBaseURL, nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n", err)
	}
	trans := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &trans,
	}
	tick := time.Tick(1000 * time.Millisecond)
	count := 0
	LastPrice := float64(0)
	go checkIsNeedCloseEURUSD()
	go writeCsvEURUSD()

	for range tick {
		for count < 3 {
			resp, err := client.Do(request) //超时设置
			if err != nil {
				count++
				log.Printf("ERROR----request eurusd failed----err:%v\n", err)
				continue
			} else {
				readBytes, _ := ioutil.ReadAll(resp.Body)
				dataSlice := strings.Split(string(readBytes), ",")
				if len(dataSlice) != 9 {
					count++
					log.Printf("ERROR----eurusd reponse err----resp:%+v\n", dataSlice)
					continue
				} else {
					//StopRunChan <- Conf.Recovery //status为true 程序启动第一次必然不是recovery 后续也不用发消息
				}
				big, _ := strconv.ParseFloat(dataSlice[2]+dataSlice[3], 64)
				pips, _ := strconv.ParseFloat(dataSlice[4]+dataSlice[5], 64)
				LastPrice = Operation.HPround(Operation.HPMul(Operation.HPAdd(big, pips), float64(0.5)), 5)
				mEURUSDChan <- LastPrice
				resp.Body.Close()
				break
			}
		}
	}
}



var mEURJPYChan = make(chan float64)
var nEURJPYChan = make(chan float64)
var EURJPYPriceCheck = "EURJPYPriceCheck"

func checkIsNeedCloseEURJPY() {
	for price := range mEURJPYChan {
		lenList,_ := Cache.RedisRPUSH(EURJPYPriceCheck,fmt.Sprintf("%v",price))
		if lenList > 12 {
			Cache.RedisLPop(EURJPYPriceCheck)
			s,_ := Cache.RedisLRange(EURJPYPriceCheck,0,-1)
			pl := []float64{}
			for _, x := range s {
				t,_ := strconv.ParseFloat(x,64)
				pl = append(pl,t)
			}


			avg := float64(0)
			for _,x := range pl {
				avg = Operation.HPAdd(avg,x)
			}
			avg = Operation.HPDivInt(avg,int64(len(pl)))

			rs := float64(0)
			for _,x := range pl {
				rs = Operation.HPAdd(rs,Operation.HPMul(Operation.HPSub(x,avg),Operation.HPSub(x,avg)))
			}

			rs2 := float64(0)
			rs2 = Operation.HPDivInt(rs,int64(len(pl)-1))

			rs3 := math.Sqrt(rs2)
			nEURJPYChan <- Operation.HPround(rs3,8)
			//log.Printf("test check price ---- pl:%+v,avg:%+v,rs2:%+v,rs3:%+v\n",pl,avg,rs2,rs3)
		}
	}
}

func writeCsvEURJPY() {
	var filepath= "/Users/ludongdong/zaqizaba/EURJPY.csv"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"波动率"})
	w.Flush()
	for price := range nEURJPYChan {
		w.Write([]string{fmt.Sprintf("%+v\n",price)})
		w.Flush()
	}
}



func EURJPY() {
	var EURJPYBaseURL= "https://webrates.truefx.com/rates/connect.html?p=anystring&f=csv&s=n&c=EUR/JPY"
	request, err := http.NewRequest("GET", EURJPYBaseURL, nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n", err)
	}
	trans := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &trans,
	}
	tick := time.Tick(1000 * time.Millisecond)
	count := 0
	LastPrice := float64(0)
	go checkIsNeedCloseEURJPY()
	go writeCsvEURJPY()

	for range tick {
		for count < 3 {
			resp, err := client.Do(request) //超时设置
			if err != nil {
				count++
				log.Printf("ERROR----request EURJPY failed----err:%v\n", err)
				continue
			} else {
				readBytes, _ := ioutil.ReadAll(resp.Body)
				dataSlice := strings.Split(string(readBytes), ",")
				if len(dataSlice) != 9 {
					count++
					log.Printf("ERROR----EURJPY reponse err----resp:%+v\n", dataSlice)
					continue
				} else {
					//StopRunChan <- Conf.Recovery //status为true 程序启动第一次必然不是recovery 后续也不用发消息
				}
				big, _ := strconv.ParseFloat(dataSlice[2]+dataSlice[3], 64)
				pips, _ := strconv.ParseFloat(dataSlice[4]+dataSlice[5], 64)
				LastPrice = Operation.HPround(Operation.HPMul(Operation.HPAdd(big, pips), float64(0.5)), 5)
				mEURJPYChan <- LastPrice
				resp.Body.Close()
				break
			}
		}
	}
}


var mUSDJPYChan = make(chan float64)
var nUSDJPYChan = make(chan float64)
var USDJPYPriceCheck = "USDJPYPriceCheck"

func checkIsNeedCloseUSDJPY() {
	for price := range mUSDJPYChan {
		lenList,_ := Cache.RedisRPUSH(USDJPYPriceCheck,fmt.Sprintf("%v",price))
		if lenList > 12 {
			Cache.RedisLPop(USDJPYPriceCheck)
			s,_ := Cache.RedisLRange(USDJPYPriceCheck,0,-1)
			pl := []float64{}
			for _, x := range s {
				t,_ := strconv.ParseFloat(x,64)
				pl = append(pl,t)
			}


			avg := float64(0)
			for _,x := range pl {
				avg = Operation.HPAdd(avg,x)
			}
			avg = Operation.HPDivInt(avg,int64(len(pl)))

			rs := float64(0)
			for _,x := range pl {
				rs = Operation.HPAdd(rs,Operation.HPMul(Operation.HPSub(x,avg),Operation.HPSub(x,avg)))
			}

			rs2 := float64(0)
			rs2 = Operation.HPDivInt(rs,int64(len(pl)-1))

			rs3 := math.Sqrt(rs2)
			nUSDJPYChan <- Operation.HPround(rs3,8)
			//log.Printf("test check price ---- pl:%+v,avg:%+v,rs2:%+v,rs3:%+v\n",pl,avg,rs2,rs3)
		}
	}
}

func writeCsvUSDJPY() {
	var filepath= "/Users/ludongdong/zaqizaba/USDJPY.csv"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"波动率"})
	w.Flush()
	for price := range nUSDJPYChan {
		w.Write([]string{fmt.Sprintf("%+v\n",price)})
		w.Flush()
	}
}



func USDJPY() {
	var USDJPYBaseURL= "https://webrates.truefx.com/rates/connect.html?p=anystring&f=csv&s=n&c=USD/JPY"
	request, err := http.NewRequest("GET", USDJPYBaseURL, nil)
	if err != nil {
		log.Panicf("ERROR----construct request error:%v\n", err)
	}
	trans := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &trans,
	}
	tick := time.Tick(1000 * time.Millisecond)
	count := 0
	LastPrice := float64(0)
	go checkIsNeedCloseUSDJPY()
	go writeCsvUSDJPY()

	for range tick {
		for count < 3 {
			resp, err := client.Do(request) //超时设置
			if err != nil {
				count++
				log.Printf("ERROR----request USDJPY failed----err:%v\n", err)
				continue
			} else {
				readBytes, _ := ioutil.ReadAll(resp.Body)
				dataSlice := strings.Split(string(readBytes), ",")
				if len(dataSlice) != 9 {
					count++
					log.Printf("ERROR----USDJPY reponse err----resp:%+v\n", dataSlice)
					continue
				} else {
					//StopRunChan <- Conf.Recovery //status为true 程序启动第一次必然不是recovery 后续也不用发消息
				}
				big, _ := strconv.ParseFloat(dataSlice[2]+dataSlice[3], 64)
				pips, _ := strconv.ParseFloat(dataSlice[4]+dataSlice[5], 64)
				LastPrice = Operation.HPround(Operation.HPMul(Operation.HPAdd(big, pips), float64(0.5)), 5)
				mUSDJPYChan <- LastPrice
				resp.Body.Close()
				break
			}
		}
	}
}

func main() {
	go EURUSD()
	go EURJPY()
	go USDJPY()
	ch := make(chan struct{})
	<- ch
}