package main

import "fmt"

var (
	space = "space"
	water = "water"
)

/*******************************
			utility
*******************************/

type transferR interface {
	transferRed(string, int) int
}

type transferG interface {
	tranferGreen(string, int) int
}

type transferW interface {
	transferWhite(string, int) int
}

type MobileSuit struct {
	weight int
	engine int
	name   string
}

type Gtype struct {
	MobileSuit
}

type Ztype struct {
	MobileSuit
}

// 速度算出
func (m *MobileSuit) calculateSpeedIndex() float32 {
	return float32(m.engine) / float32(m.weight)
}

// ザクタイプ
func (z *Ztype) transferZaku(field string, seconds int) int {
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
		  具体的な値
*******************************/

type red struct {
	Ztype
	transferR
}

type green struct {
	Ztype
	transferG
}

type white struct {
	Gtype
	transferW
}

func (w *white) transferWhite(field string, seconds int) int {
	return w.transferGtype(field, seconds)
}

// 移動距離算出
func (g *green) transferGreen(field string, seconds int) int {
	return g.transferZaku(field, seconds)
}

// 移動距離算出
func (r *red) transferRed(field string, seconds int) int {
	return r.transferZaku(field, seconds)
}

/*******************************
			utility
*******************************/

func main() {
	field := space
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
	transferWhite := white.transferWhite(field, 200)
	transferGreen := green.transferGreen(field, 200)
	transferRed := red.transferRed(field, 200)
	fmt.Printf(
		"%s: %v, %s: %v, %s: %v",
		white.name, transferWhite, green.name, transferGreen, red.name, transferRed,
	)
}
