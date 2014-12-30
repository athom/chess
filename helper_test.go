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

func TestNormalizeX(t *testing.T) {
	Convey("(3,2) (4,9) = 1", t, func() {
		So(normalizeX(Pos{3, 2}, Pos{4, 9}), ShouldEqual, 1)
	})
	Convey("(3,2) (14,99) = 1", t, func() {
		So(normalizeX(Pos{3, 2}, Pos{14, 99}), ShouldEqual, 1)
	})
	Convey("(4,2) (4,2) = 0", t, func() {
		So(normalizeX(Pos{4, 2}, Pos{4, 2}), ShouldEqual, 0)
	})
	Convey("(3,13) (2,11) = -1", t, func() {
		So(normalizeX(Pos{3, 13}, Pos{2, 11}), ShouldEqual, -1)
	})
}

func TestNormalizeY(t *testing.T) {
	Convey("(3,2) (4,9) = 1", t, func() {
		So(normalizeY(Pos{3, 2}, Pos{4, 9}), ShouldEqual, 1)
	})
	Convey("(3,2) (4,99) = 1", t, func() {
		So(normalizeY(Pos{3, 2}, Pos{4, 99}), ShouldEqual, 1)
	})
	Convey("(3,2) (4,2) = 0", t, func() {
		So(normalizeY(Pos{3, 2}, Pos{4, 2}), ShouldEqual, 0)
	})
	Convey("(3,13) (4,11) = -1", t, func() {
		So(normalizeY(Pos{3, 13}, Pos{4, 11}), ShouldEqual, -1)
	})
}

func TestSinglePathX(t *testing.T) {
	Convey("(4,0) (4,0)", t, func() {
		So(singleLinePathX(Pos{4, 0}, Pos{4, 0}), ShouldBeEmpty)
	})
	Convey("(0,0) (4,0)", t, func() {
		So(singleLinePathX(Pos{0, 0}, Pos{4, 0}), ShouldResemble, []Pos{
			Pos{1, 0},
			Pos{2, 0},
			Pos{3, 0},
			Pos{4, 0},
		})
	})
	Convey("(4,0) (1,0)", t, func() {
		So(singleLinePathX(Pos{4, 0}, Pos{1, 0}), ShouldResemble, []Pos{
			Pos{3, 0},
			Pos{2, 0},
			Pos{1, 0},
		})
	})
}

func TestSinglePathY(t *testing.T) {
	Convey("(0,4) (0,4)", t, func() {
		So(singleLinePathY(Pos{0, 4}, Pos{0, 4}), ShouldBeEmpty)
	})
	Convey("(0,0) (0,4)", t, func() {
		So(singleLinePathY(Pos{0, 0}, Pos{0, 4}), ShouldResemble, []Pos{
			Pos{0, 1},
			Pos{0, 2},
			Pos{0, 3},
			Pos{0, 4},
		})
	})
	Convey("(0,4) (0,1)", t, func() {
		So(singleLinePathY(Pos{0, 4}, Pos{0, 1}), ShouldResemble, []Pos{
			Pos{0, 3},
			Pos{0, 2},
			Pos{0, 1},
		})
	})
}

func TestShortestPath(t *testing.T) {
	Convey("(0,0) (0,0)", t, func() {
		So(shortestPaths(Pos{0, 0}, Pos{0, 0}), ShouldBeEmpty)
	})
	Convey("(0,0) (5,0)", t, func() {
		So(shortestPaths(Pos{0, 0}, Pos{5, 0}), ShouldResemble, [][]Pos{
			[]Pos{
				Pos{1, 0},
				Pos{2, 0},
				Pos{3, 0},
				Pos{4, 0},
				Pos{5, 0},
			},
		})
	})
	Convey("(0,4) (0,2)", t, func() {
		So(shortestPaths(Pos{0, 4}, Pos{0, 2}), ShouldResemble, [][]Pos{
			[]Pos{
				Pos{0, 3},
				Pos{0, 2},
			},
		})
	})

	Convey("(0,0) (1,1)", t, func() {
		So(shortestPaths(Pos{0, 0}, Pos{1, 1}), ShouldResemble, [][]Pos{
			[]Pos{
				Pos{0, 1},
				Pos{1, 1},
			},
			[]Pos{
				Pos{1, 0},
				Pos{1, 1},
			},
		})
	})

	Convey("(0,0) (2,3)", t, func() {
		So(shortestPaths(Pos{0, 0}, Pos{2, 3}), ShouldResemble, [][]Pos{
			[]Pos{
				Pos{0, 1},
				Pos{0, 2},
				Pos{0, 3},
				Pos{1, 3},
				Pos{2, 3},
			},
			[]Pos{
				Pos{0, 1},
				Pos{0, 2},
				Pos{1, 2},
				Pos{1, 3},
				Pos{2, 3},
			},
			[]Pos{
				Pos{0, 1},
				Pos{1, 1},
				Pos{1, 2},
				Pos{1, 3},
				Pos{2, 3},
			},
			[]Pos{
				Pos{1, 0},
				Pos{1, 1},
				Pos{1, 2},
				Pos{1, 3},
				Pos{2, 3},
			},
			[]Pos{
				Pos{0, 1},
				Pos{0, 2},
				Pos{1, 2},
				Pos{2, 2},
				Pos{2, 3},
			},
			[]Pos{
				Pos{0, 1},
				Pos{1, 1},
				Pos{1, 2},
				Pos{2, 2},
				Pos{2, 3},
			},
			[]Pos{
				Pos{1, 0},
				Pos{1, 1},
				Pos{1, 2},
				Pos{2, 2},
				Pos{2, 3},
			},
			[]Pos{
				Pos{0, 1},
				Pos{1, 1},
				Pos{2, 1},
				Pos{2, 2},
				Pos{2, 3},
			},
			[]Pos{
				Pos{1, 0},
				Pos{1, 1},
				Pos{2, 1},
				Pos{2, 2},
				Pos{2, 3},
			},
			[]Pos{
				Pos{1, 0},
				Pos{2, 0},
				Pos{2, 1},
				Pos{2, 2},
				Pos{2, 3},
			},
		})
	})
}
