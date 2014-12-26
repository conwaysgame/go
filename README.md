# Conways's Game of Life in Go

This is an implementation of [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway's_Game_of_Life) in Go I created as an exercise in understanding Go and Behaviour-Driven Development in Go. Created by [@basicallydan](https://github.com/basicallydan).

## The Game

Run `go run main.go` and you should see a 20x20 world starting with a glider. Each step, the glider will move and eventually die out when it hits the edge of the world. If you want to change the size of the world or starting population, please edit `main.go`.

## Tests

Run `go test ./gameOfLife` and you should see a little `ok` message

## Contributing

If you'd like to contribute, that's great and I encourage it - I do, however, also encourage reading the [contributing doc](https://github.com/conwaysgame/go/blob/master/contributing.md) first. The golden rule for contributing is that **you ensure that the tests are still passing before you make a pull request.**