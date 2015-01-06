package chess

import "strconv"

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
	Side      Side `json:"side"`
	Value     int  `json:"value"`
	JustMoved bool `json:"just_moved"`
}

func (this *Unit) Name() string {
	if this.Side == NONE {
		return this.Side.String()
	}
	return this.Side.String() + " " + strconv.Itoa(this.Value)
}

func (this *Unit) Turn(other *Unit) {
	this.Side = other.Side
	this.Value = other.Value
	this.JustMoved = true
}

func (this *Unit) Set(info UnitInfo) {
	this.Side = info.Side
	this.Value = info.Value
	this.JustMoved = info.JustMoved
}

func (this *Unit) SetNone() {
	this.Side = NONE
	this.Value = 0
	this.JustMoved = false
}

func NewUnit(side Side, v int) (r *Unit) {
	r = &Unit{
		Value: v,
		Side:  side,
	}
	return
}
