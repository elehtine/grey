package reversi

import "errors"

const (
    Height = 8
    Width = 8
)

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

    capture := false
    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dy == 0 && dx == 0 {
                continue
            }
            if board.canCapture(x, y, Direction{dx, dy}) {
                capture = true
            }
        }
    }

    if !capture {
        return errors.New("Move is not capture")
    }

    board.tiles[x][y] = board.turn
    board.nextTurn()
    return nil
}

func (board *Board) canCapture(x, y int, dir Direction) bool {
    x += dir.dx
    y += dir.dy
    if x < 0 || x >= 8 || y < 0 || y >= 8 || board.tiles[x][y] == board.turn {
        return false
    }
    for {
        if x < 0 || x >= 8 || y < 0 || y >= 8 || board.tiles[x][y] == Empty {
            return false
        }
        if board.tiles[x][y] == board.turn {
            return true
        }
        x += dir.dx
        y += dir.dy
    }
}



func (board *Board) nextTurn() {
    if board.turn == Light {
        board.turn = Dark
        return
    }
    board.turn = Light
}
