package chess

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistance(t *testing.T) {
	Convey("(0,0) (1,1) = 2", t, func() {
		So(distance(Pos{0, 0}, Pos{1, 1}), ShouldEqual, 2)
	})
	Convey("(0,0) (0,0) = 0", t, func() {
		So(distance(Pos{0, 0}, Pos{0, 0}), ShouldEqual, 0)
	})
	Convey("(0,0) (3,0) = 3", t, func() {
		So(distance(Pos{0, 0}, Pos{3, 0}), ShouldEqual, 3)
	})
	Convey("(0,0) (0,3) = 3", t, func() {
		So(distance(Pos{0, 0}, Pos{0, 3}), ShouldEqual, 3)
	})
	Convey("(3,2) (4,9) = 8", t, func() {
		So(distance(Pos{3, 2}, Pos{4, 9}), ShouldEqual, 8)
	})
}

func TestGameStateText(t *testing.T) {
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
