package main

import (
    "os"
    "bufio"

    "github.com/elehtine/grey/reversi"
)

func main() {
    board := reversi.NewBoard()
    reader := bufio.NewReader(os.Stdin)

    for {
	    Draw(board)
        file, rank := InputMove(board, reader)
        board.Move(file, rank)
    }
}
