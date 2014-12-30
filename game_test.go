package chess

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExportGameStateText(t *testing.T) {
	Convey("initial state", t, func() {
		g := NewGame(6, NewTextFormatter())
		So(
			g.ToText(),
			ShouldEqual,
			`
 6 5 4 3 2 1
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
一二三四五六
`,
		)
	})

	Convey("movable positions", t, func() {
		//g := NewGame(6, NewTextFormatter())
		//ps := g.Select(Pos{0, 0})
		//So(
		//ps,
		//ShouldEqual,
		//[]Pos{Pos{0, 1}},
		//)
	})

	Convey("move step", t, func() {
		Convey("with correct move", func() {
			g := NewGame(6, NewTextFormatter())
			g.Move(Pos{0, 0}, Pos{0, 1})
			So(
				g.ToText(),
				ShouldEqual,
				`
 6 5 4 3 2 1
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
一 0 0 0 0 0
 0二三四五六
`,
			)
		})
		Convey("with incorrect move", func() {
			g := NewGame(6, NewTextFormatter())
			Convey("at NONE unit", func() {
				err := g.Move(Pos{0, 1}, Pos{1, 1})
				So(err.Error(), ShouldEqual, "illegal move")
			})
			Convey("at too far distance", func() {
				err := g.Move(Pos{0, 0}, Pos{1, 1})
				So(err.Error(), ShouldEqual, "illegal move")
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
 0 0 5 0 0 0
 0 6 0 3 0 1
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 4 0 0 0
一二三四五六
`,
		)
	})
}
