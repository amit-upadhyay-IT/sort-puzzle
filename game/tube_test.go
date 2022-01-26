package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTube(t *testing.T) {

	tube1 := GetTube(4, RED, GREY, DARK_GREEN, LIGHT_BLUE, BROWN)
	assert.Nil(t, tube1)

	tube1 = GetTube(4, RED, GREY, DARK_GREEN, BROWN)
	assert.NotNil(t, tube1)
}

func TestTube_GetCapacity(t *testing.T) {

	tube := GetTube(5, RED)

	assert.NotNil(t, tube)
	assert.Equal(t, 5, tube.GetCapacity())
}

func TestTube_GetEmptyLength(t *testing.T) {

	tube1 := GetTube(4, RED, GREY, DARK_GREEN, LIGHT_BLUE)
	assert.NotNil(t, tube1)
	assert.Equal(t, 0, tube1.GetEmptyLength())

	tube1 = GetTube(4, LIGHT_BLUE, RED, PINK)
	assert.NotNil(t, tube1)
	assert.Equal(t, 1, tube1.GetEmptyLength())

	tube1 = GetTube(4, LIGHT_BLUE, RED)
	assert.NotNil(t, tube1)
	assert.Equal(t, 2, tube1.GetEmptyLength())

	tube1 = GetTube(4, LIGHT_BLUE)
	assert.NotNil(t, tube1)
	assert.Equal(t, 3, tube1.GetEmptyLength())

	tube1 = GetTube(4)
	assert.NotNil(t, tube1)
	assert.Equal(t, 4, tube1.GetEmptyLength())
}

func TestTube_GetLength(t *testing.T) {

	tube1 := GetTube(4, RED, GREY, DARK_GREEN, LIGHT_BLUE)
	assert.NotNil(t, tube1)
	assert.Equal(t, 4, tube1.GetLength())

	tube1 = GetTube(4, LIGHT_BLUE, RED, PINK)
	assert.NotNil(t, tube1)
	assert.Equal(t, 3, tube1.GetLength())

	tube1 = GetTube(4, LIGHT_BLUE, RED)
	assert.NotNil(t, tube1)
	assert.Equal(t, 2, tube1.GetLength())

	tube1 = GetTube(4, LIGHT_BLUE)
	assert.NotNil(t, tube1)
	assert.Equal(t, 1, tube1.GetLength())

	tube1 = GetTube(4)
	assert.NotNil(t, tube1)
	assert.Equal(t, 0, tube1.GetLength())
}

func TestTube_GetTopBlock(t *testing.T) {

	tube1 := GetTube(4, RED, GREY, DARK_GREEN, LIGHT_BLUE)
	assert.NotNil(t, tube1)
	assert.Equal(t, LIGHT_BLUE, tube1.GetTopBlock().GetColor())

	tube1 = GetTube(4, LIGHT_BLUE, RED, PINK)
	assert.NotNil(t, tube1)
	assert.Equal(t, PINK, tube1.GetTopBlock().GetColor())

	blk, isPopSuccess := tube1.Pop()
	assert.True(t, isPopSuccess)
	assert.NotNil(t, blk)
	assert.Equal(t, 1, len(blk))
	assert.Equal(t, PINK, blk[0].GetColor())
	assert.Equal(t, RED, tube1.GetTopBlock().GetColor())
}

func TestTube_GetTopBlocksSet(t *testing.T) {

	tube1 := GetTube(4, RED, GREY, DARK_GREEN, DARK_GREEN)
	assert.NotNil(t, tube1)
	assert.Equal(t, 2, len(tube1.GetTopBlocksSet()))
	assert.Equal(t, DARK_GREEN, tube1.GetTopBlocksSet()[0].GetColor())

	tube1 = GetTube(4, RED, GREY, GREY, GREY)
	assert.NotNil(t, tube1)
	assert.Equal(t, 3, len(tube1.GetTopBlocksSet()))
	assert.Equal(t, GREY, tube1.GetTopBlocksSet()[0].GetColor())

	tube1 = GetTube(4, RED, GREY, GREY, ORANGE)
	assert.NotNil(t, tube1)
	assert.Equal(t, 1, len(tube1.GetTopBlocksSet()))
	assert.Equal(t, ORANGE, tube1.GetTopBlocksSet()[0].GetColor())
}

