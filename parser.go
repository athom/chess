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

func NewPlayerStateParser() (r *PlayerStateParser) {
	r = &PlayerStateParser{}
	return
}

type PlayerStateParser struct {
}

func (this *PlayerStateParser) Parse(msg string) (r *PlayerState) {
	for _, s := range []string{
		"q",
		"Q",
		"quit",
		"QUIT",
	} {
		if strings.Contains(msg, s) {
			r = &PlayerState{
				State: IN_ABORT,
			}
			return
		}
	}

	for _, s := range []string{
		"gg",
		"GG",
	} {
		if strings.Contains(msg, s) {
			r = &PlayerState{
				State: IN_GIVEUP,
			}
			return
		}
	}

	r = &PlayerState{
		State: IN_ILLEAGLE_OPERATION,
	}
	var err error
	var p1, p2 Pos
	cmds := strings.Split(msg, ":")
	if len(cmds) != 2 {
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

	mi := NewMoveInfo(p1, p2)
	r = &PlayerState{
		State:    IN_MOVE,
		MoveInfo: mi,
	}

	return
}
