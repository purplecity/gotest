package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	r "math/rand"
	"net/http"
	"strings"
	"time"
)

/*
1 如何停止------接口时间限制。比赛界面如果是年赛后在去请求的直接就给个已经结束 期待下一季。其他接口也加相应的时间限制。定时redis设置key虽然还在也无所谓 虚拟账户没报名不能下单
2 日赛检测 --- 不算复杂的是没有自由下单 这样就意味着没有那么多竞争少了些可能奇怪的事情。 因为启动时间都是90的整数倍 所以其实还是90s段段的 有虚拟账户结算的就是比赛信息
		让100个用户 分不同标的物去下单 每个标的物 一个个周期下单双对涨跌 不同用户下不同side不同金额 。
		另起一个程序定期去调用所有接口 返回打印日志
3 周赛检测 ---- 需要统计至少一天的日赛成绩 第二天改服务器时间 改starttime 重复日赛的操作就行
			没日赛的是0 有日赛的有成绩加上去就行

4 年赛检测 ---- 服务器时间连续改2次 报名耍年赛就行

5 注意事项 --- 最好晚上测 因为放便看结算 也方便统计周赛年赛
*/


type tradeResponse struct {
	baseResponse
	Tid     string  // 0 order id
	Bal     map[string]float64
	Pe  	float64
	Ht 		int64
	Ot   	int64
	Si 		int32
	Issue   int64
	Ts  	[]interface{}
}

type hpPingPong struct {
	Operation  string `json:"op"`  // sub  ping pong
}

type SubRequest struct {
	Operation  string `json:"op"`  // sub  unsub
	Token  	   string `json:"to"`
}

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}
type loginResponse struct {
	baseResponse
	Username string
	Token  string
	InvitationCode string
	Bal     map[string]float64
	Symbol  string
	cnName  string
}

func genValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	x := len(numeric)
	r.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ r.Intn(x) ])
	}
	return sb.String()
}

func genNumber(n int) int {
	r.Seed(time.Now().UnixNano())
	return r.Intn(n)
}

func reg(token string)  {
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/dayGame"
	req, _ := http.NewRequest("POST", url2, bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR----dayGame failed----err:%+v\n", err)
	}
	data := map[string]interface{}{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data)
	log.Printf("Daygame  ---- %+v\n",data)
	req.Body.Close()
}

var (
	amountList = []int64{20,30,40,50,60}

	trans = http.Transport{
		DisableKeepAlives:true,
	}
	client = &http.Client{
		Transport:&trans,
	}
)

func trade( token,symbol string) {
	side := 2+genNumber(3) //0 1 2 3 4对应金额 20 30 40 50 60 意味着2000块钱至少可以下30次 1800秒 半小时
	amount := amountList[genNumber(5)]

	z := map[string]interface{}{}
	z["am"] = amount
	z["si"] = side
	z["in"] = 60
	z["sy"] = symbol
	z["ts"] = time.Now().Unix()
	z["ve"] = "0.7.4"
	z["at"] = 0
	z["m"] = "centralism"

	o, _ := json.Marshal(z)
	var jsonStr3= []byte(o)
	url3 := "http://app-hpoption-web-test.azfaster.com:8081/tradeSDP"
	trareq, _ := http.NewRequest("POST", url3, bytes.NewBuffer(jsonStr3))
	trareq.Header.Set("Content-Type", "application/json")
	trareq.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	traderesp,err := client.Do(trareq)
	if err != nil {
		log.Printf("ERROR----trade failed----err:%+v\n", err)
	}
	tradedata := tradeResponse{}
	tradebody, _ := ioutil.ReadAll(traderesp.Body)
	json.Unmarshal([]byte(tradebody), &tradedata)
	log.Printf("trade::%+v::%+v\n",tradedata,side)
	trareq.Body.Close()
}

func printLog(token string) {
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/gameSingleUnFinRank"
	prireq, _ := http.NewRequest("POST", url2, bytes.NewBuffer([]byte{}))
	prireq.Header.Set("Content-Type", "application/json")
	prireq.Header.Set("Authorization",fmt.Sprintf("Bearer %s",token))
	resp, err := client.Do(prireq)
	if err != nil {
		log.Printf("ERROR----gameSingleUnFinRank failed----err:%+v\n", err)
	}
	data := map[string]interface{}{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data) //打印出邀请码 另用于注册账号 打印虚拟金额是否变化
	log.Printf("gameSingleUnFinRank::%+v\n",data)
	prireq.Body.Close()
}

