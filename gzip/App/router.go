package App

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)


type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc //ServeHTTP
	Middleware mux.MiddlewareFunc
}

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}

var routes []Route

func ret2(w http.ResponseWriter, value interface{}) {
	resp, _ := json.Marshal(value)
	b64Str := base64.StdEncoding.EncodeToString([]byte(resp))

	buf := new(bytes.Buffer)
	wr := gzip.NewWriter(buf)

	leng, err := wr.Write([]byte(b64Str))
	if err != nil || leng == 0 {
		log.Printf("gzip failed err:%+v\n",err)
	}
	err = wr.Flush()
	if err != nil {
		log.Printf("gzip flush failed err:%+v\n",err)
	}
	err = wr.Close()
	if err != nil {
		log.Printf("gzip close failed err:%+v\n",err)
	}
	b := buf.Bytes()


	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Accept-Encoding", "gzip")
	w.Write([]byte(b))
}

func requestBegin2(req *http.Request)  {
	readBytes, _ := ioutil.ReadAll(req.Body)
	your_to_byte, _ := base64.StdEncoding.DecodeString(string(readBytes))
	reqStruct := map[string]interface{}{}
	json.Unmarshal([]byte(your_to_byte),&reqStruct)
	log.Printf("Info----Receive----%+v\n",reqStruct)
	/*
	your_to_byte, _ := base64.StdEncoding.DecodeString(string(readBytes))
	your_string := string(your_to_byte)
	log.Printf("string %+v\n",your_string)

	 */
	return
}

func helloworld(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	requestBegin2(req)
	ret2(w, baseResponse{Msg: "Success", Code: 0})
}

func init() {

	register("POST","/helloworld",helloworld,nil)


}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler) //其实就是等于标准库中的serverMux.Handle(pattern,Handle)
		}
	}
	return router
}
