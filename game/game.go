package trisnake

import (
	"fmt"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

// StartGame will start the game with the tilescreen.
func StartGame() {
	// Checks for UTF-8 support.
	utf8support = !strings.Contains(os.Getenv("LANG"), "C.UTF-8")
	// Initializing a new game.
	sg = tl.NewGame()

	// Calls the titlescreen function wich will return the titlescreen.
	ts := NewTitleScreen()
	// Add the titlescreentext to the screen.
	ts.AddEntity(ts.TitleText)
	ts.AddEntity(ts.OptionsText)
	// Sets the level to titlescren.
	sg.Screen().SetLevel(ts)
	// Set FPS to 10 for just the menu.
	sg.Screen().SetFps(10)
	sg.Start()
}

// NewTitleScreen will create a new titlescreen and return it
func NewTitleScreen() *Titlescreen {
	ts = new(Titlescreen)
	// Create a new base level
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: ParseUserSettingsColor(backgroundcolor),
	})
	ts.GameDifficulty = hard
	// Create a new title text
	ts.TitleText = tl.NewText(10, 5, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack)
	ts.OptionsText = tl.NewText(10, 7, "Press INSERT for options!", tl.ColorWhite, tl.ColorBlack)
	return ts
}

func NewOptionsscreen() *Gameoptionsscreen {
	gop = new(Gameoptionsscreen)
	gop.Level = tl.NewBaseLevel(tl.Cell{
		Bg: ParseUserSettingsColor(backgroundcolor),
	})
	gop.DifficultyBackground = tl.NewRectangle(5, 2, 33, 10, tl.ColorWhite)
	gop.CurrentDifficultyText = tl.NewText(6, 3, fmt.Sprintf("Current difficulty: Normal"), tl.ColorBlack, tl.ColorWhite)
	gop.DifficultyEasy = tl.NewText(6, 6, fmt.Sprintf("Press F1 for Easy (8 speed)"), tl.ColorBlack, tl.ColorWhite)
	gop.DifficultyNormal = tl.NewText(6, 8, fmt.Sprintf("Press F2 for Normal (12 speed)"), tl.ColorBlack, tl.ColorWhite)
	gop.DifficultyHard = tl.NewText(6, 10, fmt.Sprintf("Press F3 for Hard (20 speed)"), tl.ColorBlack, tl.ColorWhite)

	gop.AddEntity(gop.DifficultyBackground)
	gop.AddEntity(gop.CurrentDifficultyText)
	gop.AddEntity(gop.DifficultyEasy)
	gop.AddEntity(gop.DifficultyNormal)
	gop.AddEntity(gop.DifficultyHard)

	return gop
}

func NewGamescreen() *Gamescreen {
	gs = new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: ParseUserSettingsColor(backgroundcolor),
	})
	// Check difficulty and set standard FPS.
	CheckDiffiultyFPS()
	// Starting score is 0.
	gs.Score = 0
	// Calls the funtion to create a new snake.
	gs.SnakeEntity = NewSnake()
	// Call the function to create a new arena, given the arena width and height.
	gs.ArenaEntity = NewArena(arenawidth, arenaheight)
	// Calls the food function to create a new piece of food.
	gs.FoodEntity = NewFood()
	// Calls the function to create a new sidepanel.
	gs.SidepanelObject = NewSidepanel()

	// Adds all of the entities needed to start the game
	gs.AddEntity(gs.FoodEntity)
	gs.AddEntity(gs.SidepanelObject.Background)
	gs.AddEntity(gs.SidepanelObject.ScoreText)
	gs.AddEntity(gs.SidepanelObject.SpeedText)
	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.ArenaEntity)

	sg.Screen().SetFps(gs.FPS)

	return gs
}

// NewSidepanel will create a new sidepanel given the arena height and width.
func NewSidepanel() *Sidepanel {
	// Creates a new rectangle for the scoretext, speedtext and instructions.
	sp = new(Sidepanel)
	sp.Background = tl.NewRectangle(arenawidth+1, 0, 30, arenaheight, ParseUserSettingsColor(sidepanelcolor))
	sp.ScoreText = tl.NewText(arenawidth+2, 1, fmt.Sprintf("Score: %d", gs.Score), tl.ColorBlack, tl.ColorWhite)
	sp.SpeedText = tl.NewText(arenawidth+2, 3, fmt.Sprintf("Speed: %.0f", gs.FPS), tl.ColorBlack, tl.ColorWhite)
	return sp
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
	gos.Finalscore = tl.NewText(10, 7, fmt.Sprintf("Score: %d", gs.Score), tl.ColorWhite, tl.ColorBlack)
	gos.Gameoveroptions = tl.NewRectangle(30, 4, 30, 5, tl.ColorWhite)

	restart := tl.NewText(32, 5, "Press \"Home\" to restart!", tl.ColorBlack, tl.ColorWhite)
	quit := tl.NewText(32, 7, "Press \"Delete\" to restart!", tl.ColorBlack, tl.ColorWhite)

	// Adds the text and rectangle to the gameover screen level.
	gos.AddEntity(gos.Gameovertext)
	gos.AddEntity(gos.Finalscore)
	gos.AddEntity(gos.Gameoveroptions)
	gos.AddEntity(restart)
	gos.AddEntity(quit)

	sg.Screen().SetLevel(gos)
}

// UpdateScore updates the with the given amount of points 🐁
func UpdateScore(amount int) {
	gs.Score += amount
	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
}

// UpdateFPS updates the fps text.
func UpdateFPS() {
	sg.Screen().SetFps(gs.FPS)
	sp.SpeedText.SetText(fmt.Sprintf("Speed: %.0f", gs.FPS))
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

// RestartGame will restart the game and reset the position of the food and the snake to prevent collision issues.
func RestartGame() {
	// Removes the current snake and food from the level.
	gs.RemoveEntity(gs.SnakeEntity)
	gs.RemoveEntity(gs.FoodEntity)

	// Generate a new snake and food.
	gs.SnakeEntity = NewSnake()
	gs.FoodEntity = NewFood()

	// Revert the score and fps to the standard.
	CheckDiffiultyFPS()
	gs.Score = 0

	// Update the score and fps text.
	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
	sp.SpeedText.SetText(fmt.Sprintf("Speed: %.0f", gs.FPS))

	// Adds the snake and food back and sets them to the standard position.
	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.FoodEntity)
	sg.Screen().SetFps(gs.FPS)
	sg.Screen().SetLevel(gs)
}

func CheckDiffiultyFPS() {
	switch ts.GameDifficulty {
	case easy:
		gs.FPS = 8
	case normal:
		gs.FPS = 12
	case hard:
		gs.FPS = 25
	}
}
