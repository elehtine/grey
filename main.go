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
        InputMove(board, reader)
    }
}
