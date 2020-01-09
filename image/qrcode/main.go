package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)


func main() {
	err := qrcode.WriteFile("https://app-hpoption-download.azfaster.com?code=HPOPTIONcode33333", qrcode.Medium, 280, "/Users/ludongdong/go/src/gotest/image/qr.png")
	if err != nil {
		fmt.Printf("failed ==== %+v\n",err)
	}

}
