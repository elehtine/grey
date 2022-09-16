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

