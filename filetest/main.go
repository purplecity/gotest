package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
)


var HistoryFilePath = "/root/hd/SHCI.txt"

type HistoryData struct {
	P float64	`json:"hp"`
	C int64		`json:"hc"`
	T int64		`json:"ht"`
}


func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func main() {
	end :=  int64(57420)  //1566869460
	start := int64(53820)
	f, err := os.OpenFile(HistoryFilePath,os.O_CREATE|os.O_RDONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}
	buf := make([]byte, (end-start+1)*24)

	n, err :=f.ReadAt(buf,start*24)
	if err != nil {
		log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
	}

	var m int64= 0
	var datalist []HistoryData

	for m < (end-start+1)*24 {
		price := ByteToFloat64(buf[m:m+8])
		cnt := BytesToInt64(buf[m+8:m+16])
		ts := BytesToInt64(buf[m+16:m+24])
		ds := HistoryData{P:price,C:cnt,T:ts}
		datalist = append(datalist,ds)
		m += 24
	}
	for _,x := range datalist {
		fmt.Printf("%+v\n",x)
	}

}


/*
var HistoryFilePath = "/Users/ludongdong/zaqizaba/BTCUSDT.txt"

type HistoryData struct {
	P float64	`json:"hp"`
	C int64		`json:"hc"`
	T int64		`json:"ht"`
}

func WriteHistoryFile(price float64, count,ts int64) {
	f, err := os.OpenFile(HistoryFilePath,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}
	p := Float64ToByte(price)
	c := Int64ToBytes(count)
	t := Int64ToBytes(ts)
	f.Write(p)
	f.Write(c)
	f.Write(t)
}

func ReadHistoryFile(start,end int64)  []HistoryData {
	f, err := os.OpenFile(HistoryFilePath,os.O_CREATE|os.O_RDONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}
	buf := make([]byte, (end-start+1)*24)

	n, err :=f.ReadAt(buf,start*24)
	if err != nil {
		log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
	}

	var m int64= 0
	var datalist []HistoryData

	for m < (end-start+1)*24 {
		price := ByteToFloat64(buf[m:m+8])
		cnt := BytesToInt64(buf[m+8:m+16])
		ts := BytesToInt64(buf[m+16:m+24])
		ds := HistoryData{P:price,C:cnt,T:ts}
		datalist = append(datalist,ds)
		m += 24
	}
	return datalist
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}


func testReadHistoryFile(cnt int64)  (price float64){
	f, err := os.OpenFile(HistoryFilePath,os.O_CREATE|os.O_RDONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}

	buf := make([]byte, 24)

	n, err :=f.ReadAt(buf,cnt*24)
	if err != nil {
		log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
	}

	price = ByteToFloat64(buf[0:8])
	fmt.Println(price)
	return
}

func main() {

	for i:=1;i<100;i++{
		go WriteHistoryFile(1.34,1,1112)
		go testReadHistoryFile(5)
		time.Sleep(time.Second)
	}

	done := make(chan struct{})
	<-done
	/*
	f, err := os.OpenFile("/root/hd/BTCUSDT.txt",os.O_CREATE|os.O_RDONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}

	buf := make([]byte, 24)

	n, err :=f.ReadAt(buf,166*24)
	if err != nil {
		log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
	}

	price := ByteToFloat64(buf[0:8])
	cnt := BytesToInt64(buf[8:16])
	ts := BytesToInt64(buf[16:24])
	fmt.Printf("%+v,%+v,%+v\n",price,cnt,ts)
	return

	/*
		WriteHistoryFile(1.33,0,1111)
	WriteHistoryFile(1.34,1,1112)
	WriteHistoryFile(1.35,2,1113)
	WriteHistoryFile(1.36,3,1114)
	f, err := os.OpenFile(HistoryFilePath,os.O_CREATE|os.O_RDONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}
	//跳转到文本中的某处，并返回此处的偏移量 读取或写入会改变偏移
	offset,_:= f.Seek(0,2)
	fmt.Printf("PCTEST----%+v\n",offset)
	buf := make([]byte, 24)
	n, err :=f.ReadAt(buf,offset-24)
	if err != nil {
		log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
	}
	p := ByteToFloat64(buf[0:8])
	var sp float64
	if offset > 24*24*60*60  {
		n, err :=f.ReadAt(buf,offset-24*60*60*60)
		if err != nil {
			log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
		}
		sp = ByteToFloat64(buf[0:8])
	} else {
		n, err :=f.ReadAt(buf,0)
		if err != nil {
			log.Printf("ERROR----read byte:%v----err:%v\n",n,err)
		}
		sp = ByteToFloat64(buf[0:8])
	}
	fmt.Printf("PCTEST---sp:%+v\n",sp)
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Abs(sp-p)/sp), 64)
	fmt.Printf("%+v\n",v)
	*/

	//WriteHistoryFile(1.35,1113)
	//[91 123 34 80 34 58 49 46 51 53 44 34 84 34 58 49 49 49 51 125 93]
	//ReadHistoryFile(1,2)
	/*
		var hd = historyData{P:1.333,T:12222}
		data,_ := json.Marshal(hd)
		fmt.Println(unsafe.Sizeof(data))
	*/

	/*
		var testData []historyData
		testData = append(testData,historyData{1.33,1111},historyData{1.34,1112},historyData{1.35,1113})
		x,_ := json.Marshal(testData)
		fmt.Println(x)
	*/
	//[91 123 34 80 34 58 49 46 51 51 44 34 84 34 58 49 49 49 49 125 44 123 34 80 34 58 49 46 51 52 44 34 84 34 58 49 49 49 50 125 93]

	//fmt.Println(byte(91))
