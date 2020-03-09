package  main

import (
	"fmt"
	"log"
	"net/http"
)

type baseResponse struct {
	Code	uint	   // 0 success  others error
	Msg 	string  // success  errorMsg
}

/*
func main() {
	url := "https://app-hpoption-webapi.azfaster.com:8081/loginByPassword"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("hpoption", "1688")
	trans := http.Transport{
		DisableKeepAlives:true,
	}
	client := &http.Client{
		Transport:&trans,
	}
	registrs,err := client.Do(req)
	if err != nil {
		log.Printf("ERROR----regist failed----err:%+v\n", err)
	}
	body, err := gzip.NewReader(registrs.Body)
	registbody, _ := ioutil.ReadAll(body)
	your_to_byte, _ := base64.StdEncoding.DecodeString(string(registbody))
	your_string := string(your_to_byte)
	log.Printf("string %+v\n",your_string)
	req.Body.Close()
}

 */

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":1210", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	/*
	x := map[string]interface{}{}
	readbytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(readbytes,&x)

	 */

	err := r.ParseForm()
	if err != nil {
		fmt.Printf("%+v\n",err)
	}
	fmt.Printf("hehe %+v\n",r.FormValue("secret_key"))

	w.Write([]byte("bye bye ,this is v1 httpServer"))
}
