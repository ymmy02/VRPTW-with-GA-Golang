package ut

import (
	"../ga"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EPSILON = 1e-5
)

//*********//
// Private //
//*********//
func addSuffix(file string, suffix int) string {
	var filename string
	if suffix == 0 {
		filename = file
	} else {
		s := fmt.Sprintf("_%03d", suffix)
		filename = file + s
	}
	return filename
}

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

func CalcFitnessAverage(indvList []*ga.Individual) float64 {
	var avg float64 = 0.0
	for _, indv := range indvList {
		avg += indv.Fitness
	}
	return avg / float64(len(indvList))
}

func RemoveDuplication(indvList []*ga.Individual) []*ga.Individual {
	noduplList := make([]*ga.Individual, 0)
	noduplList = append(noduplList, indvList[0])
	for i, indv1 := range indvList[1:] {
		flagAdd := true
		for _, indv2 := range noduplList {
			nvehicle1 := indv1.NVehicle()
			distance1 := indv1.Distance
			nvehicle2 := indv2.NVehicle()
			distance2 := indv2.Distance
			if indv1.IsEqual(indv2) {
				flagAdd = false
				break
			}
			if nvehicle1 == nvehicle2 &&
				distance1-distance2 < EPSILON {
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

func WriteResults(selection string, generations []int, nvehicleAvgs, distanceAvgs, fitnessAvgs,
	nvehicleBests, distanceBests, fitnessBests []float64, outputPath string, suffix int) {
	outputPath += "/"
	filename := addSuffix("output", suffix) + ".dat"
	file, err := os.Create(outputPath + filename)
	if err != nil {
		fmt.Println("File Open Error")
	}
	defer file.Close()

	switch selection {
	case "pareto":
		for i := 0; i < len(generations); i++ {
			ge := strconv.Itoa(generations[i]) + " "
			na := strconv.FormatFloat(nvehicleAvgs[i], 'f', 4, 64) + " "
			da := strconv.FormatFloat(distanceAvgs[i], 'f', 4, 64) + " "
			nb := strconv.FormatFloat(nvehicleBests[i], 'f', 4, 64) + " "
			db := strconv.FormatFloat(distanceBests[i], 'f', 4, 64) + "\n"
			line := ge + na + da + nb + db
			file.Write(([]byte)(line))
		}
	default:
		for i := 0; i < len(generations); i++ {
			ge := strconv.Itoa(generations[i]) + " "
			na := strconv.FormatFloat(nvehicleAvgs[i], 'f', 4, 64) + " "
			da := strconv.FormatFloat(distanceAvgs[i], 'f', 4, 64) + " "
			fa := strconv.FormatFloat(fitnessAvgs[i], 'f', 4, 64) + " "
			nb := strconv.FormatFloat(nvehicleBests[i], 'f', 4, 64) + " "
			db := strconv.FormatFloat(distanceBests[i], 'f', 4, 64) + " "
			fb := strconv.FormatFloat(fitnessBests[i], 'f', 4, 64) + "\n"
			line := ge + na + da + fa + nb + db + fb
			file.Write(([]byte)(line))
		}
	}
}

func WriteBestSolutions(selection string, bestSolutions []*ga.Individual,
	outputPath string, suffix int) {
	outputPath += "/"
	filename := addSuffix("best_solutions", suffix) + ".dat"

	// Number of Vehicles, Total Distance
	file, err := os.Create(outputPath + filename)
	if err != nil {
		fmt.Println("File Open Error")
	}
	defer file.Close()
	switch selection {
	case "pareto":
		for _, indv := range bestSolutions {
			nvehicle := strconv.Itoa(indv.NVehicle())
			distance := strconv.FormatFloat(indv.Distance, 'f', 4, 64)
			line := nvehicle + " " + distance + "\n"
			file.Write(([]byte)(line))
		}
	default:
		for _, indv := range bestSolutions {
			nvehicle := strconv.Itoa(indv.NVehicle())
			distance := strconv.FormatFloat(indv.Distance, 'f', 4, 64)
			fitness := strconv.FormatFloat(indv.Fitness, 'f', 4, 64)
			line := nvehicle + " " + distance + " " + fitness + "\n"
			file.Write(([]byte)(line))
		}
	}

	// Rkutings
	for i, indv := range bestSolutions {
		s := fmt.Sprintf("%03d", i)
		filename = addSuffix("routing"+s, suffix) + ".txt"
		f, e := os.Create(outputPath + filename)
		if e != nil {
			fmt.Println("File Open Error")
		}
		for _, route := range indv.Chromosome {
			line := make([]string, len(route))
			for j, node := range route {
				line[j] = strconv.Itoa(node)
			}
			f.Write(([]byte)(strings.Join(line, " ") + "\n"))
		}
	}
}
