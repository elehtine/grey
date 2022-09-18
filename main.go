package main

import (
    "os"
    "bufio"
    "fmt"

    "github.com/elehtine/grey/reversi"
    "github.com/elehtine/grey/ui"
)

func main() {
    board := reversi.NewBoard()
    scanner := bufio.NewScanner(os.Stdin)
    uiBuilder := ui.NewUserInterfaceBuilder(board, scanner)

    fmt.Print("Choose dark player (user, bot) ")
    scanner.Scan()
    darkPlayer := scanner.Text()

    fmt.Print("Choose light player (user, bot) ")
    scanner.Scan()
    lightPlayer := scanner.Text()

    uiBuilder.DarkPlayer(darkPlayer)
    uiBuilder.LightPlayer(lightPlayer)
    userInterface := uiBuilder.GetUserInterface()
    userInterface.PlayGame()
}
