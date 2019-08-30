package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"



	c.OnRequest(func(r *colly.Request) {

		fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))

	})
	c.OnHTML("<span."first_span col_green" '="">2886.24</span>)
}