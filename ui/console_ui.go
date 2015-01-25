package ui

import (
	"fmt"
	"strconv"

	"github.com/athom/chess"
)


const (
	UI_WAITING = `Welcome to yeer's chess
Waiting for other player...`
	UI_READY             = `Found opponent, game start!`
	UI_BOARD_UPDATED     = ``
	UI_ILLEGAL_OPERATION = `Illegal operation!`
	UI_OPPOENENT_ABORT   = `Oppenent leave suddently!`
	UI_OPPOENENT_GIVEUP  = `Oppenent give up, you win!`
	UI_GAME_OVER_WIN     = `Game over, you win!`
	UI_GAME_OVER_LOSE    = `Game over, you lose...`
)

var tipsMap map[chess.OutState]string = map[chess.OutState]string{
	chess.OUT_WAIT:              UI_WAITING,
	chess.OUT_READY:             UI_READY,
	chess.OUT_ILLEGAL_OPERATION: UI_ILLEGAL_OPERATION,
	chess.OUT_BOARD_UPDATED:     UI_BOARD_UPDATED,
	chess.OUT_OPPOENENT_ABORT:   UI_OPPOENENT_ABORT,
	chess.OUT_OPPOENENT_GIVEUP:  UI_OPPOENENT_GIVEUP,
	chess.OUT_GAMEOVER_WIN:      UI_GAME_OVER_WIN,
	chess.OUT_GAMEOVER_LOSE:     UI_GAME_OVER_LOSE,
}

var whiteUnitView map[int]string = map[int]string{
	1: "一",
	2: "二",
	3: "三",
	4: "四",
	5: "五",
	6: "六",
	7: "七",
	8: "八",
	9: "九",
}

func NewConsoleUI(gs *chess.GameState) (r *ConsoleUI) {
	r = &ConsoleUI{gs}
	return
}

type ConsoleUI struct {
	gameState *chess.GameState
}

func (this *ConsoleUI) Render() (r string) {
	r = tipsMap[this.gameState.State]
	if this.gameState.State != chess.OUT_ILLEGAL_OPERATION && this.gameState.MyBoardInfo != nil {
		r += "\n"
		r += this.printBoard()
	}

	fmt.Println(r)
	return
}

func (this *ConsoleUI) printBoard() (r string) {
	myBoard := this.gameState.MyBoardInfo
	size := myBoard.Size
	height := size
	width := size

	y := height - 1
	for y >= 0 {
		line := ``
		x := 0
		for x < width {
			u := myBoard.FindUnit(chess.Pos{x, y})
			line += this.unitAppearance(u.Unit)
			x += 1
		}
		line += "\n"
		r += line
		y -= 1
	}
	return
}

func (this *ConsoleUI) unitAppearance(u chess.Unit) string {
	if u.Side == chess.NONE {
		return " 0"
	}
	side := this.gameState.MyBoardInfo.Side

	if u.Side == side {
		return " " + strconv.Itoa(u.Value)
	}

	return whiteUnitView[u.Value]
}
