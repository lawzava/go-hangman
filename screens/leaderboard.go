package screens

import (
	_ "github.com/mattn/go-sqlite3"
)

type LeaderboardBoard struct {
	ID       int
	Place    string
	Word     string
	Guesses  string
	Points   string
	Finished string
}

func (s *Switch) Leaderboard() {
	s.CurrentState = GameLeaderboard
	s.ShowLeaderboard()
}

func (s *Switch) ShowLeaderboard() {
	//boardData := s.DB.GetBoard()
}
