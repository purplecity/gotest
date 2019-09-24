package main

import (
	"github.com/go-redis/redis"
	"log"
)

type hpRedisClient struct {
	redisClient    *redis.Client
}

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
	RedLine = "redline"
)


var hporderclient *hpRedisClient

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
			return nil, err
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
	_, err = client.redisClient.HMSet(key,value).Result()
	if err != nil {
		return err
	}
	return nil
}

func SetRedLine(symbol,value string) error{
	client,err := GetRedisOrderClient()
	if err != nil {
		return err
	}
	_, err = client.redisClient.Set(symbol+RedLine,value,0).Result()
	if err != nil {
		return err
	}
	return nil
}

var(
	BTCOddsInfo = "BTCOddsInfo"
	SHCIOddsInfo = "SHCIOddsInfo"
	SZCIOddsInfo = "SZCIOddsInfo"

	BTCLevelOneMaxDv = 50000
	BTCLevelOneGreaterOdds = 0.9
	BTCLevelOneLessOdds = 0.9

	BTCLevelTwoMaxDv = 100000
	BTCLevelTwoGreaterOdds = 1.2
	BTCLevelTwoLessOdds = 0.6

	BTCLevelThreeMaxDv = 200000
	BTCLevelThreeGreaterOdds = 1.7
	BTCLevelThreeLessOdds = 0.1

	BTCLevelFourGreaterOdds = 1.8
	BTCLevelFourLessOdds = 0

	BTCMAP = map[string]interface{}{
		"LevelOneMaxDv":BTCLevelOneMaxDv, "LevelOneGreaterOdds":BTCLevelOneGreaterOdds,"LevelOneLessOdds":BTCLevelOneLessOdds,

		"LevelTwoMaxDv":BTCLevelTwoMaxDv, "LevelTwoGreaterOdds":BTCLevelTwoGreaterOdds,"LevelTwoLessOdds":BTCLevelTwoLessOdds,

		"LevelThreeMaxDv":BTCLevelThreeMaxDv, "LevelThreeGreaterOdds":BTCLevelThreeGreaterOdds,"LevelThreeLessOdds":BTCLevelThreeLessOdds,

		"LevelFourGreaterOdds":BTCLevelFourGreaterOdds,"LevelFourLessOdds":BTCLevelFourLessOdds,
	}

	SHCILevelOneMaxDv = 50000
	SHCILevelOneGreaterOdds = 0.9
	SHCILevelOneLessOdds = 0.9

	SHCILevelTwoMaxDv = 100000
	SHCILevelTwoGreaterOdds = 1.2
	SHCILevelTwoLessOdds = 0.6

	SHCILevelThreeMaxDv = 200000
	SHCILevelThreeGreaterOdds = 1.7
	SHCILevelThreeLessOdds = 0.1

	SHCILevelFourGreaterOdds = 1.8
	SHCILevelFourLessOdds = 0

	SHCIMAP = map[string]interface{}{
		"LevelOneMaxDv":SHCILevelOneMaxDv, "LevelOneGreaterOdds":SHCILevelOneGreaterOdds,"LevelOneLessOdds":SHCILevelOneLessOdds,

		"LevelTwoMaxDv":SHCILevelTwoMaxDv, "LevelTwoGreaterOdds":SHCILevelTwoGreaterOdds,"LevelTwoLessOdds":SHCILevelTwoLessOdds,

		"LevelThreeMaxDv":SHCILevelThreeMaxDv, "LevelThreeGreaterOdds":SHCILevelThreeGreaterOdds,"LevelThreeLessOdds":SHCILevelThreeLessOdds,

		"LevelFourGreaterOdds":SHCILevelFourGreaterOdds,"LevelFourLessOdds":SHCILevelFourLessOdds,
	}

	SZCILevelOneMaxDv = 50000
	SZCILevelOneGreaterOdds = 0.9
	SZCILevelOneLessOdds = 0.9

	SZCILevelTwoMaxDv = 100000
	SZCILevelTwoGreaterOdds = 1.2
	SZCILevelTwoLessOdds = 0.6

	SZCILevelThreeMaxDv = 200000
	SZCILevelThreeGreaterOdds = 1.7
	SZCILevelThreeLessOdds = 0.1

	SZCILevelFourGreaterOdds = 1.8
	SZCILevelFourLessOdds = 0

	SZCIMAP = map[string]interface{}{
		"LevelOneMaxDv":SZCILevelOneMaxDv, "LevelOneGreaterOdds":SZCILevelOneGreaterOdds,"LevelOneLessOdds":SZCILevelOneLessOdds,

		"LevelTwoMaxDv":SZCILevelTwoMaxDv, "LevelTwoGreaterOdds":SZCILevelTwoGreaterOdds,"LevelTwoLessOdds":SZCILevelTwoLessOdds,

		"LevelThreeMaxDv":SZCILevelThreeMaxDv, "LevelThreeGreaterOdds":SZCILevelThreeGreaterOdds,"LevelThreeLessOdds":SZCILevelThreeLessOdds,

		"LevelFourGreaterOdds":SZCILevelFourGreaterOdds,"LevelFourLessOdds":SZCILevelFourLessOdds,
	}

	BTCCurOdds = "BTCCurOdds"
	SHCICurOdds = "SHCICurOdds"
	SZCICurOdds = "SZCICurOdds"

	BTCCURMAP = map[string]interface{}{"UpOdds":0.9,"DownOdds":0.9,"Count":0}
	SHCICURMAP = map[string]interface{}{"UpOdds":0.9,"DownOdds":0.9,"Count":0}
	SZCICURMAP = map[string]interface{}{"UpOdds":0.9,"DownOdds":0.9,"Count":0}
)

