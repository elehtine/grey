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


func (board *Board) Move(x, y int) error {
    if board.tiles[x][y] != Empty {
        return errors.New("Square not empty")
    }

    captureCount := 0
    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dy == 0 && dx == 0 {
                continue
            }
            captureCount += board.capture(x, y, Direction{dx, dy})
        }
    }

    if captureCount == 0 {
        return errors.New("Move is not capture")
    }

    board.tiles[x][y] = board.turn
    board.nextTurn()
    return nil
}

func (board *Board) capture(x, y int, dir Direction) int {
    startX, startY := x + dir.dx, y + dir.dy
    if offBoard(startX, startY) || board.tiles[startX][startY] == board.turn {
        return 0
    }

    x, y = startX, startY
    for board.tiles[x][y] != board.turn {
        x += dir.dx
        y += dir.dy
        if offBoard(x, y) || board.tiles[x][y] == Empty {
            return 0
        }
    }

    count := 0
    for x, y = startX, startY; board.tiles[x][y] != board.turn; count++ {
        board.tiles[x][y] = board.turn
        x += dir.dx
        y += dir.dy
    }
    return count
}



func (board *Board) nextTurn() {
    if board.turn == Light {
        board.turn = Dark
        return
    }
    board.turn = Light
}
