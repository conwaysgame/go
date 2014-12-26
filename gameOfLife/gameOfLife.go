package gameOfLife

import (
  "math"
)

var world [][]bool
var width, height int

func SetWorldSize(w int, h int) {
  width = w
  height = h
  world = make([][]bool, height)
  for y := 0; y < height; y++ {
    world[y] = make([]bool, width)
  }
}

func ResetWorld() {
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      world[y][x] = false
    }
  }
}

func Populate(x int, y int) {
  world[y][x] = true
}

func Toggle(x int, y int) {
  world[y][x] = !world[y][x]
}

func NumberOfLivingNeighbours(x int, y int) int {
  var neighbours int = 0
  for nx := int(math.Max(float64(x - 1), 0)); nx < int(math.Min(float64(x + 2), float64(width))); nx++ {
    for ny := int(math.Max(float64(y - 1), 0)); ny < int(math.Min(float64(y + 2), float64(height))); ny++ {
      if ((nx != x || ny != y) && Living(nx, ny)) {
        neighbours++
      }
    }
  }
  return neighbours
}

func Step() {
  var newWorld [][]bool = make([][]bool, height)
  for y := 0; y < height; y++ {
    newWorld[y] = make([]bool, width)
  }
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      var neighbours int = NumberOfLivingNeighbours(x, y)
      if (Living(x, y)) {
        if (neighbours < 2) {         // Under-population/loneliness :()
          newWorld[y][x] = false
        } else if (neighbours > 3) {  // Over-population
          newWorld[y][x] = false
        } else {                      // Keep it the same as it was before
          newWorld[y][x] = true
        }
        // If 2 or 3, do nowt!
      } else { // It's dead at the mo
        if (neighbours == 3) {
          newWorld[y][x] = true
        }
      }
    }
  }
  world = newWorld
}

func ToString() string {
  var worldString string = ""
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      if (world[y][x]) {
        worldString += "◼"
      } else {
        worldString += "◻"
      }
    }
    if (y != height - 1) {
      worldString += "\n"
    }
  }
  return worldString
}

func Living(x int, y int) bool {
  return world[y][x] == true
}
