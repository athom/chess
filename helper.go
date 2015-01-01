package chess

import (
	"math"

	"github.com/athom/goset"
)

func left(p Pos) Pos {
	return p.Move(-1, 0)
}
func right(p Pos) Pos {
	return p.Move(1, 0)
}
func up(p Pos) Pos {
	return p.Move(0, 1)
}
func down(p Pos) Pos {
	return p.Move(0, -1)
}

func distance(p1 Pos, p2 Pos) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

func normalize(x1, x2 int) int {
	a := x2 - x1
	b := int(math.Abs(float64(a)))
	if b == 0 {
		return 0
	}
	return a / b
}

func normalizeX(p1 Pos, p2 Pos) int {
	return normalize(p1.X, p2.X)
}

func normalizeY(p1 Pos, p2 Pos) int {
	return normalize(p1.Y, p2.Y)
}

func singleLinePathX(p1 Pos, p2 Pos) (r []Pos) {
	direction := normalizeX(p1, p2)
	return singleLinePath(p1.X, p2.X, direction, func(v int) Pos {
		return Pos{v, p1.Y}
	})
}
func singleLinePathY(p1 Pos, p2 Pos) (r []Pos) {
	direction := normalizeY(p1, p2)
	return singleLinePath(p1.Y, p2.Y, direction, func(v int) Pos {
		return Pos{p1.X, v}
	})
}
func singleLinePath(a, b, d int, f func(x int) Pos) (r []Pos) {
	x := a
	for x != b {
		x += d
		r = append(r, f(x))
	}
	return
}

func shortestPaths(p1 Pos, p2 Pos) (r [][]Pos) {
	// at the same point return none paths
	if p1 == p2 {
		return
	}

	// same x value pararell with y axis, single line path
	if p1.X == p2.X {
		r = append(r, singleLinePathY(p1, p2))
		return
	}

	// same y value pararell with x axis, single line path
	if p1.Y == p2.Y {
		r = append(r, singleLinePathX(p1, p2))
		return
	}

	// cauculate paths recursively
	stepX := normalizeX(p1, p2)
	stepY := normalizeY(p1, p2)
	pathX := shortestPaths(p1, p2.Move(-stepX, 0))
	for i, _ := range pathX {
		pathX[i] = append(pathX[i], p2)
	}
	pathY := shortestPaths(p1, p2.Move(0, -stepY))
	for i, _ := range pathY {
		pathY[i] = append(pathY[i], p2)
	}

	r = append(pathX, pathY...)
	return
}

func reachablePaths(p1 Pos, p2 Pos, steps int) (r [][]Pos) {
	// at the same point return none paths
	if p1 == p2 {
		return
	}
	// too far, can not reach
	distance := distance(p1, p2)
	if distance > steps {
		return
	}

	// exactly steps to reach to end point
	if distance == steps {
		return shortestPaths(p1, p2)
	}

	paths := pathsFromNextStep(p1, p2, steps, func(oldPos Pos) Pos {
		return oldPos.Move(-1, 0)
	})
	r = append(r, paths...)
	paths = pathsFromNextStep(p1, p2, steps, func(oldPos Pos) Pos {
		return oldPos.Move(1, 0)
	})
	r = append(r, paths...)
	paths = pathsFromNextStep(p1, p2, steps, func(oldPos Pos) Pos {
		return oldPos.Move(0, -1)
	})
	r = append(r, paths...)
	paths = pathsFromNextStep(p1, p2, steps, func(oldPos Pos) Pos {
		return oldPos.Move(0, 1)
	})
	r = append(r, paths...)
	return
}

func pathsFromNextStep(p1 Pos, p2 Pos, steps int, nextMove func(Pos) Pos) (r [][]Pos) {
	p := nextMove(p1)
	if p == p2 {
		return
	}
	paths := reachablePaths(p, p2, steps-1)
	for _, path := range paths {
		// ensure visit same place only once
		if goset.IsIncluded(path, p1) {
			continue
		}
		path = append([]Pos{p}, path...)
		r = append(r, path)
	}
	return
}

func reachRange(p Pos, steps int) (r []Pos) {
	if steps < 1 {
		return
	}
	if steps == 1 {
		r = append(r, p.Move(-1, 0))
		r = append(r, p.Move(0, 1))
		r = append(r, p.Move(1, 0))
		r = append(r, p.Move(0, -1))
		return
	}

	if steps == 2 {
		r = append(r, p.Move(-1, -1))
		r = append(r, p.Move(-2, 0))
		r = append(r, p.Move(-1, 1))
		r = append(r, p.Move(0, 2))
		r = append(r, p.Move(1, 1))
		r = append(r, p.Move(2, 0))
		r = append(r, p.Move(1, -1))
		r = append(r, p.Move(0, -2))
		return
	}

	r = reachRange(p, steps-2)
	for _, pos := range r {
		r = append(r, reachRange(pos, 2)...)
	}
	r = goset.Uniq(r).([]Pos)
	r = goset.RemoveElement(r, p).([]Pos)
	return
}

func insideReachRange(p Pos, steps int, length int) (r []Pos) {
	fullRange := reachRange(p, steps)
	for _, pos := range fullRange {
		// filter the outside positions
		if pos.IsOutside(length) {
			continue
		}
		r = append(r, pos)
	}
	return
}
