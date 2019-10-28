package trisnake

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Variable insideborderW and insideborderH are variables consisting of the arenawidth and height and subtract both with 1
// in order to account for the arena border.
var insideborderW = 70 - 1
var insideborderH = 25 - 1

// NewFood will create a new piece of food, this will only happen once when the game has started.
func NewFood() *Food {
	food := new(Food)
	// Create a new entity food with a standard position and 1x1 size
	food.Entity = tl.NewEntity(1, 1, 1, 1)
	// Call function MoveFood to move the food to a random position.
	food.MoveFood()

	return food
}

// MoveFood moves the food into a new random position.
func (food *Food) MoveFood() {

	// Calls the RandomInsideArena function to make sure that the foods spawns inside the arena.
	NewX := RandomInsideArena(insideborderW, 1)
	NewY := RandomInsideArena(insideborderH, 1)

	// Changes the X and Y coordinates of the food.
	food.Foodposition.X = NewX
	food.Foodposition.Y = NewY
	food.Emoji = RandomFood()

	// Set the new position of the food.
	food.SetPosition(food.Foodposition.X, food.Foodposition.Y)
}

// RandomFood will use the ASCII-charset to pick a random rune from the slice and print it out as food.
func RandomFood() rune {
	// This slice contains all of the possible food icons.
	emoji := []rune{
		'R', // Favourite dish, extra points!!!
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'S', // You do not want to eat the skull
	}

	rand.Seed(time.Now().UnixNano())

	// Return a random rune picked from the slice
	return emoji[rand.Intn(len(emoji))]
}

// Draw will print out the food on the screen.
func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Foodposition.X, food.Foodposition.Y, &tl.Cell{
		Ch: food.Emoji,
	})
}

// Contains checks if food contains the coordinates, if so this will return a bool.
func (food *Food) Contains(c Coordinates) bool {
	return c.X == food.Foodposition.X && c.Y == food.Foodposition.Y
}

// RandomInsideArena will the minimal, which is just inside the border and the maximal, being the arena width or height.
func RandomInsideArena(iMax int, iMin int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(iMax-iMin) + iMin
}
