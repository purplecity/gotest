package  main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//url := "https://webrates.truefx.com/rates/connect.html?u=jsTrader&p=anystring&q=ozrates&c=AUD/USD,USD/JPY&f=html&s=n"
	url := "https://webrates.truefx.com/rates/connect.html?p=anystring&c=EUR/USD&s=n"
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
	/*
		rs := map[string]interface{}{}
		json.Unmarshal(rs)
		fmt.Printf("%+T",body)

	*/
}
