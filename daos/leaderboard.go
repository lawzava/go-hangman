package daos

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Leaderboard struct {
	Database *sql.DB
}

type LeaderboardBoardData struct {
	ID       int
	Word     string
	Guesses  string
	Points   int
	Finished bool
}

const (
	createTable              string = "CREATE TABLE IF NOT EXISTS leaderboard (id INTEGER PRIMARY KEY, word TEXT, guesses []INTEGER, points INTEGER, finished BOOLEAN)"
	insertGame               string = "INSERT INTO leaderboard (word, guesses, points, finished) VALUES (?, ?, ?, ?)"
	updateGame               string = "UPDATE leaderboard SET guesses=?, points=?, finished=? WHERE id=?"
	getLeaderboard           string = "SELECT * FROM leaderboard ORDER BY points DESC"
	getLeaderboardFinished   string = "SELECT * FROM leaderboard WHERE finished = 1 ORDER BY points DESC"
	getLeaderboardUnfinished string = "SELECT * FROM leaderboard WHERE finished = 0 ORDER BY id DESC"
)

func (l *Leaderboard) InitBoard() {
	l.Database, _ = sql.Open("sqlite3", "./leaderboard.db")
	statement, _ := l.Database.Prepare(createTable)
	statement.Exec()
}

func (l *Leaderboard) GetBoard() []LeaderboardBoardData {
	var oneRow LeaderboardBoardData
	var fullBoard []LeaderboardBoardData
	rows, _ := l.Database.Query(getLeaderboard)
	for rows.Next() {
		rows.Scan(&oneRow.ID, &oneRow.Word, &oneRow.Guesses, &oneRow.Points, &oneRow.Finished)
		fullBoard = append(fullBoard, oneRow)
	}
	return fullBoard
}

func (l *Leaderboard) GetBoardFinished() []LeaderboardBoardData {
	var oneRow LeaderboardBoardData
	var fullBoard []LeaderboardBoardData
	rows, _ := l.Database.Query(getLeaderboardFinished)
	for rows.Next() {
		rows.Scan(&oneRow.ID, &oneRow.Word, &oneRow.Guesses, &oneRow.Points, &oneRow.Finished)
		fullBoard = append(fullBoard, oneRow)
	}
	return fullBoard
}

func (l *Leaderboard) GetBoardUnfinished() []LeaderboardBoardData {
	var oneRow LeaderboardBoardData
	var fullBoard []LeaderboardBoardData
	rows, _ := l.Database.Query(getLeaderboardUnfinished)
	for rows.Next() {
		rows.Scan(&oneRow.ID, &oneRow.Word, &oneRow.Guesses, &oneRow.Points, &oneRow.Finished)
		fullBoard = append(fullBoard, oneRow)
	}
	return fullBoard
}

func (l *Leaderboard) InsertGame(word string, guesses []rune) int {
	var newRow LeaderboardBoardData
	newRow.Word = word
	newRow.Guesses = string(guesses)
	newRow.Points = calculatePoints(word, guesses)
	newRow.Finished = isFinished(word, guesses)
	statement, _ := l.Database.Prepare(insertGame)
	result, err := statement.Exec(newRow.Word, newRow.Guesses, newRow.Points, newRow.Finished)
	if err != nil {
		fmt.Println(err)
	}
	lastID, _ := result.LastInsertId()
	return int(lastID)
}

func (l *Leaderboard) UpdateGame(word string, guesses []rune, id int) {
	var newRow LeaderboardBoardData
	newRow.Guesses = string(guesses)
	newRow.Points = calculatePoints(word, guesses)
	newRow.Finished = isFinished(word, guesses)
	statement, _ := l.Database.Prepare(updateGame)
	_, err := statement.Exec(newRow.Guesses, newRow.Points, newRow.Finished, id)
	if err != nil {
		fmt.Println(err)
	}
}

func calculatePoints(word string, guesses []rune) int {
	points := len(word) * 10

	for i := range guesses {
		penalty := 10
		for _, c := range word {
			if guesses[i] == c {
				penalty = 0
				break
			}
		}
		points = points - penalty
	}

	return points
}

func isFinished(word string, guesses []rune) bool {
	wordGuessed := true
	for _, c := range word {
		printing := '_'
		for i := range guesses {
			if guesses[i] == c {
				printing = c
				break
			}
		}
		if printing == '_' {
			wordGuessed = false
		}

	}
	return wordGuessed
}
