package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func testfile() {
	time.Sleep(time.Second*1)
	fmt.Println("hehe")
}

func main() {
	go func() {
		f, err := os.Create("/Users/ludongdong/Downloads/cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}()
	testfile()
}
