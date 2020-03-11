package main

import (
	"HPOptionServer/Common/CommonConf"
	"HPOptionServer/Common/HPSQL"
	"log"
)

func main() {
	pulltimeinfo := []HPSQL.Fypulltime{}
	HPSQL.GetTopByCondStruct(CommonConf.TableFypulltime,"Pulltime",2,map[string]interface{}{},&pulltimeinfo)
	log.Printf("%+v,%+v\n",pulltimeinfo[0],pulltimeinfo)
}
