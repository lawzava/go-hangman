package screens

import (
	termbox "github.com/nsf/termbox-go"
)

type state int

const (
	gameMenu state = iota
	gameLeaderboard
	gameStarted
)

const (
	menuNewGameText                = "1) Start New Game"
	menuContinueLastGameText       = "2) Continue Last Game"
	menuLeaderboardText            = "3) Leaderboard"
	menuNewGame              state = iota
	menuContinueLastGame
	menuLeaderboard
)

type Switch struct {
	CurrentState state
	MenuState    state
	GoalWord     string
	X            int
	Y            int
	Screen       [][]int
}

type menuItem struct {
}

func (s *Switch) ShowMenu() {
	s.CurrentState = gameMenu
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	switch s.MenuState {
	case 0:
		for i, c := range menuNewGameText {
			termbox.SetCell(i+3, 3, c, termbox.ColorBlack, termbox.ColorWhite)
		}
		for i, c := range menuContinueLastGameText {
			termbox.SetCell(i+3, 5, c, termbox.ColorWhite, termbox.ColorBlack)
		}
		for i, c := range menuLeaderboardText {
			termbox.SetCell(i+3, 7, c, termbox.ColorWhite, termbox.ColorBlack)
		}
	case 1:
		for i, c := range menuNewGameText {
			termbox.SetCell(i+3, 3, c, termbox.ColorWhite, termbox.ColorBlack)
		}
		for i, c := range menuContinueLastGameText {
			termbox.SetCell(i+3, 5, c, termbox.ColorBlack, termbox.ColorWhite)
		}
		for i, c := range menuLeaderboardText {
			termbox.SetCell(i+3, 7, c, termbox.ColorWhite, termbox.ColorBlack)
		}
	case 2:
		for i, c := range menuNewGameText {
			termbox.SetCell(i+3, 3, c, termbox.ColorWhite, termbox.ColorBlack)
		}
		for i, c := range menuContinueLastGameText {
			termbox.SetCell(i+3, 5, c, termbox.ColorWhite, termbox.ColorBlack)
		}
		for i, c := range menuLeaderboardText {
			termbox.SetCell(i+3, 7, c, termbox.ColorBlack, termbox.ColorWhite)
		}
	}

	termbox.Sync()

}

func (s *Switch) MenuUp() {
	if s.MenuState > 0 {
		s.MenuState = s.MenuState - 1
	}
	s.ShowMenu()
}

func (s *Switch) MenuDown() {
	if s.MenuState < 2 {
		s.MenuState = s.MenuState + 1
	}
	s.ShowMenu()
}
