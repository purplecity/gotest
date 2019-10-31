package main

import (
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"time"
)

func main() {
	now := time.Now()
	et := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	cond := map[string]interface{}{"Settletime__lt": et}
	scoreLastDayInfo := []mysql.Scorerecord{}
	mysql.GetAllRecord("Scorerecord", cond, &scoreLastDayInfo)

	var scoreMap = map[string]float64{}
	for _, v := range scoreLastDayInfo {
		if v.Playerid == "" {
			if _, ok := scoreMap[v.Directorid]; ok {
				scoreMap[v.Directorid] = Operation.HPAdd(scoreMap[v.Directorid], Operation.HPMul(v.Amount, 0.002))
			} else {
				scoreMap[v.Directorid] = Operation.HPMul(v.Amount, 0.002)
			}
		} else {
			if _, ok := scoreMap[v.Playerid]; ok {
				scoreMap[v.Playerid] = Operation.HPAdd(scoreMap[v.Playerid], Operation.HPMul(v.Amount, 0.002))
			} else {
				scoreMap[v.Playerid] = Operation.HPMul(v.Amount, 0.002)
			}
		}
	}

	for k, v := range scoreMap {
		ta := float64(0)
		takeScoreInfo2 := []mysql.Takescorerecord{}
		mysql.GetAllRecord("Takescorerecord", map[string]interface{}{"Uid": k, "Handletime__lt": now.Unix()}, &takeScoreInfo2)
		for _, x := range takeScoreInfo2 {
			ta = Operation.HPAdd(ta, x.Amount)
		}

		mysql.UpdateByCond("Score", map[string]interface{}{"Uid": k}, map[string]interface{}{
			"Total":  v,
			"Remain": Operation.HPSub(v, ta),
		})
	}
}
