package bots

type EnglishBot struct{
	Name string
}

func (e EnglishBot) Hello() string{
	return "Hello, I'm "+e.Name+"!"
}
func (e EnglishBot) Bye() string{
	return "Bye, "+e.Name+" was with you!"
}
func (e EnglishBot) Time() string{
	return "Now is "+GetTime()
}
func (e EnglishBot) Date() string{
	return "Date is "+GetDate()
}
func (e EnglishBot) Weekday() string{
	return "Today is "+GetWeekday()
}