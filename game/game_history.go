package game

import "fmt"

type GameHistory struct {
	stateBlueprints []*GameArena
	movesCounter    int
}

func GetGameHistory() *GameHistory {
	return &GameHistory{
		stateBlueprints: make([]*GameArena, 100000),
		movesCounter:    -1,
	}
}

func (gh *GameHistory) AddSuccessfulMove(gameState interface{}, gameType GameType) error {
	//if cap(gh.stateBlueprints) < gh.movesCounter {
	//	// increase size of stateBlueprints
	//	temp := gh.stateBlueprints
	//	gh.stateBlueprints = make([]*GameArena, 2*gh.movesCounter)
	//	for i, item := range temp {
	//		gh.stateBlueprints[i] = item
	//	}
	//}

	switch gameType {
	case SortPuzzle:
		newGS := gameState.(*GameArena).DeepCopyGameArena()
		gh.movesCounter++
		if gh.movesCounter >= 9999 {
			fmt.Print("wtf")
		}
		gh.stateBlueprints[gh.movesCounter] = newGS
	}
	return nil
}

// todo: return error as well
func (gh *GameHistory) UndoAMove() *GameArena {
	if gh.movesCounter <= 0 || len(gh.stateBlueprints) <= 0 || gh.movesCounter >= len(gh.stateBlueprints) {
		return nil
	}
	gh.movesCounter--
	return gh.stateBlueprints[gh.movesCounter]
}

type GameType string

const (
	SortPuzzle GameType = "sort_puzzle"
)
