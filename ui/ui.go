/*
Package ui is user interface between player moves and reversi game.

The ui package can be replaced with own user interface. In other words ui
package doesn't affect reversi package in any way.
*/
package ui

import (
    "fmt"
    "bufio"
    "errors"
    "math/rand"

    "github.com/elehtine/grey/reversi"
)

type player interface {
    move()
}

type user struct {
    board *reversi.Board
    scanner *bufio.Scanner
}

type bot struct {
    board *reversi.Board
    easy bool
}

/*
A UserInterfaceBuilder is builder for userInterface object used to set dark and
light players.
*/
type UserInterfaceBuilder struct {
    board *reversi.Board
    scanner *bufio.Scanner
    darkPlayer player
    lightPlayer player
}

/*
Create new instance of ui builder. Dark and light players set to bots as
default.

It is good practice to call DarkPlayer and LightPlayer methods both once.
*/
func NewUserInterfaceBuilder(board *reversi.Board, scanner *bufio.Scanner) *UserInterfaceBuilder {
    return &UserInterfaceBuilder{
        board: board,
        scanner: scanner,
        darkPlayer: &bot{board: board, easy: false},
        lightPlayer: &bot{board: board, easy: false},
    }
}

// Set value for dark player of object to be build.
func (uiBuilder *UserInterfaceBuilder) DarkPlayer(player string) {
    uiBuilder.darkPlayer = uiBuilder.createPlayer(player)
}

// Set value for light player of object to be build.
func (uiBuilder *UserInterfaceBuilder) LightPlayer(player string) {
    uiBuilder.lightPlayer = uiBuilder.createPlayer(player)
}

// Return
func (uiBuilder *UserInterfaceBuilder) GetUserInterface() *userInterface {
    return &userInterface{
        board: uiBuilder.board,
        darkPlayer: uiBuilder.darkPlayer,
        lightPlayer: uiBuilder.lightPlayer,
    }
}

func (uiBuilder *UserInterfaceBuilder) createPlayer(player string) player {
    if player == "bot" {
        return &bot{
            board: uiBuilder.board,
            easy: false,
        }
    }
    if player == "user" {
        return &user{
            board: uiBuilder.board,
            scanner: uiBuilder.scanner,
        }
    }
    return &bot{
        board: uiBuilder.board,
        easy: true,
    }
}

func (user *user) move() {
    fmt.Print("Give move: ")
    for {
        user.scanner.Scan()
        move := user.scanner.Text()

        file, rank, err := parseMove(move)
        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        m := reversi.MakeMove(file, rank)
        err = user.board.Move(m)
        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        break
    }
}

func (bot *bot) move() {
    legalMoves := bot.board.Moves()
    if bot.easy {
        index := rand.Intn(len(legalMoves))
        bot.board.Move(legalMoves[index])
        return
    }

    index, best := 0, score(legalMoves[0])
    for i, move := range legalMoves {
        if best > score(move) {
            best = score(move)
            index = i
        }
    }
    bot.board.Move(legalMoves[index])
}

func score(move reversi.Move) int {
    file, rank := move.File, move.Rank
    if file >= 4 {
        file = 7 - file
    }
    if rank >= 4 {
        rank = 7 - rank
    }
    return file*file + rank*rank
}

type userInterface struct {
    board *reversi.Board
    darkPlayer player
    lightPlayer player
}

func (ui *userInterface) PlayGame() {
    for ui.board.Status().Turn != reversi.Empty {
        turn := ui.board.Status().Turn
        draw(ui.board)
        if turn == reversi.Dark {
            ui.darkPlayer.move()
        } else if turn == reversi.Light {
            ui.lightPlayer.move()
        }
    }
    result(ui.board)
}

func draw(board *reversi.Board) {
    fmt.Println("  abcdefgh")
    fmt.Println(" +--------+")
    for rank := 0; rank < reversi.Height; rank++ {
        rankNum := rank + 1
        fmt.Printf("%d|", rankNum)
        for file := 0; file < reversi.Width; file++ {
            tile := "."
            switch board.Get(file, rank) {
            case reversi.Dark:
                tile = "x"
            case reversi.Light:
                tile = "o"
            }
            fmt.Print(tile)
        }
        fmt.Printf("|%d\n", rankNum)
    }
    fmt.Println(" +--------+")
    fmt.Println("  abcdefgh")
}

func result(board *reversi.Board) {
    st := board.Status()
    draw(board)
    fmt.Printf("Dark: %d, Light: %d\n", st.DarkPoints, st.LightPoints)
}

func parseMove(moveString string) (int, int, error) {
    if len(moveString) != 2 {
        return -1, -1, errors.New("Wrong number of characters in move")
    }

    file := int(moveString[0] - 97)
    rank := int(moveString[1] - 49)

    if rank < 0 || rank >= 8 || file < 0 || file >= 8 {
        return -1, -1, errors.New("Malformatted file or rank")
    }

    return file, rank, nil
}
