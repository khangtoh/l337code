package tictactoe

import "testing"

var testCases = []struct {
	game   Game
	winner string
}{
	{
		Game{
			board: [3][3]string{
				{"X", "X", "X"},
				{"O", "-", "-"},
				{"-", "-", "-"},
			},
		},
		"X",
	},
}

func TestCheckWinner(t *testing.T) {
	for _, tc := range testCases {
		winner := CheckWinner(tc.game)
		if winner != tc.winner {
			t.Fatalf("Winner is not X and should be X")
		}
	}
}
