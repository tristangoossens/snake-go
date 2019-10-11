package trisnake

import (
	tl "github.com/JoelOtter/termloop"
)

func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(5, 5, 1, 1)
	snake.Direction = right
	snake.Length = len(snake.Bodylength)
	snake.Bodylength = []Coordinates{
		{1, 6},
		{2, 6},
		{3, 6},
	}

	return snake
}

func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength)-1]
}

func (snake *Snake) BorderCollison() bool {
	return arena.Contains(*snake.Head())
}

func (snake *Snake) FoodCollision() bool {
	return food.Contains(*snake.Head())
}

func (snake *Snake) Draw(screen *tl.Screen) {
	nHead := *snake.Head()
	switch snake.Direction {
	case up:
		nHead.Y--
	case down:
		nHead.Y++
	case left:
		nHead.X--
	case right:
		nHead.X++
	}

	if snake.FoodCollision() {
		switch food.Emoji {
		case '🐁':
			UpdateScore(5)
		case 'R':
			UpdateScore(5)
		default:
			UpdateScore(1)
			snake.Bodylength = append(snake.Bodylength, nHead)
		}
		food.MoveFood()
	} else {
		snake.Bodylength = append(snake.Bodylength[1:], nHead)
	}

	snake.SetPosition(nHead.X, nHead.Y)

	if snake.BorderCollison() {
		Gameover()
	}

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: tl.ColorWhite,
			Ch: '▒',
		})
	}
}
