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

func StartGame() {
	
}

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
	})

	snake := s.NewSnake()
	level.AddEntity(snake)

	level.AddEntity(snake)
	game.Screen().SetLevel(level)
	game.Screen().SetFps(10)
	game.Start()
}