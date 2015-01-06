package chess

import (
	"testing"

	"github.com/athom/goset"
	. "github.com/smartystreets/goconvey/convey"
)

func TestImportGameStateText(t *testing.T) {
	Convey("load valid manual", t, func() {
		g := NewGame(6, NewTextFormatter())
		g.LoadBoardInfo(BoardInfo{
			Units{
				UnitInfo{Unit{BLACK, 1, false}, Pos{0, 0}},
				UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
				UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
				UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
				UnitInfo{Unit{BLACK, 5, false}, Pos{4, 0}},
				UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

				UnitInfo{Unit{WHITE, 1, false}, Pos{5, 4}},
				UnitInfo{Unit{WHITE, 3, false}, Pos{3, 4}},
				UnitInfo{Unit{WHITE, 4, false}, Pos{2, 1}},
				UnitInfo{Unit{WHITE, 5, false}, Pos{2, 5}},
				UnitInfo{Unit{WHITE, 6, false}, Pos{1, 4}},
			},
			6,
		})
		ui := NewConsoleUI(&GameState{
			State:       OUT_BOARD_UPDATED,
			MyBoardInfo: NewMyBoardInfo(g.BoardInfo(), BLACK),
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
		ui = NewConsoleUI(&GameState{
			State:       OUT_BOARD_UPDATED,
			MyBoardInfo: NewMyBoardInfo(g.BoardInfo(), WHITE),
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

func TestGameSelectUnitAndShoePoints(t *testing.T) {
	Convey("default movable positions", t, func() {
		g := NewGame(6, NewTextFormatter())
		var ps []Pos
		ps = g.Select(Pos{0, 0})
		So(
			ps,
			ShouldResemble,
			[]Pos{
				Pos{0, 1},
			},
		)
		ps = g.Select(Pos{1, 0})
		So(
			ps,
			ShouldResemble,
			[]Pos{
				Pos{0, 1},
				Pos{1, 2},
				Pos{2, 1},
			},
		)
		ps = g.Select(Pos{2, 0})
		So(
			ps,
			ShouldResemble,
			[]Pos{
				Pos{0, 1},
				Pos{1, 2},
				Pos{2, 3},
				Pos{3, 2},
				Pos{4, 1},
			},
		)
		ps = g.Select(Pos{3, 0})
		So(goset.IsEqual(ps, []Pos{
			Pos{0, 1},
			Pos{1, 2},
			Pos{2, 1},
			Pos{2, 3},
			Pos{3, 2},
			Pos{3, 4},
			Pos{4, 1},
			Pos{4, 3},
			Pos{5, 2},
		}), ShouldBeTrue)
		ps = g.Select(Pos{4, 0})
		So(goset.IsEqual(ps, []Pos{
			Pos{0, 1},
			Pos{1, 2},
			Pos{2, 1},
			Pos{2, 3},
			Pos{3, 2},
			Pos{3, 4},
			Pos{4, 3},
			Pos{4, 5},
			Pos{5, 2},
			Pos{5, 4},
		}), ShouldBeTrue)
		ps = g.Select(Pos{5, 0})
		So(goset.IsEqual(ps, []Pos{
			Pos{0, 1},
			Pos{1, 2},
			Pos{2, 1},
			Pos{2, 3},
			Pos{3, 2},
			Pos{3, 4},
			Pos{4, 1},
			Pos{4, 3},
			Pos{4, 5},
			Pos{5, 2},
			Pos{5, 4},
		}), ShouldBeTrue)
	})

	Convey("ceitain situation movable positions", t, func() {
		Convey("white 5 is surounded closely", func() {
			g := NewGame(6, NewTextFormatter())
			g.LoadBoardInfo(BoardInfo{
				Units{
					UnitInfo{Unit{BLACK, 1, false}, Pos{0, 0}},
					UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
					UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
					UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
					UnitInfo{Unit{BLACK, 5, false}, Pos{4, 0}},
					UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

					UnitInfo{Unit{WHITE, 1, false}, Pos{5, 5}},
					UnitInfo{Unit{WHITE, 2, false}, Pos{4, 5}},
					UnitInfo{Unit{WHITE, 3, false}, Pos{1, 4}},
					UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
					UnitInfo{Unit{WHITE, 5, false}, Pos{1, 5}},
					UnitInfo{Unit{WHITE, 6, false}, Pos{0, 5}},
				},
				6,
			})
			var ps []Pos
			ps = g.Select(Pos{1, 5})
			So(ps, ShouldBeEmpty)
		})

		Convey("white 5 is surounded loosely", func() {
			g := NewGame(6, NewTextFormatter())
			g.LoadBoardInfo(BoardInfo{
				Units{
					UnitInfo{Unit{BLACK, 1, false}, Pos{1, 2}},
					UnitInfo{Unit{BLACK, 2, false}, Pos{0, 1}},
					UnitInfo{Unit{BLACK, 3, false}, Pos{1, 3}},
					UnitInfo{Unit{BLACK, 4, false}, Pos{3, 2}},
					UnitInfo{Unit{BLACK, 5, false}, Pos{2, 3}},
					UnitInfo{Unit{BLACK, 6, false}, Pos{5, 2}},

					UnitInfo{Unit{WHITE, 1, false}, Pos{4, 4}},
					UnitInfo{Unit{WHITE, 2, false}, Pos{4, 3}},
					UnitInfo{Unit{WHITE, 3, false}, Pos{1, 5}},
					UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
					UnitInfo{Unit{WHITE, 5, false}, Pos{1, 4}},
					UnitInfo{Unit{WHITE, 6, false}, Pos{5, 4}},
				},
				6,
			})
			var ps []Pos
			ps = g.Select(Pos{1, 4})
			So(ps, ShouldResemble, []Pos{
				Pos{5, 5},
			})
		})
	})
}
