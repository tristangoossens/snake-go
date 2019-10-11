package trisnake

import (
	"fmt"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
	// s "github.com/tristangoossens/snake-go/game"
)

var Snakegame *tl.Game
var arenawidth int = 70
var arenaheight int = 25
var arena *Arena
var food *Food
var score int
var scoretxt *tl.Text
var sidepanel *tl.Rectangle
var utf8support bool

func StartGame() {
	utf8support = !strings.Contains(os.Getenv("LANG"), "C.UTF-8")
	Snakegame = tl.NewGame()

	ts := NewTitleScreen()
	ts.AddEntity(ts.TitleText)
	Snakegame.Screen().SetLevel(ts)
	Snakegame.Screen().SetFps(12)
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
	scoretxt = tl.NewText(arenawidth+2, 1, fmt.Sprintf("Score: %d", score), tl.ColorBlack, tl.ColorWhite)
	return sidepanel, scoretxt
}

func Gameover() {
	gos := new(Gameoverscreen)
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gos.Gameovertext = tl.NewText(10, 5, "Gameover!", tl.ColorWhite, tl.ColorBlack)
	gos.Finalscore = tl.NewText(10, 8, fmt.Sprintf("Score: %d", score), tl.ColorWhite, tl.ColorBlack)

	gos.AddEntity(gos.Gameovertext)
	gos.AddEntity(gos.Finalscore)

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
		sidepanel, scoretxt = NewSidepanel()

		level.AddEntity(food)
		level.AddEntity(sidepanel)
		level.AddEntity(scoretxt)
		level.AddEntity(snake)
		level.AddEntity(arena)

		Snakegame.Screen().SetLevel(level)
	}
}

// func IsUTF8Supported() bool {
// 	return !strings.Contains(os.Getenv("LANG"), "C.UTF-8")
// }

func UpdateScore(amount int) {
	score += amount
	scoretxt.SetText(fmt.Sprintf("Score: %d", score))
}
