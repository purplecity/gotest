package main

import (
	"gotest/gzip/App"
	"log"
	"net/http"
)

func main() {
	router := App.NewRouter() // router object has ServeHTTP method
	log.Fatal(http.ListenAndServe("127.0.0.1:8888",router))
	//log.Fatal(http.ListenAndServeTLS("127.0.0.1:8888",CommonConf.PemFile,CommonConf.KeyFile,router))
}
