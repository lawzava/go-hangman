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
				case e.Key == termbox.KeyF2:
					if h.CurrentState == screens.GameLeaderboard {
						h.LeaderboardState.Board = screens.ShowLeaderboard
						h.Leaderboard()
					}
				case e.Key == termbox.KeyF3:
					if h.CurrentState == screens.GameLeaderboard {
						h.LeaderboardState.Board = screens.ShowLeaderboardFinished
						h.Leaderboard()
					}
				case e.Key == termbox.KeyF4:
					if h.CurrentState == screens.GameLeaderboard {
						h.LeaderboardState.Board = screens.ShowLeaderboardUnfinished
						h.Leaderboard()
					}
				case e.Key == termbox.KeyBackspace2:
					h.ShowMenu()
				case e.Key == termbox.KeyArrowDown:
					if h.CurrentState == screens.GameMenu {
						h.MenuDown()
					} else if h.CurrentState == screens.GameLeaderboard {
						h.LeaderboardDown()
					}
				case e.Key == termbox.KeyArrowUp:
					if h.CurrentState == screens.GameMenu {
						h.MenuUp()
					} else if h.CurrentState == screens.GameLeaderboard {
						h.LeaderboardUp()
					}
				case e.Key == termbox.KeyEnter:
					if h.CurrentState == screens.GameMenu {
						h.MenuEnter(h.MenuState)
					} else if h.CurrentState == screens.GameLeaderboard {
						//	h.Leaderboard()
						// CONTINUE GAME WITH ID
					}
				case e.Ch > 0:
					if h.CurrentState == screens.GameStarted {
						if unicode.IsLetter(e.Ch) {
							h.AddGuess(e.Ch)
						}
					}
				}
			}
		default:
		}
	}
}
