package trisnake

import tl "github.com/JoelOtter/termloop"

func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.Direction != left {
				snake.Direction = right
			}
		case tl.KeyArrowLeft:
			if snake.Direction != right {
				snake.Direction = left
			}
		case tl.KeyArrowUp:
			if snake.Direction != down {
				snake.Direction = up
			}
		case tl.KeyArrowDown:
			if snake.Direction != up {
				snake.Direction = down
			}
		}
	}
}

