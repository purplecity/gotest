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

	var x = map[string]interface{}{}
	json.Unmarshal([]byte(data),&x)
	fmt.Printf("%+v\n",x["compilerOptions"].(map[string]interface{})["tt"].(float64))

}
