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
    tiles := [Width][Height]Square {}
    tiles[3][3] = Light
    tiles[4][4] = Light
    tiles[4][3] = Dark
    tiles[3][4] = Dark
    return tiles
}

type Board struct {
    tiles [Width][Height]Square
    turn Square
}

type Direction struct {
    dx, dy int
}

func NewBoard() *Board {
    board := new(Board)
    board.tiles = startPosition()
    board.turn = Dark
    return board
}

func (board *Board) Get(x, y int) Square {
    return board.tiles[x][y]
}

type Status struct {
    turn Square
    darkPoints, lightPoints int
}

func (board *Board) Status() Status {
    return Status{
        turn: board.turn,
        darkPoints: board.count(Dark),
        lightPoints: board.count(Light),
    }
}

func (board *Board) count(turn Square) int {
    count := 0
    for rank := 0; rank < Height; rank++ {
        for file := 0; file < Width; file++ {
            if board[x][y] == turn {
                count++
            }
        }
    }
    return count
}

func (board *Board) Move(x, y int) error {
    if board.turn == Empty {
        return errors.New("Game over")
    }
    if board.tiles[x][y] != Empty {
        return errors.New("Square not empty")
    }

    captureCount := 0
    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dy == 0 && dx == 0 {
                continue
            }
            captureCount += board.capture(x, y, Direction{dx, dy}, board.turn)
        }
    }

    if captureCount == 0 {
        return errors.New("Move is not capture")
    }

    board.tiles[x][y] = board.turn
    board.nextTurn()
    return nil
}

func (board *Board) canCapture(x, y int, dir Direction, turn Square) bool {
    if board.tiles[x][y] != Empty {
        return false
    }

    x += dir.dx
    y += dir.dy
    if offBoard(x, y) || board.tiles[x][y] == Empty || board.tiles[x][y] == turn {
        return false
    }

    for board.tiles[x][y] != turn {
        x += dir.dx
        y += dir.dy
        if offBoard(x, y) || board.tiles[x][y] == Empty {
            return false
        }
    }
    return true
}

func (board *Board) capture(x, y int, dir Direction, turn Square) int {
    if !board.canCapture(x, y, dir, turn) {
        return 0
    }

    count := 0
    startX, startY := x + dir.dx, y + dir.dy
    for x, y = startX, startY; board.tiles[x][y] != turn; count++ {
        board.tiles[x][y] = turn
        x += dir.dx
        y += dir.dy
    }
    return count
}


func (board *Board) nextTurn() {
    lightCapture, darkCapture := false, false
    for file := 0; file < Width; file++ {
        for rank := 0; rank < Height; rank++ {
            for dy := -1; dy <= 1; dy++ {
                for dx := -1; dx <= 1; dx++ {
                    if dy == 0 && dx == 0 {
                        continue
                    }
                    dir := Direction{dx, dy}
                    if board.canCapture(file, rank, dir, Light) {
                        lightCapture = true
                    }
                    if board.canCapture(file, rank, dir, Dark) {
                        darkCapture = true
                    }
                }
            }
        }
    }


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
