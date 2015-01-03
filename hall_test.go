package chess

import (
	"net"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

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
	Convey("player jion hall", t, func() {
		h := NewHall()
		h.joins <- NewConnMocker()
		So(len(h.Players()), ShouldEqual, 1)
		h.joins <- NewConnMocker()
		time.Sleep(100 * time.Millisecond)
		So(len(h.Players()), ShouldEqual, 2)
	})
}
