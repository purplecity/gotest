package main

import (
	"fmt"
	"github.com/go-redis/redis" // thread-safe. client like database/HPSQL DB, represent a conn pool
	"gotest/redis/Snowflake"
	"log"
	"time"
)

var (
	RedisAddr = "127.0.0.1:6379"
	RedisPassword = "k"
	//RedisPassword = "HP@123"
	RedisDB = 0
	RedisOrderDB = 1
	RedisLockDB = 2
	RedisMaxRetries = 1
	RedisLockExpireTime = 2
	RedisLockSleepTime = 40
	RedLine = "redline"
)


type hpRedisClient struct {
	redisClient    *redis.Client
}

var hpclient *hpRedisClient
var hporderclient *hpRedisClient
var hplockclient *hpRedisClient

//for 验证码
func GetRedisClient() (*hpRedisClient,error) {
	if hpclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Printf("ERROR----connect Cache failed----err:%v\n",err)
			return nil,err
		}
		hpclient = &hpRedisClient{redisClient:client}
	}
	return hpclient,nil
}


func RedisRPUSH(key string, value float64) (int64,error) {
	client,err := GetRedisClient()
	if err != nil {
		return 0,err
	}
	vc, err  := client.redisClient.RPush(key,value).Result()
	if err != nil {
		log.Printf("ERROR----Not have this phonenumber validcode----err:%+v\n",err)
		return 0,err
	}
	fmt.Printf("rpush %+v\n",vc)
	return vc,nil
}


func RedisLPop(key string) error {
	client,err := GetRedisClient()
	if err != nil {
		return err
	}
	_, err  = client.redisClient.LPop(key).Result()
	if err != nil {
		log.Printf("ERROR----Not have this phonenumber validcode----err:%+v\n",err)
		return err
	}
	return nil
}

func RedisLRange(key string,start,stop int64) ([]string,error) {
	client,err := GetRedisClient()
	if err != nil {
		return nil,err
	}
	vc, err  := client.redisClient.LRange(key,start,stop).Result()
	if err != nil {
		log.Printf("ERROR----Not have this list----err:%+v\n",err)
		return nil,err
	}
	fmt.Printf("%+v\n",vc)
	return vc,nil
}



func GetRedisValidcode(ph string) (string,error) {
	client,err := GetRedisClient()
	if err != nil {
		return "",err
	}
	vc, err  := client.redisClient.Get("ValidCode:"+ph).Result()
	if err != nil {
		log.Printf("ERROR----Not have this phonenumber validcode----err:%+v\n",err)
		return "",err
	}
	return vc,nil
}

func SetRedisValidcode(key,value string,ext time.Duration) error {
	client,err := GetRedisClient()
	if err != nil {
		return err
	}
	client.redisClient.Set(key,value,ext)
	return nil
}

func ExistSMSLimit(key string) (int64,error) {
	client,err := GetRedisClient()
	if err != nil {
		return 0,err
	}
	vc, err := client.redisClient.Exists(key).Result()
	if err != nil {
		log.Printf("ERROR----ExistLimit failed----err:%+v\n",err)
		return 0,err
	}
	return vc,err
}

func GetRedisLimit(key string) (map[string]string,error) {
	client,err := GetRedisClient()
	if err != nil {
		return nil,err
	}
	vc, err  := client.redisClient.HGetAll(key).Result()
	if err != nil {
		log.Printf("ERROR----GetRedisLimit failed----err:%+v\n",err)
		return nil,err
	}
	return vc,nil
}


func HsetSMSTime(key string) error {
	client,err := GetRedisClient()
	if err != nil {
		return err
	}
	client.redisClient.HSet(key,fmt.Sprintf("%+v",time.Now().UnixNano()),"0")
	return nil
}




//for 存单子为了赔率检测 和 亏损检测
func GetRedisOrderClient() (*hpRedisClient,error) {
	if hporderclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisOrderDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Printf("ERROR----connect Cache failed----err:%v\n", err)
			return nil,err
		}
		hporderclient = &hpRedisClient{redisClient: client}
	}
	return hporderclient,nil
}


func HPOddsHMSet(key string, value map[string]interface{}) error {
	client,err := GetRedisOrderClient()
	if err != nil {
		return err
	}
	client.redisClient.HMSet(key,value)
	return nil
}

