package ga

import (
	"../node"
	"../ut"
	"math/rand"
)

//*********//
// Private //
//*********//

// ===== <Uniform Order Crossover (UOX)> ===== //
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

//********//
// Public //
//********//
func Crossover(method string, nodes *node.NodeList,
	offsprings []*Individual, rate float64) []*Individual {
	population := len(offsprings)
	newOffsprings := make([]*Individual, 0, offsprings)
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
