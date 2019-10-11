package trisnake

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

func NewFood() *Food {
	rand.Seed(time.Now().UnixNano())

	food := new(Food)
	food.Foodposition.X = rand.Intn(arenawidth - 5)
	food.Foodposition.Y = rand.Intn(arenaheight - 2)
	food.Entity = tl.NewEntity(food.Foodposition.X, food.Foodposition.Y, 1, 1)
	if IsUTF8Supported() {
		food.Emoji = RandomFoodUTF8()
	} else {
		food.Emoji = RandomFood()
	}
	return food
}

func (food *Food) MoveFood() {
	rand.Seed(time.Now().UnixNano())
	newx := rand.Intn(arenawidth - 5)
	newy := rand.Intn(arenaheight - 2)

	food.Foodposition.X = newx
	food.Foodposition.Y = newy

	if IsUTF8Supported() {
		food.Emoji = RandomFoodUTF8()
	} else {
		food.Emoji = RandomFood()
	}

	food.SetPosition(food.Foodposition.X, food.Foodposition.Y)
}

func RandomFood() rune {
	emoji := []rune{
		'R', // Favourite dish, extra points!!! 😋
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'S', // You do not want to eat the skull 💀
	}

	rand.Seed(time.Now().UnixNano())

	return emoji[rand.Intn(len(emoji))]
}

func RandomFoodUTF8() rune {
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
		'💀', // You do not want to eat the skull 💀
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