func HPOddsHGetAll(key string) (map[string]string,error) {
	client,err := GetRedisOrderClient()
	if err != nil {
		return nil,err
	}
	vc, err := client.redisClient.HGetAll(key).Result()
	if err != nil {
		log.Printf("ERROR----HPOddsHGetAll failed----err:%+v\n",err)
		return nil,err
	}
	return vc,nil
}


func SetRedisOrder(key string) error {
	client,err := GetRedisOrderClient()
	if err != nil {
		return err
	}
	client.redisClient.Set(key,0,time.Minute*1)
	return nil
}

func GetOrderKeys(pattern string) ([]string,error) {
	client,err := GetRedisOrderClient()
	if err != nil {
		return nil,err
	}
	v,err:= client.redisClient.Keys(pattern).Result()
	if err != nil {
		log.Printf("ERROR----GetOrderKeys failed----err:%+v\n",err)
		return nil,err
	}
	return v,nil
}


func SetRedLine(symbol,value string) (string,error){
	client,err := GetRedisOrderClient()
	if err != nil {
		return "",err
	}
	v,err := client.redisClient.Set(symbol+RedLine,value,0).Result()
	fmt.Printf("%s,%+v\n",v,err)
	return v,nil
}

func GetRedLine(symbol string) (bool,error) {
	client,err := GetRedisOrderClient()
	if err != nil {
		return false,err
	}
	s, err  := client.redisClient.Get(symbol+RedLine).Result()
	if err != nil {
		log.Printf("ERROR----get redline  failed ----err:%+v\n",err)
		return false,err
	}
	return s == "1",nil
}



// for 锁
func GetRedisLockClient() (*hpRedisClient,error) {
	if hplockclient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPassword,
			DB:         RedisLockDB,
			MaxRetries: RedisMaxRetries,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Printf("ERROR----connect Cache failed----err:%v\n", err)
			return nil,err
		}
		hplockclient = &hpRedisClient{redisClient: client}
	}
	return hplockclient,nil
}

func getRedisLock(key string) (string,error) {
	client,err := GetRedisLockClient()
	if err != nil {
		return "",err
	}
	x,err  := client.redisClient.Get(key).Result()
	if  err!= nil {
		log.Printf("ERROR----GetRedisLock failed ----err:%+v\n",err)
		return "",err
	}
	return x,nil
}

func delRedisLock(key string) error {
	client,err := GetRedisLockClient()
	if err != nil {
		return err
	}
	client.redisClient.Del(key)
	return nil
}

//暂时先不设置强制抢锁

func RedisLock(key string) (string,error) {
	v := Snowflake.GenID()
	client,err := GetRedisLockClient()
	if err != nil {
		return "",err
	}
	for {
		x,err := client.redisClient.SetNX(key,v,time.Second*time.Duration(RedisLockExpireTime)).Result()
		if err!= nil {
			log.Printf("ERROR----SetRedisLock failed ----err:%+v\n",err)
			return "",err
		} else if x {
			break
		} else {
			time.Sleep(time.Millisecond*time.Duration(RedisLockSleepTime)) //中间执行打印出来结果大概30s毫秒左右
		}
	}
	return v,nil
}

func RedisUnlock(key string) error{
	/*
		if uuid == Cache.GetRedisLock(uid) {
			Cache.DelRedisLock(uid)  //
		}
	*/
	// 大概120ms 判断的时间都要几倍于锁住块的执行的时间
	return delRedisLock(key)
}


func SetSMSLimit(key string) error {
	client,err := GetRedisClient()
	if err != nil {
		return err
	}
	luaScript := redis.NewScript(`
	redis.call("HSET", KEYS[1], ARGV[1], ARGV[2])
	redis.call("EXPIRE", KEYS[1], ARGV[3])
	return 1
	`)
	_,err = luaScript.Run(client.redisClient,[]string{key},"0","0",10).Result()
	if err != nil {
		log.Printf("ERROR----run lua script failed----err:%+v\n",err)
		return err
	}
	return nil
}


func main () {
	RedisRPUSH("testlist",5.77)
	RedisLPop("testlist")
	s,_ := RedisLRange("testlist",0,-1)
	for _, x := range s {
		fmt.Printf("%+v,%T\n",x,x)
	}
}