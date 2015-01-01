package chess

import "errors"

type Snapshot map[Pos]Unit

func NewGame(size int, formatter Formatter) (r *Game) {
	r = &Game{size: size}
	r.formatter = formatter
	r.reset()

	return
}

type Game struct {
	size      int
	unitMap   map[Pos]*Unit
	formatter Formatter
}

func (this *Game) clear() {
	for _, v := range this.unitMap {
		v.SetNone()
	}
}
func (this *Game) reset() {
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

func (this *Game) Move(srcPos, desPos Pos) (err error) {
	srcUnit, ok := this.unitMap[srcPos]
	if !ok || srcUnit.Side == NONE {
		err = errors.New("illegal move")
		return
	}
	desUnit, ok := this.unitMap[desPos]
	if !ok {
		err = errors.New("illegal move")
		return
	}
	if srcUnit.Value != distance(srcPos, desPos) {
		err = errors.New("illegal move")
		return
	}

	desUnit.Turn(srcUnit)
	srcUnit.SetNone()
	return
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

func (this *Game) ToJson() (r []byte) {
	return
}

func (this *Game) ToText() (r string) {
	r += "\n"
	y := this.size - 1
	for y >= 0 {
		line := ``
		x := 0
		for x < this.size {
			u := this.unitMap[Pos{x, y}]
			line += this.formatter.Fmt(u)
			x += 1
		}
		line += "\n"
		r += line
		y -= 1
	}
	return
}

func (this *Game) LoadSnapshot(s Snapshot) (err error) {
	this.clear()
	for k, v := range s {
		this.unitMap[k].Turn(&v)
	}
	return
}
