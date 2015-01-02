package chess

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCmd(t *testing.T) {
	Convey("parse cmd to positions", t, func() {
		parser := NewCmdParser()
		p1, p2, _ := parser.Parse("0,0:0,1")
		So(p1, ShouldResemble, Pos{0, 0})
		So(p2, ShouldResemble, Pos{0, 1})
	})
}
