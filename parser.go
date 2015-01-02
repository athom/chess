package chess

import (
	"errors"
	"strconv"
	"strings"
)

var CommandError = errors.New("wrong command")

func NewCmdParser() (r *CmdParser) {
	r = &CmdParser{}
	return
}

type CmdParser struct {
}

func (this *CmdParser) Parse(cmd string) (p1 Pos, p2 Pos, err error) {
	cmds := strings.Split(cmd, ":")
	if len(cmds) != 2 {
		err = CommandError
		return
	}
	p1, err = parsePoint(cmds[0])
	if err != nil {
		return
	}
	p2, err = parsePoint(cmds[1])
	if err != nil {
		return
	}
	return
}

func parsePoint(subCmd string) (r Pos, err error) {
	subCmd = strings.TrimRight(subCmd, "\n")
	ps := strings.Split(subCmd, ",")
	if len(ps) != 2 {
		err = CommandError
		return
	}

	x, err := strconv.Atoi(ps[0])
	if err != nil {
		err = CommandError
		return
	}
	y, err := strconv.Atoi(ps[1])
	if err != nil {
		err = CommandError
		return
	}

	r = Pos{x, y}
	return
}
