package main

import (
	"fmt"
	"gotest/sql/mysql"
)

func main() {
	userlist := []mysql.AdminUsers{}
	mysql.GetAllRecord("AdminUsers", map[string]interface{}{"Phonenumber__startswith":"2"},&userlist)
	for _,x := range userlist {
		tradeInfo := []mysql.Realtrade{}
		mysql.GetAllRecord("Realtrade", map[string]interface{}{"Uid":x.Uid,"Orderresult":1},&tradeInfo)
		fmt.Println(len(tradeInfo))
		for _,y := range tradeInfo {
			fmt.Printf("%+v\n",y)
		}

	}
}
