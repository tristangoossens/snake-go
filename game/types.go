package trisnake

import tl "github.com/JoelOtter/termloop"

//Game Object Variables.
var sg *tl.Game
var sp *Sidepanel
var gs *Gamescreen
var ts *Titlescreen
var gop *Gameoptionsscreen

// Own created types.
type direction int
type difficulty int
type colorobject int

// Game options
var Difficulty = "Normal"
var ColorObject = "Snake"

const (
	easy difficulty = iota
	normal
	hard
)

const (
	snake colorobject = iota
	food
	arena
)

const (
	up direction = iota
	down
	left
	right
)

type Titlescreen struct {
	tl.Level
	Logo           *tl.Entity
	GameDifficulty difficulty
	OptionsText    []*tl.Text
}

type Gameoverscreen struct {
	tl.Level
	Logo              *tl.Entity
	Finalstats        []*tl.Text
	OptionsBackground *tl.Rectangle
	OptionsText       []*tl.Text
}

type Gameoptionsscreen struct {
	tl.Level
	StartText *tl.Text

	CurrentColorObjectText *tl.Text
	ObjectBackground       *tl.Rectangle
	ColorObjectOptions     []*tl.Text

	CurrentDifficultyText *tl.Text
	DifficultyBackground  *tl.Rectangle
	DifficultyOptions     []*tl.Text

	ColorPanelBackground *tl.Rectangle
	ColorPanelOptions    []string
	ColorSelectedIcon    *tl.Text
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
	Background     *tl.Rectangle
	Instructions   []string
	ScoreText      *tl.Text
	SpeedText      *tl.Text
	DifficultyText *tl.Text
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
