package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	ageList := []string{"account_type", "app_id", "app_user", "coin_type", "out_trade_sn", "quantity"}

	//排序，实现比较方法即可
	sort.Slice(ageList, func(i, j int) bool {
		return ageList[i] < ageList[j]
	})
	fmt.Printf("after sort:%v\n", ageList)
	waitString := fmt.Sprintf("account_type=%+v&app_id=%+v&app_user=%+v&coin_type=%+v&out_trade_sn=%+v&quantity=%+v",
		2,"771e22a66bd74ff3ba574fc9cdc760be","appusertest",107,"sn_order_c982807294421",50,)
	//waitString = `"account_type"=2&"app_id"="771e22a66bd74ff3ba574fc9cdc760be"&"app_user"="appusertest"&"coin_type"="107"&"out_trade_sn"="sn_order_c982807294421"&"quantity"="50"`
	waitString = `account_type=2&app_id=771e22a66bd74ff3ba574fc9cdc760be&app_user=appusertest&coin_type=107&out_trade_sn=sn_order_c982807294421&quantity=50`

	h := hmac.New(sha256.New, []byte("d5f68ea66bd74ff3ba574fc9cdc12345"))
	h.Write([]byte(waitString))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(sha)))

}


