package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}

func main() {
	for i:=0;i<=10000;i++ {
		ph := "8"+genValidateCode(10)
		x := map[string]string{}
		x["pn"] = ph
		x["pw"] = ph
		x["ic"] = "8ykZqi"
		y,_ := json.Marshal(x)
		var jsonStr = []byte(y)

		url := "http://47.244.212.51:8888/register"

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {

			panic(err)

		}
		reqStruct := baseResponse{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(body),&reqStruct)
		fmt.Printf("%s,%s,%+v,response Status:%v\n",ph,time.Now(),reqStruct,resp.Status,)


		defer resp.Body.Close()
		time.Sleep(time.Duration(10)*time.Millisecond)
	}
}
