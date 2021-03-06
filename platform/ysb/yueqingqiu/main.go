package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type request struct  {
	Action string `xml:"action,attr"`
	Ele reqYSBELE `xml:"element"`
}

type reqYSBELE struct {
	Id string `xml:"id,attr"` // 读取flag属性
	Pro []reqYSBpp `xml:"properties"`
}

type reqYSBpp struct {
	Name string `xml:"name,attr"` // 读取flag属性
	Proterties interface{} `xml:",chardata"`
}

//充值第一步回复
type respproperties struct {
	Action string `xml:"action,attr"` // 读取flag属性
	Element respele `xml:"element"`
}

type respele struct {
	Id string `xml:"id,attr"` // 读取flag属性
	Pro []resppp `xml:"properties"`
}

type resppp struct {
	Name string `xml:"name,attr"` // 读取flag属性
	Proterties string `xml:",chardata"`
}

var baseurl = "http://testsportapi.a1sport88.com/XMLExchange.aspx?TrancType=1&ThirdPartyID=XY&CURID=RMB"

func main() {
	hreq := request{Action:"EBaccountbalance"}
	ppList := []reqYSBpp{}
	ppList = append(ppList,reqYSBpp{Name:"UN",Proterties:"XY_666666"})
	ppList = append(ppList,reqYSBpp{Name:"SN",Proterties:6666668})

	hele := reqYSBELE{Id:"100440328",Pro:ppList}
	hreq.Ele = hele

	jsonStr,err:=xml.Marshal(hreq)
	rn, _ := http.NewRequest("POST", baseurl, bytes.NewBuffer(jsonStr))
	rn.Header.Set("Content-Type", "text/xml")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	resp1,err := client.Do(rn)
	if err != nil {
		log.Printf("ERROR----yue 1 req----err:%+v\n", err)
	}

	defer rn.Body.Close()
	defer resp1.Body.Close()

	con, _ := ioutil.ReadAll(resp1.Body)
	test1 := respproperties{}
	err = xml.Unmarshal(con,&test1)
	if err != nil {
		log.Printf("ERROR---yue 1 resp----err:%+v\n", err)
	}

	resp1map := map[string]interface{}{}
	for _,v := range test1.Element.Pro {
		resp1map[v.Name] = v.Proterties
	}
	log.Printf("resp1map :%+v\n",resp1map)
}