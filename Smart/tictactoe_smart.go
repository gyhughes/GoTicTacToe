
package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
)

func main() {
	// Initialize variables.
	board := [][]int{{0, 0, 0},{0, 0, 0},{0, 0, 0}}
	var finished int = -1
	shooter1 := make(chan bool)
	shooter2 := make(chan bool)
	var s1acc int64 = 60
	var s2acc int64 = 80
	if (len(os.Args) == 3) {
		// Use cmd arguments for custom shot accuracy.
		s1acc, _ = strconv.ParseInt(os.Args[1], 10, 0)
		s2acc, _ = strconv.ParseInt(os.Args[2], 10, 0)
	}

	// Play game.
	go shootBasketball(int(s1acc), shooter1)
	go shootBasketball(int(s2acc), shooter2)
	for (finished < 0) {
		var playerNum int = getMove(shooter1, shooter2)
		placePiece(playerNum, board)
		printBoard(board)
		// Check if game is over.
		finished = isGameOver(board)
	}
	printFinished(finished)
}

// Notifies shooter when shot is made.
func shootBasketball(accuracy int, s chan bool) {
	for {
		for (rand.Intn(100) + 1 > accuracy) {
		}
		s <- true
	}
}

// Waits for player to make shot to place piece.
func getMove(shooter1, shooter2 chan bool) int {
	// Wait for action.
	select {
	case <-shooter1:
		return 1
	case <-shooter2:
		return 2
	}
}

// Prints the board in ascii.
func printBoard(board [][]int) {
	XO := [][]string{{" ", " ", " "},{" ", " ", " "},{" ", " ", " "}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (board[i][j] == 1) {
				XO[i][j] = "X"
			} else if (board[i][j] == 2) {
				XO[i][j] = "O"
			}
		}
	}

	fmt.Printf("+---+---+---+\n")
	fmt.Printf("| %s | %s | %s |\n", XO[0][0], XO[0][1], XO[0][2])
	fmt.Printf("+---+---+---+\n")
	fmt.Printf("| %s | %s | %s |\n", XO[1][0], XO[1][1], XO[1][2])
	fmt.Printf("+---+---+---+\n")
	fmt.Printf("| %s | %s | %s |\n", XO[2][0], XO[2][1], XO[2][2])
	fmt.Printf("+---+---+---+\n\n")
}

// Prints finished message.
func printFinished(finished int) {
	if (finished == 1) {
		fmt.Println("Player 1 wins!")
	} else if (finished == 2) {
		fmt.Println("Player 2 wins!")
	} else {
		fmt.Println("Cat's game! Play again!")
	}
}

// Returns true if board is completely filled, else false.
func isBoardFull(board [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (board[i][j] == 0) {
				return false
			}
		}
	}
	return true
}

// Places player's piece on board.
func placePiece(playerNumber int, board [][]int) {
	// Sanity check. Possibly unneeded.
	if (isBoardFull(board)) {
		return
	}

	// Check if a move wins the game.
	if (placeWinningMove(playerNumber, board)) {
		return
	}

	// Check if a move stops the opponent from winning.
	if (placeBlockingMove(playerNumber, board)) {
		return
	}

	// No strategies to use, place randomly.
	for {
		i := rand.Intn(3)
		j := rand.Intn(3)
		if (board[i][j] == 0) {
			board[i][j] = playerNumber
			return
		}
	}
}

// Returns player number of the winner.
// Returns -1 if game is not over.
// Returns 0 if cat's game. (Tie)
func isGameOver(board [][]int) int {
	if (board[0][0] != 0) {
		if ((board[0][0] == board[0][1] && board[0][1] == board[0][2]) ||
			(board[0][0] == board[1][0] && board[1][0] == board[2][0])) {
			return board[0][0]
		}
	}
	if (board[1][1] != 0) {
		if ((board[0][1] == board[1][1] && board[1][1] == board[2][1]) ||
			(board[1][0] == board[1][1] && board[1][1] == board[1][2]) ||
			(board[0][0] == board[1][1] && board[1][1] == board[2][2]) ||
			(board[2][0] == board[1][1] && board[1][1] == board[0][2])) {
			return board[1][1]
		}
	}
	if(board[2][2] != 0) {
		if ((board[2][0] == board[2][1] && board[2][1] == board[2][2]) ||
			(board[2][0] == board[2][1] && board[2][1] == board[2][2])) {
			return board[2][2]
		}
	}
	if (isBoardFull(board)) {
		// No winner.
		return 0
	}
	// Game not over.
	return -1
}

