package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlock_GetColor(t *testing.T) {

	blk := GetBlock(BROWN)
	assert.NotNil(t, blk)
	assert.Equal(t, BROWN, blk.GetColor())
}
