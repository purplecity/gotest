package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var  (
	YSBPrefix = "1429"
	YSBsecretkey = "N37Wq9cfPSz24fyb"
	YSBDATAURL = "http://webapi.a1go.org/WhiteLabelApi/api/Vendor"
)

func main() {
	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}


	type ysbparam struct {
		MsgId string
		DateFrom string
		DateTo string
		VendorId string
		HashPassword string
	}
	bgst := int64(1583757650) //030910
	bget := int64(1583844590)  //030915
	data := map[string]interface{}{}
	param := ysbparam{}
	utcLoc,_ := time.LoadLocation("")
	fromtime := time.Unix(bgst,0).In(utcLoc).Format("06-01-02T15:04")
	totime := time.Unix(bget,0).In(utcLoc).Format("06-01-02T15:04")

	param.MsgId = "BD"
	param.DateFrom = fromtime
	param.DateTo = totime
	param.VendorId = YSBPrefix
	srcbytes := []byte("BD"+fromtime+totime+YSBPrefix+YSBsecretkey)
	log.Println(string(srcbytes))
	md5str := fmt.Sprintf("%x", md5.Sum(srcbytes))
	param.HashPassword = strings.ToUpper(md5str)
	data["Params"] = param
	log.Printf("%+v\n",param)
	jsonstr, _ := json.Marshal(data)

	r, _ := http.NewRequest("POST", YSBDATAURL,bytes.NewBuffer(jsonstr)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("get ysb  order xmlfileread failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)
	y := map[string]interface{}{}
	json.Unmarshal(readBytes,&y)
	for k,v := range y["Data"].([]interface{})[0].(map[string]interface{}) {
		log.Printf("%+v,%T",k,v)
	}
	log.Printf("get  ysb  order xmlfileread return: %+v\n",string(readBytes))

	/*
	2020/03/09 21:10:38 LoginID,string
	2020/03/09 21:10:38 TeamType,string
	2020/03/09 21:10:38 SelectionPlace,string
	2020/03/09 21:10:38 Score,string
	2020/03/09 21:10:38 TeasedPoint,string
	2020/03/09 21:10:38 TrxDate,string
	2020/03/09 21:10:38 SportType,string
	2020/03/09 21:10:38 OddFormat,string
	2020/03/09 21:10:38 EventNameCH,string
	2020/03/09 21:10:38 CombinationId,string
	2020/03/09 21:10:38 BonusAmount,string
	2020/03/09 21:10:38 PostedPrice,string
	2020/03/09 21:10:38 SelectionEN,string
	2020/03/09 21:10:38 ReSettlement,string
	2020/03/09 21:10:38 Bookings,string
	2020/03/09 21:10:38 EventNameEN,string
	2020/03/09 21:10:38 ParentEventId,string
	2020/03/09 21:10:38 SelectionValue,string
	2020/03/09 21:10:38 EventResults,string
	2020/03/09 21:10:38 ProcessDate,string
	2020/03/09 21:10:38 Corners,string
	2020/03/09 21:10:38 SelectionCH,string
	2020/03/09 21:10:38 RaceToResult,string
	2020/03/09 21:10:38 VendorId,string
	2020/03/09 21:10:38 BetAmount,string
	2020/03/09 21:10:38 Profit,string
	2020/03/09 21:10:38 CompetitionNameEN,string
	2020/03/09 21:10:38 Status,string
	2020/03/09 21:10:38 Void,string
	2020/03/09 21:10:38 DecimalPrice,string
	2020/03/09 21:10:38 YellowCards,string
	2020/03/09 21:10:38 SportsTypeCH,string
	2020/03/09 21:10:38 SettlementAmount,string
	2020/03/09 21:10:38 CashoutID,string
	2020/03/09 21:10:38 EventSession,string
	2020/03/09 21:10:38 BetTypeId,string
	2020/03/09 21:10:38 BetMode,string
	2020/03/09 21:10:38 BetTypeNameCH,string
	2020/03/09 21:10:38 SixMinResult,string
	2020/03/09 21:10:38 Overtime,string
	2020/03/09 21:10:38 BetId,string
	2020/03/09 21:10:38 EventId,string
	2020/03/09 21:10:38 Correction,string
	2020/03/09 21:10:38 CompetitionNameCH,string
	2020/03/09 21:10:38 BetTypeNameEN,string
	2020/03/09 21:10:38 GMTEventDate,string
	2020/03/09 21:10:38 SelectionId,string
	2020/03/09 21:10:38 PayoutPercentage,string
	2020/03/09 21:10:38 Inning,string
	*/

}
