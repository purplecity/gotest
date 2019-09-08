package main

import (
	"gotest/sql/redis"
)
var(
	BTCODDS = "BTCODDS"
	SHCIODDS = "SHCIODDS"
	SZCIODDS = "SZCIODDS"

	BTCLevelOneMinDv = 0
	BTCLevelOneMaxDv = 50000
	BTCLevelOneGreaterOdds = 0.9
	BTCLevelOneLessOdds = 0.9

	BTCLevelTwoMinDv = 50000
	BTCLevelTwoMaxDv = 100000
	BTCLevelTwoGreaterOdds = 1.2
	BTCLevelTwoLessOdds = 0.6

	BTCLevelThreeMinDv = 100000
	BTCLevelThreeMaxDv = 200000
	BTCLevelThreeGreaterOdds = 1.7
	BTCLevelThreeLessOdds = 0.1

	BTCLevelFourMinDv = 200000
	BTCLevelFourGreaterOdds = 1.8
	BTCLevelFourLessOdds = 0

	BTCMAP = map[string]interface{}{
		"LevelOneMinDv":BTCLevelOneMinDv,"LevelOneMaxDv":BTCLevelOneMaxDv,
		"LevelOneGreaterOdds":BTCLevelOneGreaterOdds,"LevelOneLessOdds":BTCLevelOneLessOdds,

		"LevelTwoMinDv":BTCLevelTwoMinDv,"LevelTwoMaxDv":BTCLevelTwoMaxDv,
		"LevelTwoGreaterOdds":BTCLevelTwoGreaterOdds,"LevelTwoLessOdds":BTCLevelTwoLessOdds,

		"LevelThreeMinDv":BTCLevelThreeMinDv,"LevelThreeMaxDv":BTCLevelThreeMaxDv,
		"LevelThreeGreaterOdds":BTCLevelThreeGreaterOdds,"LevelThreeLessOdds":BTCLevelThreeLessOdds,

		"LevelFourMinDv":BTCLevelFourMinDv,
		"LevelFourGreaterOdds":BTCLevelFourGreaterOdds,"LevelFourLessOdds":BTCLevelFourLessOdds,
	}

	SHCILevelOneMinDv = 0
	SHCILevelOneMaxDv = 50000
	SHCILevelOneGreaterOdds = 0.9
	SHCILevelOneLessOdds = 0.9

	SHCILevelTwoMinDv = 50000
	SHCILevelTwoMaxDv = 100000
	SHCILevelTwoGreaterOdds = 1.2
	SHCILevelTwoLessOdds = 0.6

	SHCILevelThreeMinDv = 100000
	SHCILevelThreeMaxDv = 200000
	SHCILevelThreeGreaterOdds = 1.7
	SHCILevelThreeLessOdds = 0.1

	SHCILevelFourMinDv = 200000
	SHCILevelFourGreaterOdds = 1.8
	SHCILevelFourLessOdds = 0

	SHCIMAP = map[string]interface{}{
		"LevelOneMinDv":SHCILevelOneMinDv,"LevelOneMaxDv":SHCILevelOneMaxDv,
		"LevelOneGreaterOdds":SHCILevelOneGreaterOdds,"LevelOneLessOdds":SHCILevelOneLessOdds,

		"LevelTwoMinDv":SHCILevelTwoMinDv,"LevelTwoMaxDv":SHCILevelTwoMaxDv,
		"LevelTwoGreaterOdds":SHCILevelTwoGreaterOdds,"LevelTwoLessOdds":SHCILevelTwoLessOdds,

		"LevelThreeMinDv":SHCILevelThreeMinDv,"LevelThreeMaxDv":SHCILevelThreeMaxDv,
		"LevelThreeGreaterOdds":SHCILevelThreeGreaterOdds,"LevelThreeLessOdds":SHCILevelThreeLessOdds,

		"LevelFourMinDv":SHCILevelFourMinDv,
		"LevelFourGreaterOdds":SHCILevelFourGreaterOdds,"LevelFourLessOdds":SHCILevelFourLessOdds,
	}

	SZCILevelOneMinDv = 0
	SZCILevelOneMaxDv = 50000
	SZCILevelOneGreaterOdds = 0.9
	SZCILevelOneLessOdds = 0.9

	SZCILevelTwoMinDv = 50000
	SZCILevelTwoMaxDv = 100000
	SZCILevelTwoGreaterOdds = 1.2
	SZCILevelTwoLessOdds = 0.6

	SZCILevelThreeMinDv = 100000
	SZCILevelThreeMaxDv = 200000
	SZCILevelThreeGreaterOdds = 1.7
	SZCILevelThreeLessOdds = 0.1

	SZCILevelFourMinDv = 200000
	SZCILevelFourGreaterOdds = 1.8
	SZCILevelFourLessOdds = 0

	SZCIMAP = map[string]interface{}{
		"LevelOneMinDv":SZCILevelOneMinDv,"LevelOneMaxDv":SZCILevelOneMaxDv,
		"LevelOneGreaterOdds":SZCILevelOneGreaterOdds,"LevelOneLessOdds":SZCILevelOneLessOdds,

		"LevelTwoMinDv":SZCILevelTwoMinDv,"LevelTwoMaxDv":SZCILevelTwoMaxDv,
		"LevelTwoGreaterOdds":SZCILevelTwoGreaterOdds,"LevelTwoLessOdds":SZCILevelTwoLessOdds,

		"LevelThreeMinDv":SZCILevelThreeMinDv,"LevelThreeMaxDv":SZCILevelThreeMaxDv,
		"LevelThreeGreaterOdds":SZCILevelThreeGreaterOdds,"LevelThreeLessOdds":SZCILevelThreeLessOdds,

		"LevelFourMinDv":SZCILevelFourMinDv,
		"LevelFourGreaterOdds":SZCILevelFourGreaterOdds,"LevelFourLessOdds":SZCILevelFourLessOdds,
	}

)

func main() {
	redis.HPHMSetOdds(BTCODDS,BTCMAP)
	redis.HPHMSetOdds(SHCIODDS,SHCIMAP)
	redis.HPHMSetOdds(SZCIODDS,SZCIMAP)
}
