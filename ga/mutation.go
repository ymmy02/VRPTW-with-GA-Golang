package ga

import (
	"../node"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

//*********//
// Private //
//*********//
func inversion(nodes *node.NodeList, chromosome [][]int) [][]int {
	flattench := Flatten(chromosome)
	size := len(flattench)
	cut1 := rand.Intn(size - 1)
	cut2 := rand.Intn(size)
	for cut1 == cut2 {
		cut2 = rand.Intn(size)
	}
	if cut1 > cut2 {
		cut1, cut2 = cut2, cut1
	}
	reversePart := flattench[cut1:cut2]
	sort.Sort(sort.Reverse(sort.IntSlice(reversePart)))
	newChromosome := shapeFlatToVehicles(nodes, flattench)
	return newChromosome
}

//********//
// Public //
//********//
func Mutation(method string, nodes *node.NodeList,
	offsprings []*Individual, rate float64) []*Individual {
	population := len(offsprings)
	newOffsprings := make([]*Individual, 0, population)
	for i := 0; i < population; i++ {
		rand.Seed(time.Now().UnixNano())
		uniform := rand.Float64()
		tmp := copyIndividual(offsprings[i])
		if uniform < rate {
			switch method {
			case "inversion":
				tmp.Chromosome = inversion(nodes, tmp.Chromosome)
			default:
				fmt.Println("!!!!! [ga/Mutation] switch doesn't has such paramerter:",
					method, "!!!!!")
				os.Exit(0)
			}
		}
		tmp.Chromosome = removeNullRoute(tmp.Chromosome)
		newOffsprings = append(newOffsprings, tmp)
	}
	return newOffsprings
}
