/*
Package reversi implements game of reversi. One game can be played with single
instance.

The reversi can be played via ui package. However reversi can be used without
ui package.
*/
package reversi

import "errors"

// Height and Width of the game board
const (
	Height = 8
	Width  = 8
)

func offBoard(file, rank int) bool {
	if file < 0 || file >= 8 || rank < 0 || rank >= 8 {
		return true
	}
	return false
}

// Enum for squares of the board
type Square int

const (
	Empty Square = iota // Empty square
	Dark                // Square with dark disk
	Light               // Square with light disk
)

func startPosition() [Width][Height]Square {
	squares := [Width][Height]Square{}
	squares[3][3] = Light
	squares[4][4] = Light
	squares[4][3] = Dark
	squares[3][4] = Dark
	return squares
}

type direction struct {
	dx, dy int
}

// A Move contains file and rank for move.
type Move struct {
	File, Rank int // File is x-coordinate, Rank is y-coordinate.
}

// Constructor for Move object
func MakeMove(file, rank int) Move {
	return Move{File: file, Rank: rank}
}

// A Board contains current state of the game. It stores turn and all squares.
type Board struct {
	squares [Width][Height]Square
	turn    Square
}

// Initialize board with starting position
func NewBoard() *Board {
	board := new(Board)
	board.squares = startPosition()
	board.turn = Dark
	return board
}

// A Status stores object for returning game turn and points.
type Status struct {
	Turn                    Square // If game is ended turn is Empty, else Dark or Light.
	DarkPoints, LightPoints int    // Points for each player.
}

// Calculate and return current status of the game
func (board *Board) Status() Status {
	return Status{
		Turn:        board.turn,
		DarkPoints:  board.count(Dark),
		LightPoints: board.count(Light),
	}
}

// Returns value of square in given file and rank
func (board *Board) Get(file, rank int) Square {
	return board.squares[file][rank]
}

// Calculate and return number of certain disks in the board
func (board *Board) count(turn Square) int {
	count := 0
	for file := 0; file < Width; file++ {
		for rank := 0; rank < Height; rank++ {
			if board.squares[file][rank] == turn {
				count++
			}
		}
	}
	return count
}

/*
If illegal move, throw error

Otherwise put new disk in the board. Flip caputerd disks. Change turn.
*/
func (board *Board) Move(move Move) error {
	file, rank := move.File, move.Rank
	if board.turn == Empty {
		return errors.New("Game over")
	}
	if !board.isLegalMove(move) {
		return errors.New("Illegal move")
	}

	captureCount := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			captureCount += board.capture(file, rank, direction{dx, dy}, board.turn)
		}
	}

	if captureCount == 0 {
		return errors.New("Move is not capture")
	}

	board.squares[file][rank] = board.turn
	board.nextTurn()
	return nil
}

func (board *Board) isLegalMove(move Move) bool {
	for _, m := range board.legalMoves(board.turn) {
		if m == move {
			return true
		}
	}
	return false
}

// Return all legal moves for player in turn
func (board *Board) Moves() []Move {
	return board.legalMoves(board.turn)
}

func (board *Board) legalMoves(turn Square) []Move {
	moves := make([]Move, 0)
	for file := 0; file < Width; file++ {
		for rank := 0; rank < Height; rank++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dy == 0 && dx == 0 {
						continue
					}
					dir := direction{dx, dy}
					if board.canCapture(file, rank, dir, turn) {
						moves = append(moves, MakeMove(file, rank))
					}
				}
			}
		}
	}
	return moves
}

func (board *Board) canCapture(file, rank int, dir direction, turn Square) bool {
	if board.squares[file][rank] != Empty {
		return false
	}

	file += dir.dx
	rank += dir.dy
	if offBoard(file, rank) {
		return false
	}

	if square := board.squares[file][rank]; square == Empty || square == turn {
		return false
	}

	for board.squares[file][rank] != turn {
		file += dir.dx
		rank += dir.dy
		if offBoard(file, rank) || board.squares[file][rank] == Empty {
			return false
		}
	}
	return true
}

func (board *Board) capture(file, rank int, dir direction, turn Square) int {
	if !board.canCapture(file, rank, dir, turn) {
		return 0
	}

	count := 0
	startFile, startRank := file+dir.dx, rank+dir.dy
	for file, rank = startFile, startRank; board.squares[file][rank] != turn; count++ {
		board.squares[file][rank] = turn
		file += dir.dx
		rank += dir.dy
	}
	return count
}

func (board *Board) nextTurn() {
	lightMoves := board.legalMoves(Light)
	darkMoves := board.legalMoves(Dark)
	lightCapture := len(lightMoves) > 0
	darkCapture := len(darkMoves) > 0

	if !lightCapture && !darkCapture {
		board.turn = Empty
		return
	}

	if board.turn == Light && darkCapture {
		board.turn = Dark
		return
	}

	if lightCapture {
		board.turn = Light
	} else {
		board.turn = Dark
	}
}
