package bots

import(
	"time"
)

func GetTime() string{
	name :="Europe/Minsk"
	t:=time.Now()
	loc,_:=time.LoadLocation(name)
	t = t.In(loc)
	return t.Format("15:04")
}

func GetDate() string{
	t:=time.Now()
	return t.Format("January 2, 2006")
}
func GetWeekday() string{
	t:=time.Now()
	return t.Format("Monday")
}
