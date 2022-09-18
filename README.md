# Grey

[![Go Reference](https://pkg.go.dev/badge/github.com/elehtine/grey.svg)](https://pkg.go.dev/github.com/elehtine/grey)
[![build and test](https://github.com/elehtine/grey/actions/workflows/go.yml/badge.svg)](https://github.com/elehtine/grey/actions/workflows/go.yml)

Command line Reversi game with Golang.


## Usage

Clone project with git.

`go` commands:
- Install compile and install project: `go install`
- Run project: `go run main.go`
- Run tests: `go run test ./reversi -v`
- Read documentation with `go doc`

### Example

```
$ go run main.go -dark=bot -light=user
user
bot
bot user
  abcdefgh
 +--------+
1|........|1
2|........|2
3|........|3
4|...ox...|4
5|...xo...|5
6|........|6
7|........|7
8|........|8
 +--------+
  abcdefgh
  abcdefgh
 +--------+
1|........|1
2|........|2
3|........|3
4|..xxx...|4
5|...xo...|5
6|........|6
7|........|7
8|........|8
 +--------+
  abcdefgh
Give move: c5
  abcdefgh
 +--------+
1|........|1
2|........|2
3|........|3
4|..xxx...|4
5|..ooo...|5
6|........|6
7|........|7
8|........|8
 +--------+
  abcdefgh
  abcdefgh
 +--------+
1|........|1
2|........|2
3|........|3
4|..xxx...|4
5|..xoo...|5
6|.x......|6
7|........|7
8|........|8
 +--------+
  abcdefgh
Give move: 
...
```

## Structure

Folder `reversi` contains game which is independent of user interface. Folder `ui` contains user interface for game. Game can be played via command line input. Bots can also be chosen to play a game.
