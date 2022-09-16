package reversi

const (
    Height = 8
    Width = 8
)

type Tile int
const (
    Empty Tile = iota
    Black
    White
)


type Board struct {
    tiles [Width][Height]Tile
}

func NewBoard() *Board {
    board := new(Board)
    board.tiles[3][3] = White
    board.tiles[4][4] = White
    board.tiles[4][3] = Black
    board.tiles[3][4] = Black
    return board
}

func (board *Board) Get(x, y int) Tile {
    return board.tiles[x][y]
}
