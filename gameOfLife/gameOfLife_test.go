package gameOfLife

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "fmt"
)

func TestSetWorldSize(t *testing.T) {
  SetWorldSize(10, 10)
  ResetWorld()
  Populate(9, 9)
  assert.Equal(t, Living(9, 9), true)
  SetWorldSize(5, 5)
}

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
  assert.Equal(t, T◻String(), "◻◻◻◻◻\n◻◻◻◻◻\n◼◼◻◻◻\n◼◻◻◻◻\n◻◻◻◻◻")
}

func TestGlider(t *testing.T) {
  SetWorldSize(6, 6)
  ResetWorld()
  // This is a glider in it's initial form
  Populate(2, 1)
  Populate(3, 2)
  Populate(1, 3)
  Populate(2, 3)
  Populate(3, 3)
  Step()
  assert.Equal(t, Living(1, 2), true)
  assert.Equal(t, Living(3, 2), true)
  assert.Equal(t, Living(2, 3), true)
  assert.Equal(t, Living(3, 3), true)
  assert.Equal(t, Living(2, 4), true)
}
