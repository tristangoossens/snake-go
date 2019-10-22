package trisnake

import tl "github.com/JoelOtter/termloop"

// Variables.
var Snakegame *tl.Game
var level tl.Level
var arena *Arena
var fps float64 = 12
var food *Food
var snake *Snake
var score int
var fpstext *tl.Text
var scoretxt *tl.Text
var sidepanel *tl.Rectangle
var utf8support bool

// Own created types.
type direction int
type gamemode int

// Const for gamemode is not yet working in version 1
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

type Titlescreen struct {
	tl.Level
	TitleText *tl.Text
}

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

type Gameoverscreen struct {
	tl.Level
	Gameovertext    *tl.Text
	Finalscore      *tl.Text
	Gameoveroptions *tl.Rectangle
}

type Food struct {
	*tl.Entity
	Foodposition Coordinates
	Emoji        rune
}
