package chess

type Formatter interface {
	Fmt(*Unit, Side) string
}

func NewTextFormatter() (r *TextFormatter) {
	r = &TextFormatter{
		blackLooking: map[int]string{
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

		whiteLooking: map[int]string{
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

func (this *TextFormatter) Fmt(u *Unit, sideView Side) (r string) {
	switch u.Side {
	case BLACK:
		if sideView == BLACK {
			r = this.blackLooking[u.Value]
		} else {
			r = this.whiteLooking[u.Value]
		}
	case WHITE:
		if sideView == BLACK {
			r = this.whiteLooking[u.Value]
		} else {
			r = this.blackLooking[u.Value]
		}
	default:
		r = " 0"
	}
	return
}
