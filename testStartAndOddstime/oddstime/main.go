package main

import (
	"HPOptionServer/Common/CommonConf"
	"log"
	"time"
)

func ForinitOdds(ts int64) {
	done := make(chan bool)
	go func() {
		now := time.Unix(ts,0)
		Et1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.BTCFirstclosehour, CommonConf.BTCFirstclosemin, 0, 0, now.Location()).Unix()
		Bt1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.BTCFirstopenhour, CommonConf.BTCFirstopenmin, 0, 0, now.Location()).Unix()
		if CommonConf.Starttime >= Bt1 {
			n := (CommonConf.Starttime - Bt1)/90
			m := (CommonConf.Starttime - Bt1) % 90
			t := int64(0)
			if m <= 30 {
				t = n*90 + Bt1 + 30
			} else {
				t = (n+1)*90 + Bt1 + 30
			}
			st := time.Unix(t,0)
			log.Printf("BTC 1 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else if CommonConf.Starttime >Et1 && CommonConf.Starttime < Bt1 {
			t := Bt1 + 30
			st := time.Unix(t,0)
			log.Printf("BTC 2 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else  {
			last := now.Add(time.Hour*-24)
			bt1 := time.Date(last.Year(), last.Month(), last.Day(), CommonConf.BTCFirstopenhour, CommonConf.BTCFirstopenmin, 0, 0, last.Location()).Unix()
			n := (CommonConf.Starttime - bt1)/90
			m := (CommonConf.Starttime - bt1) % 90
			t := int64(0)
			if m <= 30 {
				t = n*90 + bt1 + 30
			} else {
				t = (n+1)*90 + bt1 + 30
			}
			st := time.Unix(t,0)
			log.Printf("BTC 3 %+v\n",st)
			time.Sleep(st.Sub(now))
		}

		tick := time.Tick(time.Second * 90)
		for range tick {
			log.Println("BTC")
		}
	}()

	go func() {
		now := time.Unix(ts,0)
		Et1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.ForexFirstclosehour, CommonConf.ForexFirstclosemin, 0, 0, now.Location()).Unix()
		Bt1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.ForexFirstopenhour, CommonConf.ForexFirstopenmin, 0, 0, now.Location()).Unix()
		if CommonConf.Starttime >= Bt1 {
			n := (CommonConf.Starttime - Bt1)/90
			m := (CommonConf.Starttime - Bt1) % 90
			t := int64(0)
			if m <= 30 {
				t = n*90 + Bt1 + 30
			} else {
				t = (n+1)*90 + Bt1 + 30
			}
			st := time.Unix(t,0)
			log.Printf("Forex 1 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else if CommonConf.Starttime >Et1 && CommonConf.Starttime < Bt1 {
			t := Bt1 + 30
			st := time.Unix(t,0)
			log.Printf("Forex 2 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else  {
			last := now.Add(time.Hour*-24)
			bt1 := time.Date(last.Year(), last.Month(), last.Day(), CommonConf.ForexFirstopenhour, CommonConf.ForexFirstopenmin, 0, 0, last.Location()).Unix()
			n := (CommonConf.Starttime - bt1)/90
			m := (CommonConf.Starttime - bt1) % 90
			t := int64(0)
			if m <= 30 {
				t = n*90 + bt1 + 30
			} else {
				t = (n+1)*90 + bt1 + 30
			}
			st := time.Unix(t,0)
			log.Printf("Forex 3 %+v\n",st)
			time.Sleep(st.Sub(now))
		}

		tick := time.Tick(time.Second * 90)
		for range tick {
			log.Println("Forex")

		}
	}()

	go func() {
		now := time.Unix(ts,0)

		Et1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.IndexFirstclosehour, CommonConf.IndexFirstclosemin, 0, 0, now.Location()).Unix()
		Et2 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.IndexSecondclosehour, CommonConf.IndexSecondclosemin, 0, 0, now.Location()).Unix()
		Bt1 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.IndexFirstopenhour, CommonConf.IndexFirstopenmin, 0, 0, now.Location()).Unix()
		Bt2 := time.Date(now.Year(), now.Month(), now.Day(), CommonConf.IndexSecondopenhour, CommonConf.IndexSecondopenmin, 0, 0, now.Location()).Unix()

		//因为初始化是0 所以边界可以相等 直接等到30s后
		if CommonConf.Starttime <= Bt1 {
			t := Bt1 + 30
			st := time.Unix(t, 0)
			log.Printf("Index 0 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else if CommonConf.Starttime > Bt1 && CommonConf.Starttime <= Et1 {
			n := (CommonConf.Starttime - Bt1) / 90
			m := (CommonConf.Starttime - Bt1) % 90
			t := int64(0)
			if m <= 30 {
				t = n*90 + Bt1 + 30
			} else {
				t = (n+1)*90 + Bt1 + 30
			}
			st := time.Unix(t, 0)
			log.Printf("Index 1 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else if  CommonConf.Starttime >Et1 && CommonConf.Starttime <= Bt2 {
			t := Bt2 + 30
			st := time.Unix(t,0)
			log.Printf("Index 2 %+v\n",st)
			time.Sleep(st.Sub(now))
		} else if CommonConf.Starttime > Bt2 && CommonConf.Starttime <= Et2 {
			n := (CommonConf.Starttime - Bt2)/90
			m := (CommonConf.Starttime - Bt2) % 90
			t := int64(0)
			if m <= 30 {
				t = n*90 + Bt2 + 30
			} else {
				t = (n+1)*90 + Bt2 + 30
			}
			st := time.Unix(t,0)
			log.Printf("Index 3 %+v\n",st)
			time.Sleep(st.Sub(now))
		}  else  {
			next := now.Add(time.Hour*24)
			bt1 := time.Date(next.Year(), next.Month(), next.Day(), CommonConf.IndexFirstopenhour, CommonConf.IndexFirstopenmin, 0, 0, next.Location()).Unix()
			t := bt1 + 30
			st := time.Unix(t, 0)
			log.Printf("Index 4 %+v\n",st)
			time.Sleep(st.Sub(now))
		}

		tick := time.Tick(time.Second * 90)
		for range tick {
			log.Println("Index")

		}
	}()
	<- done

}

func main() {
	ForinitOdds(1574332020)
}