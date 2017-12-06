package ga

import (
	"fmt"
	"os"
)

const (
	SAME = iota
	LEFT
	RIGHT
)

//*********//
// Private //
//*********//
func contains(indvList []*Individual, id int) {
	for _, indv := range indvList {
		if indv.ID() == id {
			return true
		}
	}
	return false
}

func doesLeftDominateRight(candidata *Individual, counterpart *Individual) int {
	nvehicles1 = candidate.NVehicle()
	nvehicles2 = counterpart.NVehicle()
	distance1 = candidate.Distance
	distance2 = counterpart.Distance

	if nvehicles1 == numofvehicle2 && distance1 == distance2 {
		return SAME
	}
	if nvehicles1 <= numofvehicle2 && distance1 <= distance2 {
		return LEFT
	}
	if nvehicles1 >= numofvehicle2 && distance1 >= distance2 {
		return RIGHT
	}
	return SAME
}

func makeParetoRankingList(indvList []*Individual) [][]*Individual {
	rankingList := make([][]*Individual)
	for len(indvList) > 0 {
		currentRankList, indvList = MakeCurrentRankingList(indvList)
		rankingList = append(rankingList, currentRankList)
	}
	return rankingList
}

//********//
// Public //
//********//
func MakeCurrentRankingList(currentRankCandidates []*Individual) ([]*Individual, []*Individual) {
	dominatedList := make([]*Individual, 0, len(currentRankCandidates))
	nondominatedList := make([]*Individual, 0, len(currentRankCandidates))

	for i, candidate := range currentRankCandidates {
		id := candidate.ID()
		isDominated := false
		if contains(dominatedList, id) {
			continue
		}
		for _, counterpart := range current_rank_candidates[i+1:] {
			id = counterpart.ID()
			if contains(dominatedList, id) {
				continue
			}
			switch doesLeftDominateRight(candidate, counterpart) {
			case SAME:
			case LEFT:
				dominatedList = append(dominatedList, counterpart)
			case RIGHT:
				dominatedList = append(dominated_list, candidate)
				isDominated = true
				break
			default:
				fmt.Println("ga/functions/MakeParetoRankingList")
				os.Exit(0)
			}
		}
		if !isDominated {
			nondominatedList = append(nondominatedList, candidate)
		}
	}

	return nondominatedList, doesLeftDominateRight
}
