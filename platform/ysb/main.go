package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

  //电脑版测试环境 ysb体育竞猜大厅

/*
<?xml version='1.0' encoding='utf-8'?>
<request action='clogin'>
    <element id='id001'>
        <properties name='UN'>A1GO_ysb020</properties>
        <properties name='VID'>A1GO</properties>
		<properties	name='SG'>002f974ec1381a9e609a982a938690d7</properties>
        <properties name='SN'>92645ddcdd04c9ab2148b65ba97101c</properties>
		<properties name ="OPRID">A1GO_ysb</properties>
    </element>
</request>
*/

type reqproperties struct {
	Action string `xml:"action,attr"` // 读取flag属性
	Element reqele `xml:"element"`
}

type reqele struct {
	Id string `xml:"id,attr"` // 读取flag属性
	Pro []reqpp `xml:"properties"`
}

type reqpp struct {
	Name string `xml:"name,attr"` // 读取flag属性
	Proterties string `xml:",chardata"`
}


type response struct  {
	Action string `xml:"action,attr"`
	Ele respYSBELE `xml:"element"`
}

type respYSBELE struct {
	Id string `xml:"id,attr"` // 读取flag属性
	Pro []respYSBpp `xml:"properties"`
}

type respYSBpp struct {
	Name string `xml:"name,attr"` // 读取flag属性
	Proterties interface{} `xml:",chardata"`
}

func YSBLoginValidate(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	con, _ := ioutil.ReadAll(req.Body) //获取post的数据
	fmt.Println(string(con))
	test1 := reqproperties{}
	err := xml.Unmarshal(con,&test1)
	if err != nil {
		log.Printf("%+v\n",err)
	}
	log.Printf("%+v\n",test1)
	//暂时不做校验 直接成功 返回

	resp1map := map[string]interface{}{}
	for _,v := range test1.Element.Pro {
		resp1map[v.Name] = v.Proterties
	}
	log.Printf("resp1map :%+v\n",resp1map)


	hresp := response{Action:"clogin"}
	ppList := []respYSBpp{}
	ppList = append(ppList,respYSBpp{Name:"UN",Proterties:resp1map["UN"]})
	ppList = append(ppList,respYSBpp{Name:"UID",Proterties:resp1map["SG"]})
	ppList = append(ppList,respYSBpp{Name:"S",Proterties:0})
	ppList = append(ppList,respYSBpp{Name:"CC",Proterties:0})
	ppList = append(ppList,respYSBpp{Name:"ED",Proterties:""})

	hele := respYSBELE{Id:"id001",Pro:ppList}
	hresp.Ele = hele

	output,err:=xml.Marshal(hresp)
	if err != nil {
		log.Printf("marshal xml error: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	headerBytes := []byte(xml.Header)
	headerBytes = append(headerBytes,output...)
	w.Write(headerBytes)
}

/*
<?xml version='1.0' encoding='utf-8'?><request action='clogin'><element id='id001'><properties name='UN'>XY_666666</properties><properties name='VID'>XY</properties><properties name='SG'>666666</properties><properties name='SN'>1d5efb00-f43d-4a4d-83b5-1e33b1922306</properties><properties name='OPRID'></properties></element></request>
2020/03/02 18:09:03 {Action:clogin Element:{Id:id001 Pro:[{Name:UN Proterties:XY_666666} {Name:VID Proterties:XY} {Name:SG Proterties:666666} {Name:SN Proterties:1d5efb00-f43d-4a4d-83b5-1e33b1922306} {Name:OPRID Proterties:}]}}
2020/03/02 18:09:03 resp1map :map[OPRID: SG:666666 SN:1d5efb00-f43d-4a4d-83b5-1e33b1922306 UN:XY_666666 VID:XY]
*/

func main() {
	http.HandleFunc("/", YSBLoginValidate) //设置访问的路由
	err := http.ListenAndServe(":8889", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}