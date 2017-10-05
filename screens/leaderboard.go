package screens

import (
	"fmt"

	"github.com/hitchnsmile/go-hangman/daos"
	_ "github.com/mattn/go-sqlite3"
	termbox "github.com/nsf/termbox-go"
)

type LeaderboardStates int

const (
	ShowLeaderboard LeaderboardStates = iota
	ShowLeaderboardFinished
	ShowLeaderboardUnfinished
)

const (
	leaderboardTitle        = "LEADERBOARD"
	leaderboardSubtitle     = "Choose game you wish to continue"
	leaderboardInstructions = "F2: Show all games | F3: Show completed games | F4: Show games in progress"
)

var leaderboardTableHeader = []string{"PLACE", "POINTS", "WORD", "GUESSSES", "STATUS"}

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
	switch s.LeaderboardState.Board {
	case ShowLeaderboard:
		s.ShowLeaderboard()
	case ShowLeaderboardFinished:
		s.ShowLeaderboardFinished()
	case ShowLeaderboardUnfinished:
		s.ShowLeaderboardUnfinished()
	}
}

func (s *Switch) ShowLeaderboard() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	boardRaw := s.DB.GetBoard()
	boardData := convertToPreparedResponse(boardRaw)
	drawBoard(boardData)
	termbox.Sync()
}

func (s *Switch) ShowLeaderboardFinished() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	boardRaw := s.DB.GetBoardFinished()
	boardData := convertToPreparedResponse(boardRaw)
	drawBoard(boardData)
	termbox.Sync()
}

func (s *Switch) ShowLeaderboardUnfinished() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	boardRaw := s.DB.GetBoardUnfinished()
	boardData := convertToPreparedResponse(boardRaw)
	drawBoard(boardData)
	termbox.Sync()
}

func drawBoard(data []LeaderboardBoard) {
	printSimpleText(leaderboardTitle, 2, 5)
	printSimpleText(leaderboardSubtitle, 3, 5)
	printSimpleText(leaderboardInstructions, 4, 5)
	drawTable(6, 5, data)
}

func drawTable(y, x int, rows []LeaderboardBoard) {
	newPosition := x

	for i := 0; i < len(leaderboardTableHeader); i++ {
		maxPosition := newPosition

		if lastPosition := printTableElement(newPosition, y, leaderboardTableHeader[i]); lastPosition > maxPosition {
			maxPosition = lastPosition
		}

		switch leaderboardTableHeader[i] {
		case leaderboardTableHeader[0]:
			for j := range rows {
				if lastPosition := printTableElement(newPosition, y+j+1, rows[j].Place); lastPosition > maxPosition {
					maxPosition = lastPosition
				}
			}
		case leaderboardTableHeader[1]:
			for j := range rows {
				if lastPosition := printTableElement(newPosition, y+j+1, rows[j].Points); lastPosition > maxPosition {
					maxPosition = lastPosition
				}
			}
		case leaderboardTableHeader[2]:
			for j := range rows {
				if lastPosition := printTableElement(newPosition, y+j+1, rows[j].Word); lastPosition > maxPosition {
					maxPosition = lastPosition
				}
			}
		case leaderboardTableHeader[3]:
			for j := range rows {
				if lastPosition := printTableElement(newPosition, y+j+1, rows[j].Guesses); lastPosition > maxPosition {
					maxPosition = lastPosition
				}
			}
		case leaderboardTableHeader[4]:
			for j := range rows {
				if lastPosition := printTableElement(newPosition, y+j+1, rows[j].Finished); lastPosition > maxPosition {
					maxPosition = lastPosition
				}
			}
		}

		newPosition = maxPosition + 1
	}
}

func printTableElement(x, y int, text string) int {
	for _, c := range text {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
		x++
	}
	return x
}

func convertToPreparedResponse(raw []daos.LeaderboardBoardData) []LeaderboardBoard {
	var response []LeaderboardBoard
	for i := range raw {
		var tempResp LeaderboardBoard
		tempResp.ID = raw[i].ID
		tempResp.Place = fmt.Sprintf("%d. ", i+1)
		tempResp.Word = convertCurrentWord(raw[i].Word, raw[i].Guesses)
		tempResp.Points = fmt.Sprint(raw[i].Points)
		tempResp.Guesses = raw[i].Guesses
		tempResp.Finished = finishedStatus(raw[i].Finished)
		response = append(response, tempResp)
	}
	return response
}

func finishedStatus(finished bool) string {
	if finished {
		return "COMPLETED"
	} else {
		return "IN PROGRESS"
	}
}

func convertCurrentWord(word string, guesses string) string {
	newWord := ""
	for _, c := range word {
		needToChange := true
		for _, g := range guesses {
			if c == g {
				needToChange = false
			}
		}
		if needToChange {
			newWord = newWord + "_"
		} else {
			newWord = newWord + string(c)
		}
	}

	return newWord
}
