package main

import (
    "os"
    "bufio"

    "github.com/elehtine/grey/reversi"
    "github.com/elehtine/grey/ui"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    board := reversi.NewBoard()
    cli := ui.NewCommandLineInterface(board, reader)

    cli.PlayGame()
}
