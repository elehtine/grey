package main

import (
    "os"
    "bufio"

    "github.com/elehtine/grey/reversi"
    "github.com/elehtine/grey/ui"
)

func botGame(board *reversi.Board) ui.UserInterface {
    return ui.NewAutoPlay(board)
}

func twoPlayerGame(board *reversi.Board) ui.UserInterface {
    reader := bufio.NewReader(os.Stdin)
    return ui.NewCommandLineInterface(board, reader)
}

func main() {
    board := reversi.NewBoard()
    ui := botGame(board)
    ui.PlayGame()
}
