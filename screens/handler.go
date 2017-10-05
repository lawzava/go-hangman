package screens

import "github.com/hitchnsmile/go-hangman/daos"

type GameState int

const (
	GameMenu GameState = iota
	GameStarted
	GameFinished
	GameLeaderboard
)

type Switch struct {
	CurrentState     GameState
	MenuState        MenuState
	LeaderboardState struct {
		Board      LeaderboardStates
		Selection  int
		SelectedID int
	}
	GameID   int
	GoalWord string
	Guesses  []rune
	X        int
	Y        int
	Screen   [][]int
	DB       daos.Leaderboard
}
