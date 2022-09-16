package main

import (
    "fmt"
    "github.com/elehtine/grey/reversi"
)

func Draw(b *reversi.Board) {
    fmt.Printf("Board %dx%d\n", reversi.Width, reversi.Height)
}
