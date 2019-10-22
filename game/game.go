package trisnake

import (
	"fmt"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
	// s "github.com/tristangoossens/snake-go/game"
)

// StartGame will start the game with the tilescreen.
func StartGame() {
	// Checks for UTF-8 support.
	utf8support = !strings.Contains(os.Getenv("LANG"), "C.UTF-8")
	// Initializing a new game.
	Snakegame = tl.NewGame()

	// Calls the titlescreen function wich will return the screen.
	ts := NewTitleScreen()
	// Add the titlescreentext to the screen.
	ts.AddEntity(ts.TitleText)
	// Sets the level to titlescren.
	Snakegame.Screen().SetLevel(ts)
	// Set FPS to 12.
	Snakegame.Screen().SetFps(fps)
	Snakegame.Start()
}

// NewTitleScreen will create a new titlescreen and return it
func NewTitleScreen() *Titlescreen {
	ts := new(Titlescreen)
	// Create a new base level
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: ParseUserSettingsColor(backgroundcolor),
	})
	// Create a new title text
	ts.TitleText = tl.NewText(10, 5, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack)
	return ts
}

// NewSidepanel will create a new sidepanel given the arena height and width.
func NewSidepanel() (*tl.Rectangle, *tl.Text, *tl.Text) {
	// Creates a new rectangle for the scoretext and instructions
	sidepanel := tl.NewRectangle(arenawidth+1, 0, 30, arenaheight, ParseUserSettingsColor(sidepanelcolor))
	scoretxt = tl.NewText(arenawidth+2, 1, fmt.Sprintf("Score: %d", score), tl.ColorBlack, tl.ColorWhite)
	fpstext = tl.NewText(arenawidth+2, 3, fmt.Sprintf("FPS: %.0f", fps), tl.ColorBlack, tl.ColorWhite)
	return sidepanel, scoretxt, fpstext
}

// Gameover is initialized when the snake has died.
func Gameover() {
	gos := new(Gameoverscreen)
	// Creates a new baselevel for the gameoverscreen.
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: ParseUserSettingsColor(backgroundcolor),
	})
	// Creates a new gameover text.
	gos.Gameovertext = tl.NewText(10, 5, "Gameover!", tl.ColorWhite, tl.ColorBlack)
	// Creates a score text for your final score.
	gos.Finalscore = tl.NewText(10, 7, fmt.Sprintf("Score: %d", score), tl.ColorWhite, tl.ColorBlack)
	gos.Gameoveroptions = tl.NewRectangle(30, 4, 30, 5, tl.ColorWhite)

	restart := tl.NewText(32, 5, "Press \"Home\" to restart!", tl.ColorBlack, tl.ColorWhite)
	quit := tl.NewText(32, 7, "Press \"Delete\" to restart!", tl.ColorBlack, tl.ColorWhite)

	// Adds the text and rectangle to the gameover screen level.
	gos.AddEntity(gos.Gameovertext)
	gos.AddEntity(gos.Finalscore)
	gos.AddEntity(gos.Gameoveroptions)
	gos.AddEntity(restart)
	gos.AddEntity(quit)

	Snakegame.Screen().SetLevel(gos)
}

// Tick will listen for a keypress to initiate the game.
func (ts *Titlescreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event and the key pressed is the enter key.
	if event.Type == tl.EventKey && event.Key == tl.KeyEnter {
		// Creates a new baselevel ffor the snake game.
		level = tl.NewBaseLevel(tl.Cell{
			Bg: ParseUserSettingsColor(backgroundcolor),
		})

		// Calls the funtion to create a new snake.
		snake = NewSnake()
		// Call the function to create a new arena, given the arena width and height.
		arena = NewArena(arenawidth, arenaheight)
		// Calls the food function to create a new piece of food.
		food = NewFood()
		// Calls the function to create a new sidepanel
		sidepanel, scoretxt, fpstext = NewSidepanel()

		// Adds all of the entities needed to start the game
		level.AddEntity(food)
		level.AddEntity(sidepanel)
		level.AddEntity(scoretxt)
		level.AddEntity(fpstext)
		level.AddEntity(snake)
		level.AddEntity(arena)

		// Sets the gamelevel to the snakegame! 🐍
		Snakegame.Screen().SetLevel(level)
	}
}

// UpdateScore updates the with the given amount of points 🐁
func UpdateScore(amount int) {
	score += amount
	scoretxt.SetText(fmt.Sprintf("Score: %d", score))
}

// UpdateFPS updates the fps text.
func UpdateFPS() {
	Snakegame.Screen().SetFps(fps)
	fpstext.SetText(fmt.Sprintf("Speed: %.0f", fps))
}

// ParseUserSettingsColor will parse the users input in gamefiles.go for all of the colors
func ParseUserSettingsColor(color string) tl.Attr {
	switch color {
	case "Black":
		return tl.ColorBlack
	case "White":
		return tl.ColorWhite
	case "Red":
		return tl.ColorRed
	case "Yellow":
		return tl.ColorYellow
	case "Green":
		return tl.ColorGreen
	case "Blue":
		return tl.ColorBlue
	case "Cyan":
		return tl.ColorCyan
	case "Magenta":
		return tl.ColorMagenta
	default:
		return tl.ColorDefault
	}
}

// Tick is a method for the gameoverscreen wich listens for either a restart or a quit input from the user.
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
		}
	}
}

// RestartGame will restart the game and reset the position of the food and the snake to prevent collision issues.
func RestartGame() {
	// Removes the current snake and food from the level.
	level.RemoveEntity(snake)
	level.RemoveEntity(food)

	// Generate a new snake and food.
	snake = NewSnake()
	food = NewFood()

	// Revert the score and fps to the standard.
	fps = 12
	score = 0

	// Update the score and fps text.
	scoretxt.SetText(fmt.Sprintf("Score: %d", score))
	fpstext.SetText(fmt.Sprintf("Speed: %.0f", fps))

	// Adds the snake and food back and sets them to the standard position.
	level.AddEntity(snake)
	level.AddEntity(food)
	Snakegame.Screen().SetLevel(level)
}
