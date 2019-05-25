package match

import (
	//"fmt"
	"sync"
)

type Player struct {
	Name  string
	Skill int
}
func (p *Player) Play(wg *sync.WaitGroup, c chan int, w chan *Player){
	defer wg.Done()
	for {
		ball, ok := <-c
		if !ok {
			//fmt.Printf("\n%v won the match!", p.Name)
			w<-p
			break
		}
		if ball <= p.Skill {
			nb := RandBall.Intn(100)
			//fmt.Printf("\nPlayer %v return ball %v", p.Name, nb)
			c <- nb
		}
		if ball > p.Skill {
			//fmt.Printf("\n%v lose ball %v", p.Name, ball)
			close(c)
			break
		}
	}
}
