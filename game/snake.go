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

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: 'o',
		})
	}
}