func main() {

	HPOddsHMSet(BTCCurOdds,BTCCURMAP)
	HPOddsHMSet(SHCICurOdds,SHCICURMAP)
	HPOddsHMSet(SZCICurOdds,SZCICURMAP)
	SetRedLine("BTCredline","0")

	/*
	btccurodds := redis.HPOddsHGetAll(BTCCurOdds)
	log.Printf("%+v\n",btccurodds)
	count,_ := strconv.Atoi(btccurodds["Count"])
	UpOdds,_ := strconv.ParseFloat(btccurodds["UpOdds"],64)
	DownOdds,_ := strconv.ParseFloat(btccurodds["DownOdds"],64)
	log.Printf("%+v,%+v,%+v\n",count,UpOdds,DownOdds)

	btcoddsinfo := redis.HPOddsHGetAll(BTCOddsInfo)
	LevelOneMaxDv,_ := strconv.ParseFloat(btcoddsinfo["LevelOneMaxDv"],64)
	LevelOneGreaterOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelOneGreaterOdds"],64)
	LevelOneLessOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelOneLessOdds"],64)
	LevelTwoMaxDv,_ := strconv.ParseFloat(btcoddsinfo["LevelTwoMaxDv"],64)
	LevelTwoGreaterOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelTwoGreaterOdds"],64)
	LevelTwoLessOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelTwoLessOdds"],64)
	LevelThreeMaxDv,_ := strconv.ParseFloat(btcoddsinfo["LevelThreeMaxDv"],64)
	LevelThreeGreaterOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelThreeGreaterOdds"],64)
	LevelThreeLessOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelThreeLessOdds"],64)
	LevelFourGreaterOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelFourGreaterOdds"],64)
	LevelFourLessOdds,_ := strconv.ParseFloat(btcoddsinfo["LevelFourLessOdds"],64)
	log.Printf("%+v\n",btcoddsinfo)
	log.Printf("%+v,%+v,%+v\n%+v,%+v,%+v\n%+v,%+v,%+v\n%+v,%+v\n",
		LevelOneMaxDv,LevelOneGreaterOdds,LevelOneLessOdds,
		LevelTwoMaxDv,LevelTwoGreaterOdds,LevelTwoLessOdds,
		LevelThreeMaxDv,LevelThreeGreaterOdds,LevelThreeLessOdds,
		LevelFourGreaterOdds,LevelFourLessOdds)
	*/
}
