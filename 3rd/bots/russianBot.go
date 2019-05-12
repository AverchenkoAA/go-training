package bots

type RussianBot struct{
	Name string
}
func (r RussianBot) Command(command string) string{
	switch command{
	case "Привет":
		return r.Hello()

	case "Время":
		return r.Time()

	case "Дата":
		return r.Date()

	case "День":
		return r.Weekday()

	case "Пока":
		return r.Bye()

	default:
		return "Прости, но я не понимаю."
	}
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
func (r RussianBot) ExitWord() string{
	return "Пока"
}