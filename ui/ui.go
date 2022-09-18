package ui

import (
    "fmt"
    "bufio"
    "errors"
    "math/rand"

    "github.com/elehtine/grey/reversi"
)

type Player interface {
    move()
}

type User struct {
    board *reversi.Board
    reader *bufio.Reader
}

type Bot struct {
    board *reversi.Board
    easy bool
}

type PlayerGenerator struct {
    board *reversi.Board
    reader *bufio.Reader
}

func NewPlayerGenerator(board *reversi.Board, reader *bufio.Reader) *PlayerGenerator {
    return &PlayerGenerator{board: board, reader: reader}
}

func (playerGenerator *PlayerGenerator) CreateUser() Player {
    return &User{board: playerGenerator.board, reader: playerGenerator.reader}
}

func (playerGenerator *PlayerGenerator) CreateBot() Player {
    return &Bot{board: playerGenerator.board, easy: false}
}

func (user *User) move() {
    fmt.Print("Give move: ")
    var file, rank int

    for {
        move, _, err := user.reader.ReadLine()
        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        file, rank, err = parseMove(move)
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

func (bot *Bot) move() {
    legalMoves := bot.board.Moves()
    if bot.easy {
        n := len(legalMoves)
        bot.board.Move(legalMoves[rand.Intn(n)])
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

type UserInterface struct {
    board *reversi.Board
    darkPlayer Player
    lightPlayer Player
}

func NewUserInterface(board *reversi.Board, dark, light Player) *UserInterface {
    return &UserInterface{board: board, darkPlayer: dark, lightPlayer: light}
}

func (ui *UserInterface) PlayGame() {
    for ui.board.Status().Turn != reversi.Empty {
        draw(ui.board)
        if ui.board.Status().Turn == reversi.Dark {
            ui.darkPlayer.move()
        } else {
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

func parseMove(moveString []byte) (int, int, error) {
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
