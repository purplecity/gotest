package pac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type tradeInfo struct {
	UserId 		string		`json:u`
	Amount 		float64		`json:a`
	Side   		int			`json:s`
	Interval	int			`json:i`
	Price		float64		`json:p`
}

func main() {
	//http.HandleFunc("/trade",aaa)
	http.HandleFunc("/trade",bbb)
	log.Fatal(http.ListenAndServe("127.0.0.1:8888",nil))
}

//两种方式都行但是 但是都不能实现 如果req的json body是tradeInfo类型的标签键值对的时候 没能写入结构体变量

func aaa(w http.ResponseWriter, req *http.Request)  {
	a, _ := ioutil.ReadAll(req.Body)
	//var tInfo map[string]interface{}
	var tInfo tradeInfo
	err := json.Unmarshal([]byte(a),&tInfo)
	if err != nil {
		fmt.Println("error unmarshal")
	}
	fmt.Println(tInfo.Price)
}


func bbb(w http.ResponseWriter, req *http.Request)  {
	//var tInfo map[string]interface{}
	var tInfo tradeInfo
	json.NewDecoder(req.Body).Decode(&tInfo)
	fmt.Printf("%T",tInfo.Price)

}

