package events

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

func EventHandler(event chan termbox.Event) {
	for {
		select {
		case e := <-event:
			if e.Type == termbox.EventKey {
				switch {
				case e.Key == termbox.KeyEsc || e.Key == termbox.KeyCtrlC:
					return
				case e.Ch > 0:
					fmt.Print(string(e.Ch))
				}
			}
		default:
		}
	}
}
