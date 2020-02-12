package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

/*
a) 1 = GetBalance(玩家余额请求),
b) 2 = Deposit(充值请求),
c) 3 = Confirm Deposit(充值确定回复),
d) 4 = Withdraw(兑换请求)
*/

/*
<?xml version='1.0' encoding='utf8'?>
<request action="EBdeposit">
    <element id="100440318">
        <properties name="UID">ysb020</properties>
        <properties name="UN">A1GO _ ysb020</properties>
        <properties name="SN">99999999</properties>
        <properties name="VID">A1GO</properties>
        <properties name="CI">RMB></properties>
        <properties name="LI">1</properties>
        <properties name="AMT">10</properties>
        <properties name="RN">0</properties>
        <properties name="OPRID">[OPRID]</properties>
    </element>
</request>
*/

//充值第一 3步
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

//充值第二部回复


var baseurl = "http://testsportapi.a1sport88.com/XMLExchange.aspx?TrancType=2&ThirdPartyID=XY&CURID=RMB"
func main () {

	hreq := request{Action:"EBdeposit"}
	ppList := []reqYSBpp{}
	ppList = append(ppList,reqYSBpp{Name:"UN",Proterties:"55555"})
	ppList = append(ppList,reqYSBpp{Name:"UID",Proterties:"55555"})
	ppList = append(ppList,reqYSBpp{Name:"SN",Proterties:55555})
	ppList = append(ppList,reqYSBpp{Name:"VID",Proterties:"XY"})
	ppList = append(ppList,reqYSBpp{Name:"CI",Proterties:"RMB"})
	ppList = append(ppList,reqYSBpp{Name:"LI",Proterties:1})
	ppList = append(ppList,reqYSBpp{Name:"AMT",Proterties:500})
	ppList = append(ppList,reqYSBpp{Name:"RN",Proterties:"55555"})
	ppList = append(ppList,reqYSBpp{Name:"OPRID",Proterties:0})  //要问清的

	hele := reqYSBELE{Id:"100440318",Pro:ppList}
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
		log.Printf("ERROR----chongzhi 1 req----err:%+v\n", err)
	}

	defer rn.Body.Close()
	defer resp1.Body.Close()

	con, _ := ioutil.ReadAll(resp1.Body)
	test1 := respproperties{}
	err = xml.Unmarshal(con,&test1)
	if err != nil {
		log.Printf("ERROR---chongzhi 1 resp----err:%+v\n", err)
	}

	resp1map := map[string]interface{}{}
	for _,v := range test1.Element.Pro {
		resp1map[v.Name] = v.Proterties
	}
	log.Printf("resp1map :%+v\n",resp1map)



	hreq2 := request{Action:"EBdepositConfirm"}
	ppList2 := []reqYSBpp{}
	ppList2 = append(ppList,reqYSBpp{Name:"UN",Proterties:"55555"})
	ppList2 = append(ppList,reqYSBpp{Name:"ST",Proterties:0})
	ppList2 = append(ppList,reqYSBpp{Name:"SN",Proterties:55555})
	ppList2 = append(ppList,reqYSBpp{Name:"PI",Proterties:resp1map["PI"]})
	ppList2 = append(ppList,reqYSBpp{Name:"ED",Proterties:"-"}) //下划线还是中线

	hele2 := reqYSBELE{Id:"100440318",Pro:ppList2}
	hreq2.Ele = hele2

	jsonStr2,err :=xml.Marshal(hreq2)
	rn2, _ := http.NewRequest("POST", baseurl, bytes.NewBuffer(jsonStr2))
	rn2.Header.Set("Content-Type", "text/xml")

	resp2,err := client.Do(rn)
	if err != nil {
		log.Printf("ERROR----chongzhi 2 req----err:%+v\n", err)
	}


	con2, _ := ioutil.ReadAll(resp2.Body)
	test2 := respproperties{}
	err = xml.Unmarshal(con2,&test2)
	if err != nil {
		log.Printf("ERROR---chongzhi 2 resp----err:%+v\n", err)
	}

	resp2map := map[string]interface{}{}
	for _,v := range test2.Element.Pro {
		resp2map[v.Name] = v.Proterties
	}
	log.Printf("resp2map :%+v\n",resp2map)

}
