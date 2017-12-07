package ga

import (
	"../node"
	"fmt"
	"math"
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
func calcDistanceBetweenNodes(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func calcOneVehicleDistance(nodes *node.NodeList, route []int) float64 {
	depot := nodes.Depot()
	x_depot, y_depot := depot.Position()
	var totalDistance float64 = 0.0

	x, y := nodes.Position(route[0])
	totalDistance += calcDistanceBetweenNodes(x_depot, y_depot, x, y)
	for i := 0; i < len(route)-1; i++ {
		x1, y1 := nodes.Position(route[i])
		x2, y2 := nodes.Position(route[i+1])
		totalDistance += calcDistanceBetweenNodes(x1, y1, x2, y2)
	}
	x, y = nodes.Position(route[len(route)-1])
	totalDistance += calcDistanceBetweenNodes(x, y, x_depot, y_depot)

	return totalDistance
}

func clacDistance(nodes, chromosome [][]int) float64 {
	var totalDistance float64 = 0.0
	for _, route := range chromosome {
		totalDistance += calcOneVehicleDistance(nodes, route)
	}
	return totalDistance
}

func shuffle(data []int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func shapeFlatToVehicles(nodes *node.NodeList, flattench []int) [][]int {
	chromosome := make([][]int)
	size := len(flattench)
	var cut1, cut2 int = 0, 0
	shuffle(flattench)
	for cut1 < size {
		breakFlag := false
		for cut2 = cut1; cut2 < size+1; cut2++ {
			route := flattench[cut1:cut2]
			if !nodes.is_feasible(route) {
				cut1 = cut2 - 1
				route = route[:len(route)-1]
				breakFlag = true
				break
			}
		}
		if breakFlag {
			cut1 = cut2
		}
		chromosome = append(chromosome, route)
	}
	return chromosome
}

func copyIndividual(indv Individual) *Individual {
	newch := make([][]int, len(indv))
	copy(newch, indv.Chromosome)
	// Debug
	fmt.Println(newch)
	distance = indv.Distance
	fitness = indv.Fitness
	newIndv = &Individual{
		Chromosome: newch,
		Distance:   distance,
		Fitness:    fitness,
	}
	return newIndv
}

func containsIndividual(indvList []*Individual, id int) {
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
func CreateIndividualList(population int, nodes *node.NodeList) []*Individual {
	indvList := make([]*Individual, 0, population)
	for i := 0; i < population; i++ {
		indv = CreateIndividualList(nodes)
		indvList = append(indvList, indv)
	}
	return indvList
}

func MakeCurrentRankingList(currentRankCandidates []*Individual) ([]*Individual, []*Individual) {
	dominatedList := make([]*Individual, 0, len(currentRankCandidates))
	nondominatedList := make([]*Individual, 0, len(currentRankCandidates))

	for i, candidate := range currentRankCandidates {
		id := candidate.ID()
		isDominated := false
		if containsIndividual(dominatedList, id) {
			continue
		}
		for _, counterpart := range current_rank_candidates[i+1:] {
			id = counterpart.ID()
			if containsIndividual(dominatedList, id) {
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
