package gameOfLife

import (
  "math"
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

func Step() {
  for x := 0; x < 5; x++ {
    for y := 0; y < 5; y++ {
      var neighbours int = NumberOfLivingNeighbours(x, y)
      if (Living(x, y)) {
        if (neighbours < 2) {         // Under-population/loneliness :()
          Toggle(x, y)
        } else if (neighbours > 3) {  // Over-population
          Toggle(x, y)
        }
        // If 2 or 3, do nowt!
      } else { // It's dead at the mo
        if (neighbours == 3) {
          Populate(x, y)
        }
      }
    }
  }
}

func ToString() string {
  var worldString string = ""
  for y := 0; y < 5; y++ {
    for x := 0; x < 5; x++ {
      if (world[x][y]) {
        worldString += "X"
      } else {
        worldString += "O"
      }
    }
    if (y != 4) {
      worldString += "\n"
    }
  }
  return worldString
}

func Living(x int, y int) bool {
  return world[x][y]
}
