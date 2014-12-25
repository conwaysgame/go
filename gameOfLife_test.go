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

// OOOOO
// OOOOO
// XXOOO
// XOOOO
// OOOOO
// Should go to
// OOOOO
// OOOOO
// XXOOO
// XXOOO
// OOOOO

func TestStep(t *testing.T) {
	ResetWorld()
  Populate(0, 2)
  Populate(0, 3)
  Populate(1, 2)
  assert.Equal(t, Living(1, 3), false)
  assert.Equal(t, Living(0, 0), false)
  assert.Equal(t, Living(1, 0), false)
  assert.Equal(t, Living(0, 1), false)
  assert.Equal(t, Living(1, 1), false)
  Step()
  assert.Equal(t, Living(0, 2), true)
  assert.Equal(t, Living(0, 3), true)
  assert.Equal(t, Living(1, 2), true)
  assert.Equal(t, Living(1, 3), true)
  assert.Equal(t, Living(0, 0), false)
  assert.Equal(t, Living(1, 0), false)
  assert.Equal(t, Living(0, 1), false)
  assert.Equal(t, Living(1, 1), false)
}

func TestToString(t *testing.T) {
	ResetWorld()
  Populate(0, 2)
  Populate(0, 3)
  Populate(1, 2)
  assert.Equal(t, ToString(), "OOOOO\nOOOOO\nXXOOO\nXOOOO\nOOOOO")
}
