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
	prefixlist := []string{"131hp","132hp","133hp","134hp","135hp","136hp","137hp","138hp","139hp",
	"151hp","152hp","153hp","154hp","155hp","156hp","157hp","158hp","159hp","181hp","182hp"}
	phlist := []string{}
	for _,x := range prefixlist {
		for i:=1;i<=5;i++ {
			ivcode := "7nvHW3" //我们自己的邀请码
			ph := x + genValidateCode(6)
			regmap := map[string]string{}
			regmap["pn"] = ph
			regmap["pw"] = ph
			regmap["ic"] = ivcode
			regbina, _ := json.Marshal(regmap)
			var jsonStr= []byte(regbina)
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
			phlist = append(phlist,fmt.Sprintf("\"%+v\"",ph))
			time.Sleep(time.Millisecond*500)
		}
	}

	fmt.Printf("%+v\n",phlist)

	//已经注册查询
	//登录
	//报名
	//每天玩一下
}

/*
["131hp109129" "131hp812713" "131hp858503" "131hp273928" "131hp461275" "132hp084686" "132hp481077" "132hp839996" "132hp082105" "132hp664886" "133hp567607" "133hp512998" "133hp517997" "133hp763788" "133hp357678" "134hp179002" "134hp003749" "134hp665195" "134hp590498" "134hp800339" "135hp308503" "135hp604499" "135hp136519" "135hp825924" "135hp740706" "136hp650721" "136hp640962" "136hp229471" "136hp572097" "136hp687956" "137hp915871" "137hp395729" "137hp580089" "137hp589130" "137hp727018" "138hp754660" "138hp712811" "138hp935371" "138hp908154" "138hp358278" "139hp998789" "139hp224581" "139hp340287" "139hp891518" "139hp145515" "151hp815810" "151hp033290" "151hp180955" "151hp461551" "151hp994346" "152hp885874" "152hp876406" "152hp401093" "152hp964674" "152hp975522" "153hp770807" "153hp712162" "153hp862930" "153hp265629" "153hp004381" "154hp335936" "154hp822785" "154hp024429" "154hp297032" "154hp499050" "155hp580724" "155hp506537" "155hp565016" "155hp818594" "155hp672168" "156hp524615" "156hp503365" "156hp602317" "156hp458618" "156hp257574" "157hp679148" "157hp128119" "157hp805353" "157hp524117" "157hp015194" "158hp255859" "158hp110111" "158hp133611" "158hp304696" "158hp195730" "159hp957293" "159hp817809" "159hp283192" "159hp591958" "159hp813940" "181hp964502" "181hp060121" "181hp918588" "181hp314646" "181hp915317" "182hp551132" "182hp060482" "182hp327878" "182hp367022" "182hp110205"]
*/

// 0 必须下单 否则用户发奖排名不一致
//1  不能跟已有账号重合 必然不重合 反正没关系
//2  既然1000个报名 那么至少得制造50个上榜成绩 不然说不过去。 程序别太复杂 具体来说就是别跑太久别搞的进程太多 然后只要保证就算没有人玩也有50个上榜就行 也就是说 每天日赛和年赛只要有50个有成绩就行
//3 做成定时任务 去操作  放在别的机器上比如国内那台阿里云
//4 人数上限设置为500比较合适 如果真的达到500个报名真实人数那也就可以了
//5 设置20成绩8人 50 100 200 500 1000 2000成绩各7人 这个各个成绩人数可配置 因为正反下单 所以需要100个人下单保证ok 根据复杂度安排每轮下单人数 顶多50轮*90s 75分钟完事
//6 为了不造成每次 机器人总是上限给人造成破绽 还可以实际注册1000人每次从1000里面选1000人 但是实际好像没太必要 但是这个问题不是很大 可能没必要 就相当于常胜将军也没事
//7 以命令行参数的形式去启动  参数包括手机号 side 金额 时间这个4个 这样一个程序通过命令行跑多少个进程就行  就只用日赛和年赛 写进定时任务这种 如果错了要把错误手机号打印 然后要把所有成功的手机号打印 甚至直接2000下单