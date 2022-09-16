package reversi


import (
    "testing"
)

func TestBoardStartPosition(t *testing.T) {
    board := NewBoard()
    for y := 0; y < Height; y++ {
        for x := 0; x < Width; x++ {
            if (x == 3 || x == 4) && (y == 3 || y == 4) {
                tile := board.Get(x, y)
                if tile == Empty {
                    t.Errorf("Get(%d, %d) = %d, want not Empty", x, y, tile)
                }
                continue
            }
            tile := board.Get(x, y)
            if tile != Empty {
                t.Errorf("Get(%d, %d) = %d, want Empty", x, y, tile)
            }
        }
    }
}

type Move struct {
    x, y int
    turn Tile
}

func TestTwoIllegalMoves(t *testing.T) {
    board := NewBoard()

    moves := []Move {
        Move{x: 5, y: 3, turn: Dark},
        Move{x: 4, y: 2, turn: Light},
    }

    for _, m := range moves {
        board.Move(m.x, m.y)
        tile := board.Get(m.x, m.y)
        if tile != m.turn {
            t.Errorf("Get(%d, %d) = %d, want %d", m.x, m.y, tile, m.turn)
        }
    }
}

