package chess

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExportGameStateText(t *testing.T) {
	Convey("initial state", t, func() {
		g := NewGame(6, NewTextFormatter())
		So(
			g.TextView(BLACK),
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
			g.TextView(WHITE),
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
				g.TextView(BLACK),
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
				g.TextView(WHITE),
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
