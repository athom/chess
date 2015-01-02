package chess

type looper interface {
	Side() Side
	innerInit() int
	innerCondition(int) bool
	innerChange(int) int
	outterInit() int
	outterCondition(int) bool
	outterChange(int) int
}

type blackLooper struct {
	game *Game
}

func (this *blackLooper) Side() Side {
	return BLACK
}
func (this *blackLooper) innerInit() int {
	return 0
}
func (this *blackLooper) innerCondition(i int) bool {
	return i < this.game.size
}
func (this *blackLooper) innerChange(i int) int {
	return i + 1
}
func (this *blackLooper) outterInit() int {
	return this.game.size - 1
}
func (this *blackLooper) outterCondition(i int) bool {
	return i >= 0
}
func (this *blackLooper) outterChange(i int) int {
	return i - 1
}

type whiteLooper struct {
	game *Game
}

func (this *whiteLooper) Side() Side {
	return WHITE
}
func (this *whiteLooper) innerInit() int {
	return this.game.size - 1
}
func (this *whiteLooper) innerCondition(i int) bool {
	return i >= 0
}
func (this *whiteLooper) innerChange(i int) int {
	return i - 1
}
func (this *whiteLooper) outterInit() int {
	return 0
}
func (this *whiteLooper) outterCondition(i int) bool {
	return i < this.game.size
}
func (this *whiteLooper) outterChange(i int) int {
	return i + 1
}

func (this *Game) renderBoardText(lp looper) (r string) {
	r += "\n"
	y := lp.outterInit()
	for lp.outterCondition(y) {
		line := ``
		x := lp.innerInit()
		for lp.innerCondition(x) {
			u := this.unitMap[Pos{x, y}]
			line += this.formatter.Fmt(u, lp.Side())
			x = lp.innerChange(x)
		}
		line += "\n"
		r += line
		y = lp.outterChange(y)
	}
	return
}

func (this *Game) TextView(side Side) (r string) {
	if side == BLACK {
		return this.renderBoardText(&blackLooper{this})
	}
	return this.renderBoardText(&whiteLooper{this})
}
