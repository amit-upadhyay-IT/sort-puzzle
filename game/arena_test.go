package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameArena_GetTubes(t *testing.T) {

	tubeOne := GetTube(4, PINK)
	tubeTwo := GetTube(4, GREY, ORANGE, ORANGE)

	ga := GetGameArena([]*Tube{tubeOne, tubeTwo})
	assert.NotNil(t, ga)
	assert.Equal(t, 2, len(ga.GetTubes()))
	assert.Equal(t, PINK, ga.GetTubes()[0].GetTopBlock().GetColor())
	assert.Equal(t, ORANGE, ga.GetTubes()[1].GetTopBlock().GetColor())
}

func TestGameArena_GetTubesCount(t *testing.T) {

	tubeOne := GetTube(4, ORANGE)
	tubeTwo := GetTube(4, ORANGE, ORANGE, ORANGE)
	tubeThree := GetTube(4, ORANGE, ORANGE, ORANGE)

	ga := GetGameArena([]*Tube{tubeOne, tubeTwo, tubeThree})
	assert.NotNil(t, ga)
	assert.Equal(t, 3, ga.GetTubesCount())
}

func TestGameArena_IsArenaSorted(t *testing.T) {
	tubeOne := GetTube(3, )
	tubeTwo := GetTube(3, ORANGE, ORANGE, ORANGE)
	tubeThree := GetTube(3, ORANGE, ORANGE, ORANGE)
	ga := GetGameArena([]*Tube{tubeOne, tubeTwo, tubeThree})
	assert.NotNil(t, ga)
	assert.True(t, ga.IsArenaSorted())


	tubeOne = GetTube(4, )
	tubeTwo = GetTube(4, ORANGE, ORANGE, ORANGE, ORANGE)
	tubeThree = GetTube(4, GREY, GREY, GREY, GREY)
	ga = GetGameArena([]*Tube{tubeTwo, tubeThree, tubeOne})
	assert.NotNil(t, ga)
	assert.True(t, ga.IsArenaSorted())


	tubeOne = GetTube(4, )
	tubeTwo = GetTube(4, ORANGE, ORANGE, ORANGE, ORANGE)
	tubeThree = GetTube(4, GREY, GREY, GREY, PINK)
	ga = GetGameArena([]*Tube{tubeTwo, tubeThree, tubeOne})
	assert.NotNil(t, ga)
	assert.False(t, ga.IsArenaSorted())

	tubeOne = GetTube(4, )
	tubeTwo = GetTube(4, )
	tubeThree = GetTube(4, )
	ga = GetGameArena([]*Tube{tubeTwo, tubeThree, tubeOne})
	assert.NotNil(t, ga)
	assert.True(t, ga.IsArenaSorted())
}

func TestGameArena_Move(t *testing.T) {
	tubeOne := GetTube(4, ORANGE)

	tubeTwo := GetTube(4, ORANGE, ORANGE, ORANGE)

	ga := GetGameArena([]*Tube{tubeOne, tubeTwo})
	ga.Move(0, 1)

	assert.True(t, ga.IsArenaSorted())
}

func TestGameArena_Move2(t *testing.T) {
	t1 := GetTube(4, DARK_BLUE, ORANGE, DARK_BLUE, ORANGE)
	t2 := GetTube(4, ORANGE, DARK_BLUE, ORANGE, DARK_BLUE)
	t3 := GetTube(4)
	ga := GetGameArena([]*Tube{t1, t2, t3})

	ga.Move(0, 2)
	ga.Move(1, 0)
	ga.Move(1, 2)
	ga.Move(0, 1)
	ga.Move(0, 2)
	ga.Move(1, 0)
	ga.Move(1, 2)

	assert.True(t, ga.IsArenaSorted())
}

func TestGameArena_Move3(t *testing.T) {
	t1 := GetTube(4, DARK_BLUE, ORANGE, RED, DARK_BLUE)
	t2 := GetTube(4, ORANGE, ORANGE, RED, DARK_BLUE)
	t3 := GetTube(4, RED, DARK_BLUE, ORANGE, RED)
	t4 := GetTube(4)
	t5 := GetTube(4)
	ga := GetGameArena([]*Tube{t1, t2, t3, t4, t5})

	ga.Move(0, 3)
	ga.Move(2, 0)
	ga.Move(1, 3)
	ga.Move(1, 4)
	ga.Move(0, 4)
	ga.Move(0, 1)
	ga.Move(2, 1)
	ga.Move(0, 3)
	ga.Move(2, 3)
	ga.Move(2, 4)
	assert.True(t, ga.IsArenaSorted())
}


func TestGameArena_Move4(t *testing.T) {
	t1 := GetTube(4, RED, RADIUM, RADIUM, RADIUM)
	t2 := GetTube(4, ORANGE, RED, PINK, RADIUM)
	t3 := GetTube(4, PINK, ORANGE, RED, ORANGE)
	t4 := GetTube(4, DARK_BLUE, PINK, ORANGE, PINK)
	t5 := GetTube(4, DARK_BLUE, DARK_BLUE, DARK_BLUE, RED)
	t6 := GetTube(4)
	t7 := GetTube(4)
	ga := GetGameArena([]*Tube{t1, t2, t3, t4, t5, t6, t7})

	ga.Move(0, 5)
	ga.Move(1, 5)
	ga.Move(3, 1)
	ga.Move(2, 3)
	ga.Move(2, 0)
	ga.Move(4, 0)
	ga.Move(3, 6)
	ga.Move(2, 6)
	ga.Move(3, 2)
	ga.Move(1, 2)
	ga.Move(1, 0)
	ga.Move(1, 6)
	assert.False(t, ga.IsArenaSorted())
	ga.Move(3, 4)
	assert.True(t, ga.IsArenaSorted())
}

