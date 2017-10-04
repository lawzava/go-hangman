package events

import (
	"fmt"

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
					h.ShowMenu()
				case e.Key == termbox.KeyArrowDown:
					h.MenuDown()
				case e.Key == termbox.KeyArrowUp:
					h.MenuUp()
				case e.Ch > 0:
					fmt.Print(string(e.Ch))
				}
			}
		default:
		}
	}
}
