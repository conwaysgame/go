package gameOfLife

import (
  "math"
  // "fmt"
)

var world [5][5]bool

func ResetWorld() {
  for x := 0; x < 5; x++ {
    for y := 0; y < 5; y++ {
      world[x][y] = false
    }
  }
}

func Populate(x int, y int) {
  world[x][y] = true
}

func Toggle(x int, y int) {
  world[x][y] = !world[x][y]
}

func NumberOfLivingNeighbours(x int, y int) int {
  var neighbours int = 0
  // fmt.Printf("For %d, %d\n", x, y)
  for nx := int(math.Max(float64(x - 1), 0)); nx < int(math.Min(float64(x + 2), 5)); nx++ {
    for ny := int(math.Max(float64(y - 1), 0)); ny < int(math.Min(float64(y + 2), 5)); ny++ {
      // fmt.Printf("Checking %d, %d...", nx, ny)
      if ((nx != x || ny != y) && Living(nx, ny)) {
        // fmt.Printf("Yup!")
        neighbours++
      }
      // fmt.Printf("\n")
    }
  }
  return neighbours
}

func Living(x int, y int) bool {
  return world[x][y]
}

