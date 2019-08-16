package main

import (
	"fmt"
	"gotest/sql/mysql"
)

func  main() {
	rs :=  []mysql.Realtrade{}
	//mysql.GetAllRecord("Realtrade",map[string]interface{}{"Uid":"1161953288443645952"},&rs)
	mysql.GetGroupOneByCondList("Realtrade","Symbol","Settletime",map[string]interface{}{"Uid":"1161953288443645952"},&rs)
	fmt.Printf("%+v,%+v\n",len(rs),rs)
}
