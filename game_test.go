package chess

import (
	"testing"

	"github.com/athom/goset"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExportGameStateText(t *testing.T) {
	Convey("initial state", t, func() {
		g := NewGame(6, NewTextFormatter())
		So(
			g.ToText(),
			ShouldEqual,
			`
六五四三二一
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 1 2 3 4 5 6
`,
		)
		So(
			g.ToText2(),
			ShouldEqual,
			`
六五四三二一
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 1 2 3 4 5 6
`,
		)
	})

	Convey("move step", t, func() {
		Convey("with correct move", func() {
			g := NewGame(6, NewTextFormatter())
			g.Move(Pos{0, 0}, Pos{0, 1}, BLACK)
			So(
				g.ToText(),
				ShouldEqual,
				`
六五四三二一
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 1 0 0 0 0 0
 0 2 3 4 5 6
`,
			)
			So(
				g.ToText2(),
				ShouldEqual,
				`
六五四三二 0
 0 0 0 0 0一
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 1 2 3 4 5 6
`,
			)
		})
		Convey("with incorrect move", func() {
			g := NewGame(6, NewTextFormatter())
			Convey("at NONE unit", func() {
				err := g.Move(Pos{0, 1}, Pos{1, 1}, BLACK)
				So(err, ShouldEqual, ErrIllegalMove)
			})
			Convey("at opponent's unit", func() {
				err := g.Move(Pos{5, 5}, Pos{5, 4}, BLACK)
				So(err, ShouldEqual, ErrIllegalMove)
			})
			Convey("at too far distance", func() {
				err := g.Move(Pos{0, 0}, Pos{1, 1}, BLACK)
				So(err, ShouldEqual, ErrIllegalMove)
			})
		})

	})
}

func TestImportGameStateText(t *testing.T) {
	Convey("load valid manual", t, func() {
		g := NewGame(6, NewTextFormatter())
		g.LoadSnapshot(Snapshot{
			Pos{0, 0}: Unit{BLACK, 1},
			Pos{1, 0}: Unit{BLACK, 2},
			Pos{2, 0}: Unit{BLACK, 3},
			Pos{3, 0}: Unit{BLACK, 4},
			Pos{4, 0}: Unit{BLACK, 5},
			Pos{5, 0}: Unit{BLACK, 6},

			Pos{5, 4}: Unit{WHITE, 1},
			Pos{3, 4}: Unit{WHITE, 3},
			Pos{2, 1}: Unit{WHITE, 4},
			Pos{2, 5}: Unit{WHITE, 5},
			Pos{1, 4}: Unit{WHITE, 6},
		})
		So(
			g.ToText(),
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
		So(
			g.ToText2(),
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
			g.LoadSnapshot(Snapshot{
				Pos{0, 0}: Unit{BLACK, 1},
				Pos{1, 0}: Unit{BLACK, 2},
				Pos{2, 0}: Unit{BLACK, 3},
				Pos{3, 0}: Unit{BLACK, 4},
				Pos{4, 0}: Unit{BLACK, 5},
				Pos{5, 0}: Unit{BLACK, 6},

				Pos{5, 5}: Unit{WHITE, 1},
				Pos{4, 5}: Unit{WHITE, 2},
				Pos{1, 4}: Unit{WHITE, 3},
				Pos{2, 5}: Unit{WHITE, 4},
				Pos{1, 5}: Unit{WHITE, 5},
				Pos{0, 5}: Unit{WHITE, 6},
			})
			var ps []Pos
			ps = g.Select(Pos{1, 5})
			So(ps, ShouldBeEmpty)
		})

		Convey("white 5 is surounded loosely", func() {
			g := NewGame(6, NewTextFormatter())
			g.LoadSnapshot(Snapshot{
				Pos{1, 2}: Unit{BLACK, 1},
				Pos{0, 1}: Unit{BLACK, 2},
				Pos{1, 3}: Unit{BLACK, 3},
				Pos{3, 2}: Unit{BLACK, 4},
				Pos{2, 3}: Unit{BLACK, 5},
				Pos{5, 2}: Unit{BLACK, 6},

				Pos{4, 4}: Unit{WHITE, 1},
				Pos{4, 3}: Unit{WHITE, 2},
				Pos{1, 5}: Unit{WHITE, 3},
				Pos{2, 5}: Unit{WHITE, 4},
				Pos{1, 4}: Unit{WHITE, 5},
				Pos{5, 4}: Unit{WHITE, 6},
			})
			var ps []Pos
			ps = g.Select(Pos{1, 4})
			So(ps, ShouldResemble, []Pos{
				Pos{5, 5},
			})
		})
	})
}
