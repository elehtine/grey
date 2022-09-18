package reversi

import "errors"

const (
    Height = 8
    Width = 8
)

func offBoard(file, rank int) bool {
    if file < 0 || file >= 8 || rank < 0 || rank >= 8 {
        return true
    }
    return false
}

type Square int
const (
    Empty Square = iota
    Dark
    Light
)

func startPosition() [Width][Height]Square {
    squares := [Width][Height]Square {}
    squares[3][3] = Light
    squares[4][4] = Light
    squares[4][3] = Dark
    squares[3][4] = Dark
    return squares
}

type direction struct {
    dx, dy int
}
type Move struct {
    File, Rank int
}

func MakeMove(file, rank int) Move {
    return Move{File: file, Rank: rank}
}

type Board struct {
    squares [Width][Height]Square
    turn Square
}
func NewBoard() *Board {
    board := new(Board)
    board.squares = startPosition()
    board.turn = Dark
    return board
}

type Status struct {
    Turn Square
    DarkPoints, LightPoints int
}

func (board *Board) Status() Status {
    return Status{
        Turn: board.turn,
        DarkPoints: board.count(Dark),
        LightPoints: board.count(Light),
    }
}

func (board *Board) Get(file, rank int) Square {
    return board.squares[file][rank]
}

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
    startFile, startRank := file + dir.dx, rank + dir.dy
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
