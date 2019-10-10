package main

import (
	tl "github.com/JoelOtter/termloop"
	s "github.com/tristangoossens/snake-go/game"
)

var snakegame *tl.Game
var arenawidth int = 70
var arenaheight int = 25
var score int

func main() {
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

func NewSidepanel() (*tl.Rectangle, *tl.Text) {
	sidepanel := tl.NewRectangle(arenawidth+1, 0, 30, arenaheight, tl.ColorWhite)
	scoretxt := tl.NewText(arenawidth+2, 1, "Score: 0", tl.ColorBlack, tl.ColorWhite)
	return sidepanel, scoretxt
}
