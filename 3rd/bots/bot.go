package bots

type Bot interface{
	Hello() string
	Bye() string
	Time() string
	Date() string
	Weekday() string
	Command(command string) string
	ExitWord() string
}