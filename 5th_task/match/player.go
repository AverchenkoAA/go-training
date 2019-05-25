package match

import (
	//"fmt"
	"sync"
)
//Player имитирует игрока 
type Player struct {
	Name  string
	Skill int
	PlayerScore int
	PlayerGames int
	PlayerSets int
}
//Play (wg *sync.WaitGroup, c chan int, w chan *Player) имитирует одну партию
func (p *Player) Play(wg *sync.WaitGroup, c chan int, w chan *Player){
	defer wg.Done()
	for {
		ball, ok := <-c
		if !ok {
			w<-p
			break
		}
		if ball <= p.Skill {
			nb := RandBall.Intn(100)
			c <- nb
		}
		if ball > p.Skill {
			close(c)
			break
		}
	}
}
