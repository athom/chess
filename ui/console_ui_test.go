package ui

import (
	"testing"

	"github.com/athom/chess"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGameStates(t *testing.T) {
	g := chess.NewGame(6)
	g.Move(chess.Pos{0, 0}, chess.Pos{0, 1}, chess.BLACK)
	g.Move(chess.Pos{4, 0}, chess.Pos{5, 4}, chess.WHITE)

	Convey("waiting", t, func() {
		gs := &chess.GameState{State: chess.OUT_WAIT}
		ui := NewConsoleUI(gs)
		So(ui.Render(), ShouldEqual, UI_WAITING)
	})

	Convey("ready", t, func() {
		gs := &chess.GameState{State: chess.OUT_READY}
		ui := NewConsoleUI(gs)
		So(ui.Render(), ShouldEqual, UI_READY)
	})

	Convey("moved", t, func() {
		Convey("in black player's view", func() {
			gs := &chess.GameState{
				State:       chess.OUT_BOARD_UPDATED,
				MyBoardInfo: chess.NewMyBoardInfo(g.BoardInfo(), chess.BLACK),
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
			gs := &chess.GameState{
				State:       chess.OUT_BOARD_UPDATED,
				MyBoardInfo: chess.NewMyBoardInfo(g.BoardInfo(), chess.WHITE),
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

func TestImportGameStateText(t *testing.T) {
	Convey("load valid manual", t, func() {
		g := chess.NewGame(6)
		g.LoadBoardInfo(chess.BoardInfo{
			chess.Units{
				chess.UnitInfo{chess.Unit{chess.BLACK, 1, false}, chess.Pos{0, 0}},
				chess.UnitInfo{chess.Unit{chess.BLACK, 2, false}, chess.Pos{1, 0}},
				chess.UnitInfo{chess.Unit{chess.BLACK, 3, false}, chess.Pos{2, 0}},
				chess.UnitInfo{chess.Unit{chess.BLACK, 4, false}, chess.Pos{3, 0}},
				chess.UnitInfo{chess.Unit{chess.BLACK, 5, false}, chess.Pos{4, 0}},
				chess.UnitInfo{chess.Unit{chess.BLACK, 6, false}, chess.Pos{5, 0}},

				chess.UnitInfo{chess.Unit{chess.WHITE, 1, false}, chess.Pos{5, 4}},
				chess.UnitInfo{chess.Unit{chess.WHITE, 3, false}, chess.Pos{3, 4}},
				chess.UnitInfo{chess.Unit{chess.WHITE, 4, false}, chess.Pos{2, 1}},
				chess.UnitInfo{chess.Unit{chess.WHITE, 5, false}, chess.Pos{2, 5}},
				chess.UnitInfo{chess.Unit{chess.WHITE, 6, false}, chess.Pos{1, 4}},
			},
			6,
		})
		ui := NewConsoleUI(&chess.GameState{
			State:       chess.OUT_BOARD_UPDATED,
			MyBoardInfo: chess.NewMyBoardInfo(g.BoardInfo(), chess.BLACK),
		})
		So(
			ui.Render(),
			ShouldEqual,
			`
 0 0五 0 0 0
 0六 0三 0一
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0四 0 0 0
 1 2 3 4 5 6
`,
		)
		ui = NewConsoleUI(&chess.GameState{
			State:       chess.OUT_BOARD_UPDATED,
			MyBoardInfo: chess.NewMyBoardInfo(g.BoardInfo(), chess.WHITE),
		})
		So(
			ui.Render(),
			ShouldEqual,
			`
六五四三二一
 0 0 0 4 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 1 0 3 0 6 0
 0 0 0 5 0 0
`,
		)
	})
}
