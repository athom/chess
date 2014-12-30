package chess

import "math"

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
	pathX := shortestPaths(p1, Pos{p2.X - stepX, p2.Y})
	for i, _ := range pathX {
		pathX[i] = append(pathX[i], p2)
	}
	pathY := shortestPaths(p1, Pos{p2.X, p2.Y - stepY})
	for i, _ := range pathY {
		pathY[i] = append(pathY[i], p2)
	}

	r = append(pathX, pathY...)
	return
}
