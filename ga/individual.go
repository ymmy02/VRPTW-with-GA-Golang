package ga

import (
	"../node"
	"reflect"
)

type Individual struct {
	Chromosome [][]int
	Distance   float64
	Fitness    float64
}

//***********************//
// Methods of Individual //
//***********************//
func (indv *Individual) IsEqual(counterpart *Individual) bool {
	result := reflect.DeepEqual(indv.Chromosome, counterpart.Chromosome)
	return result
}

func (indv *Individual) NVehicle() int {
	nvehicle := len(indv.Chromosome)
	return nvehicle
}

//***********//
// Functions //
//***********//
func CreateIndividual(nodes *node.NodeList) *Individual {
	flattench := make([]int, len(nodes.CusotmersIDList()))
	chromosome := make([][]int, 0)
	customersIDList := nodes.CusotmersIDList()
	copy(flattench, customersIDList)
	shuffle(flattench)
	chromosome = shapeFlatToVehicles(nodes, flattench)
	indv := &Individual{Chromosome: make([][]int, 0)}
	indv.Chromosome = chromosome
	return indv
}
