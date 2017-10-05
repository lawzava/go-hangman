package screens

type GameState int

const (
	GameMenu GameState = iota
	GameStarted
	GameFinished
	GameLeaderboard
)

type Switch struct {
	CurrentState GameState
	MenuState    MenuState
	GoalWord     string
	Guesses      []rune
	X            int
	Y            int
	Screen       [][]int
}
