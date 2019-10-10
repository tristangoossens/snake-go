package trisnake

import tl "github.com/JoelOtter/termloop"

func NewArena(w, h int) *Arena {
	arena := new(Arena)
	arena.Width = w - 1
	arena.Height = h - 1
	arena.Entity = tl.NewEntity(1, 1, 1, 1)

	arena.ArenaBorder = make(map[Coordinates]int)

	for x := 0; x < arena.Width; x++ {
		arena.ArenaBorder[Coordinates{x, 0}] = 1
		arena.ArenaBorder[Coordinates{x, arena.Height}] = 1
	}

	for y := 0; y < arena.Height+1; y++ {
		arena.ArenaBorder[Coordinates{0, y}] = 1
		arena.ArenaBorder[Coordinates{arena.Width, y}] = 1
	}
	return arena
}

func (arena *Arena) Contains(c Coordinates) bool {
	_, exists := arena.ArenaBorder[c]
	return exists
}

func (arena *Arena) Draw(screen *tl.Screen) {
	for i := range arena.ArenaBorder {
		if arena == nil {
			return
		}

		screen.RenderCell(i.X, i.Y, &tl.Cell{
			Bg: tl.ColorWhite,
		})
	}
}
