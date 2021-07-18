package main

import "fmt"

var (
	space = "space"
	water = "water"
)

/*******************************
	utility
*******************************/

type MobileSuit struct {
	weight int
	engine int
	name   string
}

type transfer interface {
	transferGtype(string, int) int
	transferZtype(string, int) int
}

type Gtype struct {
	MobileSuit
	transfer
}

type Ztype struct {
	MobileSuit
	transfer
}

// 速度算出
func (m *MobileSuit) calculateSpeedIndex() float32 {
	return float32(m.engine) / float32(m.weight)
}

// Ztype 移動距離算出
func (z *Ztype) transferZtype(field string, seconds int) int {
	var transfer int
	switch field {
	case space:
		index := z.calculateSpeedIndex()
		transfer = int(index * 1.2 * float32(seconds))
	case water:
		index := z.calculateSpeedIndex()
		transfer = int(index * .4 * float32(seconds))
	default:
		index := z.calculateSpeedIndex()
		transfer = int(index * float32(seconds))
	}
	return transfer
}

// Gtype 移動距離算出
func (g *Gtype) transferGtype(field string, seconds int) int {
	var transfer int
	switch field {
	case space:
		index := g.calculateSpeedIndex()
		transfer = int(index * 1.2 * float32(seconds))
	case water:
		index := g.calculateSpeedIndex()
		transfer = int(index * .7 * float32(seconds))
	default:
		index := g.calculateSpeedIndex()
		transfer = int(index * float32(seconds))
	}
	return transfer
}

/*******************************
	instance
*******************************/

type red struct {
	Ztype
}

type green struct {
	Ztype
}

type white struct {
	Gtype
}

/*******************************
	functions
*******************************/

func main() {
	field := space
	seconds := 200

	white := &white{
		Gtype: Gtype{
			MobileSuit: MobileSuit{
				weight: 43,
				engine: 55,
				name:   "白い悪魔",
			},
		},
	}

	green := &green{
		Ztype: Ztype{
			MobileSuit: MobileSuit{
				weight: 58,
				engine: 43,
				name:   "緑の脇役",
			},
		},
	}

	red := &red{
		Ztype: Ztype{
			MobileSuit: MobileSuit{
				weight: 58,
				engine: 43,
				name:   "赤いの",
			},
		},
	}

	transferWhite := white.transferGtype(field, seconds)
	transferGreen := green.transferZtype(field, seconds)
	transferRed := red.transferZtype(field, seconds)
	fmt.Printf(
		"%s: %v, %s: %v, %s: %v",
		white.name, transferWhite, green.name, transferGreen, red.name, transferRed,
	)
}
