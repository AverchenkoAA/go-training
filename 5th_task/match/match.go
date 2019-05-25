package match

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)
type TennisMatch struct {
	Player1, Player2 Player
	//Score int
}

var RandBall *rand.Rand

func init(){
	S3 := rand.NewSource(time.Now().UnixNano())
	RandBall = rand.New(S3)
}
func (tm *TennisMatch) Start(wg *sync.WaitGroup){
	c := make(chan int)
	winner := make(chan *Player)

	go tm.Player1.Play(wg,c,winner)
	go tm.Player2.Play(wg,c,winner)

	c <- RandBall.Intn(100)
	
	for {
		w, ok := <-winner
		if w!=nil {
			fmt.Printf("\n%v won the match!",w.Name)
			close(winner)
		} 
		if !ok {
			break
		}		
	}

}