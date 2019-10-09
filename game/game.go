package trisnake

import (
	tl "github.com/JoelOtter/termloop"
	s "github.com/tristangoossens/snake-go/game"
)

var snakegame *tl.Game
var arenawidth int = 70
var arenaheight int = 25
var score int

func NewSidepanel() (*tl.Rectangle, *tl.Text) {
	sidepanel := tl.NewRectangle(arenawidth+1, 0, 30, arenaheight, tl.ColorWhite)
	scoretxt := tl.NewText(arenawidth+2, 1, "Score: 0", tl.ColorBlack, tl.ColorWhite)
	return sidepanel, scoretxt
}

func (game *Game) Mainmenu() {
	startmenu := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
	})

	starttxt := tl.NewText(45, 10, "Snake - by tristangoossens", tl.ColorBlack, tl.ColorGreen)

	startmenu.AddEntity(starttxt)
	snakegame.Screen().SetLevel(startmenu)
}

func (game *Game) StartGame() {

	snakegame = tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	snake := s.NewSnake()
	arena := s.NewArena(arenawidth, arenaheight)
	sidepanel, titletext := NewSidepanel()

	level.AddEntity(sidepanel)
	level.AddEntity(titletext)
	level.AddEntity(snake)
	level.AddEntity(arena)

	snakegame.Screen().SetLevel(level)
	snakegame.Screen().SetFps(10)
	snakegame.Start()
}