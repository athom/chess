package chess

type Pair struct {
	Player1 *Player
	Player2 *Player

	cmdParser *CmdParser
	game      *Game
	//incoming  chan *IncommingMessage
	outgoing  chan string
}

func NewPair() (r *Pair) {
	r = &Pair{}
	return
}

func (this *Pair) Waiting() {
}

func (this *Pair) StartGame() {
}
