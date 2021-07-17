package main

import "fmt"

var (
	space = "space"
	water = "water"
)

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

type red struct {
	MobileSuit
	transferR
}

type green struct {
	MobileSuit
	transferG
}

type white struct {
	MobileSuit
	transferW
}

// 速度算出
func (m *MobileSuit) calculateSpeedIndex() float32 {
	return float32(m.engine) / float32(m.weight)
}

// ザクタイプ
func (m *MobileSuit) transferZaku(field string, seconds int) int {
	var transfer int
	switch field {
	case space:
		index := m.calculateSpeedIndex()
		transfer = int(index * 1.2 * float32(seconds))
	case water:
		index := m.calculateSpeedIndex()
		transfer = int(index * .4 * float32(seconds))
	default:
		index := m.calculateSpeedIndex()
		transfer = int(index * float32(seconds))
	}
	return transfer
}

// 白い悪魔移動距離算出
func (w *white) transferWhite(field string, seconds int) int {
	var transfer int
	switch field {
	case space:
		index := w.calculateSpeedIndex()
		transfer = int(index * 1.2 * float32(seconds))
	case water:
		index := w.calculateSpeedIndex()
		transfer = int(index * .7 * float32(seconds))
	default:
		index := w.calculateSpeedIndex()
		transfer = int(index * float32(seconds))
	}
	return transfer
}

// 移動距離算出
func (g *green) transferGreen(field string, seconds int) int {
	return g.transferZaku(field, seconds)
}

// 移動距離算出
func (r *red) transferRed(field string, seconds int) int {
	return r.transferZaku(field, seconds)
}

func main() {
	field := space
	white := &white{
		MobileSuit: MobileSuit{
			weight: 43,
			engine: 55,
			name:   "白い悪魔",
		},
	}
	green := &green{
		MobileSuit: MobileSuit{
			weight: 58,
			engine: 43,
			name:   "緑の脇役",
		},
	}
	red := &red{
		MobileSuit: MobileSuit{
			weight: 58,
			engine: 43,
			name:   "赤いの",
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
