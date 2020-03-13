package main

import (
	"encoding/json"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"sort"
	"time"
)

var (
	ftpurl = "xg.gdcapi.com"
	ftpusername = "GI5.xingyu"
	ftppwd = "9S1gmjk65P"
)
func main() {

	tick := time.Tick(time.Minute*10)
	for range tick {
		//下载

	}
	//连接
	c, err := ftp.Dial(ftpurl+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(ftpusername, ftppwd)
	if err != nil {
		log.Fatal(err)
	}


	namelist,err := c.NameList("/AGIN")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",namelist)

	// Do something with the FTP conn
	r, err := c.Retr("/AGIN/20200312/202003120119.xml")
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(r)
	dirnamelist := []string{}
	json.Unmarshal(buf,&dirnamelist)

	println(string(buf))
	sort.Sort(dirnamelist)


	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
	//登陆

	//读取
	//存储
}