// Tries to find a spot on the board to win the game.
// Returns true if a spot was found, else false.
func placeWinningMove(playerNumber int, board [][]int) bool {
	if (board[0][0] == 0) {
		if ((board[0][1] == playerNumber && board[0][1] == board[0][2]) ||
			(board[1][0] == playerNumber && board[1][0] == board[2][0]) ||
			(board[1][1] == playerNumber && board[1][1] == board[2][2])) {
			board[0][0] = playerNumber
			return true
		}
	}
	if (board[0][1] == 0) {
		if ((board[0][0] == playerNumber && board[0][0] == board[0][2]) ||
			(board[1][1] == playerNumber && board[1][1] == board[2][1])) {
			board[0][1] = playerNumber
			return true
		}
	}
	if (board[0][2] == 0) {
		if ((board[0][0] == playerNumber && board[0][0] == board[0][1]) ||
			(board[1][2] == playerNumber && board[1][2] == board[2][2]) ||
			(board[1][1] == playerNumber && board[1][1] == board[2][0])) {
			board[0][2] = playerNumber
			return true
		}
	}
	if (board[1][0] == 0) {
		if ((board[0][0] == playerNumber && board[0][0] == board[2][0]) ||
			(board[1][1] == playerNumber && board[1][1] == board[1][2])) {
			board[1][0] = playerNumber
			return true
		}
	}
	if (board[1][1] == 0) {
		if ((board[0][0] == playerNumber && board[0][0] == board[2][2]) ||
			(board[0][2] == playerNumber && board[0][2] == board[2][0]) ||
			(board[1][0] == playerNumber && board[1][0] == board[1][2]) ||
			(board[0][1] == playerNumber && board[0][1] == board[2][1])) {
			board[1][1] = playerNumber
			return true
		}
	}
	if (board[1][2] == 0) {
		if ((board[1][0] == playerNumber && board[1][0] == board[1][1]) ||
			(board[0][2] == playerNumber && board[0][2] == board[2][2])) {
			board[1][2] = playerNumber
			return true
		}
	}
	if (board[2][0] == 0) {
		if ((board[0][0] == playerNumber && board[0][0] == board[1][0]) ||
			(board[2][1] == playerNumber && board[2][1] == board[2][2]) ||
			(board[1][1] == playerNumber && board[1][1] == board[0][2])) {
			board[1][0] = playerNumber
			return true
		}
	}
	if (board[2][1] == 0) {
		if ((board[2][0] == playerNumber && board[2][0] == board[2][2]) ||
			(board[0][1] == playerNumber && board[0][1] == board[1][1])) {
			board[2][1] = playerNumber
			return true
		}
	}
	if (board[2][2] == 0) {
		if ((board[0][0] == playerNumber && board[0][0] == board[1][1]) ||
			(board[2][0] == playerNumber && board[2][0] == board[2][1]) ||
			(board[0][2] == playerNumber && board[0][2] == board[1][2])) {
			board[2][2] = playerNumber
			return true
		}
	}
	// No winning moves available.
	return false
}

// Tries to find a spot on the board to prevent a win.
// Returns true if a spot was found, else false.
func placeBlockingMove(playerNumber int, board [][]int) bool {
	var opponentNumber int
	if (playerNumber == 1) {
		opponentNumber = 2
	} else {
		opponentNumber = 1
	}
	if (board[0][0] == 0) {
		if ((board[0][1] == opponentNumber && board[0][1] == board[0][2]) ||
			(board[1][0] == opponentNumber && board[1][0] == board[2][0]) ||
			(board[1][1] == opponentNumber && board[1][1] == board[2][2])) {
			board[0][0] = playerNumber
			return true
		}
	}
	if (board[0][1] == 0) {
		if ((board[0][0] == opponentNumber && board[0][0] == board[0][2]) ||
			(board[1][1] == opponentNumber && board[1][1] == board[2][1])) {
			board[0][1] = playerNumber
			return true
		}
	}
	if (board[0][2] == 0) {
		if ((board[0][0] == opponentNumber && board[0][0] == board[0][1]) ||
			(board[1][2] == opponentNumber && board[1][2] == board[2][2]) ||
			(board[1][1] == opponentNumber && board[1][1] == board[2][0])) {
			board[0][2] = playerNumber
			return true
		}
	}
	if (board[1][0] == 0) {
		if ((board[0][0] == opponentNumber && board[0][0] == board[2][0]) ||
			(board[1][1] == opponentNumber && board[1][1] == board[1][2])) {
			board[1][0] = playerNumber
			return true
		}
	}
	if (board[1][1] == 0) {
		if ((board[0][0] == opponentNumber && board[0][0] == board[2][2]) ||
			(board[0][2] == opponentNumber && board[0][2] == board[2][0]) ||
			(board[1][0] == opponentNumber && board[1][0] == board[1][2]) ||
			(board[0][1] == opponentNumber && board[0][1] == board[2][1])) {
			board[1][1] = playerNumber
			return true
		}
	}
	if (board[1][2] == 0) {
		if ((board[1][0] == opponentNumber && board[1][0] == board[1][1]) ||
			(board[0][2] == opponentNumber && board[0][2] == board[2][2])) {
			board[1][2] = playerNumber
			return true
		}
	}
	if (board[2][0] == 0) {
		if ((board[0][0] == opponentNumber && board[0][0] == board[1][0]) ||
			(board[2][1] == opponentNumber && board[2][1] == board[2][2]) ||
			(board[1][1] == opponentNumber && board[1][1] == board[0][2])) {
			board[1][0] = playerNumber
			return true
		}
	}
	if (board[2][1] == 0) {
		if ((board[2][0] == opponentNumber && board[2][0] == board[2][2]) ||
			(board[0][1] == opponentNumber && board[0][1] == board[1][1])) {
			board[2][1] = playerNumber
			return true
		}
	}
	if (board[2][2] == 0) {
		if ((board[0][0] == opponentNumber && board[0][0] == board[1][1]) ||
			(board[2][0] == opponentNumber && board[2][0] == board[2][1]) ||
			(board[0][2] == opponentNumber && board[0][2] == board[1][2])) {
			board[2][2] = playerNumber
			return true
		}
	}
	// No blocking moves available.
	return false
}