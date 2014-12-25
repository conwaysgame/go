package main

import(
  "time"
  "fmt"
  "./gameOfLife"
)

func main() {
  gameOfLife.ResetWorld()
  gameOfLife.Populate(0, 2)
  gameOfLife.Populate(1, 0)
  gameOfLife.Populate(1, 2)
  gameOfLife.Populate(2, 1)
  gameOfLife.Populate(2, 2)
  fmt.Println("Running a world of 5, 5")
  for {
    fmt.Println("World at:\n")
    fmt.Println(gameOfLife.ToString(), "\n\n")
    gameOfLife.Step()
    time.Sleep(500 * time.Millisecond)
  }
}
