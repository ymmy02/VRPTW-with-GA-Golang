package vrptw

import (
	"../ga"
	"../node"
	"../ut"
	"fmt"
	"sort"
)

type VRPTW struct {
	isOptimized   bool
	bestSolutions []*ga.Individual
	generations   []int
	nvehicleAvgs  []float64
	distanceAvgs  []float64
	nvehicleBests []float64
	distanceBests []float64
}

//******************//
// Methods of VRPTW //
//******************//
// = Private = //
func (v *VRPTW) record(selection string, generation int, indvList []*ga.Individual) {
	bestIdnvList := ut.PickUpBestIndvs(selection, indvList)

	nvehicleAvg := ut.CalcNvehicleAverage(indvList)
	distanceAvg := ut.CalcDistanceAverage(indvList)
	nvehicleBest := ut.CalcNvehicleAverage(bestIdnvList)
	distanceBest := ut.CalcDistanceAverage(bestIdnvList)

	v.generations = append(v.generations, generation)
	v.nvehicleAvgs = append(v.nvehicleAvgs, nvehicleAvg)
	v.distanceAvgs = append(v.distanceAvgs, distanceAvg)
	v.nvehicleBests = append(v.nvehicleBests, nvehicleBest)
	v.distanceBests = append(v.distanceBests, distanceBest)
}

// = Public = //
func (v *VRPTW) GAOptimize(nodes *node.NodeList, population int,
	generationSpan int, selection string, crossover string,
	mutation string, wNvehicle float64, wDistance float64,
	eliteSize int, tournamentSize int, cxRate float64, muRate float64) {
	//************//
	// Initialize //
	//************//
	parents := ga.CreateIndividualList(population, nodes)
	//offsprings := make([]*Individual, 0, population)
	ga.SetDistance(nodes, parents)
	// Evaluate Fitness
	switch selection {
	case "wsum":
		for i, indv := range parents {
			parents[i].Fitness =
				ga.WsumEvaluate(indv.NVehicle(),
					indv.Distance, wNvehicle, wDistance)
		}
	case "ranksum":
		nvehicleListTmp := make([]int, population)
		distanceListTmp := make([]float64, population)
		for i, indv := range parents {
			nvehicleListTmp[i] = indv.NVehicle()
			distanceListTmp[i] = indv.Distance
		}
		sort.Sort(sort.IntSlice(nvehicleListTmp))
		sort.Sort(sort.Float64Slice(distanceListTmp))
		nvehicleList := removeDuplicateInt(nvehicleListTmp)
		distanceList := removeDuplicateFloat64(distanceListTmp)
		for i, indv := range parents {
			nvehicle := indv.NVehicle()
			distance := indv.Distance
			parents[i].Fitness =
				float64(ut.FindIndexInt(nvehicleList, nvehicle) + 1)
			parents[i].Fitness +=
				float64(ut.FindIndexFloat64(distanceList, distance) + 1)
		}
	default:
	}
	v.record(selection, 0, parents)

	//***********//
	// Main Loop //
	//***********//
	for n := 1; n <= generationSpan; n++ {
		// Selection
		offsprings := ga.Selection(selection, parents, tournamentSize, eliteSize)
		// Crossover
		offsprings = ga.Crossover(crossover, nodes, offsprings, cxRate)
		// Mutation
		offsprings = ga.Mutation(mutation, nodes, offsprings, muRate)
		// Change Generation
		copy(parents, offsprings)
		ga.SetDistance(nodes, parents)
		// Evaluate Fitness
		switch selection {
		case "wsum":
			for i, indv := range parents {
				parents[i].Fitness =
					ga.WsumEvaluate(indv.NVehicle(),
						indv.Distance, wNvehicle, wDistance)
			}
		case "ranksum":
			nvehicleListTmp := make([]int, population)
			distanceListTmp := make([]float64, population)
			for i, indv := range parents {
				nvehicleListTmp[i] = indv.NVehicle()
				distanceListTmp[i] = indv.Distance
			}
			sort.Sort(sort.IntSlice(nvehicleListTmp))
			sort.Sort(sort.Float64Slice(distanceListTmp))
			nvehicleList := removeDuplicateInt(nvehicleListTmp)
			distanceList := removeDuplicateFloat64(distanceListTmp)
			for i, indv := range parents {
				nvehicle := indv.NVehicle()
				distance := indv.Distance
				parents[i].Fitness =
					float64(ut.FindIndexInt(nvehicleList, nvehicle) + 1)
				parents[i].Fitness +=
					float64(ut.FindIndexFloat64(distanceList, distance) + 1)
			}
		default:
		}
		v.bestSolutions = ut.PickUpBestIndvs(selection, parents)
		v.record(selection, n, parents)
		v.printLog(n)
	}

	v.isOptimized = true
}

//***********//
// Functions //
//***********//

// = Private = //
func removeDuplicateInt(list []int) []int {
	results := make([]int, 0, len(list))
	encountered := map[int]bool{}
	for i := 0; i < len(list); i++ {
		if !encountered[list[i]] {
			encountered[list[i]] = true
			results = append(results, list[i])
		}
	}
	return results
}

func removeDuplicateFloat64(list []float64) []float64 {
	results := make([]float64, 0, len(list))
	encountered := map[float64]bool{}
	for i := 0; i < len(list); i++ {
		if !encountered[list[i]] {
			encountered[list[i]] = true
			results = append(results, list[i])
		}
	}
	return results
}

func (v *VRPTW) printLog(generation int) {
	fmt.Println("### Best Solutions of Generation", generation, "###")
	for _, bestIndv := range v.bestSolutions {
		vehicles := bestIndv.NVehicle()
		distance := bestIndv.Distance
		fmt.Println("Vehicles :", vehicles, "Distance :", distance)
	}
}

// = Public = //
func CreateInstance(population, generationSpan int) *VRPTW {
	v := &VRPTW{
		isOptimized:   false,
		bestSolutions: make([]*ga.Individual, 0, 10),
		generations:   make([]int, 0, generationSpan+1),
		nvehicleAvgs:  make([]float64, 0, generationSpan+1),
		distanceAvgs:  make([]float64, 0, generationSpan+1),
		nvehicleBests: make([]float64, 0, generationSpan+1),
		distanceBests: make([]float64, 0, generationSpan+1),
	}
	return v
}
