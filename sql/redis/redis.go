package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"gotest/sql/Snowflake"
	"log"
	"strconv"
	"time"
)

var (
	//RedisAddr = "127.0.0.1:6379"
	RedisAddr = "47.244.212.51:6379"
	//RedisPassword = "k"
	RedisPassword = "7U'G~1LzI+]3_~D"
	RedisDB = 0 //验证码以及次数
	RedisOrderDB = 1 //下单
	RedisLockDB = 2 //锁
	RedisOddrDB = 3//赔率
	RedisMaxRetries = 1
	RedisLockSleepTime = 40
	RedisLockExpireTime = 2
)

type hpRedisClient struct {
	redisClient    *redis.Client
}

var hpclient *hpRedisClient
var hporderclient *hpRedisClient
var hplockclient *hpRedisClient
var hpoddsclient *hpRedisClient

func GetRedisClient() *hpRedisClient {
	if hpclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Panicf("ERROR----connect Cache failed----err:%v\n",err)
		}
		hpclient = &hpRedisClient{redisClient:client}
	}
	return hpclient
}


func SetRedisPrice(key string, value float64) {
	client := GetRedisClient()
	client.redisClient.Set(key,value,0)
}

func GetRedisPrice(key string) (v float64) {
	client := GetRedisClient()
	x,err  := client.redisClient.Get(key).Result()
	if  err!= nil {
		log.Printf("ERROR----GetRedisPrice failed ----err:%+v\n",err)
	}
	v,_ = strconv.ParseFloat(x,64)
	return
}

func GetRedisValidcode(ph string) (vc string) {
	client := GetRedisClient()
	vc, err  := client.redisClient.Get("ValidCode:"+ph).Result()
	if err != nil {
		log.Printf("ERROR----Not have this phonenumber validcode----err:%+v\n",err)
	}
	return
}

func SetRedisValidcode(key,value string,ext time.Duration) {
	client := GetRedisClient()
	client.redisClient.Set(key,value,ext)
}

func ExistSMSLimit(key string) (vc int64) {
	client := GetRedisClient()
	vc, err := client.redisClient.Exists(key).Result()
	if err != nil {
		log.Printf("ERROR----ExistLimit failed----err:%+v\n",err)
	}
	return
}

func GetRedisLimit(key string) (vc map[string]string) {
	client := GetRedisClient()
	vc, err  := client.redisClient.HGetAll(key).Result()
	if err != nil {
		log.Printf("ERROR----GetRedisLimit failed----err:%+v\n",err)
	}
	return
}

func SetSMSLimit(key string) {
	client := GetRedisClient()
	client.redisClient.HSet(key,"0","0")
	client.redisClient.Expire(key,time.Hour*24)
}

func HsetSMSTime(key string) {
	client := GetRedisClient()
	client.redisClient.HSet(key,fmt.Sprintf("%+v",time.Now().UnixNano()),"0")
}





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
			log.Panicf("ERROR----connect Cache failed----err:%v\n", err)
		}
		hporderclient = &hpRedisClient{redisClient: client}
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



func GetRedisLockClient() *hpRedisClient {
	if hplockclient == nil {
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
		hplockclient = &hpRedisClient{redisClient: client}
	}
	return hplockclient
}


func setRedisLock(key,value string) (x bool) {
	client := GetRedisLockClient()
	x,err := client.redisClient.SetNX(key,value,time.Second*time.Duration(RedisLockExpireTime)).Result()
	if  err!= nil {
		log.Printf("ERROR----SetRedisLock failed ----err:%+v\n",err)
	}
	return
}

func getRedisLock(key string) string {
	client := GetRedisLockClient()
	x,err  := client.redisClient.Get(key).Result()
	if  err!= nil {
		log.Printf("ERROR----GetRedisLock failed ----err:%+v\n",err)
	}
	return x
}

func delRedisLock(key string) {
	client := GetRedisLockClient()
	client.redisClient.Del(key)
}

//暂时先不设置强制抢锁

func RedisLock(key string) string {
	v := Snowflake.GenID()
	for {
		if setRedisLock(key,v) {
			break
		} else {
			time.Sleep(time.Millisecond*time.Duration(RedisLockSleepTime)) //中间执行打印出来结果大概30s毫秒左右
		}
	}
	return v
}

func RedisUnlock(key string) {
	/*
		if uuid == Cache.GetRedisLock(uid) {
			Cache.DelRedisLock(uid)  //
		}
	*/
	// 大概120ms 判断的时间都要几倍于锁住块的执行的时间
	delRedisLock(key)
}


func GetRedisOddsClient() *hpRedisClient {
	if hpoddsclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisOddrDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Panicf("ERROR----connect Cache failed----err:%v\n", err)
		}
		hpoddsclient = &hpRedisClient{redisClient: client}
	}
	return hpoddsclient
}


func HPOddsHMSet(key string, value map[string]interface{}) {
	client := GetRedisOddsClient()
	client.redisClient.HMSet(key,value)
}

func HPOddsHGetAll(key string) map[string]string {
	client := GetRedisOddsClient()
	vc, err := client.redisClient.HGetAll(key).Result()
	if err != nil {
		log.Printf("ERROR----HPOddsHGetAll failed----err:%+v\n",err)
	}
	return vc
}