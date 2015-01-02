package chess

type Side int

const (
	NONE  Side = 0
	BLACK Side = iota
	WHITE
)

func (this Side) String() string {
	switch this {
	case NONE:
		return "none"
	case BLACK:
		return "black"
	case WHITE:
		return "white"
	}
	return ""
}

type Unit struct {
	Side  Side
	Value int
}

func (this *Unit) Turn(other *Unit) {
	this.Side = other.Side
	this.Value = other.Value
}

func (this *Unit) SetNone() {
	this.Side = NONE
	this.Value = 0
}

func NewUnit(side Side, v int) (r *Unit) {
	r = &Unit{
		Value: v,
		Side:  side,
	}
	return
}
