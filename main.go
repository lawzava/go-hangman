package main

import (
	"github.com/hitchnsmile/go-hangman/events"
	"github.com/hitchnsmile/go-hangman/screens"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	var h screens.Switch

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event := make(chan termbox.Event)
	go func() {
		for {
			event <- termbox.PollEvent()
		}
	}()

	h.X, h.Y = termbox.Size()
	h.GameID = h.DB.InitBoard()
	h.ShowMenu()
	events.EventHandler(event, &h)
}