func main() {

	//注册
	ph := "0105" + genValidateCode(10)
	x := map[string]string{}
	x["pn"] = ph
	x["pw"] = ph
	x["ic"] = "wzW2hg"
	m, _ := json.Marshal(x)
	var jsonStr= []byte(m)
	url := "http://app-hpoption-web-test.azfaster.com:8081/register"
	regreq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	regreq.Header.Set("Content-Type", "application/json")


	registrs,err := client.Do(regreq)
	if err != nil {
		log.Printf("ERROR----regist failed----err:%+v\n", err)
	}
	registdata := baseResponse{}
	registbody, _ := ioutil.ReadAll(registrs.Body)
	json.Unmarshal([]byte(registbody), &registdata)
	log.Printf("regist %+v,%+v\n",ph,registdata)
	regreq.Body.Close()
	log.Printf("================RegistFinished===============\n")



 	//登录
	y := map[string]string{}
	y["pn"] = ph
	y["pw"] = ph
	y["v"] = "0.7.4"
	n, _ := json.Marshal(y)
	var jsonStr2= []byte(n)
	url2 := "http://app-hpoption-web-test.azfaster.com:8081/loginByPassword"
	logreq, _ := http.NewRequest("POST", url2, bytes.NewBuffer(jsonStr2))
	logreq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(logreq)
	if err != nil {
		log.Printf("ERROR----login failed----err:%+v\n", err)
	}
	data := loginResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &data) //打印出邀请码 另用于注册账号 打印虚拟金额是否变化
	token := data.Token
	logreq.Body.Close()
	doneC := make(chan struct{})
	log.Printf("================LoginFinished===============\n")

	//ws
	/*
	go func() {
		endpoint := "ws://app-hpoption-ws-test.azfaster.com:55555/ws/BTCUSDT"
		hpdial := &websocket.Dialer{}
		wsConn, _, err := hpdial.Dial(endpoint, nil)
		if err != nil {
			log.Printf("ERROR----dial  ws failed----err:%+v\n", err)
			os.Exit(1)
		}


		subdata := SubRequest{Operation: "sub", Token: token}
		dataByte, _ := json.Marshal(subdata)
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(dataByte))
		if err != nil {
			log.Printf("ERROR----write sub failed----err:%+v\n", err)
			os.Exit(1)
		}

		log.Printf("TESTTEST=============3333=======\n")


		defer wsConn.Close()
		defer close(doneC)

		log.Printf("TESTTEST=============11111=======\n")
		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Printf("ERROR----read  message failed----err:%+v\n", err.Error())
				time.Sleep(time.Second * 1)
				return
			}
			if message != nil {
				var hpresp map[string]interface{}
				json.Unmarshal(message, &hpresp)
				log.Printf("look this whta %+v\n", hpresp)
				if v, ok := hpresp["op"]; ok && v.(string) == "ping" {
					log.Printf("INFO----receive ping\n")
					v, _ := json.Marshal(hpPingPong{Operation: "pong"})
					wsConn.WriteMessage(websocket.TextMessage, v)
				} else if v, ok := hpresp["op"]; ok && v.(string) == "BTCCenOddsNotify" {
					log.Printf("INFO----receive BTCCenOddsNotify\n")
					if hpresp["ct"] == 0 {
						btcmutex.Lock()
						btcupodds = hpresp["cuo"].(float64)
						btcdownodds = hpresp["cdo"].(float64)
						btcmutex.Unlock()
					}
				}
			}
			log.Printf("TESTTEST=============2222=======\n")
		}
	}()


	 */

	//报名
	reg(token)
	log.Printf("================GameRegFinished===============\n")

	//到固定的时间 不同标的物下不同的单 循环
	now := time.Now()
	st := time.Unix(1576218790,0)
	time.Sleep(st.Sub(now))


	//下单 不同标的物 下不同side 不同amount来制造差异 选定btc shci usdjpy 算了不然还要多个ws连接着没必要
	go func() {
		for {
			trade(token,"BTC")
			//trade(token,"SHCI")
			trade(token,"USDJPY")
			log.Printf("================TradeFinished===============\n")
			time.Sleep(time.Second*90)
		}
	}()

	<-doneC
}
