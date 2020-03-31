package main


import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

var HPTokenSecret = "HPOPTION1688"

func main() {
	tokenstr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODc3MjI5OTUsImh0eXBlIjoiUGxheWVyIiwidWlkIjoiMTI0MDE1NDEwMDk2NTI1MzEyMCJ9.A2wDoanF_DkRrapV3l5dw2QwMPyoclrusvB1PsVKoJk"
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{},error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(HPTokenSecret),nil
	})
	if err != nil {
		log.Printf("ERROR-----Token invalid--- %v",err)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["uid"].(string))
}
