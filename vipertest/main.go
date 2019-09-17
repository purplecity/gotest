package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)
func main() {
	v := viper.New()
	v.SetConfigName("test")
	v.AddConfigPath("$GOPATH/src/gotest/vipertest/")
	v.SetConfigType("toml")
	if err := v.ReadInConfig();err != nil {
		fmt.Printf("err:%+v\n",err)
	}
	fmt.Printf("%+T\n",v.Get("amount"))

	go func() {
		for {
			time.Sleep(time.Second*3)
			fmt.Println("look",v.Get("KeepaliveInterval"))

		}

	}()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config is change :%s \n", e.String())
		//cancel()
	})
	//开始监听
	v.WatchConfig()
	done := make(chan struct{})
	<- done
}
