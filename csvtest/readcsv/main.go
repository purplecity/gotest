package main

import (
	"encoding/csv"
	"fmt"
	"gotest/csvtest/Operation"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("/Users/ludongdong/zaqizaba/eurusd.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat[:])))

	record, err := r.Read()
	if err == io.EOF {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n",record)

	sum := float64(0)
	count := int64(0)
	for   {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		a := strings.Split(record[0],"\n")
		x,_ := strconv.ParseFloat(a[0],64)
		sum = Operation.HPAdd(sum,x)
		count ++
	}
	fmt.Println(sum,count)
	fmt.Println(Operation.HPDivInt(sum,count))
}