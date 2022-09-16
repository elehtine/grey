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

func TestTwoLegalMoves(t *testing.T) {
    type Move struct {
        x, y int
        turn Square
    }

    board := NewBoard()
    moves := []Move {
        Move{x: 4, y: 5, turn: Dark},
        Move{x: 5, y: 3, turn: Light},
    }

    for _, m := range moves {
        _ = board.Move(m.x, m.y)
        tile := board.Get(m.x, m.y)
        if tile != m.turn {
            t.Errorf("Get(%d, %d) = %d, want %d", m.x, m.y, tile, m.turn)
        }
    }
}

func TestTryTwoIllegalMoves(t *testing.T) {
    type Move struct {
        x, y int
    }

    board := NewBoard()
    moves := []Move {
        Move{x: 4, y: 4},
        Move{x: 2, y: 2},
    }

    for _, m := range moves {
        tileBefore := board.Get(m.x, m.y)

        err := board.Move(m.x, m.y)
        if err == nil {
            t.Error("Error expected from illegal move")
        }

        tile := board.Get(m.x, m.y)
        if tile != tileBefore {
            t.Errorf("Get(%d, %d) = %d, want %d", m.x, m.y, tile, tileBefore)
        }
    }
}

func TestTwoMovesCheckCapture(t *testing.T) {
    type Move struct {
        x, y int
        flipX, flipY int
    }

    board := NewBoard()
    moves := []Move {
        Move{x: 2, y: 3, flipX: 3, flipY: 3},
        Move{x: 2, y: 4, flipX: 3, flipY: 4},
    }

    for _, m := range moves {
        _ = board.Move(m.x, m.y)
        tile := board.Get(m.x, m.y)
        flipTile := board.Get(m.flipX, m.flipY)
        if tile != flipTile {
            t.Errorf("Not equal %d == %d", tile, flipTile)
        }
    }
}

func TestSkipMoves(t *testing.T) {
    board := new(Board)
    tiles := [Width][Height]Square {}
    tiles[3][1] = Light
    tiles[3][2] = Dark
    tiles[3][3] = Light
    board.tiles = tiles
    board.turn = Dark

    board.Move(3, 0)
    if board.turn != Dark {
        t.Errorf("Turn should be %d but it is %d", Dark, board.turn)
    }
}
