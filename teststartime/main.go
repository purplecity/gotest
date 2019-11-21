package main

import (
	"HPOptionServer/Common/CommonConf"
	"log"
	"time"
)

func getstarttime(ts int64) (starttime int64) {
	now := time.Unix(ts,0)
	Et1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.ForexFirstclosehour, CommonConf.ForexFirstclosemin, 0, 0, now.Location()).Unix()
	Bt1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.ForexFirstopenhour, CommonConf.ForexFirstopenmin, 0, 0, now.Location()).Unix()
	last := now.Add(time.Hour*-24)
	bt1 := time.Date(last.Year(), last.Month(), last.Day(), CommonConf.ForexFirstopenhour, CommonConf.ForexFirstopenmin, 0, 0, last.Location()).Unix()

	if now.Weekday() >= 2 && now.Weekday() <= 5 {
		if ts >= Bt1 {
			starttime = Bt1 + ((ts - Bt1)/90)*90
		} else if ts <= Et1 {
			starttime = bt1 + ((ts - bt1)/90)*90
		} else {
			starttime = Bt1
		}
	} else if now.Weekday() == 1 {
		if ts >= Bt1 {
			starttime = Bt1 + ((ts - Bt1)/90)*90
		}  else {
			starttime = Bt1
		}
	} else if now.Weekday() == 6 {
		if ts <= Et1 - 3600 {
			starttime = bt1 + ((ts - bt1)/90)*90
		} else {
			m := now.Add(time.Hour*48)
			mt := time.Date(m.Year(), m.Month(), m.Day(), CommonConf.ForexFirstopenhour, CommonConf.ForexFirstopenmin, 0, 0, m.Location()).Unix()
			starttime = mt
		}
	} else {
		m := now.Add(time.Hour*24)
		mt := time.Date(m.Year(), m.Month(), m.Day(), CommonConf.ForexFirstopenhour, CommonConf.ForexFirstopenmin, 0, 0, m.Location()).Unix()
		starttime = mt
	}
	return
}

func main() {
	log.Printf("%+v\n",time.Unix(getstarttime(1573845817),0))
}
