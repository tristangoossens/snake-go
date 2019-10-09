package main

import (
	tl "github.com/JoelOtter/termloop"
	s "github.com/tristangoossens/snake-go/game"
)

const (
	easy gamemode = iota
	medium
	hard
)

type gamemode int

var width int = 70
var height int = 25

func NewSidepanel() (*tl.Rectangle, *tl.Text) {
	sidepanel := tl.NewRectangle(width+1,0,30,height, tl.ColorWhite)
	titletext := tl.NewText(width +2, 1, "Score: 0", tl.ColorBlack, tl.ColorWhite)
	return sidepanel, titletext
}

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	snake := s.NewSnake()
	arena := s.NewArena(width, height)
	sidepanel, titletext := NewSidepanel()		

	level.AddEntity(sidepanel)
	level.AddEntity(titletext)
	level.AddEntity(snake)
	level.AddEntity(arena)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(10)
	game.Start()
}