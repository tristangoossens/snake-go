package trisnake

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

func NewFood() *Food {
	rand.Seed(time.Now().UnixNano())

	food := new(Food)
	food.Foodposition.X = rand.Intn(arenawidth - 1)
	food.Foodposition.Y = rand.Intn(arenaheight - 1)
	food.Entity = tl.NewEntity(food.Foodposition.X, food.Foodposition.Y, 1, 1)
	food.Emoji = food.RandomFood()

	return food
}

func (food *Food) RandomFood() rune {
	emoji := []rune{
		'🐁', // Favourite dish, extra points!!! 😋
		'🍔',
		'🍆',
		'🍑',
		'🍏',
		'🍊',
		'🍝',
		'🍌',
		'🍞',
		'🍟',
		'🍎',
		'💀', // You do not want to eat the skull
	}

	rand.Seed(time.Now().UnixNano())

	return emoji[rand.Intn(len(emoji))]
}

func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Foodposition.X, food.Foodposition.Y, &tl.Cell{
		Ch: food.Emoji,
	})
}

func (food *Food) Contains(c Coordinates) bool {
	return c.X == food.Foodposition.X && c.Y == food.Foodposition.Y
}
