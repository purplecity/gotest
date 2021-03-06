package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	apikey = "d35b9daaa34042469093826186e511da"
	baseurl = "http://api.bw-gaming.com/"
)

func main() {
	data := url.Values{}
	bgst := int64(1583812800) //030910
	bget := int64(1583834400)  //030915
	data.Set("Type","UpdateAt")
	data.Set("StartAt",time.Unix(bgst,0).Format("2006-01-02 15:04:05"))
	data.Set("EndAt",time.Unix(bget,0).Format("2006-01-02 15:04:05"))
	data.Set("PageSize","500")



	r, _ := http.NewRequest("POST", baseurl+"api/log/get", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization",apikey)

	trans := http.Transport{
		DisableKeepAlives:true,
	}

	client := &http.Client{
		Transport: &trans,
	}

	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("get  fy  order xmlfileread failed %+v\n",err)
		return
	}


	readBytes, _ := ioutil.ReadAll(resp.Body)
	y := map[string]interface{}{}
	json.Unmarshal(readBytes,&y)
	for k,v := range y["info"].(map[string]interface{})["list"].([]interface{})[0].(map[string]interface{}) {
		log.Printf("%+v,%T",k,v)
	}
	log.Printf("get  fy  order xmlfileread return: %+v\n",string(readBytes))
	/*
	2020/03/10 15:24:51 Money,string
	2020/03/10 15:24:51 StartAt,string
	2020/03/10 15:24:51 CateID,string
	2020/03/10 15:24:51 League,string
	2020/03/10 15:24:51 MatchID,string
	2020/03/10 15:24:51 BetAmount,string
	2020/03/10 15:24:51 IP,string
	2020/03/10 15:24:51 Category,string
	2020/03/10 15:24:51 LeagueID,string
	2020/03/10 15:24:51 Result,string
	2020/03/10 15:24:51 Odds,string
	2020/03/10 15:24:51 Status,string
	2020/03/10 15:24:51 CreateAt,string
	2020/03/10 15:24:51 UpdateAt,string
	2020/03/10 15:24:51 ResultAt,string
	2020/03/10 15:24:51 BetID,string
	2020/03/10 15:24:51 Bet,string
	2020/03/10 15:24:51 Content,string
	2020/03/10 15:24:51 BetMoney,string
	2020/03/10 15:24:51 RewardAt,string
	2020/03/10 15:24:51 OrderID,string
	2020/03/10 15:24:51 UserName,string
	2020/03/10 15:24:51 Match,string
	2020/03/10 15:24:51 EndAt,string
	*/
}
