package gameOfLife

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestPopulate(t *testing.T) {
	ResetWorld()
  Populate(0, 2)
  assert.Equal(t, Living(0, 2), true)
  assert.Equal(t, Living(0, 1), false)
}

func TestToggle(t *testing.T) {
	ResetWorld()
  Toggle(0, 2)
  assert.Equal(t, Living(0, 2), true)
  assert.Equal(t, Living(0, 1), false)
}

// OOOOO
// OOOOO
// XXOOO
// XOOOO
// OOOOO
// 

func TestNumberOfLivingNeighbours(t *testing.T) {
	ResetWorld()
  Toggle(0, 2)
  Toggle(0, 3)
  Toggle(1, 2)
  assert.Equal(t, NumberOfLivingNeighbours(1, 3), 3)
  assert.Equal(t, NumberOfLivingNeighbours(1, 1), 2)
  assert.Equal(t, NumberOfLivingNeighbours(2, 2), 1)
  assert.Equal(t, NumberOfLivingNeighbours(0, 0), 0)
}

func TestNumberOfLivingNeighboursForLiveCell(t *testing.T) {
	ResetWorld()
  Toggle(0, 2)
  Toggle(0, 3)
  Toggle(1, 2)
  assert.Equal(t, NumberOfLivingNeighbours(0, 2), 2)
}
