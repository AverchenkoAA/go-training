package bots

type EnglishBot struct{
	Name string
}
func (e EnglishBot) Command(command string) string{
	switch command{
	case "Hello":
		return e.Hello()

	case "Time":
		return e.Time()

	case "Date":
		return e.Date()

	case "Day":
		return e.Weekday()

	case "Bye":
		return e.Bye()

	default:
		return "I didn't catch it. Sorry..."
	}
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
func (e EnglishBot) ExitWord() string{
	return "Bye"
}