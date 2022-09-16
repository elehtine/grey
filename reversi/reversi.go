package reversi

const (
    Height = 8
    Width = 8
)

type Tile int
const (
    Empty Tile = iota
    Dark
    Light
)


type Board struct {
    tiles [Width][Height]Tile
}

func NewBoard() *Board {
    board := new(Board)
    board.tiles[3][3] = Light
    board.tiles[4][4] = Light
    board.tiles[4][3] = Dark
    board.tiles[3][4] = Dark
    return board
}

func (board *Board) Get(x, y int) Tile {
    return board.tiles[x][y]
}
