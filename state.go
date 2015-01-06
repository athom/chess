package chess

import "encoding/json"

type InState int
type OutState int

const (
	IN_READY InState = 0
	IN_MOVE  InState = iota
	IN_ABORT
	IN_GIVEUP
	IN_ILLEAGLE_OPERATION
)

const (
	OUT_WAIT  OutState = 0
	OUT_READY OutState = iota
	OUT_ILLEGAL_OPERATION
	OUT_BOARD_UPDATED
	OUT_OPPOENENT_ABORT
	OUT_OPPOENENT_GIVEUP
	OUT_GAMEOVER_WIN
	OUT_GAMEOVER_LOSE
	OUT_GAMEOVER_FOR_WATCHER
)

func (this InState) String() string {
	switch this {
	case IN_READY:
		return "ready"
	case IN_MOVE:
		return "move"
	case IN_ABORT:
		return "abort"
	case IN_GIVEUP:
		return "giveup"
	case IN_ILLEAGLE_OPERATION:
		return "illegal operation"
	}
	return ""
}

func (this OutState) String() string {
	switch this {
	case OUT_WAIT:
		return "wait"
	case OUT_READY:
		return "ready"
	case OUT_ILLEGAL_OPERATION:
		return "illegal operation"
	case OUT_BOARD_UPDATED:
		return "moved"
	case OUT_OPPOENENT_GIVEUP:
		return "opponent give up"
	case OUT_OPPOENENT_ABORT:
		return "opponent abort game"
	case OUT_GAMEOVER_WIN:
		return "you win"
	case OUT_GAMEOVER_LOSE:
		return "you lose"
	case OUT_GAMEOVER_FOR_WATCHER:
		return "game over"
	}
	return ""
}

func NewMoveInfo(fromPos, toPos Pos) (r *MoveInfo) {
	r = &MoveInfo{fromPos, toPos}
	return
}

type MoveInfo struct {
	FromPos Pos `json:"fromPos"`
	ToPos   Pos `json:"toPos"`
}

type PlayerState struct {
	Id       string    `json:"id"`
	Side     Side      `json:"side"`
	State    InState   `json:"state"`
	MoveInfo *MoveInfo `json:"moveInfo"`
}

type GameState struct {
	State       OutState     `json:"state"`
	MyBoardInfo *MyBoardInfo `json:"boardInfo"`
}

func (this GameState) ToJson() (r []byte) {
	r, _ = json.Marshal(this)
	return
}
