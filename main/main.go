package main

import (
	"SortPuzzleSolver/game"
	"fmt"
)

var (
	tubesCount          = 14
	tubeLength          = 4
	totalBlocks         = (tubesCount - 2) * tubeLength
	blockVarietiesCount = 12
)

type Tuple struct {
	fromIndex, toIndex int
}

func main() {
	t2()
}

func t1() {
	t1 := game.GetTube(4, game.DARK_BLUE, game.ORANGE, game.DARK_BLUE, game.ORANGE)
	t2 := game.GetTube(4, game.ORANGE, game.DARK_BLUE, game.ORANGE, game.DARK_BLUE)
	t3 := game.GetTube(4)
	ga := game.GetGameArena([]*game.Tube{t1, t2, t3})

	steps := make([]Tuple, 0)
	res := make([][]Tuple, 0)
	foundARes := false
	solveSortPuzzle(ga, steps, &res, &foundARes)
	if foundARes {
		fmt.Println(res[0])
	}
}

func t2() {
	t1 := game.GetTube(4, game.RADIUM, game.BROWN, game.LIGHT_GREEN, game.PURPLE)
	t2 := game.GetTube(4, game.PURPLE, game.DARK_GREEN, game.DARK_GREEN, game.YELLOW)
	t3 := game.GetTube(4, game.PINK, game.RADIUM, game.BROWN, game.RED)
	t4 := game.GetTube(4, game.GREY, game.DARK_BLUE, game.RED, game.LIGHT_BLUE)
	t5 := game.GetTube(4, game.ORANGE, game.PURPLE, game.RADIUM, game.GREY)
	t6 := game.GetTube(4, game.YELLOW, game.DARK_GREEN, game.LIGHT_BLUE, game.LIGHT_BLUE)
	t7 := game.GetTube(4, game.RED, game.DARK_BLUE, game.PURPLE, game.ORANGE)
	t8 := game.GetTube(4, game.RED, game.YELLOW, game.DARK_BLUE, game.LIGHT_GREEN)
	t9 := game.GetTube(4, game.PINK, game.DARK_BLUE, game.DARK_GREEN, game.PINK)
	t10 := game.GetTube(4, game.PINK, game.BROWN, game.LIGHT_BLUE, game.RADIUM)
	t11 := game.GetTube(4, game.LIGHT_GREEN, game.BROWN, game.ORANGE, game.GREY)
	t12 := game.GetTube(4, game.LIGHT_GREEN, game.GREY, game.ORANGE, game.YELLOW)
	t13 := game.GetTube(4)
	t14 := game.GetTube(4)
	t15 := game.GetTube(4)
	ga := game.GetGameArena([]*game.Tube{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15})

	steps := make([]Tuple, 0)
	res := make([][]Tuple, 0)
	foundARes := false
	solveSortPuzzle(ga, steps, &res, &foundARes)
	if foundARes {
		fmt.Println(res[0])
	} else {
		fmt.Println("no solution found")
	}
}

func t3() {
	t1 := game.GetTube(4, game.RADIUM, game.ORANGE, game.RADIUM, game.DARK_BLUE)
	t2 := game.GetTube(4, game.ORANGE, game.PINK, game.PINK, game.ORANGE)
	t3 := game.GetTube(4, game.PINK, game.RED, game.DARK_BLUE, game.RED)
	t4 := game.GetTube(4, game.DARK_BLUE, game.RED, game.RADIUM, game.PINK)
	t5 := game.GetTube(4, game.DARK_BLUE, game.RADIUM, game.RED, game.ORANGE)
	t6 := game.GetTube(4)
	t7 := game.GetTube(4)
	ga := game.GetGameArena([]*game.Tube{t1, t2, t3, t4, t5, t6, t7})

	steps := make([]Tuple, 0)
	res := make([][]Tuple, 0)
	foundARes := false
	solveSortPuzzle(ga, steps, &res, &foundARes)
	if foundARes {
		fmt.Println(res[0])
	}
}


func solveSortPuzzle(ga *game.GameArena, steps []Tuple, res *[][]Tuple, foundAResult *bool) {
	// base case: if game arena is all sorted
	if ga.IsArenaSorted() {
		// log the result
		*res = append(*res, steps)
		*foundAResult = true
		return
	}

	// try out every possible placements
	for i, _ := range ga.GetTubes() {
		for j, _ := range ga.GetTubes() {
			if i != j && !*foundAResult {
				isSuccess, _ := ga.Move(i, j)
				if isSuccess {
					steps = append(steps, Tuple{fromIndex: i, toIndex: j})

					solveSortPuzzle(ga, steps, res, foundAResult)

					steps = steps[:len(steps)-1]
					ga.Undo()
				}
			}
		}
	}
}
