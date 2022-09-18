package main

import (
    "os"
    "bufio"

    "github.com/elehtine/grey/reversi"
    "github.com/elehtine/grey/ui"
)

func main() {
    board := reversi.NewBoard()
    reader := bufio.NewReader(os.Stdin)

    playerGenerator := ui.NewPlayerGenerator(board, reader)
    dark := playerGenerator.CreateUser()
    light := playerGenerator.CreateBot()
    userInterface := ui.NewUserInterface(board, dark, light)
    userInterface.PlayGame()
}