func TestTube_IsEmpty(t *testing.T) {
	tube1 := GetTube(4, RED, GREY, DARK_GREEN, DARK_GREEN)
	assert.NotNil(t, tube1)
	assert.False(t, tube1.IsEmpty())

	tube1 = GetTube(5)
	assert.NotNil(t, tube1)
	assert.True(t, tube1.IsEmpty())
}

func TestTube_IsFull(t *testing.T) {
	tube1 := GetTube(4, RED, GREY, DARK_GREEN, DARK_GREEN)
	assert.NotNil(t, tube1)
	assert.True(t, tube1.IsFull())

	tube1 = GetTube(4, GREY)
	assert.NotNil(t, tube1)
	assert.False(t, tube1.IsFull())
}

func TestTube_IsTubeSorted(t *testing.T) {
	tube1 := GetTube(4, GREY, GREY, GREY, GREY)
	assert.NotNil(t, tube1)
	assert.True(t, tube1.IsTubeSorted())

	tube1 = GetTube(4, )
	assert.NotNil(t, tube1)
	assert.True(t, tube1.IsTubeSorted())

	tube1 = GetTube(3, PINK, PINK, PINK)
	assert.NotNil(t, tube1)
	assert.True(t, tube1.IsTubeSorted())

	tube1 = GetTube(3, PINK, PINK)
	assert.NotNil(t, tube1)
	assert.False(t, tube1.IsTubeSorted())
}

func TestTube_Pop(t *testing.T) {

	tube1 := GetTube(4, GREY, GREY, GREY, GREY)
	assert.NotNil(t, tube1)
	blks, isOpSuccess := tube1.Pop()
	assert.True(t, isOpSuccess)
	assert.NotNil(t, blks)
	assert.Equal(t, 4, len(blks))
	assert.Equal(t, GREY, blks[0].GetColor())

	tube1 = GetTube(4, GREY, DARK_GREEN, DARK_GREEN, RED)
	assert.NotNil(t, tube1)
	blks, isOpSuccess = tube1.Pop()
	assert.True(t, isOpSuccess)
	assert.NotNil(t, blks)
	assert.Equal(t, 1, len(blks))
	assert.Equal(t, RED, blks[0].GetColor())
	blks, isOpSuccess = tube1.Pop()
	assert.True(t, isOpSuccess)
	assert.NotNil(t, blks)
	assert.Equal(t, 2, len(blks))
	assert.Equal(t, DARK_GREEN, blks[0].GetColor())

	tube1 = GetTube(4, GREY, DARK_GREEN, LIGHT_GREEN, RED)
	assert.NotNil(t, tube1)
	blks, isOpSuccess = tube1.Pop()
	assert.True(t, isOpSuccess)
	assert.NotNil(t, blks)
	assert.Equal(t, 1, len(blks))
	assert.Equal(t, RED, blks[0].GetColor())
	blks, isOpSuccess = tube1.Pop()
	assert.True(t, isOpSuccess)
	assert.NotNil(t, blks)
	assert.Equal(t, 1, len(blks))
	assert.Equal(t, LIGHT_GREEN, blks[0].GetColor())

	tube1 = GetTube(4)
	assert.NotNil(t, tube1)
	blks, isOpSuccess = tube1.Pop()
	assert.False(t, isOpSuccess)
	assert.Equal(t, 0, len(blks))
}

func TestTube_Push(t *testing.T) {
	tube1 := GetTube(4, GREY, DARK_GREEN, DARK_GREEN, RED)
	assert.NotNil(t, tube1)
	isOpSuccess := tube1.Push([]*Block{GetBlock(PINK)})
	assert.False(t, isOpSuccess)

	tube1 = GetTube(4, GREY, DARK_GREEN)
	assert.NotNil(t, tube1)
	isOpSuccess = tube1.Push([]*Block{GetBlock(DARK_GREEN)})
	assert.True(t, isOpSuccess)
	assert.Equal(t, 3, tube1.GetLength())
	isOpSuccess = tube1.Push([]*Block{GetBlock(RED)})
	assert.True(t, isOpSuccess)
	assert.Equal(t, 4, tube1.GetLength())
}
