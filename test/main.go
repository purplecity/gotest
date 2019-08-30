package main

import (
	"fmt"
	"math/rand"
)



func main() {
	for i:=0;i<=100;i++{
		y := rand.Intn(9)
		fmt.Println(y)
	}

}