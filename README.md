Conways's Game of Life in Go
===

This is an implementation of [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway's_Game_of_Life) in Go I created as an exercise in understanding Go and Behaviour-Driven Development in Go. Created by [@basicallydan](https://github.com/basicallydan).

**It isn't finished yet. See the [Issues](https://github.com/conwaysgame/go/issues) to see what I have left to do.**

The Game
---

Run `go run main.go` and you should see a 5x5 world starting with a glider. Each step, the glider will move and eventually die out when it hits the edge of the world. In future versions, the world will be customiseable somehow.

Tests
---

Run `go test ./gameOfLife` and you should see a little `ok` message
