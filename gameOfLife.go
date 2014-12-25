package gameOfLife

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

func Living (x int, y int) bool {
  return world[x][y]
}

