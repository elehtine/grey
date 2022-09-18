# Grey

Command line Reversi game with Golang

## Structure

Folder `reversi` contains game which is independent of user interface. Folder `ui` contains user interface for game. Game can be played via command line input or use bots to play a game.



## Usage

Clone project with git.

- Run project: `go run .`
- Run tests: `go run test ./...`

## Example

```
$ go run .
Choose dark player (user, bot) user
Choose light player (user, bot) user
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
Give move: c4
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
Give move: 
...
```
