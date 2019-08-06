package main

import "fmt"

var a = []interface{}{}
var b = []int{1,2,3,4,5,6}

func main() {
	c := [10]map[string]interface{}{}
	for index,v := range b {
		c
	}
	fmt.Printf("%+v\n",a[0:len(a)])

}

