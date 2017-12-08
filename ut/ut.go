package ut

import (
	"../ga"
	"fmt"
	"os"
	"strings"
)

//********//
// Public //
//********//
func FindIndexInt(list []int, target int) int {
	index := -1
	for i, value := range list {
		if value == target {
			index = i
			break
		}
	}
	return index
}

func FindIndexFloat64(list []float64, target float64) int {
	index := -1
	for i, value := range list {
		if value == target {
			index = i
			break
		}
	}
	return index
}

func VcFilename(filename string) string {
	paths := strings.Split(filename, "/")
	paths[len(paths)-1] = "vehicle_capacity.txt"
	vcFilename := strings.Join(paths, "/")
	return vcFilename
}

func CalcNvehicleAverage(indvList []*ga.Individual) float64 {
	var avg float64 = 0.0
	for _, indv := range indvList {
		avg += float64(indv.NVehicle())
	}
	return avg / float64(len(indvList))
}

func CalcDistanceAverage(indvList []*ga.Individual) float64 {
	var avg float64 = 0.0
	for _, indv := range indvList {
		avg += indv.Distance
	}
	return avg / float64(len(indvList))
}

func RemoveDuplication(indvList []*ga.Individual) []*ga.Individual {
	noduplList := make([]*ga.Individual, 0)
	noduplList = append(noduplList, indvList[0])
	for i, indv1 := range indvList[1:] {
		flagAdd := true
		for _, indv2 := range noduplList {
			if indv1.IsEqual(indv2) {
				flagAdd = false
				break
			}
		}
		if flagAdd {
			noduplList = append(noduplList, indvList[i])
		}
	}
	return noduplList
}

func PickUpBestIndvs(selection string, indvList []*ga.Individual) []*ga.Individual {
	bestSolutions := make([]*ga.Individual, 0)
	switch selection {
	case "pareto":
		bestSolutions, _ = ga.MakeCurrentRankingList(indvList)
	case "wsum", "ranksum":
		bestIndv := indvList[0]
		for i, indv := range indvList {
			if indv.Fitness < bestIndv.Fitness {
				bestIndv = indvList[i]
			}
		}
		bestSolutions = append(bestSolutions, bestIndv)
	default:
		fmt.Println("!!!!! [ut/PickUpBestIndvs] switch doesn't has such paramerter:",
			selection, "!!!!!")
		os.Exit(0)
	}

	return RemoveDuplication(bestSolutions)
}
