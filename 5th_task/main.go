package main

import (
	"5th_task/match"
)

func main() {
	pl1 := match.Player{Name: "Pl1", Skill: 98}
	pl2 := match.Player{Name: "Pl2", Skill: 98}

	m:=match.Match{Player1: pl1, Player2: pl2}
	m.PlayMatch()
}
