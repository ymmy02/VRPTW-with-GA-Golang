package main

import (
	"./node"
	"./ut"
	"./vrptw"
	"os"
	"strconv"
)

func main() {

	filename := os.Args[1]
	outputPath := os.Args[2]
	population, _ := strconv.Atoi(os.Args[3])
	generationSpan, _ := strconv.Atoi(os.Args[4])
	selection := os.Args[5]
	crossover := os.Args[6]
	mutation := os.Args[7]
	wNvehicle, _ := strconv.ParseFloat(os.Args[8], 64)
	wDistance, _ := strconv.ParseFloat(os.Args[9], 64)
	eliteSize, _ := strconv.Atoi(os.Args[10])
	tournamentSize, _ := strconv.Atoi(os.Args[11])
	cxRate, _ := strconv.ParseFloat(os.Args[12], 64)
	muRate, _ := strconv.ParseFloat(os.Args[13], 64)
	suffix := 0
	if len(os.Args) > 14 {
		suffix, _ = strconv.Atoi(os.Args[14])
	}

	vcFilename := ut.VcFilename(filename)
	nodeFilename := filename
	nodes := node.LoadData(vcFilename, nodeFilename)

	v := vrptw.CreateInstance(population, generationSpan)
	v.GAOptimize(nodes, population, generationSpan,
		selection, crossover, mutation, wNvehicle,
		wDistance, eliteSize, tournamentSize, cxRate, muRate)

	generations, nvehicleAvgs, distanceAvgs,
		nvehicleBests, distanceBests := v.Records()
	bestSolutions := v.BestSolutions()
	ut.WriteResults(generations, nvehicleAvgs, distanceAvgs,
		nvehicleBests, distanceBests, outputPath, suffix)
	ut.WriteBestSolutions(bestSolutions, outputPath, suffix)
}
