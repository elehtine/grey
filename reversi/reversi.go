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

func startPosition() [Width][Height]Tile {
    tiles := [Width][Height]Tile {}
    tiles[3][3] = Light
    tiles[4][4] = Light
    tiles[4][3] = Dark
    tiles[3][4] = Dark
    return tiles
}

type Board struct {
    tiles [Width][Height]Tile
}


func NewBoard() *Board {
    board := new(Board)
    board.tiles = startPosition()
    return board
}

func (board *Board) Get(x, y int) Tile {
    return board.tiles[x][y]
}

func (board *Board) Move(x, y int) {
}
