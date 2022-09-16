package main

import (
    "fmt"
    "github.com/elehtine/grey/reversi"
)


func Draw(b *reversi.Board) {
    fmt.Printf("Board %dx%d\n", reversi.Width, reversi.Height)
    for rank := 0; rank < reversi.Height; rank++ {
        fmt.Println("   ox   ")
    }
    fmt.Println("--------")
    fmt.Println("ABCDEFGH")
}
