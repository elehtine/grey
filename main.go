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

    uiBuilder := ui.NewUserInterfaceBuilder(board, reader)
    uiBuilder.DarkPlayer("bot")
    uiBuilder.LightPlayer("easy")
    userInterface := uiBuilder.GetUserInterface()
    userInterface.PlayGame()
}
