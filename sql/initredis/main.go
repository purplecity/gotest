package main

import (
	"gotest/sql/redis"
	"log"
	"strconv"
)
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

	redis.HPOddsHMSet(BTCOddsInfo,BTCMAP)
	redis.HPOddsHMSet(SHCIOddsInfo,SHCIMAP)
	redis.HPOddsHMSet(SZCIOddsInfo,SZCIMAP)
	redis.HPOddsHMSet(BTCCurOdds,BTCCURMAP)
	redis.HPOddsHMSet(SHCICurOdds,SHCICURMAP)
	redis.HPOddsHMSet(SZCICurOdds,SZCICURMAP)

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
	log.Printf("%+v,%+v,%+v,%+v\n%+v,%+v,%+v,%+v\n%+v,%+v,%+v,%+v\n%+v,%+v,%+v\n",
		LevelOneMaxDv,LevelOneGreaterOdds,LevelOneLessOdds,
		LevelTwoMaxDv,LevelTwoGreaterOdds,LevelTwoLessOdds,
		LevelThreeMaxDv,LevelThreeGreaterOdds,LevelThreeLessOdds,
		LevelFourGreaterOdds,LevelFourLessOdds)

}
