package ui

import (
    "fmt"
    "bufio"
    "errors"

    "github.com/elehtine/grey/reversi"
)

type UserInterface interface {
    PlayGame()
}

type CommandLineInterface struct {
    board *reversi.Board
    reader *bufio.Reader
}

type AutoPlay struct {
    board *reversi.Board
}

func NewAutoPlay(board *reversi.Board) *AutoPlay {
    return &AutoPlay{board: board}
}

func (play *AutoPlay) PlayGame() {
    for play.board.Status().Turn != reversi.Empty {
        draw(play.board)
        play.pickMove()
    }
    result(play.board)
}

func (play *AutoPlay) pickMove() {
    legalMoves := play.board.Moves()
    play.board.Move(legalMoves[0])
}

func NewCommandLineInterface(b *reversi.Board, r *bufio.Reader) *CommandLineInterface {
    cli := CommandLineInterface{board: b, reader: r}
    return &cli
}

func (cli *CommandLineInterface) PlayGame() {
    for cli.board.Status().Turn != reversi.Empty {
        draw(cli.board)
        cli.inputMove()
    }
    result(cli.board)
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

func (cli *CommandLineInterface) inputMove() {
    fmt.Print("Give move: ")
    var file, rank int

    for {
        move, _, err := cli.reader.ReadLine()
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
        err = cli.board.Move(m)
        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        break
    }
}

func result(board *reversi.Board) {
    st := board.Status()
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
