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

var (
	trans = http.Transport{
		DisableKeepAlives:true,
	}
	client = &http.Client{
		Transport:&trans,
	}
)

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
	enName  string
	Bgm     int
	Butttonsound int
	Tradehint    int
	Index   interface{}
	IndexList interface{}
	Mode   	string
}



func main() {
	//注册
	prefixlist := []string{"131hp","132hp","133hp","134hp","135hp","136hp","137hp","138hp","139hp"}
	//"151hp","152hp","153hp","154hp","155hp","156hp","157hp","158hp","159hp","181hp","182hp"
	for _,x := range prefixlist {
		for i:=1;i<=50;i++ {
			ivcode := "g4QUpE" //我们自己的邀请码
			ph := x + genValidateCode(6)
			regmap := map[string]string{}
			regmap["pn"] = ph
			regmap["pw"] = ph
			regmap["ic"] = ivcode
			regbina, _ := json.Marshal(regmap)
			var jsonStr= []byte(regbina)
			url := "https://app-hpoption-web.azfaster.com:8081/register"
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
			time.Sleep(time.Millisecond*500)
		}
	}



	//已经注册查询
	//登录
	//报名
	//每天玩一下

}

// 0 必须下单 否则用户发奖排名不一致
//1  不能跟已有账号重合 必然不重合 反正没关系
//2  既然1000个报名 那么至少得制造50个上榜成绩 不然说不过去。 程序别太复杂 具体来说就是别跑太久别搞的进程太多 然后只要保证就算没有人玩也有50个上榜就行 也就是说 每天日赛和年赛只要有50个有成绩就行
//3 做成定时任务 去操作  放在别的机器上比如国内那台阿里云
//4 人数上限设置为500比较合适 如果真的达到500个报名真实人数那也就可以了
//5 设置20成绩8人 50 100 200 500 1000 2000成绩各7人 这个各个成绩人数可配置 因为正反下单 所以需要100个人下单保证ok 根据复杂度安排每轮下单人数 顶多50轮*90s 75分钟完事
//6 为了不造成每次 机器人总是上限给人造成破绽 还可以实际注册1000人每次从1000里面选1000人 但是实际好像没太必要 但是这个问题不是很大 可能没必要 就相当于常胜将军也没事
//7 以命令行参数的形式去启动  参数包括手机号 side 金额 时间这个4个 这样一个程序通过命令行跑多少个进程就行  就只用日赛和年赛 写进定时任务这种 如果错了要把错误手机号打印 然后要把所有成功的手机号打印 甚至直接2000下单