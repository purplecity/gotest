package FileOperation

import (
	"encoding/binary"
	"log"
	"math"
	"os"
)

type HistoryData struct {
	P float64	`json:"hp"`
	C int64		`json:"hc"`
	T int64		`json:"ht"`
}

func WriteHistoryFile(sy string,price float64, count,ts int64) error {
	f, err := os.OpenFile(sy,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	defer f.Close()
	if err != nil {
		log.Panicf("ERROR----OpenFile failed----err:%v\n",err)
	}
	p := Float64ToByte(price)
	c := Int64ToBytes(count)
	t := Int64ToBytes(ts)
	_, err = f.Write(p)
	_,err = f.Write(c)
	_, err = f.Write(t)
	return err
}

func ReadHistoryFile(sy string,start,end int64)  []HistoryData {
	f, err := os.OpenFile(sy,os.O_CREATE|os.O_RDONLY,0666)
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
