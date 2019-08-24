package main

import (
	"fmt"
	"math"
	"strconv"
)

func test(x string) bool {
	return x
}

func main() {
	fmt.Println(strconv.ParseFloat(fmt.Sprintf("%.3f", math.Abs(5.666-3.44)/5.666), 64))
}