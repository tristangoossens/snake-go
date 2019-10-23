package trisnake

import tl "github.com/JoelOtter/termloop"

//Game Object Variables.
var sg *tl.Game
var sp *Sidepanel
var gs *Gamescreen
var ts *Titlescreen
var gop *Gameoptionsscreen

var utf8support bool

// Own created types.
type direction int
type difficulty int

// Const for gamemode is not yet working in version 1
const (
	easy difficulty = iota
	normal
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
	GameDifficulty difficulty
	TitleText      *tl.Text
	OptionsText    *tl.Text
}

type Gameoverscreen struct {
	tl.Level
	Gameovertext    *tl.Text
	Finalscore      *tl.Text
	Gameoveroptions *tl.Rectangle
}

type Gameoptionsscreen struct {
	tl.Level
	CurrentDifficultyText *tl.Text
	DifficultyBackground  *tl.Rectangle
	DifficultyEasy        *tl.Text
	DifficultyNormal      *tl.Text
	DifficultyHard        *tl.Text
}

type Gamescreen struct {
	tl.Level
	FPS             float64
	Score           int
	SnakeEntity     *Snake
	FoodEntity      *Food
	ArenaEntity     *Arena
	SidepanelObject *Sidepanel
}

type Sidepanel struct {
	Background   *tl.Rectangle
	Instructions []string
	ScoreText    *tl.Text
	SpeedText    *tl.Text
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

type Food struct {
	*tl.Entity
	Foodposition Coordinates
	Emoji        rune
}

type Coordinates struct {
	X int
	Y int
}
