package game

import (
	"SortPuzzleSolver/utils/errors"
	"fmt"
)

type GameArena struct {
	tubes []*Tube

	gameHistory *GameHistory
}

func GetGameArena(tubeList []*Tube) *GameArena {
	return &GameArena{
		tubes: tubeList,
		gameHistory: GetGameHistory(),
	}
}

// DeepCopyGameArena will create and return a new copy of arena passed in the argument
// this is not a good idea to do, but atm I do not have solution to maintain state in
// recursive method, so writing a custom method for this, I hope this is better than
// using json serialization and deserialization. TODO: do some benchmarking bro xD
func (ga *GameArena) DeepCopyGameArena() *GameArena {

	arenaTubes := make([]*Tube, 0)
	for _, tub := range ga.GetTubes() {
		clrs := make([]Color, 0)
		for _, blk := range tub.blocks {
			clrs = append(clrs, blk.GetColor())
		}
		arenaTubes = append(arenaTubes, GetTube(tub.GetCapacity(), clrs...))
	}

	return &GameArena{
		tubes: arenaTubes,
		gameHistory: ga.gameHistory,
	}
}

func (ga *GameArena) GetTubes() []*Tube {
	return ga.tubes
}

func (ga *GameArena) GetTubesCount() int {
	return len(ga.tubes)
}

// Move takes two tube numbers, this tries to move the top block from tube
// number one to tube number two, if the movement is successful it returns
// true else false
// NOTE: the indices passed here should start from 0
func (ga *GameArena) Move(fromTubeIdx, toTubeIdx int) (isMoveSuccess bool, vErr error) {

	// validation
	// tube index passed should not exceed total count of tubes present in game
	if fromTubeIdx >= ga.GetTubesCount() || toTubeIdx >= ga.GetTubesCount() {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"tube index exceeds total tube count, "+
			"fromTubeIdx: %d and toTubeIdx: %d", toTubeIdx, toTubeIdx))
		return
	}

	// fromTube should at-least have some blocks present in it
	if ga.tubes[fromTubeIdx].IsEmpty() {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"from tube should have atleast one blocks present,"+
			" fromTubeIdx: %d and toTubeIdx: %d", toTubeIdx, toTubeIdx))
		return
	}

	// toTube should have some empty space present in it
	if ga.tubes[toTubeIdx].IsFull() {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"to tube should have atleast one empty space,"+
			" fromTubeIdx: %d and toTubeIdx: %d", toTubeIdx, toTubeIdx))
		return
	}

	// top color of both tubes should match
	if !ga.tubes[fromTubeIdx].IsEmpty() &&
		!ga.tubes[toTubeIdx].IsEmpty() &&
		(ga.tubes[fromTubeIdx].GetTopBlock().GetColor() != ga.tubes[toTubeIdx].GetTopBlock().GetColor()) {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"the top of two tubes passed doesn't have same color blocks on top of them"+
			" fromTubeIdx: %d and toTubeIdx: %d", toTubeIdx, toTubeIdx))
		return
	}

	// we should only pour lesser or equal number of blocks from the source tube to destination tube
	if len(ga.tubes[fromTubeIdx].GetTopBlocksSet()) > ga.tubes[toTubeIdx].GetEmptyLength() {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"can not transfer %d blocks from tube number %d to tube number %d"+
			" \n", len(ga.tubes[fromTubeIdx].GetTopBlocksSet()), fromTubeIdx, toTubeIdx))
		return
	}

	// pouring whole of a tube to another tube also does not make sense, add validation for that too
	if ga.tubes[fromTubeIdx].IsTubeSorted() {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"should not pour a sorted tube, fromtube index: %d and to tube index %d"+
			" \n", fromTubeIdx, toTubeIdx))
		return
	}

	// from tube has same color blocks and destination tube is empty, avoid this scenario
	// because moving would not make any sense as game state would be same
	if ga.tubes[fromTubeIdx].HasSameColor() && ga.tubes[toTubeIdx].IsEmpty() {
		vErr = errors.Errorf(fmt.Sprintf("[wrong arguments passed] "+
			"there is no point pouring same colored tube to other empty tube" +
			", fromtube index: %d and to tube index %d"+
			" \n", fromTubeIdx, toTubeIdx))
		return
	}

	isMoveSuccess = true

	// now that validations are successful, we can transfer blocks
	// todo: receive err from Push and Pop in order to log them
	poppedBlocks, isPopSuccess := ga.tubes[fromTubeIdx].Pop()
	if isPopSuccess {
		if ga.tubes[toTubeIdx].Push(poppedBlocks) {
			err := ga.gameHistory.AddSuccessfulMove(ga, SortPuzzle)
			if err != nil {
				fmt.Printf("error: ", err.Error())
				return false, err
			}
			return
		} else {
			// this is not suppose to execute according to logic I understand
			// I will log this in case it happens
			fmt.Printf("something weird happened, please think why")
			// inserting what has been removed
			ga.tubes[fromTubeIdx].Push(poppedBlocks)
		}
	}

	return false, vErr
}


func (ga *GameArena) IsArenaSorted() bool {

	for _, tub := range ga.tubes {
		if !tub.IsTubeSorted() {
			return false
		}
	}

	return true
}

func (ga *GameArena) Undo() bool {
	newGameState := ga.gameHistory.UndoAMove()
	if newGameState != nil {
		*ga = *newGameState
		return true
	} else {
		return false
	}
}
