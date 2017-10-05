package events

import (
	"unicode"

	"github.com/hitchnsmile/go-hangman/screens"
	termbox "github.com/nsf/termbox-go"
)

func EventHandler(event chan termbox.Event, h *screens.Switch) {
	for {
		select {
		case e := <-event:
			if e.Type == termbox.EventKey {
				switch {
				case e.Key == termbox.KeyEsc || e.Key == termbox.KeyCtrlC:
					return
				case e.Key == termbox.KeyBackspace:
					if h.CurrentState != screens.GameMenu {
						h.ShowMenu()
					}
				case e.Key == termbox.KeyArrowDown:
					if h.CurrentState == screens.GameMenu {
						h.MenuDown()
					}
				case e.Key == termbox.KeyArrowUp:
					if h.CurrentState == screens.GameMenu {
						h.MenuUp()
					}
				case e.Key == termbox.KeyEnter:
					if h.CurrentState == screens.GameMenu {
						h.MenuEnter(h.MenuState)
					}
				case e.Ch > 0:
					if unicode.IsLetter(e.Ch) {
						h.AddGuess(e.Ch)
					}
				}
			}
		default:
		}
	}
}
