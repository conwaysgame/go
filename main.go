package main

import(
  "time"
  "fmt"
  "./gameOfLife"
)

func main() {
  gameOfLife.SetWorldSize(20, 20)
  gameOfLife.ResetWorld()
  gameOfLife.Populate(1, 3)
  gameOfLife.Populate(2, 1)
  gameOfLife.Populate(2, 3)
  gameOfLife.Populate(3, 2)
  gameOfLife.Populate(3, 3)
  fmt.Println("Running a world of 5, 5")
  for {
    fmt.Println("World at:\n")
    fmt.Println(gameOfLife.ToString(), "\n\n")
    gameOfLife.Step()
    time.Sleep(500 * time.Millisecond)
  }
}
