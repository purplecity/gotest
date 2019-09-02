package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"strconv"
)

func main() {

	//price := float64(0)

	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"



	c.OnRequest(func(r *colly.Request) {

		//fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))

	})
	c.OnHTML("div[id=price]", func(e *colly.HTMLElement) {
		fmt.Printf("test----%+v\n",e.Text)
		price,err := strconv.ParseFloat(e.Text,64)
		fmt.Printf("********* price----%+v\n",price)
		if err != nil {

			fmt.Printf("parse err---%+v\n",err)

		}
	})

	c.OnScraped(func(_ *colly.Response) {



		fmt.Println("OK")



	})



	err := c.Visit("https://finance.sina.com.cn/realstock/company/sh000001/nc.shtml")

	if err != nil {

		fmt.Println("visit err")

	}
}