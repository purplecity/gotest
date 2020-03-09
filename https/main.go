package main

import (
	"log"
	"net/http"
	"gotest/https/App"
)

var(
PemFile = "/root/htts/full_chain_web.pem"
KeyFile = "/root/https/private_web.key"

)


/*
func init() {
	logfile, err := os.OpenFile("/root/go/src/test/https/https.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open log file failed")
		os.Exit(1)
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func clearUnfinishedOrder() {

}

 */

func main() {
	router := App.NewRouter() // router object has ServeHTTP method
	//log.Fatal(http.ListenAndServe("0.0.0.0:8888",router))
	//log.Fatal(http.ListenAndServeTLS("127.0.0.1:8888",PemFile,KeyFile,router))
}