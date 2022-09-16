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
}

func NewBoard() *Board {
    return &Board{}
}

func (board *Board) get(x, y int) Tile {
    return Empty
}
