package main

import (
	"5th_task/match"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	pl1 := match.Player{Name: "Pl1", Skill: 72}
	pl2 := match.Player{Name: "Pl2", Skill: 72}

	m:=match.TennisMatch{pl1,pl2}
	m.Start(&wg)

	wg.Wait()  

}
