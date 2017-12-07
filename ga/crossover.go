package ga

import (
	"../node"
	"../ut"
	"fmt"
	"math/rand"
)

//*********//
// Private //
//*********//
func containsNode(flattench []int, node int) bool {
	for _, n := range flattench {
		if n == node {
			return true
		}
	}
	return false
}

func findIndex(flattench []int, node int) int {
	for i, n := range flattench {
		if n == node {
			return i
		}
	}
	return -1
}

// ===== <Uniform Order Crossover (UOX)> ===== //
func uniformOrderCrossover(nodes *node.NodeList, ch1, ch2 [][]int) ([][]int, [][]int) {
	flattench1 := ut.flatten(ch1)
	flattench2 := ut.flatten(ch2)
	size := len(flattench1)
	mask := make([]int, size)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		mask[i] = rand.Intn(2)
	}
	tmpch1 := make([]int, size)
	tmpch2 := make([]int, size)
	for i := 0; i < size; i++ {
		tmpch1[i] = flattench1[i] * mask[i]
	}
	for i := 0; i < size; i++ {
		tmpch2[i] = flattench2[i] * mask[i]
	}

	for _, node := range flattench2 {
		if !containsNode(tmpch1, node) {
			insertIndex := findIndex(tmpch1, 0)
			tmpch1[insertIndex] = node
		}
	}
	for _, node := range flattench1 {
		if !containsNode(tmpch2, node) {
			insertIndex := findIndex(tmpch2, 0)
			tmpch2[insertIndex] = node
		}
	}

	tmp1 := shapeFlatToVehicles(nodes, tmp1)
	tmp2 := shapeFlatToVehicles(nodes, tmp2)
	return tmp1, tmp2
}

// ===== </Uniform Order Crossover (UOX)> ===== //

// ===== <Partially Mapped Crossover (PMX)> ===== //
func getNoConflictList(origin []int, counterpart []int) []int {
	tmp := make([]int, len(oringin))
	copy(tmp, oringin)

	for i, node := range oringin {
		if containsNode(counterpart, node) {
			tmp[i] = 0
		}
	}
	return tmp
}

func partiallyMappedCrossover(nodes *node.NodeList,
	ch1, ch2 [][]int) ([][]int, [][]int) {
	flattench1 := ut.flatten(ch1)
	flattench2 := ut.flatten(ch2)
	size := len(flattench1)
	rand.Seed(time.Now().UnixNano())
	point1 := rand.Intn(size - 1)
	point2 := rand.Intn(size)
	for point1 == point2 {
		point2 = rand.Intn(size)
	}
	if point1 > point2 {
		point1, point2 = point2, point1
	}

	tmpch1 := make([]int, size)
	tmp := flattench2[point1:point2]
	pre := getNoConflictList(flattench1[:point1], tmp)
	suf := getNoConflictList(flattench1[point2:], tmp)
	tmpch1 = append(pre, tmp...)
	tmpch1 = append(tmpch1, suf...)

	tmpch2 := make([]int, size)
	tmp := flattench1[point1:point2]
	pre := getNoConflictList(flattench2[:point1], tmp)
	suf := getNoConflictList(flattench2[point2:], tmp)
	tmpch2 = append(pre, tmp...)
	tmpch2 = append(tmpch1, suf...)

	for _, node := range flattench2 {
		if !containsNode(tmpch1, node) {
			insertIndex := findIndex(tmpch1, 0)
			tmpch1[insertIndex] = node
		}
	}
	for _, node := range flattench1 {
		if !containsNode(tmpch2, node) {
			insertIndex := findIndex(tmpch2, 0)
			tmpch2[insertIndex] = node
		}
	}

	tmp1 := shapeFlatToVehicles(nodes, tmp1)
	tmp2 := shapeFlatToVehicles(nodes, tmp2)
	return tmp1, tmp2
}

// ===== </Partially Mapped Crossover (PMX)> ===== //

