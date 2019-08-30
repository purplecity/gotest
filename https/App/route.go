package App

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc //ServeHTTP
	Middleware mux.MiddlewareFunc
}

var routes []Route


func sendSMS(w http.ResponseWriter, req *http.Request) {
	fmt.Println("receiv msg")
	defer req.Body.Close()
	v := map[string]interface{}{"Code":0,"Msg":"success"}
	resp, _ := json.Marshal(v)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(resp)
}

func init() {
	register("POST", "/sendSMS",sendSMS,nil)
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