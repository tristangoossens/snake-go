package trisnake

import tl "github.com/JoelOtter/termloop"

func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(5, 5, 1, 1)
	snake.Direction = right
	snake.Bodylength = []Coordinates{
		{3, 5},
		{4, 5},
		{5, 5},
	}

	return snake
}

func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength)-1]
}

func (snake *Snake) isCollidingWithBorder() bool {
	return arena.Contains(*snake.Head())
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

	snake.Bodylength = append(snake.Bodylength[1:], nHead)

	snake.SetPosition(nHead.X, nHead.Y)

	if snake.isCollidingWithBorder() {
		Gameover()
	}

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: tl.ColorWhite,
			Ch: 'o',
		})
	}
}

func (snake *Snake) Collide(c tl.Physical) {
	if _, ok := c.(*Arena); ok {
		Gameover()
	}
}
