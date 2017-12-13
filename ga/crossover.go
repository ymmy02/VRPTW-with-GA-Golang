package ga

import (
	"../node"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	flattench1 := Flatten(ch1)
	flattench2 := Flatten(ch2)
	size := len(flattench1)
	mask := make([]int, size)
	for i := 0; i < size; i++ {
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

	tmp1 := shapeFlatToVehicles(nodes, tmpch1)
	tmp2 := shapeFlatToVehicles(nodes, tmpch2)
	return tmp1, tmp2
}

// ===== </Uniform Order Crossover (UOX)> ===== //

// ===== <Partially Mapped Crossover (PMX)> ===== //
func getNoConflictList(origin []int, counterpart []int) []int {
	tmp := make([]int, len(origin))
	copy(tmp, origin)

	for i, node := range origin {
		if containsNode(counterpart, node) {
			tmp[i] = 0
		}
	}
	return tmp
}

func partiallyMappedCrossover(nodes *node.NodeList,
	ch1, ch2 [][]int) ([][]int, [][]int) {
	flattench1 := Flatten(ch1)
	flattench2 := Flatten(ch2)
	size := len(flattench1)
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
	tmp = flattench1[point1:point2]
	pre = getNoConflictList(flattench2[:point1], tmp)
	suf = getNoConflictList(flattench2[point2:], tmp)
	tmpch2 = append(pre, tmp...)
	tmpch2 = append(tmpch2, suf...)

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

	tmp1 := shapeFlatToVehicles(nodes, tmpch1)
	tmp2 := shapeFlatToVehicles(nodes, tmpch2)
	return tmp1, tmp2
}

// ===== </Partially Mapped Crossover (PMX)> ===== //

// ===== <Best Cost Route Crossover (BCRC)> ===== //
func insertNodeIntoRoute(route []int, insertNode, index int) []int {
	insertedRoute := make([]int, len(route)+1)
	tmp := append(route[:index+1], route[index:]...)
	copy(insertedRoute, tmp)
	insertedRoute[index] = insertNode
	return insertedRoute
}

func insertNode(nodes *node.NodeList,
	chromosome [][]int, L []int) [][]int {
	size := len(chromosome)
	newChromosome := make([][]int, 0, len(chromosome)+5)
	// Copy Chromosome
	for i := 0; i < size; i++ {
		tmp := make([]int, len(chromosome[i]))
		copy(tmp, chromosome[i])
		newChromosome = append(newChromosome, tmp)
	}

	for _, insertNode := range L {
		feasibleListI := make([]int, 0)
		feasibleListJ := make([]int, 0)
		for i, route := range chromosome {
			for j := 0; j < len(route); j++ {
				tmp := insertNodeIntoRoute(route, insertNode, j)
				if nodes.IsFeasible(tmp) {
					feasibleListI = append(feasibleListI, i)
					feasibleListJ = append(feasibleListJ, j)
				}
			}
		}
		if len(feasibleListI) == 0 {
			newRoute := make([]int, 1)
			newRoute[0] = insertNode
			newChromosome = append(newChromosome, newRoute)
		} else {
			index := rand.Intn(len(feasibleListI))
			indexI := feasibleListI[index]
			indexJ := feasibleListJ[index]
			tmp := insertNodeIntoRoute(newChromosome[indexI], insertNode, indexJ)
			newChromosome[indexI] = tmp
		}
	}
	return newChromosome
}

func deleteNodes(chromosome [][]int, route []int) [][]int {
	chromosomeDeleted := make([][]int, 0, len(chromosome))
	for _, rt := range chromosome {
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
	index1 := rand.Intn(len(ch1))
	index2 := rand.Intn(len(ch2))
	route1 := ch1[index1]
	route2 := ch2[index2]
	ch1 = deleteNodes(ch1, route2)
	ch2 = deleteNodes(ch2, route1)
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
	rand.Seed(time.Now().UnixNano())

	halfList1 := offsprings[:half]
	halfList2 := offsprings[half:]

	for i := 0; i < half; i++ {
		tmp1 := copyIndividual(halfList1[i])
		tmp2 := copyIndividual(halfList2[i])
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
