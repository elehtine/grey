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
    board := NewBoard()
    moves := []Move {
        MakeMove(4, 5),
        MakeMove(5, 3),
    }

    for _, m := range moves {
        _ = board.Move(m)
        tile := board.Get(m.file, m.rank)
        if tile == Empty {
            t.Errorf("Get(%d, %d) = %d", m.file, m.rank, tile)
        }
    }
}

func TestTryTwoIllegalMoves(t *testing.T) {
    board := NewBoard()
    moves := []Move {
        MakeMove(4, 4),
        MakeMove(2, 2),
    }

    for _, m := range moves {
        tileBefore := board.Get(m.file, m.rank)

        err := board.Move(m)
        if err == nil {
            t.Error("Error expected from illegal move")
        }

        tile := board.Get(m.file, m.rank)
        if tile != tileBefore {
            t.Errorf("Get(%d, %d) = %d, want %d", m.file, m.rank, tile, tileBefore)
        }
    }
}

func TestTwoMovesCheckCapture(t *testing.T) {
	board := NewBoard()
	moves := []Move {
		MakeMove(2, 3),
		MakeMove(2, 4),
	}
	flips := []Move {
		MakeMove(3, 3),
		MakeMove(3, 4),
    }

    for i, m := range moves {
        _ = board.Move(m)
        f := flips[i]
        tile := board.Get(m.file, m.rank)
        flipTile := board.Get(f.file, f.rank)
        if tile != flipTile {
            t.Errorf("Not equal %d == %d", tile, flipTile)
        }
    }
}

func TestSkipMoves(t *testing.T) {
    board := new(Board)
    squares := [Width][Height]Square {}
    squares[3][1] = Light
    squares[3][2] = Dark
    squares[3][3] = Light
    board.squares = squares
    board.turn = Dark

    board.Move(MakeMove(3, 0))
    if board.turn != Dark {
        t.Errorf("Turn should be %d but it is %d", Dark, board.turn)
    }
}
