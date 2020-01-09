// Package tictactoe implements the game of Tic Tac Toe
package tictactoe

// create a Game game
// 1. two instance method
// - method to print board
// - method to insert token
// 2. implement some A.I to play against the player
// -

import (
	"fmt"
)

func main() {
	var game Game = initGame()
	PrintGame(game)
	insertToken("X", &game.board, 2, 2)
	PrintGame(game)
}

// Game game
type Game struct {
	board [3][3]string
}

// Coord represents the coordinates (X, Y) on the board
type Coord struct {
	X int
	Y int
}

func initGame() Game {
	var game = Game{}
	game.board = [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	return game
}

// Token returns the string token of a cell on the board
func Token(board [3][3]string, cell Coord) string {
	if cell.X < 3 && cell.Y < 3 {
		return board[cell.Y][cell.X]
	}
	return ""
}

func insertToken(token string, board *[3][3]string, x int, y int) {
	if x < 3 && y < 3 {
		// check if the position has a token
		if board[y][x] == "-" {
			board[y][x] = token
		}
	}
}

// PrintGame prints out the game board
func PrintGame(game Game) {
	fmt.Println("")
	for i := 0; i < 3; i++ {
		line := game.board[i]
		fmt.Println(fmt.Sprintf("%s | %s | %s", line[0], line[1], line[2]))
	}
}

// CheckWinner checks if the current board has a winner "X" or "O"
// returns "X" or "O" or "" for no winner
func CheckWinner(game Game) string {
	// Rules for winning

	// diagonals
	// [0,0],[1,1],[2,2]
	// [2,0],[1,1],[0,2]

	// horizontals
	// [0,0],[1,0],[2,0]
	// [0,1],[1,1],[2,1]
	// [0,2],[1,2],[2,2]

	// verticals
	// [0,0],[0,1],[0,2]
	// [1,0],[1,1],[1,2]
	// [2,0],[2,1],[2,2]
	var winningRules [][]Coord = [][]Coord{
		{Coord{0, 0}, Coord{1, 1}, Coord{2, 2}},
		{Coord{2, 0}, Coord{1, 1}, Coord{0, 2}},

		{Coord{0, 0}, Coord{1, 0}, Coord{2, 0}},
		{Coord{0, 1}, Coord{1, 1}, Coord{2, 1}},
		{Coord{0, 2}, Coord{1, 2}, Coord{2, 2}},

		{Coord{0, 0}, Coord{0, 1}, Coord{0, 2}},
		{Coord{1, 0}, Coord{1, 1}, Coord{1, 2}},
		{Coord{2, 0}, Coord{2, 1}, Coord{2, 2}},
	}

	if (Game{}) != game {
		// checks for "X"
		// check winning rules for "X"
		board := game.board
		for i := 0; i < len(winningRules); i++ {
			rule := winningRules[i]
			if Token(board, rule[0]) == "X" && Token(board, rule[1]) == "X" && Token(board, rule[2]) == "X" {
				return "X"
			}
			if Token(board, rule[0]) == "O" && Token(board, rule[1]) == "O" && Token(board, rule[2]) == "O" {
				return "O"
			}
		}
	}
	return ""
}
