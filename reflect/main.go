package main

import (
	"fmt"
	"reflect"

)

func main() {
	x := reflect.TypeOf(3)
	fmt.Println(x.Kind(),x.String())
	fmt.Printf("%+T\n",x)
	v := reflect.ValueOf(3)
	t := v.Type()
	fmt.Println(t.Kind(),t.String())
	fmt.Printf("%+T\n",v)

}
