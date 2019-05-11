package bots

type RussianBot struct{
	Name string
}

func (r RussianBot) Hello() string{
	return "Привет, я "+r.Name+"!"
}
func (r RussianBot) Bye() string{
	return "Пока, с Вами был "+r.Name+"!"
}
func (r RussianBot) Time() string{
	return "Сейчас время "+GetTime()
}
func (r RussianBot) Date() string{
	return "Сегодня "+GetDate()
}
func (r RussianBot) Weekday() string{
	return "День недели сегодня "+GetWeekday()
}