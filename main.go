/*
Grey is reversi game programmed using Golang.

Reversi can be played via command line. User can choose if dark and light
are players or bots.

Usage:

	grey [flags]


The flags are:

	-dark
		Dark player (user, bot, random)

	-light
		Light player (user, bot, random)

One match is played between given players.
*/
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/elehtine/grey/reversi"
	"github.com/elehtine/grey/ui"
)

type player string

func (p *player) String() string {
	return string(*p)
}

func (p *player) Set(value string) error {
	switch value {
	case "user", "bot", "random":
		*p = player(value)
	default:
		return errors.New("Malformatted player")
	}
	return nil
}

var darkPlayer player = "user"
var lightPlayer player = "bot"

func init() {
	fmt.Println(&darkPlayer)
	fmt.Println(&lightPlayer)

	flag.Var(&darkPlayer, "dark", "dark player (user, bot, random)")
	flag.Var(&lightPlayer, "light", "light player (user, bot, random)")
}

func main() {
	flag.Parse()
	fmt.Println(darkPlayer, lightPlayer)

	board := reversi.NewBoard()
	scanner := bufio.NewScanner(os.Stdin)
	uiBuilder := ui.NewUserInterfaceBuilder(board, scanner)

	uiBuilder.DarkPlayer(darkPlayer.String())
	uiBuilder.LightPlayer(lightPlayer.String())
	userInterface := uiBuilder.GetUserInterface()
	userInterface.PlayGame()
}
