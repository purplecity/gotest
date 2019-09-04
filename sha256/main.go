package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	waitString := `account_type=6&app_id=8c87527e6c494e04b890fad1791a746c&app_user=tttt&coin_type=107&out_trade_sn=55555&quantity=50`

	h := hmac.New(sha256.New, []byte("e1b6aeb75a71418ab10c04191d37a754"))
	h.Write([]byte(waitString))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha)


}


