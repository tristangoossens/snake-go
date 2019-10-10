package trisnake

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
	// s "github.com/tristangoossens/snake-go/game"
)

var Snakegame *tl.Game
var arenawidth int = 70
var arenaheight int = 25
var arena *Arena
var food *Food
var score int

func StartGame() {
	Snakegame = tl.NewGame()

	ts := NewTitleScreen()
	ts.AddEntity(ts.TitleText)
	Snakegame.Screen().SetLevel(ts)
	Snakegame.Screen().SetFps(15)
	Snakegame.Start()
}

func NewTitleScreen() *Titlescreen {
	ts := new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	ts.TitleText = tl.NewText(10, 5, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack)
	return ts
}

func NewSidepanel() (*tl.Rectangle, *tl.Text) {
	sidepanel := tl.NewRectangle(arenawidth+1, 0, 30, arenaheight, tl.ColorWhite)
	scoretxt := tl.NewText(arenawidth+2, 1, fmt.Sprintf("Score: %d", score), tl.ColorBlack, tl.ColorWhite)
	return sidepanel, scoretxt
}

func Gameover() {
	gos := new(Gameoverscreen)
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gos.Gameovertext = tl.NewText(10, 5, "Gameover!", tl.ColorWhite, tl.ColorBlack)
	gos.AddEntity(gos.Gameovertext)

	Snakegame.Screen().SetLevel(gos)
}

func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey && event.Key == tl.KeyEnter {

		level := tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlack,
		})

		snake := NewSnake()
		arena = NewArena(arenawidth, arenaheight)
		food = NewFood()
		sidepanel, scoretxt := NewSidepanel()

		level.AddEntity(food)
		level.AddEntity(sidepanel)
		level.AddEntity(scoretxt)
		level.AddEntity(snake)
		level.AddEntity(arena)

		Snakegame.Screen().SetLevel(level)
	}
}
