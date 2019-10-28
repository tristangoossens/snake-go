package trisnake

import tl "github.com/JoelOtter/termloop"

// NewArena will create a new arena with the arena with and arena height given when this function was called in game.go
// This function will create a arena using the arena struct that can be found in the types.go file.
func NewArena(w, h int) *Arena {
	arena := new(Arena)
	// Width and height of the arena are decresed by one to add corners on the arena border.
	arena.Width = w - 1
	arena.Height = h - 1
	// Each arena cell will have a width and heigth of 1.
	arena.Entity = tl.NewEntity(1, 1, 1, 1)
	// Creates a map of coordinates.
	arena.ArenaBorder = make(map[Coordinates]int)

	// This for loop will create the top and bottom borders
	for x := 0; x < arena.Width; x++ {
		arena.ArenaBorder[Coordinates{x, 0}] = 1
		arena.ArenaBorder[Coordinates{x, arena.Height}] = 1
	}

	// This for loop will create the left and right borders
	for y := 0; y < arena.Height+1; y++ {
		arena.ArenaBorder[Coordinates{0, y}] = 1
		arena.ArenaBorder[Coordinates{arena.Width, y}] = 1
	}
	return arena
}

// Contains checks if the arenaborder map contains the coordinates of the snake, if so this will return true.
func (arena *Arena) Contains(c Coordinates) bool {
	_, exists := arena.ArenaBorder[c]
	return exists
}

// Draw is a termloop function that will draw out the arena border, this is called when the game has started.
func (arena *Arena) Draw(screen *tl.Screen) {
	// This for loop will range ArenaBorder containing the coordinates of the arenaborder and will print them out on the screen.
	for i := range arena.ArenaBorder {
		screen.RenderCell(i.X, i.Y, &tl.Cell{
			Bg: CheckSelectedColor(counterArena),
		})
	}
}
