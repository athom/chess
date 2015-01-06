package chess

import "encoding/json"

type UnitInfo struct {
	Unit
	Pos Pos `json:"pos"`
	//Side      Side `json:"side"`
	//Value     int  `json:"value"`
	//JustMoved bool `json:"just_moved"`
}

type Units []UnitInfo

type BoardInfo struct {
	Units
	Size int
}

type MyBoardInfo struct {
	BoardInfo
	Side    Side `json:"side"`
	Movable bool `json:"movable"`
}

func (this *MyBoardInfo) FindUnit(pos Pos) (r UnitInfo) {
	for _, u := range this.BoardInfo.Units {
		if u.Pos == pos {
			return u
		}
	}
	return
}

func (this *MyBoardInfo) ToJson() (r []byte) {
	r, _ = json.Marshal(this)
	return
}

func GameOverBoardInfo(bi BoardInfo, side Side) (r *MyBoardInfo) {
	r = NewMyBoardInfo(bi, side)
        r.Movable = false
        return
}

func NewMyBoardInfo(bi BoardInfo, side Side) (r *MyBoardInfo) {
	r = &MyBoardInfo{bi, side, false}
	if side == NONE {
		return
	}

	if side == WHITE {
		for i, u := range r.BoardInfo.Units {
			r.BoardInfo.Units[i].Pos = flipView(u.Pos, bi.Size, bi.Size)
		}
	}

	for _, ui := range bi.Units {
		if ui.JustMoved {
			r.Movable = ui.Side != side
			return
		}
	}

	r.Movable = side == BLACK
	return
}
