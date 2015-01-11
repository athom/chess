package chess

import (
	"net"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func NewMailBoxMocker() (r *MailBoxMocker) {
	r = &MailBoxMocker{}
	return
}

type MailBoxMocker struct {
}

func (this *MailBoxMocker) Receive() (r *PlayerState) {
	r = &PlayerState{State: IN_READY}
	return
}

func (this *MailBoxMocker) Send(*GameState) {
	return
}

func NewConnMocker() (r *ConnMocker) {
	r = &ConnMocker{}
	return
}

type ConnMocker struct {
}

func (this *ConnMocker) Read(b []byte) (n int, err error) {
	return
}
func (this *ConnMocker) Write(b []byte) (n int, err error) {
	return
}
func (this *ConnMocker) Close() (err error) {
	return
}
func (this *ConnMocker) LocalAddr() (r net.Addr) {
	return
}
func (this *ConnMocker) RemoteAddr() (r net.Addr) {
	return
}
func (this *ConnMocker) SetDeadline(t time.Time) (err error) {
	return
}
func (this *ConnMocker) SetReadDeadline(t time.Time) (err error) {
	return
}
func (this *ConnMocker) SetWriteDeadline(t time.Time) (err error) {
	return
}

func TestHallJoinPlayer(t *testing.T) {
	return
	Convey("player jion hall", t, func() {
		h := NewHall()
		h.Joins <- NewMailBoxMocker()
		So(len(h.Players()), ShouldEqual, 1)
		So(len(h.Rooms()), ShouldEqual, 1)
		h.Joins <- NewMailBoxMocker()
		time.Sleep(100 * time.Millisecond)
		So(len(h.Players()), ShouldEqual, 2)
		So(len(h.Rooms()), ShouldEqual, 1)
	})
}
