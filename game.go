package chess

import "errors"

var (
	ErrIllegalMove      = errors.New("illegal move")
	ErrGameOverWhiteWin = errors.New("game over, white win")
	ErrGameOverBlackWin = errors.New("game over, black win")
)

func NewGame(size int) (r *Game) {
	r = &Game{size: size}
	r.reset()

	return
}

type Game struct {
	size    int
	unitMap map[Pos]*Unit
	over    bool
}

func (this *Game) clear() {
	for _, v := range this.unitMap {
		v.SetNone()
	}
}
func (this *Game) reset() {
	this.over = false
	this.unitMap = make(map[Pos]*Unit)
	var x = 0
	var y = 0
	for x < this.size {
		pos := Pos{X: x, Y: y}
		unit := NewUnit(BLACK, x+1)
		this.unitMap[pos] = unit
		x += 1
	}
	y = 1
	for y < this.size-1 {
		x = 0
		for x < this.size {
			pos := Pos{X: x, Y: y}
			unit := NewUnit(NONE, 0)
			this.unitMap[pos] = unit
			x += 1
		}
		y += 1
	}
	y = this.size - 1
	x = 0
	for x < this.size {
		pos := Pos{X: x, Y: y}
		unit := NewUnit(WHITE, this.size-x)
		this.unitMap[pos] = unit
		x += 1
	}

}

func (this *Game) Move(srcPos, desPos Pos, side Side) (err error) {
	if side == WHITE {
		srcPos = this.transform(srcPos)
		desPos = this.transform(desPos)
	}

	srcUnit, ok := this.unitMap[srcPos]
	if !ok || srcUnit.Side != side {
		err = ErrIllegalMove
		return
	}
	desUnit, ok := this.unitMap[desPos]
	if !ok {
		err = ErrIllegalMove
		return
	}
	// don't eat your self
	if desUnit.Side == side {
		err = ErrIllegalMove
		return
	}

	if !this.reachable(srcUnit, srcPos, desPos) {
		err = ErrIllegalMove
		return
	}

	//if this.blackKing
	if desUnit.Value == 1 {
		if desUnit.Side == WHITE {
			err = ErrGameOverBlackWin
		} else {
			err = ErrGameOverWhiteWin
		}
		this.over = true
	}

	this.clearJustMoved()
	desUnit.Turn(srcUnit)
	srcUnit.SetNone()
	return
}

func (this *Game) clearJustMoved() {
	for _, u := range this.unitMap {
		u.JustMoved = false
	}
}

func (this *Game) reachable(u *Unit, fromPos, toPos Pos) bool {
	paths := reachablePaths(fromPos, toPos, u.Value)
	badPathCount := 0
	for _, path := range paths {
		for i, point := range path {
			if point.IsOutside(this.size) {
				badPathCount += 1
				break
			}
			side := this.unitMap[point].Side
			if side != NONE {
				if i == len(path)-1 && side != u.Side {
					break
				}
				badPathCount += 1
				break
			}
		}
	}
	return badPathCount < len(paths)
}

func (this *Game) Select(pos Pos) (r []Pos) {
	unit, ok := this.unitMap[pos]
	if !ok || unit.Side == NONE {
		return
	}

	points := insideReachRange(pos, unit.Value, this.size)
	for _, p := range points {
		// filter points contains unit in my side
		if u, ok := this.unitMap[p]; ok {
			if u.Side == unit.Side {
				continue
			}
			if !this.reachable(unit, pos, p) {
				continue
			}
		}
		r = append(r, p)
	}
	return
}

func (this *Game) BoardInfo() (r BoardInfo) {
	r = BoardInfo{}
	r.Size = this.size
	for pos, unit := range this.unitMap {
		if unit.Side == NONE {
			continue
		}
		r.Units = append(
			r.Units,
			UnitInfo{*unit, pos},
		)
	}
	return
}

func (this *Game) LoadBoardInfo(bi BoardInfo) (err error) {
	this.clear()
	for _, u := range bi.Units {
		this.unitMap[u.Pos].Set(u)
	}
	return
}

func (this *Game) transform(p Pos) (r Pos) {
	r = flipView(p, this.size, this.size)
	return
}
