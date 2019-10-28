package trisnake

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)

var counterSnake = 10
var counterArena = 10

// Tick listens for a keypress and then returns a direction for the snake.
func (snake *Snake) Tick(event tl.Event) {
	// Checks if the event is a keyevent.
	if event.Type == tl.EventKey {
		switch event.Key {
		// Checks if the key is a → press.
		case tl.KeyArrowRight:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != left {
				// Changes the direction of the snake to right.
				snake.Direction = right
			}
		// Checks if the key is a ← press.
		case tl.KeyArrowLeft:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != right {
				// Changes the direction of the snake to left.
				snake.Direction = left
			}
		// Checks if the key is a ↑ press.
		case tl.KeyArrowUp:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != down {
				// Changes the direction of the snake to down.
				snake.Direction = up
			}
		// Checks if the key is a ↓ press.
		case tl.KeyArrowDown:
			// Check if the direction is not opposite to the current direction.
			if snake.Direction != up {
				// Changes the direction of the snake to up.
				snake.Direction = down
			}
		}
	}
}

// Tick is a method for the gameoverscreen which listens for either a restart or a quit input from the user.
func (gos *Gameoverscreen) Tick(event tl.Event) {
	// Check if the event is a key event.
	if event.Type == tl.EventKey {
		switch event.Key {
		// If the key pressed is backspace the game will restart!!
		case tl.KeyHome:
			// Will call the RestartGame function to restart the game.
			RestartGame()
		case tl.KeyDelete:
			// Will end the game using a fatal log. This uses the termbox package as termloop does not have a function like that.
			tb.Close()
		case tl.KeySpace:
			SaveHighScore(gs.Score, gs.FPS, Difficulty)
		}
	}
}

// Tick will listen for a keypress to initiate the game.
func (ts *Titlescreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event and the key pressed is the enter key.
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
		if event.Key == tl.KeyInsert {
			gop := NewOptionsscreen()
			sg.Screen().SetLevel(gop)
		}
	}
}

// Tick will listen for a keypress to initiate the game.
func (g *Gameoptionsscreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event.
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyF1:
			ts.GameDifficulty = easy
			Difficulty = "Easy"
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyF2:
			ts.GameDifficulty = normal
			Difficulty = "Normal"
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyArrowUp:
			switch ColorObject {
			case "Snake":
				if counterSnake <= 10 {
					return
				}
				counterSnake -= 2
				gop.ColorSelectedIcon.SetPosition(73, counterSnake)
			case "Arena":
				if counterArena <= 10 {
					return
				}
				counterArena -= 2
				gop.ColorSelectedIcon.SetPosition(73, counterArena)

			}
		case tl.KeyArrowDown:
			switch ColorObject {
			case "Snake":
				if counterSnake >= 22 {
					return
				}
				counterSnake += 2
				gop.ColorSelectedIcon.SetPosition(73, counterSnake)

			case "Arena":
				if counterArena >= 22 {
					return
				}
				counterArena += 2
				gop.ColorSelectedIcon.SetPosition(73, counterArena)

			}

		case tl.KeyF3:
			ts.GameDifficulty = hard
			Difficulty = "Hard"
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyF4:
			ColorObject = "Snake"
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyF5:
			ColorObject = "Food"
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyF6:
			ColorObject = "Arena"
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyEnter:
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
	}
}

func CheckSelectedColor(c int) tl.Attr {
	switch c {
	case 10:
		return tl.ColorWhite
	case 12:
		return tl.ColorRed
	case 14:
		return tl.ColorGreen
	case 16:
		return tl.ColorBlue
	case 18:
		return tl.ColorYellow
	case 20:
		return tl.ColorMagenta
	case 22:
		return tl.ColorCyan
	default:
		return tl.ColorDefault
	}
}
