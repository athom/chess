package chess

type Formatter interface {
	Fmt(*Unit) string
}

func NewTextFormatter() (r *TextFormatter) {
	r = &TextFormatter{
		whiteLooking: map[int]string{
			1: " 1",
			2: " 2",
			3: " 3",
			4: " 4",
			5: " 5",
			6: " 6",
			7: " 7",
			8: " 8",
			9: " 9",
		},

		blackLooking: map[int]string{
			1: "一",
			2: "二",
			3: "三",
			4: "四",
			5: "五",
			6: "六",
			7: "七",
			8: "八",
			9: "九",
		},
	}
	return
}

type TextFormatter struct {
	blackLooking map[int]string
	whiteLooking map[int]string
}

func (this *TextFormatter) Fmt(u *Unit) (r string) {
	switch u.Side {
	case BLACK:
		r = this.blackLooking[u.Value]
	case WHITE:
		r = this.whiteLooking[u.Value]
	default:
		r = " 0"
	}
	return
}
