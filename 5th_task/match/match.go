package match

import (
	"math/rand"
	"time"
	"sync"
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

	go tm.Player1.Play(wg,c)
	go tm.Player2.Play(wg,c)

	c <- RandBall.Intn(100)
}