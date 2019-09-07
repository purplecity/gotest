package main

import (
	"crypto/rand"
	"fmt"
	"github.com/go-redis/redis" // thread-safe. client like database/HPSQL DB, represent a conn pool
	"log"
	r "math/rand"
	"strconv"
	"time"
)

var (
	RedisAddr = "127.0.0.1:6379"
	RedisPassword = "k"
	//RedisPassword = "HP@123"
	RedisDB = 0
	RedisOrderDB = 2
	RedisLockDB = 2
	RedisMaxRetries = 1
	RedisLockExpireTime = 2
)

type hpRedisClient struct {
	redisClient    *redis.Client
}
var hporderclient *hpRedisClient



func GetRedisOrderClient() *hpRedisClient {
	if hporderclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisOrderDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Panicf("ERROR----connect Cache failed----err:%v\n",err)
		}
		hporderclient = &hpRedisClient{redisClient:client}
	}
	return hporderclient
}

func SetRedisOrder(key string) {
	client := GetRedisOrderClient()
	client.redisClient.Set(key,0,time.Minute*1)
}

func GetOrderKeys(pattern string) []string {
	client := GetRedisOrderClient()
	v,_ := client.redisClient.Keys(pattern).Result()
	return v
}


func genRandomInvitationCode(n int) string {
	alphabets := []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

func GetRedisLockClient() *hpRedisClient {
	if hporderclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisLockDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Panicf("ERROR----connect Cache failed----err:%v\n", err)
		}
		hporderclient = &hpRedisClient{redisClient: client}
	}
	return hporderclient
}

func SetRedisLock(key string, value interface{}) (x bool){
	client := GetRedisOrderClient()
	x,err := client.redisClient.SetNX(key,value,time.Second*time.Duration(RedisLockExpireTime)).Result()
	if  err!= nil {
		log.Printf("ERROR----GetRedisLock failed ----err:%+v\n",err)
	}
	return
}

func GetRedisLock(key string) int64 {
	client := GetRedisOrderClient()
	x,err  := client.redisClient.Get(key).Result()
	if  err!= nil {
		log.Printf("ERROR----GetRedisLock failed ----err:%+v\n",err)
	}
	v,_ := strconv.ParseInt(x,10,64)
	return v
}


func main() {

	/*
	SetOddsMap("ODDS::BTC", map[string]interface{}{
		"LevelOneMinDV":0,"LevelOneMaxDV":100000,"LevelOneMinOdds":0.9,"LevelOneMaxOdds":0.9,
		"LevelTwoMinDV":100000,"LevelTwoMaxDV":200000,"LevelTwoMinOdds":0.6,"LevelTwoMaxOdds":1.2,
		"LevelThreeMinDV":200000,"LevelThreeMaxDV":500000,"LevelThreeMinOdds":0.1,"LevelThreeMaxOdds":1.7,
		"LevelFourMinDV":500000,"LevelFourMinOdds":0,"LevelFourMaxOdds":1.8,
	})

	c := sync.WaitGroup{}
	c.Add(1200000)
	for i := 0; i< 600000;i++ {
		//有问题
		x := genRandomInvitationCode(6)
		y := genRandomInvitationCode(6)
		 go func() {
			 SetRedisOrder(fmt.Sprintf("BTCUP::%s::12345",x))
			 c.Done()
		 }()
		 go func() {
			 SetRedisOrder(fmt.Sprintf("BTCDOWN::%s::12345",y))
			 c.Done()
		 }()
	}
	c.Wait()


	nt := time.Now().UnixNano()
	l1 := GetOrderKeys("BTCUP::*::12345")
	l2 := GetOrderKeys("BTCDOWN::*::12345")
	m1 := float64(0)
	m2 := float64(0)
	cnt := sync.WaitGroup{}
	cnt.Add(2)
	go func() {
		for _,v := range l1 {
			x,_ := strconv.ParseFloat((strings.Split(v,"::"))[2],64)
			m1 += x
		}
		cnt.Done()
	}()

	go func() {
		for _,v := range l2 {
			x,_ :=  strconv.ParseFloat((strings.Split(v,"::"))[2],64)
			m2 += x
		}
		cnt.Done()
	}()
	cnt.Wait()
	et := time.Now().UnixNano()
	fmt.Printf("%+v,%+v,%+v,%+v\n",len(l1),len(l2),nt,et)

	 */
	fmt.Println(SetRedisLock("testhehe",1))
	fmt.Println(time.Now().UnixNano())
	fmt.Println(SetRedisLock("testhehe",1))
	fmt.Println(time.Now().UnixNano()) //大概是0.2毫秒
	time.Sleep(time.Second*2)
	fmt.Println(SetRedisLock("testhehe",1))

	/*
	key uid_标的物
	filed 赔率 value 下单时间
	filed 是否经过赔率变化 value 0/1


	*/


}
