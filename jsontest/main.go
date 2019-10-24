package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main()  {

	data, err := ioutil.ReadFile("/Users/ludongdong/go/src/gotest/jsontest/tsconfig.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}

	var x = []interface{}{}
	json.Unmarshal([]byte(data),&x)
	for _,m := range x {
		fmt.Printf("%+v\n",m.(map[string]interface{})["content"])
	}
}
