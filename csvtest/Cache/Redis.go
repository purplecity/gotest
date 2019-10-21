package Cache

import (
	"github.com/go-redis/redis" // thread-safe. client like database/HPSQL DB, represent a conn pool
	"log"
)

var (
	RedisAddr = "127.0.0.1:6379"
	RedisPassword = "k"
	//RedisPassword = "7U'G~1LzI+]3_~D"
	RedisDB = 0 //验证码以及次数
	RedisMaxRetries = 1
)

type hpRedisClient struct {
	redisClient    *redis.Client
}

var hpclient *hpRedisClient

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

func RedisRPUSH(key string, value interface{}) (int64,error) {
	client,err := GetRedisClient()
	if err != nil {
		return 0,err
	}
	vc, err  := client.redisClient.RPush(key,value).Result()
	if err != nil {
		log.Printf("ERROR----RPUSH failed----err:%+v\n",err)
		return 0,err
	}
	return vc,nil
}


func RedisLPop(key string) error {
	client,err := GetRedisClient()
	if err != nil {
		return err
	}
	_, err  = client.redisClient.LPop(key).Result()
	if err != nil {
		log.Printf("ERROR----LPOP failed----err:%+v\n",err)
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
	return vc,nil
}


