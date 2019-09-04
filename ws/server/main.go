

package main

import (
	"github.com/gorilla/mux"
	"gotest/ws/server/handle"
	"log"
	"net/http"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws/BTCUSDT", handle.Hpws).Methods("GET")
	log.Fatal(http.ListenAndServe("127.0.0.1:55555",router))
}