// ===== <Best Cost Route Crossover (BCRC)> ===== //
func insertNode(nodes *node.NodeList,
	chromosome [][]int, L []int) [][]int {
	size := len(chromosome)
	newChromosome := make([][]int, 0, len(chromosome))
	// Copy Chromosome
	for i := 0; i < size; i++ {
		tmp := make([]int, len(chromosome[i]))
		copy(tmp, chromosome[i])
		newChromosome = append(newChromosome, tmp)
	}

	for n, insertNode := range L {
		feasibleListI := make([]int, 0)
		feasibleListJ := make([]int, 0)
		for i, route := range chromosome {
			for j := 0; j < len(route)+1; j++ {
				tmp := append(route[:j+1], route[j:]...)
				tmp[j] = insertNode
				// Debug
				fmt.Println(tmp)
				if nodes.IsFeasible(tmp) {
					feasibleListI = append(feasibleListI, i)
					feasibleListJ = append(feasibleListJ, j)
				}
			}
			if len(feasibleListI) == 0 {
				newRoute := make([]int, 1)
				newRoute[0] = insertNode
				newChromosome = append(newChromosome, newRoute)
			} else {
				rand.Seed(time.Now().UnixNano())
				index := rand.Intn(len(feasibleListI))
				indexI := feasibleListI[index]
				indexJ := feasibleListJ[index]
				tmp := make([]int, len(newChromosome[indexI]))
				copy(tmp, newChromosome[indexI])
				tmp := append(tmp[:indexJ+1], tmp[indexJ:]...)
				tmp[indexJ] = insertNode
				newChromosome[indexI] = tmp
			}
		}
	}
	return newChromosome
}

func deleteNodes(chromosome, route) [][]int {
	chromosomeDeleted := make([][]int, 0, len(chromosome))
	for i, rt := range chromosome {
		routeDeleated := make([]int, 0)
		for _, node := range rt {
			if !containsNode(route, node) {
				routeDeleated = append(routeDeleated, node)
			}
		}
		chromosomeDeleted = append(chromosomeDeleted, routeDeleated)
	}
	return chromosomeDeleted
}

func bestCostRouteCrossover(nodes *node.NodeList,
	ch1, ch2 [][]int) ([][]int, [][]int) {
	rand.Seed(time.Now().UnixNano())
	index1 := rand.Intn(len(ch1))
	index2 := rand.Intn(len(ch2))
	ch1 := deleteNodes(ch1, route2)
	ch2 := deleteNodes(ch2, route1)
	tmp1 := insertNode(nodes, ch1, route2)
	tmp2 := insertNode(nodes, ch2, route1)
	return tmp1, tmp2
}

// ===== </Best Cost Route Crossover (BCRC)> ===== //

//********//
// Public //
//********//
func Crossover(method string, nodes *node.NodeList,
	offsprings []*Individual, rate float64) []*Individual {
	population := len(offsprings)
	newOffsprings := make([]*Individual, 0, population)
	half := int(population / 2)

	halfList1 := offsprings[:half]
	halfList2 := offsprings[half:]

	for i := 0; i < half; i++ {
		tmp1 := copyIndividual(halfList1[i])
		tmp2 := copyIndividual(halfList2[i])
		rand.Seed(time.Now().UnixNano())
		uniform := rand.Float64()
		if uniform < rate {
			switch method {
			case "uox":
				tmp1.Chromosome, tmp2.Chromosome =
					uniformOrderCrossover(nodes, tmp1.Chromosome, tmp2.Chromosome)
			case "pmx":
				tmp1.Chromosome, tmp2.Chromosome =
					partiallyMappedCrossover(nodes, tmp1.Chromosome, tmp2.Chromosome)
			case "bcrc":
				tmp1.Chromosome, tmp2.Chromosome =
					bestCostRouteCrossover(nodes, tmp1.Chromosome, tmp2.Chromosome)
			default:
				fmt.Println("!!!!! [ga/Crossover] switch doesn't has such paramerter:",
					method, "!!!!!")
				os.Exit(0)
			}
		}
		tmp1.Chromosome = removeNullRoute(tmp1.Chromosome)
		tmp2.Chromosome = removeNullRoute(tmp2.Chromosome)
		newOffsprings = append(newOffsprings, tmp1)
		newOffsprings = append(newOffsprings, tmp2)
	}

	return newOffsprings
}
