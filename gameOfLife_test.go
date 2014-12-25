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
