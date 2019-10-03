package  main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	//url := "https://webrates.truefx.com/rates/connect.html?u=jsTrader&p=anystring&q=ozrates&c=AUD/USD,USD/JPY&f=html&s=n"
	t1 := time.Now().UnixNano()
	url := "https://webrates.truefx.com/rates/connect.html?p=anystring&c=EUR/USD&f=csv&s=n"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%+v\n",err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v\n",err)
	}
	fmt.Println(string(body))
	a := strings.Split(string(body),",")
	fmt.Println(a,len(a))
	b,_ := strconv.ParseFloat(a[2],64)
	fmt.Printf("%+T,%+v\n",b,b)

	/*
		rs := map[string]interface{}{}
		json.Unmarshal(rs)
		fmt.Printf("%+T",body)

	*/
	t2 := time.Now().UnixNano()
	fmt.Println(t2/1e6-t1/1e6)
}
