package chess

type Pos struct {
	X int
	Y int
}

func (p *Pos) Move(x, y int) Pos {
	return Pos{p.X + x, p.Y + y}
}

func (p *Pos) IsInside(length int) bool {
	return p.X < 0 || p.Y < 0 || p.X > length-1 || p.Y > length-1
}
