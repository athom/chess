package chess

import (
	"testing"

	"github.com/athom/goset"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBoardInfo(t *testing.T) {
	Convey("board info size", t, func() {
		g := NewGame(6, NewTextFormatter())
		So(g.BoardInfo().Size, ShouldEqual, 6)

		g = NewGame(7, NewTextFormatter())
		So(g.BoardInfo().Size, ShouldEqual, 7)
	})

	Convey("initial game state", t, func() {
		g := NewGame(6, NewTextFormatter())
		So(
			goset.IsEqual(
				g.BoardInfo().Units,
				Units{
					UnitInfo{Unit{BLACK, 1, false}, Pos{0, 0}},
					UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
					UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
					UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
					UnitInfo{Unit{BLACK, 5, false}, Pos{4, 0}},
					UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

					UnitInfo{Unit{WHITE, 6, false}, Pos{0, 5}},
					UnitInfo{Unit{WHITE, 5, false}, Pos{1, 5}},
					UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
					UnitInfo{Unit{WHITE, 3, false}, Pos{3, 5}},
					UnitInfo{Unit{WHITE, 2, false}, Pos{4, 5}},
					UnitInfo{Unit{WHITE, 1, false}, Pos{5, 5}},
				},
			),
			ShouldBeTrue,
		)
	})

	Convey("black 1 is moved", t, func() {
		g := NewGame(6, NewTextFormatter())
		g.Move(Pos{0, 0}, Pos{0, 1}, BLACK)
		So(
			goset.IsEqual(
				g.BoardInfo().Units,
				Units{
					UnitInfo{Unit{BLACK, 1, true}, Pos{0, 1}},
					UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
					UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
					UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
					UnitInfo{Unit{BLACK, 5, false}, Pos{4, 0}},
					UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

					UnitInfo{Unit{WHITE, 6, false}, Pos{0, 5}},
					UnitInfo{Unit{WHITE, 5, false}, Pos{1, 5}},
					UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
					UnitInfo{Unit{WHITE, 3, false}, Pos{3, 5}},
					UnitInfo{Unit{WHITE, 2, false}, Pos{4, 5}},
					UnitInfo{Unit{WHITE, 1, false}, Pos{5, 5}},
				},
			),
			ShouldBeTrue,
		)
	})

	Convey("black 5 is moved", t, func() {
		g := NewGame(6, NewTextFormatter())
		g.Move(Pos{4, 0}, Pos{5, 4}, BLACK)
		So(
			goset.IsEqual(
				g.BoardInfo().Units,
				Units{
					UnitInfo{Unit{BLACK, 1, false}, Pos{0, 0}},
					UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
					UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
					UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
					UnitInfo{Unit{BLACK, 5, true}, Pos{5, 4}},
					UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

					UnitInfo{Unit{WHITE, 6, false}, Pos{0, 5}},
					UnitInfo{Unit{WHITE, 5, false}, Pos{1, 5}},
					UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
					UnitInfo{Unit{WHITE, 3, false}, Pos{3, 5}},
					UnitInfo{Unit{WHITE, 2, false}, Pos{4, 5}},
					UnitInfo{Unit{WHITE, 1, false}, Pos{5, 5}},
				},
			),
			ShouldBeTrue,
		)
	})

	Convey("white 5 eats black 1", t, func() {
		g := NewGame(6, NewTextFormatter())
		g.Move(Pos{0, 0}, Pos{0, 1}, BLACK)
		g.Move(Pos{4, 0}, Pos{5, 4}, WHITE)
		So(
			goset.IsEqual(
				g.BoardInfo().Units,
				Units{
					UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
					UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
					UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
					UnitInfo{Unit{BLACK, 5, false}, Pos{4, 0}},
					UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

					UnitInfo{Unit{WHITE, 6, false}, Pos{0, 5}},
					UnitInfo{Unit{WHITE, 5, true}, Pos{0, 1}},
					UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
					UnitInfo{Unit{WHITE, 3, false}, Pos{3, 5}},
					UnitInfo{Unit{WHITE, 2, false}, Pos{4, 5}},
					UnitInfo{Unit{WHITE, 1, false}, Pos{5, 5}},
				},
			),
			ShouldBeTrue,
		)
	})
}
func TestMyBoardInfo(t *testing.T) {
	Convey("white 5 eats black 1", t, func() {
		g := NewGame(6, NewTextFormatter())
		g.Move(Pos{0, 0}, Pos{0, 1}, BLACK)
		g.Move(Pos{4, 0}, Pos{5, 4}, WHITE)
		bi := g.BoardInfo()
		Convey("in balck player's view", func() {
			So(
				goset.IsEqual(
					NewMyBoardInfo(bi, BLACK).Units,
					Units{
						UnitInfo{Unit{BLACK, 2, false}, Pos{1, 0}},
						UnitInfo{Unit{BLACK, 3, false}, Pos{2, 0}},
						UnitInfo{Unit{BLACK, 4, false}, Pos{3, 0}},
						UnitInfo{Unit{BLACK, 5, false}, Pos{4, 0}},
						UnitInfo{Unit{BLACK, 6, false}, Pos{5, 0}},

						UnitInfo{Unit{WHITE, 6, false}, Pos{0, 5}},
						UnitInfo{Unit{WHITE, 5, true}, Pos{0, 1}},
						UnitInfo{Unit{WHITE, 4, false}, Pos{2, 5}},
						UnitInfo{Unit{WHITE, 3, false}, Pos{3, 5}},
						UnitInfo{Unit{WHITE, 2, false}, Pos{4, 5}},
						UnitInfo{Unit{WHITE, 1, false}, Pos{5, 5}},
					},
				),
				ShouldBeTrue,
			)
		})
		Convey("in white player's view", func() {
			So(
				goset.IsEqual(
					NewMyBoardInfo(bi, WHITE).Units,
					Units{
						UnitInfo{Unit{BLACK, 2, false}, Pos{4, 5}},
						UnitInfo{Unit{BLACK, 3, false}, Pos{3, 5}},
						UnitInfo{Unit{BLACK, 4, false}, Pos{2, 5}},
						UnitInfo{Unit{BLACK, 5, false}, Pos{1, 5}},
						UnitInfo{Unit{BLACK, 6, false}, Pos{0, 5}},

						UnitInfo{Unit{WHITE, 6, false}, Pos{5, 0}},
						UnitInfo{Unit{WHITE, 5, true}, Pos{5, 4}},
						UnitInfo{Unit{WHITE, 4, false}, Pos{3, 0}},
						UnitInfo{Unit{WHITE, 3, false}, Pos{2, 0}},
						UnitInfo{Unit{WHITE, 2, false}, Pos{1, 0}},
						UnitInfo{Unit{WHITE, 1, false}, Pos{0, 0}},
					},
				),
				ShouldBeTrue,
			)
		})
	})
	Convey("movable", t, func() {
		Convey("ready state", func() {
			g := NewGame(6, NewTextFormatter())
			bi := g.BoardInfo()
			So(NewMyBoardInfo(bi, WHITE).Movable, ShouldBeFalse)
			So(NewMyBoardInfo(bi, BLACK).Movable, ShouldBeTrue)
		})

		Convey("going state", func() {
			g := NewGame(6, NewTextFormatter())
			g.Move(Pos{0, 0}, Pos{0, 1}, BLACK)
			bi := g.BoardInfo()
			So(NewMyBoardInfo(bi, WHITE).Movable, ShouldBeTrue)
			So(NewMyBoardInfo(bi, BLACK).Movable, ShouldBeFalse)
		})
	})
}
