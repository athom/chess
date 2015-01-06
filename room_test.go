package chess

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlayerJoinRoom(t *testing.T) {
	Convey("players join empty room togather", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		player3 := NewPlayer(NewConnMocker())
		player4 := NewPlayer(NewConnMocker())
		r := NewRoom(player1, player2, player3, player4)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldEqual, player2)
		So(r.Watchers(), ShouldContain, player3)
		So(r.Watchers(), ShouldContain, player4)
		So(r.JoinableForPlay(), ShouldBeFalse)
	})
	Convey("player join empty room", t, func() {
		player1 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldBeNil)
		So(r.Watchers(), ShouldBeEmpty)
		So(r.JoinableForPlay(), ShouldBeTrue)
	})
	Convey("player join room with 1 black player entered", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldEqual, player2)
		So(r.Watchers(), ShouldBeEmpty)
	})
	Convey("player join room with both black and white player entered", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		player3 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Join(player3)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldEqual, player2)
		So(r.Watchers(), ShouldContain, player3)
	})
}

func TestPlayerLeaveRoom(t *testing.T) {
	Convey("wather leave room", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		player3 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Join(player3)
		r.Leave(player3)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldEqual, player2)
		So(r.Watchers(), ShouldBeEmpty)
		So(r.JoinableForPlay(), ShouldBeFalse)
	})
	Convey("white player leave room", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Leave(player2)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldBeNil)
		So(r.JoinableForPlay(), ShouldBeTrue)
	})
	Convey("black player leave room", t, func() {
		player1 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Leave(player1)
		So(r.BlackPlayer(), ShouldBeNil)
		So(r.JoinableForPlay(), ShouldBeTrue)
	})
}

func TestPlayerLeaveAndJoinRoom(t *testing.T) {
	Convey("white player leave room and rejoin", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Leave(player2)
		r.Join(player2)
		So(r.BlackPlayer(), ShouldEqual, player1)
		So(r.WhitePlayer(), ShouldEqual, player2)
	})
	Convey("black player leave room and rejoin, change player role switched", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Leave(player1)
		r.Join(player1)
		So(r.BlackPlayer(), ShouldEqual, player2)
		So(r.WhitePlayer(), ShouldEqual, player1)
	})
}

func TestGameInRoom(t *testing.T) {
	Convey("no game when there only 1 player", t, func() {
		player1 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		So(r.game, ShouldBeNil)
	})

	Convey("start game when both black and white player are ready", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		So(r.game, ShouldNotBeNil)
	})

	Convey("remove game when black player leave", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Leave(player1)
		So(r.game, ShouldBeNil)
	})

	Convey("remove game when white player leave", t, func() {
		player1 := NewPlayer(NewConnMocker())
		player2 := NewPlayer(NewConnMocker())
		r := NewRoom()
		r.Join(player1)
		r.Join(player2)
		r.Leave(player2)
		So(r.game, ShouldBeNil)
	})
}
