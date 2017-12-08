package ga

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//*********//
// Private //
//*********//
func pop(indvList []*Individual, i int) ([]*Individual, *Individual) {
	indv := indvList[i]
	indvList = append(indvList[:i], indvList[i+1:]...)
	newIndvList := make([]*Individual, len(indvList))
	copy(newIndvList, indvList)
	return newIndvList, indv
}

func choice(indvList []*Individual) *Individual {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(indvList))
	return indvList[i]
}

func tournament(parents []*Individual, tournamentSize int, eliteSize int) []*Individual {
	population := len(parents)
	offsprings := make([]*Individual, 0, population)

	// Elitism
	for i := 0; i < eliteSize; i++ {
		var minFitness float64 = 1e14
		minIndex := -1
		for i, indv := range parents {
			if fitness := indv.Fitness; fitness < minFitness {
				minIndex = i
				minFitness = fitness
			}
		}
		var tmp *Individual
		parents, tmp = pop(parents, minIndex)
		offsprings = append(offsprings, tmp)
	}

	// Selection
	samples := make([]*Individual, tournamentSize)
	var tmp *Individual
	population -= eliteSize
	for i := 0; i < population; i++ {
		var minFitness float64 = 1e14
		for j := 0; j < tournamentSize; j++ {
			samples[j] = choice(parents)
		}
		for i, indv := range samples {
			if fitness := indv.Fitness; fitness < minFitness {
				tmp = samples[i]
				minFitness = fitness
			}
		}
		offsprings = append(offsprings, tmp)
	}

	return offsprings
}

func ranksum(parents []*Individual, tournamentSize int, eliteSize int) []*Individual {
	offsprings := tournament(parents, tournamentSize, eliteSize)
	return offsprings
}

func paretoRanking(parents []*Individual, eliteSize int) []*Individual {
	population := len(parents)
	offsprings := make([]*Individual, 0, population)

	rankingList := makeParetoRankingList(parents)
	size := len(rankingList)
	npart := float64((size * (size + 1)) / 2)
	part := 1.0 / npart

	rand.Seed(time.Now().UnixNano())
	uniform := rand.Float64()

	// Elitism
	count := 0
	for _, rank := range rankingList {
		for i := 0; i < len(rank); i++ {
			if count > eliteSize {
				break
			}
			tmp := rank[i]
			offsprings = append(offsprings, tmp)
			count++
		}
	}

	// Roulette Wheel Selection
	var span float64 = 0.0
	population -= eliteSize
	for n := 0; n < population; n++ {
		rand.Seed(time.Now().UnixNano())
		uniform = rand.Float64()
		for i := 0; i < size; n++ {
			span += float64(size-i) * part
			if uniform < span {
				tmp := choice(rankingList[i])
				offsprings = append(offsprings, tmp)
				break
			}
		}
	}

	return offsprings
}

//********//
// Public //
//********//
type SLParams struct {
	tournamentSize int
	eliteSize      int
}

func Selection(method string, parents []*Individual, tournamentSize, eliteSize int) []*Individual {
	var offsprings []*Individual
	switch method {
	case "wsum":
		//tournamentSize := params.tournamentSize
		//eliteSize := params.eliteSize
		offsprings = tournament(parents, tournamentSize, eliteSize)
	case "ranksum":
		//tournamentSize := params.tournamentSize
		//eliteSize := params.eliteSize
		offsprings = ranksum(parents, tournamentSize, eliteSize)
	case "pareto":
		//eliteSize := params.eliteSize
		offsprings = paretoRanking(parents, eliteSize)
	default:
		fmt.Println("!!!!! [ga/Selection] switch doesn't has such paramerter:",
			method, "!!!!!")
		os.Exit(0)
	}
	return offsprings
}
