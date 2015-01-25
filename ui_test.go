package chess

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGameStateWating(t *testing.T) {
	g := NewGame(6)
	g.Move(Pos{0, 0}, Pos{0, 1}, BLACK)
	g.Move(Pos{4, 0}, Pos{5, 4}, WHITE)

	Convey("waiting", t, func() {
		gs := &GameState{State: OUT_WAIT}
		ui := NewConsoleUI(gs)
		So(ui.Render(), ShouldEqual, UI_WAITING)
	})

	Convey("ready", t, func() {
		gs := &GameState{State: OUT_READY}
		ui := NewConsoleUI(gs)
		So(ui.Render(), ShouldEqual, UI_READY)
	})

	Convey("moved", t, func() {
		Convey("in black player's view", func() {
			gs := &GameState{
				State:       OUT_BOARD_UPDATED,
				MyBoardInfo: NewMyBoardInfo(g.BoardInfo(), BLACK),
			}
			ui := NewConsoleUI(gs)
			So(
				ui.Render(),
				ShouldEqual,
				`
六 0四三二一
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
五 0 0 0 0 0
 0 2 3 4 5 6
`,
			)
		})

		Convey("in white player's view", func() {
			gs := &GameState{
				State:       OUT_BOARD_UPDATED,
				MyBoardInfo: NewMyBoardInfo(g.BoardInfo(), WHITE),
			}
			ui := NewConsoleUI(gs)
			So(
				ui.Render(),
				ShouldEqual,
				`
六五四三二 0
 0 0 0 0 0 5
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 1 2 3 4 0 6
`,
			)
		})
	})
}
