package ga

import (
	"../node"
	"fmt"
	"math"
	"math/rand"
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

func calcDistance(nodes *node.NodeList, chromosome [][]int) float64 {
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
	chromosome := make([][]int, 0)
	size := len(flattench)
	var cut1, cut2 int = 0, 0
	shuffle(flattench)
	for cut1 < size {
		breakFlag := false
		route := make([]int, 0, size)
		for cut2 = cut1; cut2 < size+1; cut2++ {
			route = flattench[cut1:cut2]
			if !nodes.IsFeasible(route) {
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

func copyIndividual(indv *Individual) *Individual {
	newch := make([][]int, len(indv.Chromosome))
	copy(newch, indv.Chromosome)
	// Debug
	fmt.Println(newch)
	distance := indv.Distance
	fitness := indv.Fitness
	newIndv := &Individual{
		Chromosome: newch,
		Distance:   distance,
		Fitness:    fitness,
	}
	return newIndv
}

func containsIndividual(indvList []*Individual, counterpart *Individual) bool {
	for _, indv := range indvList {
		if indv.IsEqual(counterpart) {
			return true
		}
	}
	return false
}

func doesLeftDominateRight(candidate *Individual, counterpart *Individual) int {
	nvehicles1 := candidate.NVehicle()
	nvehicles2 := counterpart.NVehicle()
	distance1 := candidate.Distance
	distance2 := counterpart.Distance

	if nvehicles1 == nvehicles2 && distance1 == distance2 {
		return SAME
	}
	if nvehicles1 <= nvehicles2 && distance1 <= distance2 {
		return LEFT
	}
	if nvehicles1 >= nvehicles2 && distance1 >= distance2 {
		return RIGHT
	}
	return SAME
}

func removeNullRoute(chromosome [][]int) [][]int {
	newChromosome := make([][]int, 0, len(chromosome))
	for i, route := range chromosome {
		if len(route) > 0 {
			newChromosome = append(newChromosome, chromosome[i])
		}
	}
	return newChromosome
}

func makeParetoRankingList(indvList []*Individual) [][]*Individual {
	rankingList := make([][]*Individual, 0)
	for len(indvList) > 0 {
		currentRankList, _ := MakeCurrentRankingList(indvList)
		rankingList = append(rankingList, currentRankList)
	}
	return rankingList
}

//********//
// Public //
//********//
func Flatten(chromosome [][]int) []int {
	flattench := make([]int, 0, 102)
	for _, route := range chromosome {
		for _, node := range route {
			flattench = append(flattench, node)
		}
	}
	return flattench
}

func CreateIndividualList(population int, nodes *node.NodeList) []*Individual {
	indvList := make([]*Individual, 0, population)
	for i := 0; i < population; i++ {
		indv := CreateIndividual(nodes)
		indvList = append(indvList, indv)
	}
	return indvList
}

func SetDistance(nodes *node.NodeList, indvList []*Individual) {
	for i := 0; i < len(indvList); i++ {
		indvList[i].Distance = calcDistance(nodes, indvList[i].Chromosome)
	}
}

func WsumEvaluate(nvehicle int, distance float64,
	wNvehicle float64, wDistance float64) float64 {
	return float64(nvehicle)*wNvehicle + distance*wDistance
}

func MakeCurrentRankingList(currentRankCandidates []*Individual) ([]*Individual, []*Individual) {
	dominatedList := make([]*Individual, 0, len(currentRankCandidates))
	nondominatedList := make([]*Individual, 0, len(currentRankCandidates))

	for i, candidate := range currentRankCandidates {
		isDominated := false
		if containsIndividual(dominatedList, candidate) {
			continue
		}
		for _, counterpart := range currentRankCandidates[i+1:] {
			if containsIndividual(dominatedList, counterpart) {
				continue
			}
			switch doesLeftDominateRight(candidate, counterpart) {
			case SAME:
			case LEFT:
				dominatedList = append(dominatedList, counterpart)
			case RIGHT:
				dominatedList = append(dominatedList, candidate)
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

	return nondominatedList, dominatedList
}
