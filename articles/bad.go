package main

import "fmt"

var (
	space = "space"
	water = "water"
)

type Transfer interface {
	transferWhite(string, int) int
	transferGreen(string, int) int
	transferRed(string, int) int
}

type MobileSuit struct {
	weight int
	engine int
	name   string
	Transfer
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
func (m *MobileSuit) transferWhite(field string, seconds int) int {
	var transfer int
	switch field {
	case space:
		index := m.calculateSpeedIndex()
		transfer = int(index * 1.2 * float32(seconds))
	case water:
		index := m.calculateSpeedIndex()
		transfer = int(index * .7 * float32(seconds))
	default:
		index := m.calculateSpeedIndex()
		transfer = int(index * float32(seconds))
	}
	return transfer
}

// 移動距離算出
func (m *MobileSuit) transferGreen(field string, seconds int) int {
	return m.transferZaku(field, seconds)
}

// 移動距離算出
func (m *MobileSuit) transferRed(field string, seconds int) int {
	return m.transferZaku(field, seconds)
}

func main() {
	field := space
	white := &MobileSuit{
		weight: 43,
		engine: 55,
		name:   "白い悪魔",
	}
	green := &MobileSuit{
		weight: 58,
		engine: 43,
		name:   "緑の脇役",
	}
	red := &MobileSuit{
		weight: 58,
		engine: 43,
		name:   "赤いの",
	}
	transferWhite := white.transferWhite(field, 200)
	transferGreen := green.transferGreen(field, 200)
	transferRed := red.transferRed(field, 200)
	fmt.Printf(
		"%s: %v, %s: %v, %s: %v",
		white.name, transferWhite, green.name, transferGreen, red.name, transferRed,
	)
}
