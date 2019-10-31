package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "gotest/testsql/mysql"
)

func main() {
	var list orm.ParamsList
	o := orm.NewOrm()
	num, err := o.Raw("SELECT DISTINCT(contributorid) FROM scorerecord WHERE contributorid != 1186469846561406976").ValuesFlat(&list)
	fmt.Printf("%+v\n",list)
	if err == nil && num > 0 {
		fmt.Println(list) // []{"1","2","3",...}
	}
	for _,x := range list {
		fmt.Println(x)
	}
}