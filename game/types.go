package trisnake

import tl "github.com/JoelOtter/termloop"

// Own created types

type direction int
type gamemode int

// Consts

const (
	easy gamemode = iota
	medium
	hard
)

const (
	up direction = iota
	down
	left
	right
)

// Structs

type Arena struct {
	*tl.Entity
	Width       int
	Height      int
	ArenaBorder map[Coordinates]int
}

type Snake struct {
	*tl.Entity
	Direction  direction
	Length     int
	Bodylength []Coordinates
	Speed      int
}

type Coordinates struct {
	X int
	Y int
}
