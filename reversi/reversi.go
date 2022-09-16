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

    board.tiles[x][y] = board.turn
    board.nextTurn()
    return nil
}

func (board *Board) nextTurn() {
    if board.turn == Light {
        board.turn = Dark
        return
    }
    board.turn = Light
}
