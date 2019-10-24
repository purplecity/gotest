package main

import (
	"encoding/csv"
	"fmt"
	"gotest/csvtest/FileOperation"
	"os"
)

func main() {
	var filepath= "/Users/ludongdong/zaqizaba/btcusdt.csv"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:%+v\n", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"价格","时间"})
	w.Flush()

	var (
		filepath2 = "/Users/ludongdong/zaqizaba/BTCUSDT.txt"
		start = int64(0)
		end = int64(999)
		interval = 74580
		)

	for end < int64(interval) {
		hd := []FileOperation.HistoryData{}
		hd = FileOperation.ReadHistoryFile(filepath2,start,end)
		if len(hd) != 1000 {
			fmt.Println("err not 1000")
			return
		}
		end += 1000
		start += 1000

		for _,x := range hd {
			w.Write([]string{fmt.Sprintf("%+v",x.P),fmt.Sprintf("%+v",x.T)})
			w.Flush()
		}
	}



}


