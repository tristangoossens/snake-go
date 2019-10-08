package trisnake

import tl "github.com/JoelOtter/termloop"

type Snake struct {
	*tl.Entity
	Direction  direction
	Length     int
	Bodylength []Coordinates
	Speed      int
}

type Coordinates struct {
	X int
	Y int
}

func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(5,5,1,1)
	snake.Direction = right
	snake.Bodylength = []Coordinates {
		{3, 5},
		{4, 5},
		{5, 5},
	}
	return snake
}

func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength) - 1]
}

func (snake *Snake) UpdateSnake(x, y int) {
	for i := 0; i < len(snake.Bodylength) - 1; i++ {
		snake.Bodylength[i] = snake.Bodylength[i + 1]
	}

	snake.SetPosition(x, y)
	snake.Head().X, snake.Head().Y = snake.Position()

}

func (snake *Snake) Draw(screen *tl.Screen) {
	x, y := snake.Position()
	switch snake.Direction {
	case up:
		snake.UpdateSnake(x, y - 1)
	case down:
		snake.UpdateSnake(x, y + 1)
	case left:
		snake.UpdateSnake(x - 1, y)
	case right:
		snake.UpdateSnake(x + 1, y)
	}
	

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: tl.ColorBlack,
			Ch: 'o',
		})
	}
}