package main

import (
	"fmt"
	"log"
)

func test() {
	log.Panicln("-----pctest")
}
func main() {
 var a int32 = 100
 var b float64 = 0.03
 fmt.Printf("%T,%v\n",float64(a)*b,float64(a)*b)
}